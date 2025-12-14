package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or more resources",
	Long:  `Get resources from the cluster`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
