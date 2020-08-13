// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aos-dev/go-storage/v2 (interfaces: Storager,Servicer,Reacher,PrefixSegmentsLister,IndexSegmenter,DirLister,PrefixLister,Statistician)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	v2 "github.com/aos-dev/go-storage/v2"
	segment "github.com/aos-dev/go-storage/v2/pkg/segment"
	types "github.com/aos-dev/go-storage/v2/types"
	info "github.com/aos-dev/go-storage/v2/types/info"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockStorager is a mock of Storager interface
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockStorager) Delete(arg0 string, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoragerMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorager)(nil).Delete), varargs...)
}

// DeleteWithContext mocks base method
func (m *MockStorager) DeleteWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithContext indicates an expected call of DeleteWithContext
func (mr *MockStoragerMockRecorder) DeleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithContext", reflect.TypeOf((*MockStorager)(nil).DeleteWithContext), varargs...)
}

// Metadata mocks base method
func (m *MockStorager) Metadata(arg0 ...*types.Pair) (info.StorageMeta, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Metadata", varargs...)
	ret0, _ := ret[0].(info.StorageMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata
func (mr *MockStoragerMockRecorder) Metadata(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockStorager)(nil).Metadata), arg0...)
}

// MetadataWithContext mocks base method
func (m *MockStorager) MetadataWithContext(arg0 context.Context, arg1 ...*types.Pair) (info.StorageMeta, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MetadataWithContext", varargs...)
	ret0, _ := ret[0].(info.StorageMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MetadataWithContext indicates an expected call of MetadataWithContext
func (mr *MockStoragerMockRecorder) MetadataWithContext(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetadataWithContext", reflect.TypeOf((*MockStorager)(nil).MetadataWithContext), varargs...)
}

// Read mocks base method
func (m *MockStorager) Read(arg0 string, arg1 ...*types.Pair) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockStoragerMockRecorder) Read(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStorager)(nil).Read), varargs...)
}

// ReadWithContext mocks base method
func (m *MockStorager) ReadWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadWithContext", varargs...)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithContext indicates an expected call of ReadWithContext
func (mr *MockStoragerMockRecorder) ReadWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithContext", reflect.TypeOf((*MockStorager)(nil).ReadWithContext), varargs...)
}

// Stat mocks base method
func (m *MockStorager) Stat(arg0 string, arg1 ...*types.Pair) (*types.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stat", varargs...)
	ret0, _ := ret[0].(*types.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockStoragerMockRecorder) Stat(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockStorager)(nil).Stat), varargs...)
}

// StatWithContext mocks base method
func (m *MockStorager) StatWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (*types.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatWithContext", varargs...)
	ret0, _ := ret[0].(*types.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatWithContext indicates an expected call of StatWithContext
func (mr *MockStoragerMockRecorder) StatWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatWithContext", reflect.TypeOf((*MockStorager)(nil).StatWithContext), varargs...)
}

// String mocks base method
func (m *MockStorager) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockStoragerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockStorager)(nil).String))
}

// Write mocks base method
func (m *MockStorager) Write(arg0 string, arg1 io.Reader, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockStoragerMockRecorder) Write(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorager)(nil).Write), varargs...)
}

// WriteWithContext mocks base method
func (m *MockStorager) WriteWithContext(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteWithContext indicates an expected call of WriteWithContext
func (mr *MockStoragerMockRecorder) WriteWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteWithContext", reflect.TypeOf((*MockStorager)(nil).WriteWithContext), varargs...)
}

// MockServicer is a mock of Servicer interface
type MockServicer struct {
	ctrl     *gomock.Controller
	recorder *MockServicerMockRecorder
}

// MockServicerMockRecorder is the mock recorder for MockServicer
type MockServicerMockRecorder struct {
	mock *MockServicer
}

