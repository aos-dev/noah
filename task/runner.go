package task

import (
	"context"
	"fmt"
	"sync"

	"github.com/aos-dev/go-storage/v3/types"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"github.com/aos-dev/noah/proto"
)

type Runner struct {
	agent *Worker
	j     *proto.Job

	queue    *nats.EncodedConn
	subject  string // All runner will share the same task subject
	storages []types.Storager

	wg     *sync.WaitGroup
	sub    *nats.Subscription
	once   sync.Once
	logger *zap.Logger
}

func NewRunner(a *Worker, j *proto.Job) *Runner {
	return &Runner{
		j:     j,
		agent: a,

		queue:    a.queue,
		subject:  a.subject, // Copy task subject from agent.
		storages: a.storages,
		wg:       &sync.WaitGroup{},
		once:     sync.Once{},
		logger:   a.logger,
	}
}

func (rn *Runner) Handle(subject, reply string) {
	rn.logger.Info("runner start handle job", zap.String("id", rn.j.Id), zap.Uint32("type", rn.j.Type))

	ctx := context.Background()

	var fn func(ctx context.Context, msg protobuf.Message) error
	var t protobuf.Message

	switch rn.j.Type {
	case TypeCopyDir:
		t = &proto.CopyDir{}
		fn = rn.HandleCopyDir
	case TypeCopyFile:
		t = &proto.CopyFile{}
		fn = rn.HandleCopyFile
	case TypeCopySingleFile:
		t = &proto.CopySingleFile{}
		fn = rn.HandleCopySingleFile
	case TypeCopyMultipartFile:
		t = &proto.CopyMultipartFile{}
		fn = rn.HandleCopyMultipartFile
	case TypeCopyMultipart:
		t = &proto.CopyMultipart{}
		fn = rn.HandleCopyMultipart
	default:
		panic("not support job type")
	}

	err := protobuf.Unmarshal(rn.j.Content, t)
	if err != nil {
		panic("unmarshal failed")
	}

	err = fn(ctx, t)
	if err != nil {
		// TODO: reply to message sender the response.
		panic(fmt.Errorf("handle task: %s", err))
	}

	// Send JobReply after the job has been handled.
	err = rn.Finish(ctx, reply)
	if err != nil {
		return
	}
}

func (rn *Runner) Async(ctx context.Context, job *proto.Job) (err error) {
	logger := rn.logger

	rn.once.Do(func() {
		// Wait for all JobReply sending to the reply subject.
		rn.sub, err = rn.queue.Subscribe(SubjectJobReply(rn.j.Id), func(job *proto.JobReply) {
			defer rn.wg.Done()

			switch job.Status {
			case JobStatusSucceed:
				rn.logger.Info("job succeed", zap.String("id", job.Id), zap.String("job", rn.j.Id))
			default:
				rn.logger.Error("job failed", zap.String("id", job.Id), zap.String("job", rn.j.Id),
					zap.String("error", job.Message),
				)
			}
		})
	})
	if err != nil {
		return err
	}

	rn.wg.Add(1)

	// Publish new job with the specific reply subject on the task subject.
	// After this job finished, the runner will send a JobReply to the reply subject.
	err = rn.queue.PublishRequest(rn.subject, SubjectJobReply(rn.j.Id), job)
	if err != nil {
		return fmt.Errorf("nats publish: %w", err)
	}

	logger.Info("publish async", zap.String("subject", rn.subject), zap.String("job", rn.j.Id))
	return
}

func (rn *Runner) Await(ctx context.Context) (err error) {
	rn.logger.Info("start await", zap.String("job", rn.j.Id))

	rn.wg.Wait()
	rn.sub.Unsubscribe()
	// Reset once so we can async again.
	rn.once = sync.Once{}

	rn.logger.Info("finish await", zap.String("job", rn.j.Id))
	return
}

func (rn *Runner) Sync(ctx context.Context, job *proto.Job) (err error) {
	logger := rn.logger

	var reply proto.JobReply

	logger.Info("start sync job", zap.String("id", job.Id), zap.String("job", rn.j.Id))

	// NATS provides the builtin request-response style API, so that we don't need to
	// care about the reply id.
	err = rn.queue.RequestWithContext(ctx, rn.subject, job, &reply)
	if err != nil {
		return fmt.Errorf("nats request: %w", err)
	}

	if reply.Status != JobStatusSucceed {
		logger.Error("job synced", zap.String("error", reply.Message))
		return fmt.Errorf("job failed: %v", reply.Message)
	}

	logger.Info("job synced", zap.String("id", job.Id), zap.String("job", rn.j.Id))
	return
}

func (rn *Runner) Finish(ctx context.Context, reply string) (err error) {
	logger := rn.logger

	logger.Info("send reply", zap.String("job", rn.j.Id), zap.String("reply", reply))
	return rn.queue.Publish(reply, &proto.JobReply{
		Id:      rn.j.Id, // Make sure JobReply sends to the parent job.
		Status:  JobStatusSucceed,
		Message: "",
	})
}
