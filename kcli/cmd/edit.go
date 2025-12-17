package cmd

import "github.com/spf13/cobra"

var editCmd = &cobra.Command{
	Use:   "edit [resource-name]",
	Short: "Edit a resource configuration",
}
