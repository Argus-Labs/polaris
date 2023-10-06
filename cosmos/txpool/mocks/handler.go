// Code generated by mockery v2.35.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type.
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields:.
func (_m *Handler) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'.
type Handler_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call.
func (_e *Handler_Expecter) Start() *Handler_Start_Call {
	return &Handler_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *Handler_Start_Call) Run(run func()) *Handler_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Start_Call) Return(_a0 error) *Handler_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Start_Call) RunAndReturn(run func() error) *Handler_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:.
func (_m *Handler) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'.
type Handler_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call.
func (_e *Handler_Expecter) Stop() *Handler_Stop_Call {
	return &Handler_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Handler_Stop_Call) Run(run func()) *Handler_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Stop_Call) Return(_a0 error) *Handler_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Stop_Call) RunAndReturn(run func() error) *Handler_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}