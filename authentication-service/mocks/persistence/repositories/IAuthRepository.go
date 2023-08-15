// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	models "github.com/shaikrasheed99/authentication-service/persistence/models"
	mock "github.com/stretchr/testify/mock"
)

// IAuthRepository is an autogenerated mock type for the IAuthRepository type
type IAuthRepository struct {
	mock.Mock
}

// FindUserByUsername provides a mock function with given fields: _a0
func (_m *IAuthRepository) FindUserByUsername(_a0 string) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: _a0
func (_m *IAuthRepository) SaveUser(_a0 *models.User) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.User) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*models.User) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(*models.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIAuthRepository creates a new instance of IAuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthRepository {
	mock := &IAuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}