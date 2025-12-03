package calculator

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2,3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 2)
	want := 5
	if got != want {
		t.Errorf("Subtract(7,2) = %d; want %d", got, want)
	}
}
