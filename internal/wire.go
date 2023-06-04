//go:build wireinject
// +build wireinject

package internal

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	"github.com/MicBun/protobuf-golang-todo/internal/domain/service"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/db"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/factory"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/grpc"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/repo"
	"github.com/google/wire"
)

func InitApp(ctx context.Context) (*App, error) {
	wire.Build(
		NewApp,

		// Infra
		db.ProviderSet,
		grpc.ProviderSet,

		repo.NewTodo,
		factory.NewTodo,
		wire.Bind(new(contract.TodoRepo), new(*repo.Todo)),

		// Domain
		service.NewTodo,
	)

	return nil, nil
}
