package main

import todo "github.com/ponyjackal/grpc-todo-list/todo"

type todoServer struct{}

func (s *todoServer) GetTodos(req *todo.TodoRequest, stream todo.TodoService_GetTodosServer) error {

}
