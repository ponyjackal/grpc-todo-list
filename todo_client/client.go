package main

import (
	"context"
	"io"
	"log"

	"github.com/ponyjackal/grpc-todo-list/todo"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect", err)
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	req := &todo.TodoRequest{UserId: "user123"}
	stream, err := client.GetTodos(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get todos: %v", err)
	}

	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving todo: %v", err)
		}
		log.Printf("Received Todo: %v", todo)
	}
}
