// Today's lesson is on Control Flow using if, else if, and else.
// This is how your program makes decisions.

package main

import "fmt"

// In Go, conditional logic is simple and readable.
// The main difference from many languages is that there are no parentheses () around
// the conditions, but the curly braces {} are always mandatory.

func main() {

	score := 75
	var grade string

	if score > 90 {
		grade = "A"

	} else if score > 60 {
		grade = "B"
	} else {
		grade = "C"
	}

	fmt.Println("Your score is :", score)
	fmt.Println("Your grade is :", grade)
}
