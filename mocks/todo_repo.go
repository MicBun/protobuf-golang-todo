// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	contract "github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	entity "github.com/MicBun/protobuf-golang-todo/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// TodoRepo is an autogenerated mock type for the TodoRepo type
type TodoRepo struct {
	mock.Mock
}

type TodoRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *TodoRepo) EXPECT() *TodoRepo_Expecter {
	return &TodoRepo_Expecter{mock: &_m.Mock}
}

// CreateOne provides a mock function with given fields: ctx, props
func (_m *TodoRepo) CreateOne(ctx context.Context, props *contract.TodoRepoCreateOneProps) error {
	ret := _m.Called(ctx, props)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoCreateOneProps) error); ok {
		r0 = rf(ctx, props)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepo_CreateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOne'
type TodoRepo_CreateOne_Call struct {
	*mock.Call
}

// CreateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - props *contract.TodoRepoCreateOneProps
func (_e *TodoRepo_Expecter) CreateOne(ctx interface{}, props interface{}) *TodoRepo_CreateOne_Call {
	return &TodoRepo_CreateOne_Call{Call: _e.mock.On("CreateOne", ctx, props)}
}

func (_c *TodoRepo_CreateOne_Call) Run(run func(ctx context.Context, props *contract.TodoRepoCreateOneProps)) *TodoRepo_CreateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*contract.TodoRepoCreateOneProps))
	})
	return _c
}

func (_c *TodoRepo_CreateOne_Call) Return(_a0 error) *TodoRepo_CreateOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepo_CreateOne_Call) RunAndReturn(run func(context.Context, *contract.TodoRepoCreateOneProps) error) *TodoRepo_CreateOne_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOne provides a mock function with given fields: ctx, props
func (_m *TodoRepo) DeleteOne(ctx context.Context, props *contract.TodoRepoDeleteOneProps) error {
	ret := _m.Called(ctx, props)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoDeleteOneProps) error); ok {
		r0 = rf(ctx, props)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepo_DeleteOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOne'
type TodoRepo_DeleteOne_Call struct {
	*mock.Call
}

// DeleteOne is a helper method to define mock.On call
//   - ctx context.Context
//   - props *contract.TodoRepoDeleteOneProps
func (_e *TodoRepo_Expecter) DeleteOne(ctx interface{}, props interface{}) *TodoRepo_DeleteOne_Call {
	return &TodoRepo_DeleteOne_Call{Call: _e.mock.On("DeleteOne", ctx, props)}
}

func (_c *TodoRepo_DeleteOne_Call) Run(run func(ctx context.Context, props *contract.TodoRepoDeleteOneProps)) *TodoRepo_DeleteOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*contract.TodoRepoDeleteOneProps))
	})
	return _c
}

func (_c *TodoRepo_DeleteOne_Call) Return(_a0 error) *TodoRepo_DeleteOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepo_DeleteOne_Call) RunAndReturn(run func(context.Context, *contract.TodoRepoDeleteOneProps) error) *TodoRepo_DeleteOne_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: ctx, props
func (_m *TodoRepo) FindAll(ctx context.Context, props *contract.TodoRepoFindAllProps) ([]entity.Todo, error) {
	ret := _m.Called(ctx, props)

	var r0 []entity.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoFindAllProps) ([]entity.Todo, error)); ok {
		return rf(ctx, props)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoFindAllProps) []entity.Todo); ok {
		r0 = rf(ctx, props)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.TodoRepoFindAllProps) error); ok {
		r1 = rf(ctx, props)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TodoRepo_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type TodoRepo_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - ctx context.Context
//   - props *contract.TodoRepoFindAllProps
func (_e *TodoRepo_Expecter) FindAll(ctx interface{}, props interface{}) *TodoRepo_FindAll_Call {
	return &TodoRepo_FindAll_Call{Call: _e.mock.On("FindAll", ctx, props)}
}

func (_c *TodoRepo_FindAll_Call) Run(run func(ctx context.Context, props *contract.TodoRepoFindAllProps)) *TodoRepo_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*contract.TodoRepoFindAllProps))
	})
	return _c
}

func (_c *TodoRepo_FindAll_Call) Return(_a0 []entity.Todo, _a1 error) *TodoRepo_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TodoRepo_FindAll_Call) RunAndReturn(run func(context.Context, *contract.TodoRepoFindAllProps) ([]entity.Todo, error)) *TodoRepo_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetOne provides a mock function with given fields: ctx, props
func (_m *TodoRepo) GetOne(ctx context.Context, props *contract.TodoRepoGetOneProps) (entity.Todo, error) {
	ret := _m.Called(ctx, props)

	var r0 entity.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoGetOneProps) (entity.Todo, error)); ok {
		return rf(ctx, props)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoGetOneProps) entity.Todo); ok {
		r0 = rf(ctx, props)
	} else {
		r0 = ret.Get(0).(entity.Todo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *contract.TodoRepoGetOneProps) error); ok {
		r1 = rf(ctx, props)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TodoRepo_GetOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOne'
type TodoRepo_GetOne_Call struct {
	*mock.Call
}

// GetOne is a helper method to define mock.On call
//   - ctx context.Context
//   - props *contract.TodoRepoGetOneProps
func (_e *TodoRepo_Expecter) GetOne(ctx interface{}, props interface{}) *TodoRepo_GetOne_Call {
	return &TodoRepo_GetOne_Call{Call: _e.mock.On("GetOne", ctx, props)}
}

func (_c *TodoRepo_GetOne_Call) Run(run func(ctx context.Context, props *contract.TodoRepoGetOneProps)) *TodoRepo_GetOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*contract.TodoRepoGetOneProps))
	})
	return _c
}

func (_c *TodoRepo_GetOne_Call) Return(_a0 entity.Todo, _a1 error) *TodoRepo_GetOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TodoRepo_GetOne_Call) RunAndReturn(run func(context.Context, *contract.TodoRepoGetOneProps) (entity.Todo, error)) *TodoRepo_GetOne_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOne provides a mock function with given fields: ctx, props
func (_m *TodoRepo) UpdateOne(ctx context.Context, props *contract.TodoRepoUpdateOneProps) error {
	ret := _m.Called(ctx, props)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *contract.TodoRepoUpdateOneProps) error); ok {
		r0 = rf(ctx, props)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepo_UpdateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOne'
type TodoRepo_UpdateOne_Call struct {
	*mock.Call
}

// UpdateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - props *contract.TodoRepoUpdateOneProps
func (_e *TodoRepo_Expecter) UpdateOne(ctx interface{}, props interface{}) *TodoRepo_UpdateOne_Call {
	return &TodoRepo_UpdateOne_Call{Call: _e.mock.On("UpdateOne", ctx, props)}
}

func (_c *TodoRepo_UpdateOne_Call) Run(run func(ctx context.Context, props *contract.TodoRepoUpdateOneProps)) *TodoRepo_UpdateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*contract.TodoRepoUpdateOneProps))
	})
	return _c
}

func (_c *TodoRepo_UpdateOne_Call) Return(_a0 error) *TodoRepo_UpdateOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepo_UpdateOne_Call) RunAndReturn(run func(context.Context, *contract.TodoRepoUpdateOneProps) error) *TodoRepo_UpdateOne_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewTodoRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepo creates a new instance of TodoRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepo(t mockConstructorTestingTNewTodoRepo) *TodoRepo {
	mock := &TodoRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}