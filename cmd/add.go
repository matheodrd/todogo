package cmd

import (
	"fmt"
	"log"

	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

var description string

var addCmd = &cobra.Command{
	Use:   "add <'title'> [-d 'description']",
	Short: "Add a new todo",
	Long:  `Creates a new todo item with a title and an optional description.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		newTodo := todo.NewTodo(title, description)

		todoList, err := todo.NewTodoList()
		if err != nil {
			log.Fatalf("Could not load todo list: %v", err)
		}

		todoList.AddTodo(newTodo)
		if err := todo.SaveTodos(todoList.Todos); err != nil {
			log.Fatalf("Could not save todos: %v", err)
		}

		fmt.Printf("Todo '%s' added successfully!\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the todo")
}
