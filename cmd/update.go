package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update {title | description | status} <arg>",
	Short: "Update an attribute of the selected todo",
	Long: `Update an attribute of the selected todo.
A todo must first be selected using the 'select' command.`,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
