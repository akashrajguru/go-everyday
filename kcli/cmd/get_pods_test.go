package cmd

import (
	"bytes"
	"io"
	"os"
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
