package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestAddWithAssert(t *testing.T) {
	got := Add(2, 3)
	assert.Equal(t, 5, got, "they should be equal")
}
