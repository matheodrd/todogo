package main

import (
	"fmt"

	"github.com/matheodrd/todogo/cmd"
	"github.com/matheodrd/todogo/todo"
)

func main() {
	todo.InitTodosFile()
	todos := todo.ReadTodosFile()

	for _, todo := range todos {
		fmt.Printf(
			"Tâche n°%d: %s | %s | Statut: %s\n",
			todo.ID, todo.Title, todo.Description, todo.Status.String(),
		)
	}

	cmd.Execute()
}
