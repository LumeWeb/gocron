// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-co-op/gocron/v2 (interfaces: Scheduler)

// Package gocronmocks is a generated GoMock package.
package gocronmocks

import (
	reflect "reflect"

	gocron "github.com/go-co-op/gocron/v2"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockScheduler is a mock of Scheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// Jobs mocks base method.
func (m *MockScheduler) Jobs() []gocron.Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].([]gocron.Job)
	return ret0
}

// Jobs indicates an expected call of Jobs.
func (mr *MockSchedulerMockRecorder) Jobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockScheduler)(nil).Jobs))
}

// NewJob mocks base method.
func (m *MockScheduler) NewJob(arg0 gocron.JobDefinition, arg1 gocron.Task, arg2 ...gocron.JobOption) (gocron.Job, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewJob", varargs...)
	ret0, _ := ret[0].(gocron.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewJob indicates an expected call of NewJob.
func (mr *MockSchedulerMockRecorder) NewJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewJob", reflect.TypeOf((*MockScheduler)(nil).NewJob), varargs...)
}

// RemoveByTags mocks base method.
func (m *MockScheduler) RemoveByTags(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RemoveByTags", varargs...)
}

// RemoveByTags indicates an expected call of RemoveByTags.
func (mr *MockSchedulerMockRecorder) RemoveByTags(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByTags", reflect.TypeOf((*MockScheduler)(nil).RemoveByTags), arg0...)
}

// RemoveJob mocks base method.
func (m *MockScheduler) RemoveJob(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveJob", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveJob indicates an expected call of RemoveJob.
func (mr *MockSchedulerMockRecorder) RemoveJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveJob", reflect.TypeOf((*MockScheduler)(nil).RemoveJob), arg0)
}

// Shutdown mocks base method.
func (m *MockScheduler) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockSchedulerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockScheduler)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockScheduler) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockSchedulerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockScheduler)(nil).Start))
}

// StopJobs mocks base method.
func (m *MockScheduler) StopJobs() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopJobs")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopJobs indicates an expected call of StopJobs.
func (mr *MockSchedulerMockRecorder) StopJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopJobs", reflect.TypeOf((*MockScheduler)(nil).StopJobs))
}

// Update mocks base method.
func (m *MockScheduler) Update(arg0 uuid.UUID, arg1 gocron.JobDefinition, arg2 gocron.Task, arg3 ...gocron.JobOption) (gocron.Job, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(gocron.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSchedulerMockRecorder) Update(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockScheduler)(nil).Update), varargs...)
}
