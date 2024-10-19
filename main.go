package main

import (
	"log"

	"github.com/matheodrd/todogo/cmd"
	"github.com/matheodrd/todogo/todo"
)

func main() {
	if err := todo.InitTodosFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}

	todoList, err := todo.NewTodoList()
	if err != nil {
		log.Fatal(err)
	}

	todoList.AddTodo(todo.NewTodo(
		"Write a RemoveTodo method",
		"Removes a todo from the list using its ID",
	))

	todoList.Display()

	cmd.Execute()
}
