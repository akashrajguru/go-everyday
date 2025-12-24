package cmd

import (
	"strings"
	"testing"
)

// We can reuse the captureOutput helper from get_pods_test.go.

func TestScaleCommand(t *testing.T) {
	// Test Success Case "kcli scale my-deployment --replicas=5"
	rootCmd.SetArgs([]string{"scale", "my-deployment", "--replicas=5"})

	output := captureOutput(func() {
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("Expeted success, but got error: %v", err)
		}
	})

	// Assertions
	if !strings.Contains(output, "Scaling deployment 'my-deployment' to 5 replicas") {
		t.Errorf("Output didn't match expectation. Got:\n%s", output)
	}
}
