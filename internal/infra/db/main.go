package db

import (
	"github.com/MicBun/protobuf-golang-todo/internal/domain/contract"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGormDB,
	NewPostgresTransactionManager,
	wire.Bind(new(contract.TransactionManager), new(*PostgresTransactionManager)),
)
