// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IAuthHandler is an autogenerated mock type for the IAuthHandler type
type IAuthHandler struct {
	mock.Mock
}

// Health provides a mock function with given fields: _a0
func (_m *IAuthHandler) Health(_a0 *gin.Context) {
	_m.Called(_a0)
}

// LoginHandler provides a mock function with given fields: _a0
func (_m *IAuthHandler) LoginHandler(_a0 *gin.Context) {
	_m.Called(_a0)
}

// SignupHandler provides a mock function with given fields: _a0
func (_m *IAuthHandler) SignupHandler(_a0 *gin.Context) {
	_m.Called(_a0)
}

// NewIAuthHandler creates a new instance of IAuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthHandler {
	mock := &IAuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}