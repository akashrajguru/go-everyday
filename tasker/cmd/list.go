package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verbose bool
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []string{"Buy Milk", "Walk Dog", "Learn Cobra"}
		fmt.Println("YOUR TASKS:")
		for i, task := range tasks {
			if verbose {
				fmt.Printf("%d. %s (Status: Pending, Priority: High)\n", i+1, task)
			} else {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed task info")
}
