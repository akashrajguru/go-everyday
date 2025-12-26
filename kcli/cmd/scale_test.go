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
	if !strings.Contains(output, "Scaling deployment 'my-deployment' to '5' replicas") {
		t.Errorf("Output didn't match expectation. Got:\n%s", output)
	}
}

func TestScaleCommand_ZeroReplicas(t *testing.T) {
	// Test 2. Login Chaeck : Test the warning for 0 replicas
	rootCmd.SetArgs([]string{"scale", "frontend", "--replicas=0"})

	output := captureOutput(func() {
		_ = rootCmd.Execute()
	})

	if !strings.Contains(output, "WARNING: Scaling to 0") {
		t.Errorf("Expected waring message for 0 replicas, but got:\n%s", output)
	}

}

func TestScaleCommand_MissingFlag(t *testing.T) {
	// 1. CRITICAL FIX: Reset the flag's "Changed" state.
	// This makes Cobra forget that the flag was set in a previous test.
	scaleCmd.Flags().Lookup("replicas").Changed = false
	// Test Case 3: Missing the required flag
	// Command "kcli scale backend" (forgot --replicas)
	rootCmd.SetArgs([]string{"scale", "backend"})
	err := rootCmd.Execute()

	if err == nil {
		t.Fatal("Expected error because 'replicas' flag is messing, but got nil")
	}

	// Verify the error
	if !strings.Contains(err.Error(), "required flag(s) \"replicas\" not set") {
		t.Errorf("Unexpected error message: %v", err)
	}
}
