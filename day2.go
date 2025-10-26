// Day 2 we will learn more formal way of using the var keyword, which gives more control.
package main

import "fmt"

func main() {

	// declaring variable without assing value using var and type
	var username string
	fmt.Println("Default username:", username)

	username = "Gopro123"
	fmt.Println("My username is :", username)

	var score int
	fmt.Println("Score:", score)
	score = 100
	fmt.Println("Score:", score)

	var graviry float64 = 9.81
	var isReady bool = true
	const appVersion = "1.0.2"
	fmt.Println("Gravity:", graviry)
	fmt.Println("Ready:", isReady)
	fmt.Println("App Version:", appVersion)
	// appVersion = "1.0.3"

	const goal = "Learn Go"
	var currentDay int = 2
	var progress float64 = 0.1
	fmt.Println("The Goal is to", goal, " today is day ", currentDay, "and progress is ", progress)
}
