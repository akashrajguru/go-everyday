package cmd

import (
	"os"
	"testing"
)

func TestEditCmd(t *testing.T) {

	// 1. Setup the "Fake Editor"
	// This command appends the text "modified" to the file passing in $1
	// 'sh -c' works on Linux/macOS.
	fackEditor := "sh -c 'echo \"modified\" >> $1'"
	// Save the original EDITOR so we can restore it later
	// We need to do this because the EDITOR environment variable is set by the system
	// and we need to restore it after the test
	originalEditor := os.Getenv("EDITOR")
	// Restore the original EDITOR after the test
	defer os.Setenv("EDITOR", originalEditor)
	// Set the EDITOR to our fake editor
	os.Setenv("EDITOR", fackEditor)
	// 2. Setup the Cobra Command arguments
	// We want to simulate running: "kcli edit my-pod"
	rootCmd.SetArgs([]string{"edit", "nginx-75b"})
	// 3. Capture the output (optional, to keep test logs clean)
	// We redirect stdout just for this test execution
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = nil // Silence the "Opening config" logs
	// Run the command
	err := rootCmd.Execute()

	// 5. Assertions
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

}

// Test what happens if the user provides NO arguments
// func TestEditCommand_NoArgs(t *testing.T) {
// 	rootCmd.SetArgs([]string{"edit"})
// 	// We expect an error here because we set Args: cobra.ExactArgs(1)
// 	err := rootCmd.Execute()
// 	if err == nil {
// 		t.Errorf("Expected an error (missing argument), but got none")
// 	}
// }
