// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeTodoServiceServer is an autogenerated mock type for the UnsafeTodoServiceServer type
type UnsafeTodoServiceServer struct {
	mock.Mock
}

type UnsafeTodoServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeTodoServiceServer) EXPECT() *UnsafeTodoServiceServer_Expecter {
	return &UnsafeTodoServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedTodoServiceServer provides a mock function with given fields:
func (_m *UnsafeTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {
	_m.Called()
}

// UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedTodoServiceServer'
type UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedTodoServiceServer is a helper method to define mock.On call
func (_e *UnsafeTodoServiceServer_Expecter) mustEmbedUnimplementedTodoServiceServer() *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call {
	return &UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedTodoServiceServer")}
}

func (_c *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call) Run(run func()) *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call) Return() *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call) RunAndReturn(run func()) *UnsafeTodoServiceServer_mustEmbedUnimplementedTodoServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUnsafeTodoServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeTodoServiceServer creates a new instance of UnsafeTodoServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeTodoServiceServer(t mockConstructorTestingTNewUnsafeTodoServiceServer) *UnsafeTodoServiceServer {
	mock := &UnsafeTodoServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