// NewMockServicer creates a new mock instance
func NewMockServicer(ctrl *gomock.Controller) *MockServicer {
	mock := &MockServicer{ctrl: ctrl}
	mock.recorder = &MockServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServicer) EXPECT() *MockServicerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServicer) Create(arg0 string, arg1 ...*types.Pair) (v2.Storager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(v2.Storager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServicerMockRecorder) Create(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServicer)(nil).Create), varargs...)
}

// CreateWithContext mocks base method
func (m *MockServicer) CreateWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (v2.Storager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWithContext", varargs...)
	ret0, _ := ret[0].(v2.Storager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithContext indicates an expected call of CreateWithContext
func (mr *MockServicerMockRecorder) CreateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithContext", reflect.TypeOf((*MockServicer)(nil).CreateWithContext), varargs...)
}

// Delete mocks base method
func (m *MockServicer) Delete(arg0 string, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServicerMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServicer)(nil).Delete), varargs...)
}

// DeleteWithContext mocks base method
func (m *MockServicer) DeleteWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithContext indicates an expected call of DeleteWithContext
func (mr *MockServicerMockRecorder) DeleteWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithContext", reflect.TypeOf((*MockServicer)(nil).DeleteWithContext), varargs...)
}

// Get mocks base method
func (m *MockServicer) Get(arg0 string, arg1 ...*types.Pair) (v2.Storager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(v2.Storager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServicerMockRecorder) Get(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServicer)(nil).Get), varargs...)
}

// GetWithContext mocks base method
func (m *MockServicer) GetWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (v2.Storager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithContext", varargs...)
	ret0, _ := ret[0].(v2.Storager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithContext indicates an expected call of GetWithContext
func (mr *MockServicerMockRecorder) GetWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithContext", reflect.TypeOf((*MockServicer)(nil).GetWithContext), varargs...)
}

// List mocks base method
func (m *MockServicer) List(arg0 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List
func (mr *MockServicerMockRecorder) List(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicer)(nil).List), arg0...)
}

// ListWithContext mocks base method
func (m *MockServicer) ListWithContext(arg0 context.Context, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListWithContext indicates an expected call of ListWithContext
func (mr *MockServicerMockRecorder) ListWithContext(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWithContext", reflect.TypeOf((*MockServicer)(nil).ListWithContext), varargs...)
}

// String mocks base method
func (m *MockServicer) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockServicerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockServicer)(nil).String))
}

// MockReacher is a mock of Reacher interface
type MockReacher struct {
	ctrl     *gomock.Controller
	recorder *MockReacherMockRecorder
}

// MockReacherMockRecorder is the mock recorder for MockReacher
type MockReacherMockRecorder struct {
	mock *MockReacher
}

// NewMockReacher creates a new mock instance
func NewMockReacher(ctrl *gomock.Controller) *MockReacher {
	mock := &MockReacher{ctrl: ctrl}
	mock.recorder = &MockReacherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReacher) EXPECT() *MockReacherMockRecorder {
	return m.recorder
}

// Reach mocks base method
func (m *MockReacher) Reach(arg0 string, arg1 ...*types.Pair) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Reach", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reach indicates an expected call of Reach
func (mr *MockReacherMockRecorder) Reach(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reach", reflect.TypeOf((*MockReacher)(nil).Reach), varargs...)
}

// ReachWithContext mocks base method
func (m *MockReacher) ReachWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReachWithContext", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReachWithContext indicates an expected call of ReachWithContext
func (mr *MockReacherMockRecorder) ReachWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReachWithContext", reflect.TypeOf((*MockReacher)(nil).ReachWithContext), varargs...)
}

// MockPrefixSegmentsLister is a mock of PrefixSegmentsLister interface
type MockPrefixSegmentsLister struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixSegmentsListerMockRecorder
}

// MockPrefixSegmentsListerMockRecorder is the mock recorder for MockPrefixSegmentsLister
type MockPrefixSegmentsListerMockRecorder struct {
	mock *MockPrefixSegmentsLister
}

