package main

import "fmt"

func main() {
	fmt.Println("This the Day 5 for Go Programming!")
	fmt.Println("Today we will look Data Types & Arithmetic Operators")

	// PART 1: Variables and data types
	var language string = "Go"
	var version = 1.22
	isLearning := true

	// zero values
	var aString string
	var anInt int
	var aBool bool

	fmt.Println(language, version, isLearning)
	fmt.Println("Zero values: ", aString, anInt, aBool)

	// PART 2: Operators

	// 2.1 Arithmetic Operators
	x := 10
	y := 3
	fmt.Println("Addition (x+y)", x+y)
	fmt.Println("Subtraction (x-y)", x-y)
	fmt.Println("Multiplication (x*y)", x*y)
	fmt.Println("Division (x/y)", x/y)
	fmt.Println("Modulus (x-y)", x%y)
	// 2.2 Comparison Operator (results in bool)
	fmt.Println("x > y", x > y)
	fmt.Println("x == 10", x == 10)
	fmt.Println("y != 3:", y != 3)

}
