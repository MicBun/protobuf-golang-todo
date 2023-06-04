package main

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal/infra/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"os"
	"time"
)

func main() {
	conn, err := grpc.Dial(os.Getenv("GRPC_TARGET"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	item := &pb.CreateOneRequest{
		Task: "Test",
	}

	// CreateItem
	createdItem, err := client.CreateOne(ctx, item)
	if err != nil {
		log.Fatalf("Failed to create item: %v", err)
	}
	log.Printf("Created item: %+v", createdItem)

	// ReadItem
	readItem, err := client.GetOne(ctx, &wrapperspb.UInt32Value{Value: createdItem.Id})
	if err != nil {
		log.Fatalf("Failed to read item: %v", err)
	}
	log.Printf("Read item: %+v", readItem)

	// ReadItems
	var limit uint32 = 999
	var offset uint32 = 0
	readItems, err := client.GetMany(ctx, &pb.GetManyRequest{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		log.Fatalf("Failed to read items: %v", err)
	}
	log.Printf("Read items: %+v", readItems)

	// UpdateItem
	updatedItem, err := client.UpdateOne(ctx, &pb.UpdateOneRequest{
		Id:     createdItem.Id,
		Status: true,
	})
	if err != nil {
		log.Fatalf("Failed to update item: %v", err)
	}
	log.Printf("Updated item: %+v", updatedItem)

	// DeleteItem
	deletedItem, err := client.DeleteOne(ctx, &wrapperspb.UInt32Value{Value: createdItem.Id})
	if err != nil {
		log.Fatalf("Failed to delete item: %v", err)
	}
	log.Printf("Deleted item: %+v", deletedItem)
}
