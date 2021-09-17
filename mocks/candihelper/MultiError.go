// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	candihelper "github.com/golangid/candi/candihelper"
	mock "github.com/stretchr/testify/mock"
)

// MultiError is an autogenerated mock type for the MultiError type
type MultiError struct {
	mock.Mock
}

// Append provides a mock function with given fields: key, err
func (_m *MultiError) Append(key string, err error) candihelper.MultiError {
	ret := _m.Called(key, err)

	var r0 candihelper.MultiError
	if rf, ok := ret.Get(0).(func(string, error) candihelper.MultiError); ok {
		r0 = rf(key, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(candihelper.MultiError)
		}
	}

	return r0
}

// Clear provides a mock function with given fields:
func (_m *MultiError) Clear() {
	_m.Called()
}

// Error provides a mock function with given fields:
func (_m *MultiError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// HasError provides a mock function with given fields:
func (_m *MultiError) HasError() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsNil provides a mock function with given fields:
func (_m *MultiError) IsNil() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Merge provides a mock function with given fields: _a0
func (_m *MultiError) Merge(_a0 candihelper.MultiError) candihelper.MultiError {
	ret := _m.Called(_a0)

	var r0 candihelper.MultiError
	if rf, ok := ret.Get(0).(func(candihelper.MultiError) candihelper.MultiError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(candihelper.MultiError)
		}
	}

	return r0
}

// ToMap provides a mock function with given fields:
func (_m *MultiError) ToMap() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}
