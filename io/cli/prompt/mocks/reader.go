// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

type Reader_Expecter struct {
	mock *mock.Mock
}

func (_m *Reader) EXPECT() *Reader_Expecter {
	return &Reader_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: _a0
func (_m *Reader) Read(_a0 []byte) (int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Reader_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - _a0 []byte
func (_e *Reader_Expecter) Read(_a0 interface{}) *Reader_Read_Call {
	return &Reader_Read_Call{Call: _e.mock.On("Read", _a0)}
}

func (_c *Reader_Read_Call) Run(run func(_a0 []byte)) *Reader_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *Reader_Read_Call) Return(_a0 int, _a1 error) *Reader_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_Read_Call) RunAndReturn(run func([]byte) (int, error)) *Reader_Read_Call {
	_c.Call.Return(run)
	return _c
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