// NewMockPrefixSegmentsLister creates a new mock instance
func NewMockPrefixSegmentsLister(ctrl *gomock.Controller) *MockPrefixSegmentsLister {
	mock := &MockPrefixSegmentsLister{ctrl: ctrl}
	mock.recorder = &MockPrefixSegmentsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrefixSegmentsLister) EXPECT() *MockPrefixSegmentsListerMockRecorder {
	return m.recorder
}

// AbortSegment mocks base method
func (m *MockPrefixSegmentsLister) AbortSegment(arg0 segment.Segment, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortSegment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortSegment indicates an expected call of AbortSegment
func (mr *MockPrefixSegmentsListerMockRecorder) AbortSegment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortSegment", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).AbortSegment), varargs...)
}

// AbortSegmentWithContext mocks base method
func (m *MockPrefixSegmentsLister) AbortSegmentWithContext(arg0 context.Context, arg1 segment.Segment, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortSegmentWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortSegmentWithContext indicates an expected call of AbortSegmentWithContext
func (mr *MockPrefixSegmentsListerMockRecorder) AbortSegmentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortSegmentWithContext", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).AbortSegmentWithContext), varargs...)
}

// CompleteSegment mocks base method
func (m *MockPrefixSegmentsLister) CompleteSegment(arg0 segment.Segment, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteSegment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteSegment indicates an expected call of CompleteSegment
func (mr *MockPrefixSegmentsListerMockRecorder) CompleteSegment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteSegment", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).CompleteSegment), varargs...)
}

// CompleteSegmentWithContext mocks base method
func (m *MockPrefixSegmentsLister) CompleteSegmentWithContext(arg0 context.Context, arg1 segment.Segment, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteSegmentWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteSegmentWithContext indicates an expected call of CompleteSegmentWithContext
func (mr *MockPrefixSegmentsListerMockRecorder) CompleteSegmentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteSegmentWithContext", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).CompleteSegmentWithContext), varargs...)
}

// ListPrefixSegments mocks base method
func (m *MockPrefixSegmentsLister) ListPrefixSegments(arg0 string, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrefixSegments", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPrefixSegments indicates an expected call of ListPrefixSegments
func (mr *MockPrefixSegmentsListerMockRecorder) ListPrefixSegments(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrefixSegments", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).ListPrefixSegments), varargs...)
}

// ListPrefixSegmentsWithContext mocks base method
func (m *MockPrefixSegmentsLister) ListPrefixSegmentsWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrefixSegmentsWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPrefixSegmentsWithContext indicates an expected call of ListPrefixSegmentsWithContext
func (mr *MockPrefixSegmentsListerMockRecorder) ListPrefixSegmentsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrefixSegmentsWithContext", reflect.TypeOf((*MockPrefixSegmentsLister)(nil).ListPrefixSegmentsWithContext), varargs...)
}

// MockIndexSegmenter is a mock of IndexSegmenter interface
type MockIndexSegmenter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSegmenterMockRecorder
}

// MockIndexSegmenterMockRecorder is the mock recorder for MockIndexSegmenter
type MockIndexSegmenterMockRecorder struct {
	mock *MockIndexSegmenter
}

// NewMockIndexSegmenter creates a new mock instance
func NewMockIndexSegmenter(ctrl *gomock.Controller) *MockIndexSegmenter {
	mock := &MockIndexSegmenter{ctrl: ctrl}
	mock.recorder = &MockIndexSegmenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexSegmenter) EXPECT() *MockIndexSegmenterMockRecorder {
	return m.recorder
}

// AbortSegment mocks base method
func (m *MockIndexSegmenter) AbortSegment(arg0 segment.Segment, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortSegment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortSegment indicates an expected call of AbortSegment
func (mr *MockIndexSegmenterMockRecorder) AbortSegment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortSegment", reflect.TypeOf((*MockIndexSegmenter)(nil).AbortSegment), varargs...)
}

