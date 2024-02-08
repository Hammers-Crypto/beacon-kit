// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	parser "github.com/itsdevbear/bolaris/config/parser"
	mock "github.com/stretchr/testify/mock"
)

// BeaconKitConfig is an autogenerated mock type for the BeaconKitConfig type
type BeaconKitConfig[T interface{}] struct {
	mock.Mock
}

type BeaconKitConfig_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *BeaconKitConfig[T]) EXPECT() *BeaconKitConfig_Expecter[T] {
	return &BeaconKitConfig_Expecter[T]{mock: &_m.Mock}
}

// Parse provides a mock function with given fields: _a0
func (_m *BeaconKitConfig[T]) Parse(_a0 parser.AppOptionsParser) (*T, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Parse")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(parser.AppOptionsParser) (*T, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(parser.AppOptionsParser) *T); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(parser.AppOptionsParser) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconKitConfig_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type BeaconKitConfig_Parse_Call[T interface{}] struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - _a0 parser.AppOptionsParser
func (_e *BeaconKitConfig_Expecter[T]) Parse(_a0 interface{}) *BeaconKitConfig_Parse_Call[T] {
	return &BeaconKitConfig_Parse_Call[T]{Call: _e.mock.On("Parse", _a0)}
}

func (_c *BeaconKitConfig_Parse_Call[T]) Run(run func(_a0 parser.AppOptionsParser)) *BeaconKitConfig_Parse_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(parser.AppOptionsParser))
	})
	return _c
}

func (_c *BeaconKitConfig_Parse_Call[T]) Return(_a0 *T, _a1 error) *BeaconKitConfig_Parse_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconKitConfig_Parse_Call[T]) RunAndReturn(run func(parser.AppOptionsParser) (*T, error)) *BeaconKitConfig_Parse_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Template provides a mock function with given fields:
func (_m *BeaconKitConfig[T]) Template() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Template")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// BeaconKitConfig_Template_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Template'
type BeaconKitConfig_Template_Call[T interface{}] struct {
	*mock.Call
}

// Template is a helper method to define mock.On call
func (_e *BeaconKitConfig_Expecter[T]) Template() *BeaconKitConfig_Template_Call[T] {
	return &BeaconKitConfig_Template_Call[T]{Call: _e.mock.On("Template")}
}

func (_c *BeaconKitConfig_Template_Call[T]) Run(run func()) *BeaconKitConfig_Template_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconKitConfig_Template_Call[T]) Return(_a0 string) *BeaconKitConfig_Template_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconKitConfig_Template_Call[T]) RunAndReturn(run func() string) *BeaconKitConfig_Template_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewBeaconKitConfig creates a new instance of BeaconKitConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconKitConfig[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconKitConfig[T] {
	mock := &BeaconKitConfig[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
