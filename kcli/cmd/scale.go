/*
This simulates kubectl scale deployment my-dep --replicas=3.
It demonstrates Integer Flags and Positional Arguments.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var replicas int

var scaleCmd = &cobra.Command{
	Use:   "scale [resource-name]",
	Short: "Set a new size for a deployment",
	// Validation : force the user to provide exactly one argument (the name)
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		fmt.Printf("Scaling deployment '%s' to '%d' replicas....\n", resourceName, replicas)
		if replicas == 0 {
			fmt.Println("WARNING: Scaling to 0 will stop all traffic to this app.")
		}

		fmt.Println("Success.")
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)

	scaleCmd.Flags().IntVarP(&replicas, "replicas", "r", 1, "Number of replicas")
	// Mark it as REQUIRED (kubectl style)
	// If user forgets --replicas, CLI will error out immediately.
	scaleCmd.MarkFlagRequired("replicas")
}
