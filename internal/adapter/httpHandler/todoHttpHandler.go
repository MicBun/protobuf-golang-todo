package httpHandler

import (
	"github.com/MicBun/protobuf-golang-todo/internal/adapter/schema"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/service"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/pb"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
)

type Todo struct {
	service *service.Todo
}

func NewTodo(
	service *service.Todo,
) *Todo {
	return &Todo{
		service,
	}
}

func (t *Todo) CreateOne(ctx echo.Context) error {
	body := new(schema.TodoCreateOneBody)
	if err := ctx.Bind(body); err != nil {
		return err
	}
	if err := ctx.Validate(body); err != nil {
		return err
	}

	_, err := t.service.CreateOne(ctx.Request().Context(), &pb.CreateOneRequest{
		Task: body.Task,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"messageId": "SUCCESS",
		"data":      nil,
	})
}

func (t *Todo) GetMany(ctx echo.Context) error {
	query := new(schema.TodoFindAllQuery)
	if err := ctx.Bind(query); err != nil {
		return err
	}
	if err := ctx.Validate(query); err != nil {
		return err
	}

	var limit *uint32
	if query.Limit != nil {
		limit = new(uint32)
		*limit = uint32(*query.Limit)
	}
	var offset *uint32
	if query.Offset != nil {
		offset = new(uint32)
		*offset = uint32(*query.Offset)
	}
	todos, err := t.service.GetMany(ctx.Request().Context(), &pb.GetManyRequest{
		Limit:  limit,
		Offset: offset,
		Status: query.Status,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"messageId": "SUCCESS",
		"data":      todos,
	})
}

func (t *Todo) GetOne(ctx echo.Context) error {
	query := new(schema.TodoGetOneQuery)
	if err := ctx.Bind(query); err != nil {
		return err
	}
	if err := ctx.Validate(query); err != nil {
		return err
	}

	todo, err := t.service.GetOne(ctx.Request().Context(), &wrapperspb.UInt32Value{
		Value: uint32(query.ID),
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"messageId": "SUCCESS",
		"data":      todo,
	})
}

func (t *Todo) UpdateOne(ctx echo.Context) error {
	query := new(schema.TodoUpdateOneQuery)
	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, query); err != nil {
		return err
	}
	if err := ctx.Validate(query); err != nil {
		return err
	}

	body := new(schema.TodoUpdateOneBody)
	if err := (&echo.DefaultBinder{}).BindBody(ctx, body); err != nil {
		return err
	}
	if err := ctx.Validate(body); err != nil {
		return err
	}

	_, err := t.service.UpdateOne(ctx.Request().Context(), &pb.UpdateOneRequest{
		Id:     uint32(query.ID),
		Status: body.Status,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"messageId": "SUCCESS",
		"data":      nil,
	})
}

func (t *Todo) DeleteOne(ctx echo.Context) error {
	query := new(schema.TodoDeleteOneQuery)
	if err := ctx.Bind(query); err != nil {
		return err
	}
	if err := ctx.Validate(query); err != nil {
		return err
	}

	_, err := t.service.DeleteOne(ctx.Request().Context(), &wrapperspb.UInt32Value{
		Value: uint32(query.ID),
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"messageId": "SUCCESS",
		"data":      nil,
	})
}
