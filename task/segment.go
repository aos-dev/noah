package task

import (
	"io/ioutil"

	"github.com/Xuanwo/storage/types/pairs"

	"github.com/qingstor/noah/pkg/progress"
	"github.com/qingstor/noah/pkg/types"
)

func (t *SegmentInitTask) new() {}
func (t *SegmentInitTask) run() {
	id, err := t.GetSegmenter().InitSegment(t.GetPath(),
		pairs.WithPartSize(t.GetPartSize()))
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	t.SetSegmentID(id)
}

func (t *SegmentFileCopyTask) new() {}
func (t *SegmentFileCopyTask) run() {
	r, err := t.GetSourceStorage().Read(t.GetSourcePath(), pairs.WithSize(t.GetSize()), pairs.WithOffset(t.GetOffset()))
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	defer r.Close()

	// TODO: Add checksum support.
	writeDone := 0
	err = t.GetDestinationSegmenter().WriteSegment(t.GetSegmentID(), t.GetOffset(), t.GetSize(), r,
		pairs.WithReadCallbackFunc(func(b []byte) {
			writeDone += len(b)
			progress.SetState(t.GetID(), progress.NewState(t.Name(), "writing", int64(writeDone), t.GetSize()))
		}),
	)
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	progress.SetState(t.GetID(), progress.FinishedState(t.Name(), int64(writeDone)))
}

func (t *SegmentStreamCopyTask) new() {}
func (t *SegmentStreamCopyTask) run() {
	// TODO: Add checksum support
	writeDone := 0
	err := t.GetDestinationSegmenter().WriteSegment(t.GetSegmentID(), t.GetOffset(), t.GetSize(), ioutil.NopCloser(t.GetContent()),
		pairs.WithReadCallbackFunc(func(b []byte) {
			writeDone += len(b)
			progress.SetState(t.GetID(), progress.NewState(t.Name(), "writing", int64(writeDone), t.GetSize()))
		}))
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	progress.SetState(t.GetID(), progress.FinishedState(t.Name(), int64(writeDone)))
}

func (t *SegmentCompleteTask) new() {}
func (t *SegmentCompleteTask) run() {
	err := t.GetSegmenter().CompleteSegment(t.GetSegmentID())
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
}
