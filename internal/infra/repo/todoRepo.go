package repo

import (
	"context"
	"fmt"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/entity"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/factory"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Todo struct {
	db      *gorm.DB
	factory *factory.Todo
}

func NewTodo(
	db *gorm.DB,
	factory *factory.Todo,
) *Todo {
	return &Todo{
		db,
		factory,
	}
}

func (r *Todo) CreateOne(ctx context.Context, props *contract.TodoRepoCreateOneProps) (entity.Todo, error) {
	query := r.db.WithContext(ctx)
	if tx, ok := props.Tx.(*gorm.DB); ok {
		query = tx.WithContext(ctx)
	}

	todo := model.Todo{
		Task: props.Task,
	}
	if err := query.
		Model(&model.Todo{}).
		Create(&todo).Error; err != nil {
		return entity.Todo{}, errors.WithStack(err)
	}

	fmt.Println("todo: ", todo)

	return r.factory.CreateFromPgModel(&todo), nil

	////
	////return query.
	////	Model(&model.Todo{}).
	////	Create(&model.Todo{
	////		Task: props.Task,
	////	}).Error
	//
	//return query.
	//	Model(&model.Todo{}).
	//	Create(&model.Todo{
	//		Task: props.Task,
	//	}).Error
}

func (r *Todo) FindAll(ctx context.Context, props *contract.TodoRepoFindAllProps) ([]entity.Todo, error) {
	var todos []model.Todo
	query := r.db.WithContext(ctx).
		Limit(props.Limit).
		Offset(props.Offset)

	if props.Status != nil {
		query = query.Where("status = ?", props.Status)
	}

	if err := query.Find(&todos).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var result []entity.Todo
	for _, todo := range todos {
		result = append(result, r.factory.CreateFromPgModel(&todo))
	}

	return result, nil
}

func (r *Todo) GetOne(ctx context.Context, props *contract.TodoRepoGetOneProps) (entity.Todo, error) {
	var todo model.Todo
	if err := r.db.WithContext(ctx).
		Where("id = ?", props.ID).
		First(&todo).Error; err != nil {
		return entity.Todo{}, errors.WithStack(err)
	}

	return r.factory.CreateFromPgModel(&todo), nil
}

func (r *Todo) UpdateOne(ctx context.Context, props *contract.TodoRepoUpdateOneProps) error {
	query := r.db.WithContext(ctx)
	if tx, ok := props.Tx.(*gorm.DB); ok {
		query = tx.WithContext(ctx)
	}

	return query.
		Model(&model.Todo{
			Model: gorm.Model{
				ID: props.ID,
			},
		}).Updates(&model.Todo{
		Status: props.Status,
	}).Error
}

func (r *Todo) DeleteOne(ctx context.Context, props *contract.TodoRepoDeleteOneProps) error {
	query := r.db.WithContext(ctx)
	if tx, ok := props.Tx.(*gorm.DB); ok {
		query = tx.WithContext(ctx)
	}

	return query.
		Where("id = ?", props.ID).
		Delete(&model.Todo{}).Error
}
