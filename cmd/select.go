package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select <id>",
	Short: "Select a todo",
	Long: `Select a todo using its ID. The ID of the selected todo is cached and
following commands (update, delete) will target the selected todo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("select called")
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
