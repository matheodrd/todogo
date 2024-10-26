package cmd

import (
	"log"
	"os"

	"github.com/matheodrd/todogo/todo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todogo",
	Short: "A CLI To-Do list tool",
	Long:  "A CLI To-Do list tool which you can use to add tasks, list them, and edit them.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todogo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	if err := todo.InitTodosFile(); err != nil {
		log.Fatalf("Failed to init todos storage file: %v", err)
	}
	if err := todo.InitCache(); err != nil {
		log.Fatalf("Failed to init cache vars file: %v", err)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
