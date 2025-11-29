package main

import "fmt"

// Single function with no parameters and return value
func sayHello() {
	fmt.Println("Hello from a function!")
}

// 2. A function with parameters
// Will take two parameters x and y
// we must specify their types
// and return single value

func add(x int, y int) int {
	sum := x + y
	return sum
}

// Go function can have multi return values
// This function returns two values: a 'string' and an 'error'.
// This is the standard way Go handles functions that might fail.
func devide(numerator int, denominator int) (int, bool) {
	if denominator == 0 {
		// We can't divide by zero.
		// Return 0 (as the int) and 'true' (signaling an error occurred).
		return 0, true
	}

	return numerator / denominator, false
}

// main function
func main() {
	sayHello()

	// result := add(5, 3)
	fmt.Println("5 + 3 is :", add(5, 3))

	// Call 'divide' and catch *both* return values.
	// We use the ':= ' syntax to declare two new variables.

	quotient, didError := devide(10, 2)
	if didError == true {
		fmt.Println("Error : Cannot divide by zero")
	} else {
		fmt.Println("10 / 2 =", quotient)
	}

	// You can use '_' (the blank identifier) to
	// ignore a return value you don't care about.
	_, err := devide(9, 0)
	if err {
		fmt.Println("second division failed")
	}
}
