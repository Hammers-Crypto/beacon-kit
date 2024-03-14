// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AppOptions is an autogenerated mock type for the AppOptions type
type AppOptions struct {
	mock.Mock
}

type AppOptions_Expecter struct {
	mock *mock.Mock
}

func (_m *AppOptions) EXPECT() *AppOptions_Expecter {
	return &AppOptions_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: _a0
func (_m *AppOptions) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// AppOptions_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type AppOptions_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 string
func (_e *AppOptions_Expecter) Get(_a0 interface{}) *AppOptions_Get_Call {
	return &AppOptions_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *AppOptions_Get_Call) Run(run func(_a0 string)) *AppOptions_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppOptions_Get_Call) Return(_a0 interface{}) *AppOptions_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppOptions_Get_Call) RunAndReturn(run func(string) interface{}) *AppOptions_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewAppOptions creates a new instance of AppOptions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppOptions(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppOptions {
	mock := &AppOptions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
