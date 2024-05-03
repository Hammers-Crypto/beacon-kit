// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	consensus "github.com/berachain/beacon-kit/mod/primitives/pkg/consensus"

	core "github.com/berachain/beacon-kit/mod/core"

	deposit "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"

	mock "github.com/stretchr/testify/mock"

	state "github.com/berachain/beacon-kit/mod/core/state"

	types "github.com/berachain/beacon-kit/mod/da/pkg/types"
)

// BeaconStorageBackend is an autogenerated mock type for the BeaconStorageBackend type
type BeaconStorageBackend struct {
	mock.Mock
}

type BeaconStorageBackend_Expecter struct {
	mock *mock.Mock
}

func (_m *BeaconStorageBackend) EXPECT() *BeaconStorageBackend_Expecter {
	return &BeaconStorageBackend_Expecter{mock: &_m.Mock}
}

// AvailabilityStore provides a mock function with given fields: ctx
func (_m *BeaconStorageBackend) AvailabilityStore(ctx context.Context) core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for AvailabilityStore")
	}

	var r0 core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars]
	if rf, ok := ret.Get(0).(func(context.Context) core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars])
		}
	}

	return r0
}

// BeaconStorageBackend_AvailabilityStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailabilityStore'
type BeaconStorageBackend_AvailabilityStore_Call struct {
	*mock.Call
}

// AvailabilityStore is a helper method to define mock.On call
//   - ctx context.Context
func (_e *BeaconStorageBackend_Expecter) AvailabilityStore(ctx interface{}) *BeaconStorageBackend_AvailabilityStore_Call {
	return &BeaconStorageBackend_AvailabilityStore_Call{Call: _e.mock.On("AvailabilityStore", ctx)}
}

func (_c *BeaconStorageBackend_AvailabilityStore_Call) Run(run func(ctx context.Context)) *BeaconStorageBackend_AvailabilityStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *BeaconStorageBackend_AvailabilityStore_Call) Return(_a0 core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars]) *BeaconStorageBackend_AvailabilityStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconStorageBackend_AvailabilityStore_Call) RunAndReturn(run func(context.Context) core.AvailabilityStore[consensus.ReadOnlyBeaconBlockBody, *types.BlobSidecars]) *BeaconStorageBackend_AvailabilityStore_Call {
	_c.Call.Return(run)
	return _c
}

// BeaconState provides a mock function with given fields: ctx
func (_m *BeaconStorageBackend) BeaconState(ctx context.Context) state.BeaconState {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeaconState")
	}

	var r0 state.BeaconState
	if rf, ok := ret.Get(0).(func(context.Context) state.BeaconState); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(state.BeaconState)
		}
	}

	return r0
}

// BeaconStorageBackend_BeaconState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeaconState'
type BeaconStorageBackend_BeaconState_Call struct {
	*mock.Call
}

// BeaconState is a helper method to define mock.On call
//   - ctx context.Context
func (_e *BeaconStorageBackend_Expecter) BeaconState(ctx interface{}) *BeaconStorageBackend_BeaconState_Call {
	return &BeaconStorageBackend_BeaconState_Call{Call: _e.mock.On("BeaconState", ctx)}
}

func (_c *BeaconStorageBackend_BeaconState_Call) Run(run func(ctx context.Context)) *BeaconStorageBackend_BeaconState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *BeaconStorageBackend_BeaconState_Call) Return(_a0 state.BeaconState) *BeaconStorageBackend_BeaconState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconStorageBackend_BeaconState_Call) RunAndReturn(run func(context.Context) state.BeaconState) *BeaconStorageBackend_BeaconState_Call {
	_c.Call.Return(run)
	return _c
}

// DepositStore provides a mock function with given fields: ctx
func (_m *BeaconStorageBackend) DepositStore(ctx context.Context) *deposit.KVStore {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DepositStore")
	}

	var r0 *deposit.KVStore
	if rf, ok := ret.Get(0).(func(context.Context) *deposit.KVStore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deposit.KVStore)
		}
	}

	return r0
}

// BeaconStorageBackend_DepositStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DepositStore'
type BeaconStorageBackend_DepositStore_Call struct {
	*mock.Call
}

// DepositStore is a helper method to define mock.On call
//   - ctx context.Context
func (_e *BeaconStorageBackend_Expecter) DepositStore(ctx interface{}) *BeaconStorageBackend_DepositStore_Call {
	return &BeaconStorageBackend_DepositStore_Call{Call: _e.mock.On("DepositStore", ctx)}
}

func (_c *BeaconStorageBackend_DepositStore_Call) Run(run func(ctx context.Context)) *BeaconStorageBackend_DepositStore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *BeaconStorageBackend_DepositStore_Call) Return(_a0 *deposit.KVStore) *BeaconStorageBackend_DepositStore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconStorageBackend_DepositStore_Call) RunAndReturn(run func(context.Context) *deposit.KVStore) *BeaconStorageBackend_DepositStore_Call {
	_c.Call.Return(run)
	return _c
}

// NewBeaconStorageBackend creates a new instance of BeaconStorageBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconStorageBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconStorageBackend {
	mock := &BeaconStorageBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
