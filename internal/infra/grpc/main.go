package grpc

import (
	"github.com/MicBun/protobuf-golang-todo/internal/domain/service"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/pb"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(New)

func New(todoService *service.Todo) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, todoService)
	return server
}
