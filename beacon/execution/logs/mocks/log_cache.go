// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	logs "github.com/itsdevbear/bolaris/beacon/execution/logs"
	mock "github.com/stretchr/testify/mock"
)

// LogCache is an autogenerated mock type for the LogCache type
type LogCache struct {
	mock.Mock
}

type LogCache_Expecter struct {
	mock *mock.Mock
}

func (_m *LogCache) EXPECT() *LogCache_Expecter {
	return &LogCache_Expecter{mock: &_m.Mock}
}

// Insert provides a mock function with given fields: log
func (_m *LogCache) Insert(log logs.LogContainer) error {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(logs.LogContainer) error); ok {
		r0 = rf(log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogCache_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type LogCache_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - log logs.LogContainer
func (_e *LogCache_Expecter) Insert(log interface{}) *LogCache_Insert_Call {
	return &LogCache_Insert_Call{Call: _e.mock.On("Insert", log)}
}

func (_c *LogCache_Insert_Call) Run(run func(log logs.LogContainer)) *LogCache_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(logs.LogContainer))
	})
	return _c
}

func (_c *LogCache_Insert_Call) Return(_a0 error) *LogCache_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LogCache_Insert_Call) RunAndReturn(run func(logs.LogContainer) error) *LogCache_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveMulti provides a mock function with given fields: index, n
func (_m *LogCache) RemoveMulti(index uint64, n uint64) ([]logs.LogContainer, error) {
	ret := _m.Called(index, n)

	if len(ret) == 0 {
		panic("no return value specified for RemoveMulti")
	}

	var r0 []logs.LogContainer
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) ([]logs.LogContainer, error)); ok {
		return rf(index, n)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) []logs.LogContainer); ok {
		r0 = rf(index, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logs.LogContainer)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(index, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogCache_RemoveMulti_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveMulti'
type LogCache_RemoveMulti_Call struct {
	*mock.Call
}

// RemoveMulti is a helper method to define mock.On call
//   - index uint64
//   - n uint64
func (_e *LogCache_Expecter) RemoveMulti(index interface{}, n interface{}) *LogCache_RemoveMulti_Call {
	return &LogCache_RemoveMulti_Call{Call: _e.mock.On("RemoveMulti", index, n)}
}

func (_c *LogCache_RemoveMulti_Call) Run(run func(index uint64, n uint64)) *LogCache_RemoveMulti_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].(uint64))
	})
	return _c
}

func (_c *LogCache_RemoveMulti_Call) Return(_a0 []logs.LogContainer, _a1 error) *LogCache_RemoveMulti_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LogCache_RemoveMulti_Call) RunAndReturn(run func(uint64, uint64) ([]logs.LogContainer, error)) *LogCache_RemoveMulti_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogCache creates a new instance of LogCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogCache {
	mock := &LogCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
