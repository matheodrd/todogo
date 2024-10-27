package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update {title | description | status} <arg>",
	Short: "Update an attribute of the selected todo",
	Long: `Update an attribute of the selected todo.
A todo must first be selected using the 'select' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

var titleSubCmd = &cobra.Command{}
var descriptionSubCmd = &cobra.Command{}
var statusSubCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(titleSubCmd, descriptionSubCmd, statusSubCmd)
}
