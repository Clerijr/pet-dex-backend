// Code generated by mockery v2.43.0. DO NOT EDIT.

package interfaces

import (
	entity "pet-dex-backend/v2/entity"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockAdressRepo is an autogenerated mock type for the AdressRepo type
type MockAdressRepo struct {
	mock.Mock
}

type MockAdressRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAdressRepo) EXPECT() *MockAdressRepo_Expecter {
	return &MockAdressRepo_Expecter{mock: &_m.Mock}
}

// FindAddressByUserID provides a mock function with given fields: ID
func (_m *MockAdressRepo) FindAddressByUserID(ID uuid.UUID) (*entity.Address, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for FindAddressByUserID")
	}

	var r0 *entity.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*entity.Address, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Address); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdressRepo_FindAddressByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAddressByUserID'
type MockAdressRepo_FindAddressByUserID_Call struct {
	*mock.Call
}

// FindAddressByUserID is a helper method to define mock.On call
//   - ID uuid.UUID
func (_e *MockAdressRepo_Expecter) FindAddressByUserID(ID interface{}) *MockAdressRepo_FindAddressByUserID_Call {
	return &MockAdressRepo_FindAddressByUserID_Call{Call: _e.mock.On("FindAddressByUserID", ID)}
}

func (_c *MockAdressRepo_FindAddressByUserID_Call) Run(run func(ID uuid.UUID)) *MockAdressRepo_FindAddressByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockAdressRepo_FindAddressByUserID_Call) Return(_a0 *entity.Address, _a1 error) *MockAdressRepo_FindAddressByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAdressRepo_FindAddressByUserID_Call) RunAndReturn(run func(uuid.UUID) (*entity.Address, error)) *MockAdressRepo_FindAddressByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAddress provides a mock function with given fields: addr
func (_m *MockAdressRepo) SaveAddress(addr *entity.Address) error {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for SaveAddress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Address) error); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdressRepo_SaveAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAddress'
type MockAdressRepo_SaveAddress_Call struct {
	*mock.Call
}

// SaveAddress is a helper method to define mock.On call
//   - addr *entity.Address
func (_e *MockAdressRepo_Expecter) SaveAddress(addr interface{}) *MockAdressRepo_SaveAddress_Call {
	return &MockAdressRepo_SaveAddress_Call{Call: _e.mock.On("SaveAddress", addr)}
}

func (_c *MockAdressRepo_SaveAddress_Call) Run(run func(addr *entity.Address)) *MockAdressRepo_SaveAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Address))
	})
	return _c
}

func (_c *MockAdressRepo_SaveAddress_Call) Return(_a0 error) *MockAdressRepo_SaveAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAdressRepo_SaveAddress_Call) RunAndReturn(run func(*entity.Address) error) *MockAdressRepo_SaveAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAdressRepo creates a new instance of MockAdressRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAdressRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAdressRepo {
	mock := &MockAdressRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
