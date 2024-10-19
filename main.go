package main

import (
	"fmt"
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

	for _, todo := range todoList.Todos {
		fmt.Printf(
			"Tâche n°%d: %s | %s | Statut: %s\n",
			todo.ID, todo.Title, todo.Description, todo.Status.String(),
		)
	}

	cmd.Execute()
}
