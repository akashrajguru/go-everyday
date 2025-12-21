package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [resource-name]",
	Short: "Edit a resource configuration",
	Long:  "Opens the default system editor to modify the resource.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		// Determine which editor to use
		// We check the environment variable "EDITOR". If empty, default to "vim".
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim" // Fallback for Mac/Linux. On Windows, you might want "notepad"
		}

		fmt.Printf("Opening config for '%s' using %s... \n", resourceName, editor)
		// we will create dummy file to edit (Simulating fetching current config)
		fileName := "temp_config.yaml"
		initialContent := []byte(fmt.Sprintf("apiVersion: v1\nkind: pod\nmetadata:\n name: %s\n", resourceName))
		_ = os.WriteFile(fileName, initialContent, 0644)

		// Lets prepare the command
		// tell os to run editor on our file
		command := exec.Command(editor, fileName)
		// The MAGIC: connect the external editor to your terminal
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		// run the editor and wait for it to close
		err := command.Run()

		if err != nil {
			fmt.Printf("Error opening editor: %v\n", err)
			return
		}

		fmt.Println("Edit complete. Configuration saved.")
		// Cleanup remove the temp file
		defer os.Remove(fileName)
		fmt.Println("File Deleted")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
