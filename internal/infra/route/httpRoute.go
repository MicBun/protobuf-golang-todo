package route

import (
	"github.com/MicBun/protobuf-golang-todo/internal/adapter/httpHandler"
	"github.com/labstack/echo/v4"
)

type HTTP struct {
	todo *httpHandler.Todo
}

func NewHTTP(
	todo *httpHandler.Todo,
) *HTTP {
	return &HTTP{
		todo,
	}
}

func (h *HTTP) Load(e *echo.Group) {
	todoGRPC := e.Group("/todo")
	todoGRPC.POST("", h.todo.CreateOne)
	todoGRPC.GET("/many", h.todo.GetMany)
	todoGRPC.GET("/one", h.todo.GetOne)
	todoGRPC.PUT("/one", h.todo.UpdateOne)
	todoGRPC.DELETE("/one", h.todo.DeleteOne)
}
