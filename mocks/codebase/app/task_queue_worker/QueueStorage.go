// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	taskqueueworker "github.com/golangid/candi/codebase/app/task_queue_worker"
	mock "github.com/stretchr/testify/mock"
)

// QueueStorage is an autogenerated mock type for the QueueStorage type
type QueueStorage struct {
	mock.Mock
}

// Clear provides a mock function with given fields: taskName
func (_m *QueueStorage) Clear(taskName string) {
	_m.Called(taskName)
}

// NextJob provides a mock function with given fields: taskName
func (_m *QueueStorage) NextJob(taskName string) string {
	ret := _m.Called(taskName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(taskName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PopJob provides a mock function with given fields: taskName
func (_m *QueueStorage) PopJob(taskName string) string {
	ret := _m.Called(taskName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(taskName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PushJob provides a mock function with given fields: job
func (_m *QueueStorage) PushJob(job *taskqueueworker.Job) {
	_m.Called(job)
}
