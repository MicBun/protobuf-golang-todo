//go:build wireinject
// +build wireinject

package internal

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/adapter/httpHandler"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/service"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/db"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/echo"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/factory"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/grpc"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/repo"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/route"
	"github.com/google/wire"
)

func InitApp(ctx context.Context) (*App, error) {
	wire.Build(
		route.NewHTTP,
		NewApp,

		// Infra
		echo.ProviderSet,
		db.ProviderSet,
		grpc.ProviderSet,

		repo.NewTodo,
		factory.NewTodo,
		wire.Bind(new(contract.TodoRepo), new(*repo.Todo)),

		// Adapter
		httpHandler.NewTodo,

		// Domain
		service.NewTodo,
	)

	return nil, nil
}
