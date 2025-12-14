package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outputFormat string
var mockData = []string{"nginx-75b", "api-gateway-2f", "db-0"}
var podsCmd = &cobra.Command{
	Use:     "pods",
	Aliases: []string{"pod", "po"},
	Short:   "List all pods",
	Long:    `List all pods in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		// Mock data
		pods := []string{"nginx-75b", "api-gateway-2f", "db-0"}
		fmt.Println("Retrieving pods...")

		// Handle the Flog logic
		if outputFormat == "json" {
			fmt.Printf("JSON OUTPUT: { \"kind\": \"PodList\", \"items\": %v }\n", pods)
		} else if outputFormat == "wide" {
			fmt.Println("NAME            IP           NODE")
			fmt.Printf("%s       10.0.0.1     node-1\n", pods[0])
			fmt.Printf("%s  10.0.0.2     node-2\n", pods[1])
		} else {
			fmt.Println("NAME")
			for _, pod := range pods {
				fmt.Println(pod)
			}
		}
	},
}

func init() {
	getCmd.AddCommand(podsCmd)
	podsCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format: text|josn|wide")
}
