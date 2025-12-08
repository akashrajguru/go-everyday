package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		fmt.Printf("ADDED: \"%s\" to your task list. \n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
