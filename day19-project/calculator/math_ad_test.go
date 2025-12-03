package calculator

import "testing"

func TestAddTableDriven(t *testing.T) {
	// Define the struct that represents a single test case.
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		// lets define the table of cases
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -2, -3},
		{"Mixed numbers", -5, 5, 0},
	}

	// Loop over the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
