package task

import (
	"github.com/qingstor/noah/pkg/types"
)

func (t *StatFileTask) new() {}
func (t *StatFileTask) run() {
	om, err := t.GetStorage().Stat(t.GetPath())
	if err != nil {
		t.TriggerFault(types.NewErrUnhandled(err))
		return
	}
	t.SetObject(om)
}
