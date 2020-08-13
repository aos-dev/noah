package task

import (
	"context"
	"fmt"
	"testing"

	"github.com/Xuanwo/navvy"
	"github.com/aos-dev/go-storage/v2"
	"github.com/aos-dev/go-storage/v2/types"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/qingstor/noah/pkg/fault"
	"github.com/qingstor/noah/pkg/mock"
)

func TestSyncTask_run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	sche := mock.NewMockScheduler(ctrl)
	srcStore := mock.NewMockStorager(ctrl)
	dstStore := mock.NewMockStorager(ctrl)
	sourcePath := uuid.New().String()
	dstPath := uuid.New().String()
	obj := &types.Object{Name: "obj-name"}

	srcStore.EXPECT().String().Do(func() {}).AnyTimes()
	dstStore.EXPECT().String().Do(func() {}).AnyTimes()

	dstStore.EXPECT().StatWithContext(gomock.Eq(ctx), gomock.Any()).
		Do(func(ctx context.Context, path string) {
			return
		}).AnyTimes()

	srcStore.EXPECT().StatWithContext(gomock.Eq(ctx), gomock.Any()).
		Do(func(ctx context.Context, path string) {
			return
		}).AnyTimes()

	t.Run("without flag", func(t *testing.T) {
		srcLister := mock.NewMockDirLister(ctrl)
		src := struct {
			storage.Storager
			storage.DirLister
		}{
			srcStore, srcLister,
		}

		task := SyncTask{}
		task.SetPool(navvy.NewPool(10))
		task.SetScheduler(sche)
		task.SetFault(fault.New())
		task.SetSourcePath(sourcePath)
		task.SetSourceStorage(src)
		task.SetDestinationStorage(dstStore)
		task.SetDestinationPath(dstPath)
		task.SetDryRun(false)
		task.SetDryRunFunc(nil)
		task.SetExisting(false)
		task.SetIgnoreExisting(false)
		task.SetRecursive(false)
		task.SetUpdate(false)
		task.SetCheckMD5(false)

		sche.EXPECT().Sync(gomock.Eq(ctx), gomock.Any()).Do(func(ctx context.Context, task navvy.Task) {
			switch v := task.(type) {
			case *ListDirTask:
				v.SetDirFunc(nil)
				v.validateInput()
				v.GetFileFunc()(obj)
			default:
				panic(fmt.Errorf("unexpected task %v", v))
			}
		}).AnyTimes()

		sche.EXPECT().Async(gomock.Eq(ctx), gomock.Any()).Do(func(ctx context.Context, task navvy.Task) {
			switch v := task.(type) {
			case *CopyFileTask:
				v.validateInput()
				assert.Equal(t, obj.Name, v.GetSourcePath())
				assert.Equal(t, obj.Name, v.GetDestinationPath())
			default:
				panic(fmt.Errorf("unexpected task %v", v))
			}
		}).AnyTimes()

		task.run(ctx)
		assert.Empty(t, task.GetFault().Error())
	})

	t.Run("with all flags", func(t *testing.T) {
		srcLister := mock.NewMockDirLister(ctrl)
		src := struct {
			storage.Storager
			storage.DirLister
		}{
			srcStore, srcLister,
		}

		task := SyncTask{}
		task.SetPool(navvy.NewPool(10))
		task.SetScheduler(sche)
		task.SetFault(fault.New())
		task.SetSourcePath(sourcePath)
		task.SetSourceStorage(src)
		task.SetDestinationStorage(dstStore)
		task.SetDestinationPath(dstPath)
		task.SetDryRun(false)
		task.SetDryRunFunc(nil)
		task.SetExisting(true)
		task.SetIgnoreExisting(true)
		task.SetRecursive(false)
		task.SetUpdate(true)
		task.SetCheckMD5(false)

		sche.EXPECT().Sync(gomock.Eq(ctx), gomock.Any()).Do(func(ctx context.Context, task navvy.Task) {
			switch v := task.(type) {
			case *ListDirTask:
				v.SetDirFunc(nil)
				v.validateInput()
				v.GetFileFunc()(obj)
			default:
				panic(fmt.Errorf("unexpected task %v", v))
			}
		}).AnyTimes()

		task.run(ctx)
		assert.Empty(t, task.GetFault().Error())
	})

	t.Run("with dry-run", func(t *testing.T) {
		srcLister := mock.NewMockDirLister(ctrl)
		src := struct {
			storage.Storager
			storage.DirLister
		}{
			srcStore, srcLister,
		}

		task := SyncTask{}
		task.SetPool(navvy.NewPool(10))
		task.SetScheduler(sche)
		task.SetFault(fault.New())
		task.SetSourcePath(sourcePath)
		task.SetSourceStorage(src)
		task.SetDestinationStorage(dstStore)
		task.SetDestinationPath(dstPath)
		task.SetDryRun(true)
		task.SetDryRunFunc(func(o *types.Object) {
			t.Log(o.Name)
		})
		task.SetExisting(false)
		task.SetIgnoreExisting(false)
		task.SetRecursive(false)
		task.SetUpdate(false)
		task.SetCheckMD5(false)

		sche.EXPECT().Sync(gomock.Eq(ctx), gomock.Any()).Do(func(ctx context.Context, task navvy.Task) {
			switch v := task.(type) {
			case *ListDirTask:
				v.SetDirFunc(nil)
				v.validateInput()
				v.GetFileFunc()(obj)
			default:
				panic(fmt.Errorf("unexpected task %v", v))
			}
		}).AnyTimes()

		task.run(ctx)
		assert.Empty(t, task.GetFault().Error())
	})

	t.Run("with recursive", func(t *testing.T) {
		srcLister := mock.NewMockDirLister(ctrl)
		src := struct {
			storage.Storager
			storage.DirLister
		}{
			srcStore, srcLister,
		}

		task := SyncTask{}
		task.SetPool(navvy.NewPool(10))
		task.SetScheduler(sche)
		task.SetFault(fault.New())
		task.SetSourcePath(sourcePath)
		task.SetSourceStorage(src)
		task.SetDestinationStorage(dstStore)
		task.SetDestinationPath(dstPath)
		task.SetDryRun(true)
		task.SetDryRunFunc(func(o *types.Object) {
			t.Log(o.Name)
		})
		task.SetExisting(false)
		task.SetIgnoreExisting(false)
		task.SetRecursive(true)
		task.SetUpdate(false)
		task.SetCheckMD5(false)

		sche.EXPECT().Sync(gomock.Eq(ctx), gomock.Any()).Do(func(ctx context.Context, task navvy.Task) {
			switch v := task.(type) {
			case *ListDirTask:
				v.SetDirFunc(nil)
				v.SetFileFunc(nil)
				v.validateInput()
			default:
				panic(fmt.Errorf("unexpected task %v", v))
			}
		}).AnyTimes()

		task.run(ctx)
		assert.Empty(t, task.GetFault().Error())
	})
}
