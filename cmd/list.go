package cmd

import (
	"fmt"
	"log"

	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  `Display the list of all the todos in the todolist.`,
	Run: func(cmd *cobra.Command, args []string) {
		todoList, err := todo.NewTodoList()
		if err != nil {
			log.Fatalf("Could not load the todo list: %v", err)
		}

		if len(todoList.Todos) == 0 {
			fmt.Println("The todo list is empty.")
		}

		for _, todo := range todoList.Todos {
			fmt.Println(todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
