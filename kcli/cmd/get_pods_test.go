package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout

func captureOutput(f func()) string {
	// Create pipe (r = read, w = write)
	r, w, _ := os.Pipe()

	// Save original stdout so we can restore it later
	originalStdout := os.Stdout
	os.Stdout = w

	// Run the function tat prints text
	f()

	// Close the write end and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read what was written to the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestGetPods(t *testing.T) {
	// Setup the arguments for the commands
	// we will simulate running "kcli get pods"
	rootCmd.SetArgs([]string{"get", "pods"})

	// Run the command inside our capture helper
	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Command executed with error: %v", err)
		}
	})

	// Assertions
	expectedPod := "nginx-75b"
	if !strings.Contains(output, expectedPod) {
		t.Errorf("Expected output to contain '%s', but got:\n%s", expectedPod, output)
	}
}

// Lets test -o json
func TestGetPodsJSON(t *testing.T) {
	// test the JSON flag: "Kcli get podss -o json"
	rootCmd.SetArgs([]string{"get", "pods", "-o", "json"})

	output := captureOutput(func() {
		_ = rootCmd.Execute()
	})

	// Check if output actually contains JSON structures
	if !strings.Contains(output, "\"kind\": \"PodList\"") {
		t.Errorf("Expected JSON output, but got:\n%s", output)
	}
}
