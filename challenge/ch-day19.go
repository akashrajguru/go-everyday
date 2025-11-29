// --- main.go ---
package main

import (
	"fmt"

	"github.com/google/uuid" // Generates unique IDs
)

func main() {
	id := uuid.New()
	fmt.Println("Generated ID:", id)
}
