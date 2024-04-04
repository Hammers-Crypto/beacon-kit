// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

type DB_Expecter struct {
	mock *mock.Mock
}

func (_m *DB) EXPECT() *DB_Expecter {
	return &DB_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: key
func (_m *DB) Delete(key []byte) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DB_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - key []byte
func (_e *DB_Expecter) Delete(key interface{}) *DB_Delete_Call {
	return &DB_Delete_Call{Call: _e.mock.On("Delete", key)}
}

func (_c *DB_Delete_Call) Run(run func(key []byte)) *DB_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DB_Delete_Call) Return(_a0 error) *DB_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Delete_Call) RunAndReturn(run func([]byte) error) *DB_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *DB) Get(key []byte) ([]byte, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type DB_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key []byte
func (_e *DB_Expecter) Get(key interface{}) *DB_Get_Call {
	return &DB_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *DB_Get_Call) Run(run func(key []byte)) *DB_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DB_Get_Call) Return(_a0 []byte, _a1 error) *DB_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_Get_Call) RunAndReturn(run func([]byte) ([]byte, error)) *DB_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: key
func (_m *DB) Has(key []byte) (bool, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (bool, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func([]byte) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type DB_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - key []byte
func (_e *DB_Expecter) Has(key interface{}) *DB_Has_Call {
	return &DB_Has_Call{Call: _e.mock.On("Has", key)}
}

func (_c *DB_Has_Call) Run(run func(key []byte)) *DB_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DB_Has_Call) Return(_a0 bool, _a1 error) *DB_Has_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DB_Has_Call) RunAndReturn(run func([]byte) (bool, error)) *DB_Has_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *DB) Set(key []byte, value []byte) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type DB_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key []byte
//   - value []byte
func (_e *DB_Expecter) Set(key interface{}, value interface{}) *DB_Set_Call {
	return &DB_Set_Call{Call: _e.mock.On("Set", key, value)}
}

func (_c *DB_Set_Call) Run(run func(key []byte, value []byte)) *DB_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].([]byte))
	})
	return _c
}

func (_c *DB_Set_Call) Return(_a0 error) *DB_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DB_Set_Call) RunAndReturn(run func([]byte, []byte) error) *DB_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
