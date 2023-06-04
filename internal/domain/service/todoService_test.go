package service_test

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/domain"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/entity"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/service"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/pb"
	"github.com/MicBun/protobuf-golang-todo/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
	"testing"
)

type todoInstance struct {
	service         *service.Todo
	mockTodoRepo    *mocks.TodoRepo
	mockTransaction *mocks.TransactionManager
}

func createInstance() *todoInstance {
	mockTodoRepo := &mocks.TodoRepo{}
	mockTransaction := &mocks.TransactionManager{}
	todoService := service.NewTodo(
		mockTodoRepo,
		mockTransaction,
	)

	return &todoInstance{
		service:         todoService,
		mockTodoRepo:    mockTodoRepo,
		mockTransaction: mockTransaction,
	}
}

func TestTodo_CreateOneGRPC(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := createInstance()
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(nil)
		instance.mockTodoRepo.EXPECT().CreateOne(mock.Anything, mock.Anything).Return(entity.Todo{}, nil)

		_, err := instance.service.CreateOne(context.Background(), &pb.CreateOneRequest{})

		assert.NoError(t, err)
	})

	t.Run("error - it should return error if todoRepo.CreateOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(assert.AnError)
		instance.mockTodoRepo.EXPECT().CreateOne(mock.Anything, mock.Anything).Return(entity.Todo{}, assert.AnError)

		_, err := instance.service.CreateOne(context.Background(), &pb.CreateOneRequest{})

		assert.EqualError(t, err, assert.AnError.Error())
	})
}

func TestTodo_GetManyGRPC(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := createInstance()
		var mockLimit uint32 = 10
		var mockOffset uint32 = 0
		instance.mockTodoRepo.EXPECT().FindAll(mock.Anything, mock.Anything).Return([]entity.Todo{
			{
				ID: 1,
			},
		}, nil)

		_, err := instance.service.GetMany(context.Background(), &pb.GetManyRequest{
			Limit:  &mockLimit,
			Offset: &mockOffset,
		})

		assert.NoError(t, err)
	})

	t.Run("error - it should return error if todoRepo.FindAll return error", func(t *testing.T) {
		instance := createInstance()
		var mockLimit uint32 = 10
		var mockOffset uint32 = 0
		instance.mockTodoRepo.EXPECT().FindAll(mock.Anything, mock.Anything).Return([]entity.Todo{}, assert.AnError)

		_, err := instance.service.GetMany(context.Background(), &pb.GetManyRequest{
			Limit:  &mockLimit,
			Offset: &mockOffset,
		})

		assert.EqualError(t, err, assert.AnError.Error())
	})
}

func TestTodo_GetOneGRPC(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{
			ID: 1,
		}, nil)

		_, err := instance.service.GetOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.NoError(t, err)
	})

	t.Run("validation - it should return ErrRecordNotFound if todoRepo.GetOne return ErrRecordNotFound", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, gorm.ErrRecordNotFound)

		_, err := instance.service.GetOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.EqualError(t, err, domain.ErrRecordNotFound.Error())
	})

	t.Run("error - it should return error if todoRepo.GetOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, assert.AnError)

		_, err := instance.service.GetOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.EqualError(t, err, assert.AnError.Error())
	})

}

func TestTodo_UpdateOneGRPC(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{
			Status: false,
		}, nil)
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(nil)
		instance.mockTodoRepo.EXPECT().UpdateOne(mock.Anything, mock.Anything).Return(nil)

		_, err := instance.service.UpdateOne(context.Background(), &pb.UpdateOneRequest{
			Id:     1,
			Status: true,
		})

		assert.NoError(t, err)
	})

	t.Run("validation - it should return ErrRecordNotFound if todoRepo.GetOne return gorm.ErrRecordNotFound", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, gorm.ErrRecordNotFound)

		_, err := instance.service.UpdateOne(context.Background(), &pb.UpdateOneRequest{
			Id:     1,
			Status: true,
		})

		assert.EqualError(t, err, domain.ErrRecordNotFound.Error())
	})

	t.Run("validation - it should return ErrNothingToUpdate if todoRepo.GetOne return same status", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{
			Status: true,
		}, nil)

		_, err := instance.service.UpdateOne(context.Background(), &pb.UpdateOneRequest{
			Id:     1,
			Status: true,
		})

		assert.EqualError(t, err, domain.ErrNothingToUpdate.Error())
	})

	t.Run("error - it should return error if todoRepo.GetOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, assert.AnError)

		_, err := instance.service.UpdateOne(context.Background(), &pb.UpdateOneRequest{
			Id:     1,
			Status: true,
		})

		assert.EqualError(t, err, assert.AnError.Error())
	})

	t.Run("error - it should return error if todoRepo.UpdateOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{
			Status: false,
		}, nil)
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(assert.AnError)
		instance.mockTodoRepo.EXPECT().UpdateOne(mock.Anything, mock.Anything).Return(assert.AnError)

		_, err := instance.service.UpdateOne(context.Background(), &pb.UpdateOneRequest{
			Id:     1,
			Status: true,
		})

		assert.EqualError(t, err, assert.AnError.Error())
	})
}

func TestTodo_DeleteOneGRPC(t *testing.T) {
	t.Run("success - it should return nil", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, nil)
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(nil)
		instance.mockTodoRepo.EXPECT().DeleteOne(mock.Anything, mock.Anything).Return(nil)

		_, err := instance.service.DeleteOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.NoError(t, err)
	})

	t.Run("validation - it should return ErrRecordNotFound if todoRepo.GetOne return gorm.ErrRecordNotFound", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, gorm.ErrRecordNotFound)

		_, err := instance.service.DeleteOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.EqualError(t, err, domain.ErrRecordNotFound.Error())
	})

	t.Run("error - it should return error if todoRepo.GetOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, assert.AnError)

		_, err := instance.service.DeleteOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.EqualError(t, err, assert.AnError.Error())
	})

	t.Run("error - it should return error if todoRepo.DeleteOne return error", func(t *testing.T) {
		instance := createInstance()
		instance.mockTodoRepo.EXPECT().GetOne(mock.Anything, mock.Anything).Return(entity.Todo{}, nil)
		instance.mockTransaction.EXPECT().Run(mock.Anything).Run(func(callback func(interface{}) error) {
			_ = callback("tx")
		}).Return(assert.AnError)
		instance.mockTodoRepo.EXPECT().DeleteOne(mock.Anything, mock.Anything).Return(assert.AnError)

		_, err := instance.service.DeleteOne(context.Background(), &wrapperspb.UInt32Value{Value: 1})

		assert.EqualError(t, err, assert.AnError.Error())
	})
}