// AbortSegmentWithContext mocks base method
func (m *MockIndexSegmenter) AbortSegmentWithContext(arg0 context.Context, arg1 segment.Segment, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortSegmentWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortSegmentWithContext indicates an expected call of AbortSegmentWithContext
func (mr *MockIndexSegmenterMockRecorder) AbortSegmentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortSegmentWithContext", reflect.TypeOf((*MockIndexSegmenter)(nil).AbortSegmentWithContext), varargs...)
}

// CompleteSegment mocks base method
func (m *MockIndexSegmenter) CompleteSegment(arg0 segment.Segment, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteSegment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteSegment indicates an expected call of CompleteSegment
func (mr *MockIndexSegmenterMockRecorder) CompleteSegment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteSegment", reflect.TypeOf((*MockIndexSegmenter)(nil).CompleteSegment), varargs...)
}

// CompleteSegmentWithContext mocks base method
func (m *MockIndexSegmenter) CompleteSegmentWithContext(arg0 context.Context, arg1 segment.Segment, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteSegmentWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteSegmentWithContext indicates an expected call of CompleteSegmentWithContext
func (mr *MockIndexSegmenterMockRecorder) CompleteSegmentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteSegmentWithContext", reflect.TypeOf((*MockIndexSegmenter)(nil).CompleteSegmentWithContext), varargs...)
}

