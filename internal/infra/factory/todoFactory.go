package factory

import (
	"github.com/MicBun/protobuf-golang-todo/internal/domain/entity"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/model"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (f *Todo) CreateFromPgModel(model *model.Todo) entity.Todo {
	return entity.Todo{
		ID:     model.ID,
		Task:   model.Task,
		Status: model.Status,
	}
}
