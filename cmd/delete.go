package cmd

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a todo",
	Long:  `Delete a todo from the todo list given its ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idTodoDelete, err := uuid.Parse(args[0])
		if err != nil {
			log.Fatalf("Error parsing given ID: %v", err)
		}

		todoList, err := todo.NewTodoList()
		if err != nil {
			log.Fatalf("Could not load the todo list: %v", err)
		}

		title, err := todoList.RemoveTodo(idTodoDelete)
		if err != nil {
			log.Fatalf("Could not remove the todo: %v", err)
		}

		if err := todo.SaveTodos(todoList.Todos); err != nil {
			log.Fatalf("Could not save todos: %v", err)
		}

		fmt.Printf("Todo '%s' deleted successfully!\n", title)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
