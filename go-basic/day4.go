package main

import "fmt"

func main() {

	// Form 1: the classic 3 parts for loop
	// This is the most common loop for a set number of iterations. similar to c and java

	fmt.Println("Starting classic loop..")
	for i := 0; i < 5; i++ {
		fmt.Println("Classic look i :", i)
	}

	// ---- Form 2: The "While" Loop ----
	fmt.Println("\nStarting 'while' style loop")
	j := 0
	for j < 3 {
		fmt.Println("While loop, j =", j)
		j++
	}
}
