package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker is a fast CLI task manager",
	Long:  `A Fast and Flexible CLI Task Manager build with love in go and Cobra.`,
	// This runs if you call tasker with no sub-commands
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Tasker! Use 'tasker --help' to see commands.")
	},
}

// Execute adds all child commands to the root command and sets the flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
