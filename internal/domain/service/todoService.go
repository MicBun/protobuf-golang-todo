package service

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/domain"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
)

type Todo struct {
	todoRepo    contract.TodoRepo
	transaction contract.TransactionManager
}

func NewTodo(
	todoRepo contract.TodoRepo,
	transaction contract.TransactionManager,
) *Todo {
	return &Todo{
		todoRepo,
		transaction,
	}
}

func (t *Todo) CreateOne(ctx context.Context, in *pb.CreateOneRequest) (*empty.Empty, error) {
	if err := t.transaction.Run(func(tx any) error {
		return t.todoRepo.CreateOne(ctx, &contract.TodoRepoCreateOneProps{
			Task: in.Task,
			Tx:   tx,
		})
	}); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (t *Todo) GetMany(ctx context.Context, in *pb.GetManyRequest) (*pb.TodoList, error) {
	limit := 10
	if in.Limit != nil {
		limit = int(*in.Limit)
	}
	offset := 0
	if in.Offset != nil {
		offset = int(*in.Offset)
	}
	todos, err := t.todoRepo.FindAll(ctx, &contract.TodoRepoFindAllProps{
		Limit:  limit,
		Offset: offset,
		Status: in.Status,
	})
	if err != nil {
		return nil, err
	}

	res := &pb.TodoList{}
	for _, todo := range todos {
		res.Todos = append(res.Todos, &pb.Todo{
			Id:     uint32(todo.ID),
			Task:   todo.Task,
			Status: todo.Status,
		})
	}

	return res, nil
}

func (t *Todo) GetOne(ctx context.Context, id *wrapperspb.UInt32Value) (*pb.Todo, error) {
	todo, err := t.todoRepo.GetOne(ctx, &contract.TodoRepoGetOneProps{
		ID: uint(id.Value),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	return &pb.Todo{
		Id:     uint32(todo.ID),
		Task:   todo.Task,
		Status: todo.Status,
	}, nil
}

func (t *Todo) UpdateOne(ctx context.Context, in *pb.UpdateOneRequest) (*empty.Empty, error) {
	todo, err := t.todoRepo.GetOne(ctx, &contract.TodoRepoGetOneProps{
		ID: uint(in.Id),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	if todo.Status == in.Status {
		return nil, domain.ErrNothingToUpdate
	}

	if errRun := t.transaction.Run(func(tx any) error {
		return t.todoRepo.UpdateOne(ctx, &contract.TodoRepoUpdateOneProps{
			ID:     uint(in.Id),
			Status: in.Status,
			Tx:     tx,
		})
	}); errRun != nil {
		return nil, errRun
	}

	return &empty.Empty{}, nil
}

func (t *Todo) DeleteOne(ctx context.Context, id *wrapperspb.UInt32Value) (*empty.Empty, error) {
	_, err := t.todoRepo.GetOne(ctx, &contract.TodoRepoGetOneProps{
		ID: uint(id.Value),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecordNotFound
		} else {
			return nil, err
		}
	}

	if errRun := t.transaction.Run(func(tx any) error {
		return t.todoRepo.DeleteOne(ctx, &contract.TodoRepoDeleteOneProps{
			ID: uint(id.Value),
			Tx: tx,
		})
	}); errRun != nil {
		return nil, errRun
	}

	return &empty.Empty{}, nil
}
