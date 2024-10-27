package cmd

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select <id>",
	Short: "Select a todo",
	Long: `Select a todo using its ID. The ID of the selected todo is cached and
following commands (update) will target the selected todo.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := uuid.Parse(args[0])
		if err != nil {
			log.Fatalf("Error parsing given ID: %v", err)
		}

		todoList, err := todo.NewTodoList()
		if err != nil {
			log.Fatalf("Could not load the todo list: %v", err)
		}

		selectedTodo, err := todoList.FindTodo(id)
		if err != nil {
			log.Fatalf("Given ID doesn't exist: %v", err)
		}

		if err := todo.SetVar("SelectedTodoID", id.String()); err != nil {
			log.Fatalf("Error caching given ID: %v", err)
		}

		fmt.Printf("Selected todo: '%s'\n", selectedTodo.Title)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
