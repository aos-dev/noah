package task

import (
	"bytes"
	"context"
	"sync"

	"github.com/aos-dev/go-storage/v2"
	_ "github.com/aos-dev/go-storage/v2/types"
	typ "github.com/aos-dev/go-storage/v2/types"
	"github.com/aos-dev/go-storage/v2/types/pairs"

	"github.com/qingstor/noah/constants"
	"github.com/qingstor/noah/pkg/progress"
	"github.com/qingstor/noah/pkg/types"
	"github.com/qingstor/noah/utils"
)

func (t *CopyDirTask) new() {}
func (t *CopyDirTask) run(ctx context.Context) error {
	x := NewListDir(t)
	err := utils.ChooseSourceStorageAsDirLister(x, t)
	if err != nil {
		return err
	}

	x.SetFileFunc(func(o *typ.Object) {
		sf := NewCopyFile(t)
		sf.SetSourcePath(o.Name)
		sf.SetDestinationPath(o.Name)
		if t.ValidateHandleObjCallbackFunc() {
			sf.SetCallbackFunc(func() {
				t.GetHandleObjCallbackFunc()(o)
			})
		}
		t.Async(ctx, sf)
	})
	x.SetDirFunc(func(o *typ.Object) {
		sf := NewCopyDir(t)
		sf.SetSourcePath(o.Name)
		sf.SetDestinationPath(o.Name)
		t.Sync(ctx, sf)
	})
	return t.Sync(ctx, x)
}

func (t *CopyFileTask) new() {}
func (t *CopyFileTask) run(ctx context.Context) error {
	check := NewBetweenStorageCheck(t)
	if err := t.Sync(ctx, check); err != nil {
		return err
	}

	// Execute check tasks
	if t.ValidateCheckTasks() {
		for _, v := range t.GetCheckTasks() {
			ct := v(check)
			if err := t.Sync(ctx, ct); err != nil {
				return err
			}
			// If either check not pass, do not copy this file.
			if result := ct.(types.ResultGetter); !result.GetResult() {
				return nil
			}
			// If all check passed, we should continue do copy works.
		}
	}

	srcSize := check.GetSourceObject().Size

	// if destination not support segmenter, we do not call multipart api
	_, ok := t.GetDestinationStorage().(storage.IndexSegmenter)
	if !ok {
		x := NewCopySmallFile(t)
		x.SetSize(srcSize)
		return t.Sync(ctx, x)
	}

	if srcSize <= t.GetPartThreshold() {
		x := NewCopySmallFile(t)
		x.SetSize(srcSize)
		return t.Sync(ctx, x)
	}

	x := NewCopyLargeFile(t)
	x.SetTotalSize(srcSize)
	return t.Sync(ctx, x)
}

func (t *CopySmallFileTask) new() {}
func (t *CopySmallFileTask) run(ctx context.Context) error {
	fileCopyTask := NewCopySingleFile(t)

	if t.ValidateCheckMD5() && t.GetCheckMD5() {
		md5Task := NewMD5SumFile(t)
		utils.ChooseSourceStorage(md5Task, t)
		md5Task.SetOffset(0)
		if err := t.Sync(ctx, md5Task); err != nil {
			return err
		}
		fileCopyTask.SetMD5Sum(md5Task.GetMD5Sum())
	} else {
		fileCopyTask.SetMD5Sum(nil)
	}

	return t.Sync(ctx, fileCopyTask)
}

func (t *CopyLargeFileTask) new() {}
func (t *CopyLargeFileTask) run(ctx context.Context) error {
	// if part size set, use it directly
	// TODO: we need to check validation of part size
	if !t.ValidatePartSize() {
		// otherwise, calculate part size.
		partSize, err := utils.CalculatePartSize(t.GetTotalSize())
		if err != nil {
			return types.NewErrUnhandled(err)
		}
		t.SetPartSize(partSize)
	}

	initTask := NewSegmentInit(t)
	err := utils.ChooseDestinationStorageAsIndexSegmenter(initTask, t)
	if err != nil {
		return err
	}

	if err := t.Sync(ctx, initTask); err != nil {
		return err
	}
	t.SetSegment(initTask.GetSegment())

	offset, part := int64(0), 0
	for {
		t.SetOffset(offset)

		x := NewCopyPartialFile(t)
		x.SetIndex(part)
		t.Async(ctx, x)
		// While GetDone is true, this must be the last part.
		if x.GetDone() {
			break
		}

		offset += x.GetSize()
		part++
	}

	// Make sure all segment upload finished.
	t.Await()
	if t.GetFault().HasError() {
		return t.GetFault()
	}
	return t.Sync(ctx, NewSegmentCompleteTask(initTask))
}

