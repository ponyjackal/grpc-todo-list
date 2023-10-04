package main

import todo "github.com/ponyjackal/grpc-todo-list/todo"

type todoServer struct{}

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
