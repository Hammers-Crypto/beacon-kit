// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"

	pkgtypes "github.com/berachain/beacon-kit/mod/da/pkg/types"

	servertypes "github.com/berachain/beacon-kit/mod/node-api/server/types"

	types "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
)

// BlockDB is an autogenerated mock type for the BlockDB type
type BlockDB struct {
	mock.Mock
}

type BlockDB_Expecter struct {
	mock *mock.Mock
}

func (_m *BlockDB) EXPECT() *BlockDB_Expecter {
	return &BlockDB_Expecter{mock: &_m.Mock}
}

// GetBlock provides a mock function with given fields:
func (_m *BlockDB) GetBlock() (*types.BeaconBlock, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlock")
	}

	var r0 *types.BeaconBlock
	var r1 error
	if rf, ok := ret.Get(0).(func() (*types.BeaconBlock, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *types.BeaconBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BeaconBlock)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlock'
type BlockDB_GetBlock_Call struct {
	*mock.Call
}

// GetBlock is a helper method to define mock.On call
func (_e *BlockDB_Expecter) GetBlock() *BlockDB_GetBlock_Call {
	return &BlockDB_GetBlock_Call{Call: _e.mock.On("GetBlock")}
}

func (_c *BlockDB_GetBlock_Call) Run(run func()) *BlockDB_GetBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockDB_GetBlock_Call) Return(_a0 *types.BeaconBlock, _a1 error) *BlockDB_GetBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlock_Call) RunAndReturn(run func() (*types.BeaconBlock, error)) *BlockDB_GetBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockBlobSidecars provides a mock function with given fields: indicies
func (_m *BlockDB) GetBlockBlobSidecars(indicies []string) ([]*pkgtypes.BlobSidecar, error) {
	ret := _m.Called(indicies)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockBlobSidecars")
	}

	var r0 []*pkgtypes.BlobSidecar
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*pkgtypes.BlobSidecar, error)); ok {
		return rf(indicies)
	}
	if rf, ok := ret.Get(0).(func([]string) []*pkgtypes.BlobSidecar); ok {
		r0 = rf(indicies)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkgtypes.BlobSidecar)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(indicies)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlockBlobSidecars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockBlobSidecars'
type BlockDB_GetBlockBlobSidecars_Call struct {
	*mock.Call
}

// GetBlockBlobSidecars is a helper method to define mock.On call
//   - indicies []string
func (_e *BlockDB_Expecter) GetBlockBlobSidecars(indicies interface{}) *BlockDB_GetBlockBlobSidecars_Call {
	return &BlockDB_GetBlockBlobSidecars_Call{Call: _e.mock.On("GetBlockBlobSidecars", indicies)}
}

func (_c *BlockDB_GetBlockBlobSidecars_Call) Run(run func(indicies []string)) *BlockDB_GetBlockBlobSidecars_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *BlockDB_GetBlockBlobSidecars_Call) Return(_a0 []*pkgtypes.BlobSidecar, _a1 error) *BlockDB_GetBlockBlobSidecars_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlockBlobSidecars_Call) RunAndReturn(run func([]string) ([]*pkgtypes.BlobSidecar, error)) *BlockDB_GetBlockBlobSidecars_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockHeader provides a mock function with given fields:
func (_m *BlockDB) GetBlockHeader() (*servertypes.BlockHeaderData, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeader")
	}

	var r0 *servertypes.BlockHeaderData
	var r1 error
	if rf, ok := ret.Get(0).(func() (*servertypes.BlockHeaderData, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *servertypes.BlockHeaderData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servertypes.BlockHeaderData)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlockHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockHeader'
type BlockDB_GetBlockHeader_Call struct {
	*mock.Call
}

// GetBlockHeader is a helper method to define mock.On call
func (_e *BlockDB_Expecter) GetBlockHeader() *BlockDB_GetBlockHeader_Call {
	return &BlockDB_GetBlockHeader_Call{Call: _e.mock.On("GetBlockHeader")}
}

func (_c *BlockDB_GetBlockHeader_Call) Run(run func()) *BlockDB_GetBlockHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockDB_GetBlockHeader_Call) Return(_a0 *servertypes.BlockHeaderData, _a1 error) *BlockDB_GetBlockHeader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlockHeader_Call) RunAndReturn(run func() (*servertypes.BlockHeaderData, error)) *BlockDB_GetBlockHeader_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockHeaders provides a mock function with given fields: slot, parent_root
func (_m *BlockDB) GetBlockHeaders(slot math.U64, parent_root bytes.B32) ([]*servertypes.BlockHeaderData, error) {
	ret := _m.Called(slot, parent_root)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeaders")
	}

	var r0 []*servertypes.BlockHeaderData
	var r1 error
	if rf, ok := ret.Get(0).(func(math.U64, bytes.B32) ([]*servertypes.BlockHeaderData, error)); ok {
		return rf(slot, parent_root)
	}
	if rf, ok := ret.Get(0).(func(math.U64, bytes.B32) []*servertypes.BlockHeaderData); ok {
		r0 = rf(slot, parent_root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*servertypes.BlockHeaderData)
		}
	}

	if rf, ok := ret.Get(1).(func(math.U64, bytes.B32) error); ok {
		r1 = rf(slot, parent_root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlockHeaders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockHeaders'
type BlockDB_GetBlockHeaders_Call struct {
	*mock.Call
}

// GetBlockHeaders is a helper method to define mock.On call
//   - slot math.U64
//   - parent_root bytes.B32
func (_e *BlockDB_Expecter) GetBlockHeaders(slot interface{}, parent_root interface{}) *BlockDB_GetBlockHeaders_Call {
	return &BlockDB_GetBlockHeaders_Call{Call: _e.mock.On("GetBlockHeaders", slot, parent_root)}
}

func (_c *BlockDB_GetBlockHeaders_Call) Run(run func(slot math.U64, parent_root bytes.B32)) *BlockDB_GetBlockHeaders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(bytes.B32))
	})
	return _c
}

