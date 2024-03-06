// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	enginev1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
	mock "github.com/stretchr/testify/mock"

	typesv1 "github.com/itsdevbear/bolaris/beacon/core/types/v1"
)

// ValsetChangeProvider is an autogenerated mock type for the ValsetChangeProvider type
type ValsetChangeProvider struct {
	mock.Mock
}

type ValsetChangeProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *ValsetChangeProvider) EXPECT() *ValsetChangeProvider_Expecter {
	return &ValsetChangeProvider_Expecter{mock: &_m.Mock}
}

// ApplyChanges provides a mock function with given fields: _a0, _a1, _a2
func (_m *ValsetChangeProvider) ApplyChanges(_a0 context.Context, _a1 []*typesv1.Deposit, _a2 []*enginev1.Withdrawal) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ApplyChanges")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*typesv1.Deposit, []*enginev1.Withdrawal) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValsetChangeProvider_ApplyChanges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyChanges'
type ValsetChangeProvider_ApplyChanges_Call struct {
	*mock.Call
}

// ApplyChanges is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []*typesv1.Deposit
//   - _a2 []*enginev1.Withdrawal
func (_e *ValsetChangeProvider_Expecter) ApplyChanges(_a0 interface{}, _a1 interface{}, _a2 interface{}) *ValsetChangeProvider_ApplyChanges_Call {
	return &ValsetChangeProvider_ApplyChanges_Call{Call: _e.mock.On("ApplyChanges", _a0, _a1, _a2)}
}

func (_c *ValsetChangeProvider_ApplyChanges_Call) Run(run func(_a0 context.Context, _a1 []*typesv1.Deposit, _a2 []*enginev1.Withdrawal)) *ValsetChangeProvider_ApplyChanges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*typesv1.Deposit), args[2].([]*enginev1.Withdrawal))
	})
	return _c
}

func (_c *ValsetChangeProvider_ApplyChanges_Call) Return(_a0 error) *ValsetChangeProvider_ApplyChanges_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValsetChangeProvider_ApplyChanges_Call) RunAndReturn(run func(context.Context, []*typesv1.Deposit, []*enginev1.Withdrawal) error) *ValsetChangeProvider_ApplyChanges_Call {
	_c.Call.Return(run)
	return _c
}

// NewValsetChangeProvider creates a new instance of ValsetChangeProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValsetChangeProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *ValsetChangeProvider {
	mock := &ValsetChangeProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
