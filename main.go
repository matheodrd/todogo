package main

import (
	"log"

	"github.com/matheodrd/todogo/cmd"
	"github.com/matheodrd/todogo/todo"
)

// Yes, I'm using main as a testing function...
func main() {
	if err := todo.InitTodosFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}

	todoList, err := todo.NewTodoList()
	if err != nil {
		log.Fatal(err)
	}

	if err := todo.SaveTodos(todoList.Todos); err != nil {
		log.Fatal(err)
	}

	todoList.Display()

	cmd.Execute()
}
