package main

import (
	"log"
	"net"

	todo "github.com/ponyjackal/grpc-todo-list/todo"
	"google.golang.org/grpc"
)

type todoServer struct {
	todo.UnimplementedTodoServiceServer
}

func (s *todoServer) GetTodos(req *todo.TodoRequest, stream todo.TodoService_GetTodosServer) error {
	todos := []*todo.Todo{
		{Id: "1", Title: "Buy groceries", Completed: false},
		{Id: "2", Title: "Clean the house", Completed: true},
		{Id: "3", Title: "Walk the dog", Completed: false},
	}

	for _, todo := range todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &todoServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
