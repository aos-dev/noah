// The following directive is necessary to make the package coherent:
// This program generates types, It can be invoked by running
// go generate
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

type task struct {
	Name        string   `json:"-"`
	Description string   `json:"description"`
	Input       Input    `json:"input,omitempty"`
	Output      []string `json:"output,omitempty"`
}

// Input include required fields and optional fields
type Input struct {
	Required []string `json:"required,omitempty"`
	Optional []string `json:"optional,omitempty"`
}

var funcs = template.FuncMap{
	"lowerFirst": func(s string) string {
		if len(s) == 0 {
			return ""
		}
		if s[0] < 'A' || s[0] > 'Z' {
			return s
		}
		return string(s[0]+'a'-'A') + s[1:]
	},
	"endwith": func(x, y string) bool {
		return strings.HasSuffix(x, y)
	},
	"merge": func(x, y []string) []string {
		a := make(map[string]struct{})
		for _, v := range x {
			a[v] = struct{}{}
		}
		for _, v := range y {
			a[v] = struct{}{}
		}
		o := make([]string, 0)
		for x := range a {
			o = append(o, x)
		}

		sort.Strings(o)
		return o
	},
	"add": func(x, y int) int {
		return x + y
	},
}

func executeTemplate(tmpl *template.Template, w io.Writer, v interface{}) {
	err := tmpl.Execute(w, v)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	tasks := make(map[string]*task)
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	// Do sort to all tasks via name.
	taskNames := make([]string, 0)
	for k := range tasks {
		sort.Strings(tasks[k].Input.Required)
		sort.Strings(tasks[k].Input.Optional)
		sort.Strings(tasks[k].Output)

		taskNames = append(taskNames, k)
	}
	sort.Strings(taskNames)

	// Format input tasks.json
	data, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("tasks.json", data, 0664)
	if err != nil {
		log.Fatal(err)
	}

	taskFile, err := os.Create("generated.go")
	if err != nil {
		log.Fatal(err)
	}
	defer taskFile.Close()

	executeTemplate(taskPageTmpl, taskFile, nil)

	testFile, err := os.Create("generated_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()

	executeTemplate(testPageTmpl, testFile, nil)

	for _, taskName := range taskNames {
		v := tasks[taskName]
		v.Name = taskName

		executeTemplate(taskTmpl, taskFile, v)
		executeTemplate(testTmpl, testFile, v)
	}
}

var taskPageTmpl = template.Must(template.New("taskPage").Parse(`// Code generated by go generate; DO NOT EDIT.
package task

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/qingstor/log"

	"github.com/qingstor/noah/pkg/schedule"
	"github.com/qingstor/noah/pkg/task"
	"github.com/qingstor/noah/pkg/types"
)

var _ = uuid.New()
`))

var taskTmpl = template.Must(template.New("task").Funcs(funcs).Parse(`
// {{ .Name }}Task will {{ .Description }}.
type {{ .Name }}Task struct {
	// Predefined value
	types.Scheduler
	types.ID
	types.CallbackFunc

	// Required Input value
{{- range $k, $v := .Input.Required }}
	types.{{$v}}
{{- end }}

	// Optional Input value
{{- range $k, $v := .Input.Optional }}
	types.{{$v}}
{{- end }}

	// Output value
{{- range $k, $v := .Output }}
	types.{{$v}}
{{- end }}
}

// New{{ .Name }} will create a {{ .Name }}Task struct and fetch inherited data from parent task.
func New{{ .Name }}(task task.Task) *{{ .Name }}Task {
	t := &{{ .Name }}Task{}
	t.SetScheduler(schedule.New())
	t.SetID(uuid.New().String())

	t.loadInput(task)

	t.new()
	return t
}

// validateInput will validate all input before run task.
func (t *{{ .Name }}Task) validateInput() {
{{- range $k, $v := .Input.Required }}
	if !t.Validate{{$v}}() {
		panic(fmt.Errorf("Task {{ $.Name }} value {{$v}} is invalid"))
	}
{{- end }}
}

// loadInput will check and load all input before new task.
func (t *{{ .Name }}Task) loadInput(task task.Task) {
	// load required fields
{{- range $k, $v := .Input.Required }}
	types.Load{{$v}}(task, t)
{{- end }}
	// load optional fields
{{- range $k, $v := .Input.Optional }}
	types.Load{{$v}}(task, t)
{{- end }}
}

// Sync run sub task directly
func (t *{{ .Name }}Task) Sync(ctx context.Context, st task.Task) error {
	return st.Run(ctx)
}

// Async run sub task asynchronously
func (t *{{ .Name }}Task) Async(ctx context.Context, st task.Task) {
	t.GetScheduler().Add(1)
	go func() {
        defer t.GetScheduler().Done()
        if err := st.Run(ctx); err != nil {
			t.TriggerFault(err)
        }
    }()
}

// Await wait sub task done
func (t *{{ .Name }}Task) Await() error {
	return t.GetScheduler().Await()
}

// Run implement task.Task
func (t *{{ .Name }}Task) Run(ctx context.Context) error {
	logger := log.FromContext(ctx)
	t.validateInput()

	logger.Debug(
		log.String("task_started", t.String()),
	)
	err := t.run(ctx)
	if err != nil {
		t.TriggerFault(err)
	}

	if err := t.Await(); err != nil {
		logger.Debug(
			log.String("task_failed", "{{ .Name }}Task"),
			log.String("err", err.Error()),
		)
		return err
	}
	if t.ValidateCallbackFunc() {
		t.GetCallbackFunc()()
	}
	logger.Debug(
		log.String("task_finished", t.String()),
	)
	return nil
}

// TriggerFault will be used to trigger a task related fault.
func (t *{{ .Name }}Task) TriggerFault(err error) {
	t.GetScheduler().AppendFault(fmt.Errorf("Failed %s: {%w}", t, err))
}

// String will implement Stringer interface.
func (t *{{ .Name }}Task) String() string {
{{- $length := add (len .Input.Required) (len .Input.Optional) }}
	s := make([]string, 0, {{$length}})
{{ range $k, $v := .Input.Required -}}
{{- if not (endwith $v "Func") }}
	s = append(s, fmt.Sprintf("{{$v}}: %s", t.{{$v}}.String()))
{{- end -}}
{{ end }}
{{- range $k, $v := .Input.Optional -}}
{{- if not (endwith $v "Func") }}
	if t.Validate{{$v}}() {
		s = append(s, fmt.Sprintf("{{$v}}: %s", t.{{$v}}.String()))
	}
{{- end -}}
{{- end }}
	return fmt.Sprintf("{{ .Name }}Task {%s}", strings.Join(s, ", "))
}

// New{{ .Name }}Task will create a {{ .Name }}Task which meets task.Task.
func New{{ .Name }}Task(task task.Task) task.Task {
	return New{{ .Name }}(task)
}
`))

var testPageTmpl = template.Must(template.New("testPage").Parse(`// Code generated by go generate; DO NOT EDIT.
package task

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qingstor/noah/pkg/schedule"
	"github.com/qingstor/noah/pkg/types"
)

`))

var testTmpl = template.Must(template.New("test").Funcs(funcs).Parse(`
func Test{{ .Name }}Task_TriggerFault(t *testing.T) {
	task := &{{ .Name }}Task{}
	task.SetScheduler(schedule.New())
	task.TriggerFault(types.NewErrUnhandled(nil))
	err := task.Await()
	assert.NotNil(t, err)
}
`))
