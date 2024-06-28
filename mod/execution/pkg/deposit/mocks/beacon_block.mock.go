// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	deposit "github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// BeaconBlock is an autogenerated mock type for the BeaconBlock type
type BeaconBlock[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	mock.Mock
}

type BeaconBlock_Expecter[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	mock *mock.Mock
}

func (_m *BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) EXPECT() *BeaconBlock_Expecter[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	return &BeaconBlock_Expecter[DepositT, BeaconBlockBodyT, ExecutionPayloadT]{mock: &_m.Mock}
}

// GetBody provides a mock function with given fields:
func (_m *BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) GetBody() BeaconBlockBodyT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 BeaconBlockBodyT
	if rf, ok := ret.Get(0).(func() BeaconBlockBodyT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BeaconBlockBodyT)
	}

	return r0
}

// BeaconBlock_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type BeaconBlock_GetBody_Call[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) GetBody() *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	return &BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]{Call: _e.mock.On("GetBody")}
}

func (_c *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) Run(run func()) *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) Return(_a0 BeaconBlockBodyT) *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) RunAndReturn(run func() BeaconBlockBodyT) *BeaconBlock_GetBody_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlock_GetSlot_Call[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload] struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) GetSlot() *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	return &BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) Run(run func()) *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) Return(_a0 math.U64) *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT]) RunAndReturn(run func() math.U64) *BeaconBlock_GetSlot_Call[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlock creates a new instance of BeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlock[DepositT interface{}, BeaconBlockBodyT deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], ExecutionPayloadT deposit.ExecutionPayload](t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT] {
	mock := &BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}