// InitIndexSegment mocks base method
func (m *MockIndexSegmenter) InitIndexSegment(arg0 string, arg1 ...*types.Pair) (segment.Segment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InitIndexSegment", varargs...)
	ret0, _ := ret[0].(segment.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitIndexSegment indicates an expected call of InitIndexSegment
func (mr *MockIndexSegmenterMockRecorder) InitIndexSegment(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitIndexSegment", reflect.TypeOf((*MockIndexSegmenter)(nil).InitIndexSegment), varargs...)
}

// InitIndexSegmentWithContext mocks base method
func (m *MockIndexSegmenter) InitIndexSegmentWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) (segment.Segment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InitIndexSegmentWithContext", varargs...)
	ret0, _ := ret[0].(segment.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitIndexSegmentWithContext indicates an expected call of InitIndexSegmentWithContext
func (mr *MockIndexSegmenterMockRecorder) InitIndexSegmentWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitIndexSegmentWithContext", reflect.TypeOf((*MockIndexSegmenter)(nil).InitIndexSegmentWithContext), varargs...)
}

// WriteIndexSegment mocks base method
func (m *MockIndexSegmenter) WriteIndexSegment(arg0 segment.Segment, arg1 io.Reader, arg2 int, arg3 int64, arg4 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteIndexSegment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteIndexSegment indicates an expected call of WriteIndexSegment
func (mr *MockIndexSegmenterMockRecorder) WriteIndexSegment(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIndexSegment", reflect.TypeOf((*MockIndexSegmenter)(nil).WriteIndexSegment), varargs...)
}

// WriteIndexSegmentWithContext mocks base method
func (m *MockIndexSegmenter) WriteIndexSegmentWithContext(arg0 context.Context, arg1 segment.Segment, arg2 io.Reader, arg3 int, arg4 int64, arg5 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteIndexSegmentWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteIndexSegmentWithContext indicates an expected call of WriteIndexSegmentWithContext
func (mr *MockIndexSegmenterMockRecorder) WriteIndexSegmentWithContext(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIndexSegmentWithContext", reflect.TypeOf((*MockIndexSegmenter)(nil).WriteIndexSegmentWithContext), varargs...)
}

// MockDirLister is a mock of DirLister interface
type MockDirLister struct {
	ctrl     *gomock.Controller
	recorder *MockDirListerMockRecorder
}

// MockDirListerMockRecorder is the mock recorder for MockDirLister
type MockDirListerMockRecorder struct {
	mock *MockDirLister
}

// NewMockDirLister creates a new mock instance
func NewMockDirLister(ctrl *gomock.Controller) *MockDirLister {
	mock := &MockDirLister{ctrl: ctrl}
	mock.recorder = &MockDirListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirLister) EXPECT() *MockDirListerMockRecorder {
	return m.recorder
}

// ListDir mocks base method
func (m *MockDirLister) ListDir(arg0 string, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDir", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListDir indicates an expected call of ListDir
func (mr *MockDirListerMockRecorder) ListDir(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDir", reflect.TypeOf((*MockDirLister)(nil).ListDir), varargs...)
}

// ListDirWithContext mocks base method
func (m *MockDirLister) ListDirWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDirWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListDirWithContext indicates an expected call of ListDirWithContext
func (mr *MockDirListerMockRecorder) ListDirWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDirWithContext", reflect.TypeOf((*MockDirLister)(nil).ListDirWithContext), varargs...)
}

// MockPrefixLister is a mock of PrefixLister interface
type MockPrefixLister struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixListerMockRecorder
}

// MockPrefixListerMockRecorder is the mock recorder for MockPrefixLister
type MockPrefixListerMockRecorder struct {
	mock *MockPrefixLister
}

// NewMockPrefixLister creates a new mock instance
func NewMockPrefixLister(ctrl *gomock.Controller) *MockPrefixLister {
	mock := &MockPrefixLister{ctrl: ctrl}
	mock.recorder = &MockPrefixListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrefixLister) EXPECT() *MockPrefixListerMockRecorder {
	return m.recorder
}

// ListPrefix mocks base method
func (m *MockPrefixLister) ListPrefix(arg0 string, arg1 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrefix", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPrefix indicates an expected call of ListPrefix
func (mr *MockPrefixListerMockRecorder) ListPrefix(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrefix", reflect.TypeOf((*MockPrefixLister)(nil).ListPrefix), varargs...)
}

// ListPrefixWithContext mocks base method
func (m *MockPrefixLister) ListPrefixWithContext(arg0 context.Context, arg1 string, arg2 ...*types.Pair) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrefixWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListPrefixWithContext indicates an expected call of ListPrefixWithContext
func (mr *MockPrefixListerMockRecorder) ListPrefixWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrefixWithContext", reflect.TypeOf((*MockPrefixLister)(nil).ListPrefixWithContext), varargs...)
}

// MockStatistician is a mock of Statistician interface
type MockStatistician struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticianMockRecorder
}

// MockStatisticianMockRecorder is the mock recorder for MockStatistician
type MockStatisticianMockRecorder struct {
	mock *MockStatistician
}

// NewMockStatistician creates a new mock instance
func NewMockStatistician(ctrl *gomock.Controller) *MockStatistician {
	mock := &MockStatistician{ctrl: ctrl}
	mock.recorder = &MockStatisticianMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatistician) EXPECT() *MockStatisticianMockRecorder {
	return m.recorder
}

// Statistical mocks base method
func (m *MockStatistician) Statistical(arg0 ...*types.Pair) (info.StorageStatistic, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Statistical", varargs...)
	ret0, _ := ret[0].(info.StorageStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Statistical indicates an expected call of Statistical
func (mr *MockStatisticianMockRecorder) Statistical(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistical", reflect.TypeOf((*MockStatistician)(nil).Statistical), arg0...)
}

// StatisticalWithContext mocks base method
func (m *MockStatistician) StatisticalWithContext(arg0 context.Context, arg1 ...*types.Pair) (info.StorageStatistic, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatisticalWithContext", varargs...)
	ret0, _ := ret[0].(info.StorageStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatisticalWithContext indicates an expected call of StatisticalWithContext
func (mr *MockStatisticianMockRecorder) StatisticalWithContext(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatisticalWithContext", reflect.TypeOf((*MockStatistician)(nil).StatisticalWithContext), varargs...)
}
