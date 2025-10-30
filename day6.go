package main

import "fmt"

func main() {
	score := 88

	//  Basic if condition
	if score > 80 {
		fmt.Println("Great job you score 1.1")
	}
	// if else lader
	if score == 100 {
		fmt.Println("Perfect score of 100")
	} else if score > 80 {
		fmt.Println("Thats a B grade")
	} else {
		fmt.Println("You passed")
	}

	// Logical operators AND (&&) OR (||)
	age := 25
	hasLicense := true

	// AND operator, both conditions must be true
	if age > 18 && hasLicense == true {
		fmt.Println("You can Drive.")
	}

	// OR operator One condition must be true
	day := "Sunday"
	if day == "Saturday" || day == "Sunday" {
		fmt.Println("Its weekend")
	}

	// 4. Go-specific: 'if' with a short statement
	// You can declare a variable right before the condition.
	// The 'limit' variable *only* exists inside this 'if/else' block.
	if limit := 50; score < limit {
		fmt.Println("Score is below the limit")
	} else {
		fmt.Println("Score is above the limit of", limit)
	}

}
