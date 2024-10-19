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

	todos, err := todo.ReadTodosFile()
	if err != nil {
		log.Fatalf("Error reading todos: %v", err)
	}

	for _, todo := range todos {
		fmt.Printf(
			"Tâche n°%d: %s | %s | Statut: %s\n",
			todo.ID, todo.Title, todo.Description, todo.Status.String(),
		)
	}

	cmd.Execute()
}
