package contract

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/entity"
)

type TodoRepo interface {
	CreateOne(ctx context.Context, props *TodoRepoCreateOneProps) error
	FindAll(ctx context.Context, props *TodoRepoFindAllProps) ([]entity.Todo, error)
	GetOne(ctx context.Context, props *TodoRepoGetOneProps) (entity.Todo, error)
	UpdateOne(ctx context.Context, props *TodoRepoUpdateOneProps) error
	DeleteOne(ctx context.Context, props *TodoRepoDeleteOneProps) error
}

type TodoRepoCreateOneProps struct {
	Task string
	Tx   any
}

type TodoRepoFindAllProps struct {
	Limit  int
	Offset int
	Status *bool
}

type TodoRepoGetOneProps struct {
	ID uint
}

type TodoRepoUpdateOneProps struct {
	ID     uint
	Status bool
	Tx     any
}

type TodoRepoDeleteOneProps struct {
	ID uint
	Tx any
}