func (t *CopyPartialFileTask) new() {
	totalSize := t.GetTotalSize()
	partSize := t.GetPartSize()
	offset := t.GetOffset()

	if totalSize <= offset+partSize {
		t.SetSize(totalSize - offset)
		t.SetDone(true)
	} else {
		t.SetSize(partSize)
		t.SetDone(false)
	}
}
func (t *CopyPartialFileTask) run(ctx context.Context) error {
	fileCopyTask := NewSegmentFileCopy(t)

	if t.ValidateCheckMD5() && t.GetCheckMD5() {
		md5Task := NewMD5SumFile(t)
		utils.ChooseSourceStorage(md5Task, t)
		if err := t.Sync(ctx, md5Task); err != nil {
			return err
		}
		fileCopyTask.SetMD5Sum(md5Task.GetMD5Sum())
	} else {
		fileCopyTask.SetMD5Sum(nil)
	}

	err := utils.ChooseDestinationIndexSegmenter(fileCopyTask, t)
	if err != nil {
		return err
	}
	return t.Sync(ctx, fileCopyTask)
}

func (t *CopyStreamTask) new() {
	bytesPool := &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, t.GetPartSize()))
		},
	}
	t.SetBytesPool(bytesPool)
}
func (t *CopyStreamTask) run(ctx context.Context) error {
	initTask := NewSegmentInit(t)
	err := utils.ChooseDestinationStorageAsIndexSegmenter(initTask, t)
	if err != nil {
		return err
	}

	// TODO: we will use expect size to calculate part size later.
	// if part size not set, use default value, otherwise load part size automatically
	if !t.ValidatePartSize() {
		t.SetPartSize(constants.DefaultPartSize)
	}

	if err := t.Sync(ctx, initTask); err != nil {
		return err
	}
	t.SetSegment(initTask.GetSegment())

	offset, part := int64(0), 0
	for {
		it := NewSegmentStreamInit(t)
		if err := t.Sync(ctx, it); err != nil {
			return err
		}

		x := NewCopyPartialStream(t)
		x.SetSize(it.GetSize())
		x.SetContent(it.GetContent())
		x.SetOffset(offset)
		x.SetIndex(part)
		t.Async(ctx, x)

		if it.GetDone() {
			break
		}

		offset += it.GetSize()
		part++
	}

	t.Await()
	if t.GetFault().HasError() {
		return t.GetFault()
	}
	return t.Sync(ctx, NewSegmentCompleteTask(initTask))
}

func (t *CopyPartialStreamTask) new() {}
func (t *CopyPartialStreamTask) run(ctx context.Context) error {
	copyTask := NewSegmentStreamCopy(t)
	if t.GetCheckMD5() {
		md5sumTask := NewMD5SumStream(t)
		if err := t.Sync(ctx, md5sumTask); err != nil {
			return err
		}
		copyTask.SetMD5Sum(md5sumTask.GetMD5Sum())
	} else {
		copyTask.SetMD5Sum(nil)
	}

	err := utils.ChooseDestinationIndexSegmenter(copyTask, t)
	if err != nil {
		return err
	}

	return t.Sync(ctx, copyTask)
}

func (t *CopySingleFileTask) new() {}
func (t *CopySingleFileTask) run(ctx context.Context) error {
	r, err := t.GetSourceStorage().ReadWithContext(ctx, t.GetSourcePath())
	if err != nil {
		return types.NewErrUnhandled(err)
	}
	defer r.Close()

	// improve progress bar's performance, do not set state for small files less than 32M
	if t.GetSize() > constants.DefaultPartSize/4 {
		progress.SetState(t.GetID(), progress.InitIncState(t.GetSourcePath(), "copy:", t.GetSize()))
	}
	// TODO: add checksum support
	writeDone := 0
	err = t.GetDestinationStorage().WriteWithContext(ctx, t.GetDestinationPath(), r, pairs.WithSize(t.GetSize()),
		pairs.WithReadCallbackFunc(func(b []byte) {
			writeDone += len(b)
			progress.UpdateState(t.GetID(), int64(writeDone))
		}))
	if err != nil {
		return types.NewErrUnhandled(err)
	}
	return nil
}