func (_c *BlockDB_GetBlockHeaders_Call) Return(_a0 []*servertypes.BlockHeaderData, _a1 error) *BlockDB_GetBlockHeaders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlockHeaders_Call) RunAndReturn(run func(math.U64, bytes.B32) ([]*servertypes.BlockHeaderData, error)) *BlockDB_GetBlockHeaders_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockPropserDuties provides a mock function with given fields: epoch
func (_m *BlockDB) GetBlockPropserDuties(epoch math.U64) ([]*servertypes.ProposerDutiesData, error) {
	ret := _m.Called(epoch)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockPropserDuties")
	}

	var r0 []*servertypes.ProposerDutiesData
	var r1 error
	if rf, ok := ret.Get(0).(func(math.U64) ([]*servertypes.ProposerDutiesData, error)); ok {
		return rf(epoch)
	}
	if rf, ok := ret.Get(0).(func(math.U64) []*servertypes.ProposerDutiesData); ok {
		r0 = rf(epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*servertypes.ProposerDutiesData)
		}
	}

	if rf, ok := ret.Get(1).(func(math.U64) error); ok {
		r1 = rf(epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlockPropserDuties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockPropserDuties'
type BlockDB_GetBlockPropserDuties_Call struct {
	*mock.Call
}

// GetBlockPropserDuties is a helper method to define mock.On call
//   - epoch math.U64
func (_e *BlockDB_Expecter) GetBlockPropserDuties(epoch interface{}) *BlockDB_GetBlockPropserDuties_Call {
	return &BlockDB_GetBlockPropserDuties_Call{Call: _e.mock.On("GetBlockPropserDuties", epoch)}
}

func (_c *BlockDB_GetBlockPropserDuties_Call) Run(run func(epoch math.U64)) *BlockDB_GetBlockPropserDuties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64))
	})
	return _c
}

func (_c *BlockDB_GetBlockPropserDuties_Call) Return(_a0 []*servertypes.ProposerDutiesData, _a1 error) *BlockDB_GetBlockPropserDuties_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlockPropserDuties_Call) RunAndReturn(run func(math.U64) ([]*servertypes.ProposerDutiesData, error)) *BlockDB_GetBlockPropserDuties_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockRewards provides a mock function with given fields:
func (_m *BlockDB) GetBlockRewards() (*servertypes.BlockRewardsData, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlockRewards")
	}

	var r0 *servertypes.BlockRewardsData
	var r1 error
	if rf, ok := ret.Get(0).(func() (*servertypes.BlockRewardsData, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *servertypes.BlockRewardsData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servertypes.BlockRewardsData)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockDB_GetBlockRewards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockRewards'
type BlockDB_GetBlockRewards_Call struct {
	*mock.Call
}

// GetBlockRewards is a helper method to define mock.On call
func (_e *BlockDB_Expecter) GetBlockRewards() *BlockDB_GetBlockRewards_Call {
	return &BlockDB_GetBlockRewards_Call{Call: _e.mock.On("GetBlockRewards")}
}

func (_c *BlockDB_GetBlockRewards_Call) Run(run func()) *BlockDB_GetBlockRewards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlockDB_GetBlockRewards_Call) Return(_a0 *servertypes.BlockRewardsData, _a1 error) *BlockDB_GetBlockRewards_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockDB_GetBlockRewards_Call) RunAndReturn(run func() (*servertypes.BlockRewardsData, error)) *BlockDB_GetBlockRewards_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlockDB creates a new instance of BlockDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockDB {
	mock := &BlockDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
