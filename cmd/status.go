package cmd

import (
	"log"

	"github.com/google/uuid"
	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status {todo | doing | done}",
	Short: "Update the status of the selected todo",
	Long: `Update the status of the selected todo.
A todo must first be selected using the 'select' command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newStatus, err := todo.ParseStatus(args[0])
		if err != nil {
			log.Fatal(err)
		}

		vars, err := todo.LoadCache()
		if err != nil {
			log.Fatalf("Error loading cached variables: %v", err)
		}

		todoList, err := todo.NewTodoList()
		if err != nil {
			log.Fatalf("Could not load the todo list: %v", err)
		}

		id, err := uuid.Parse(vars.SelectedTodoID)
		if err != nil {
			log.Fatalf("Error parsing cached ID of the selected todo: %v", err)
		}

		if err := todoList.UpdateTodoStatus(id, newStatus); err != nil {
			log.Fatal(err)
		}

		if err := todo.SaveTodos(todoList.Todos); err != nil {
			log.Fatalf("Could not save todos: %v", err)
		}
	},
}

func init() {
	updateCmd.AddCommand(statusCmd)
}
