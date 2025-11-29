// Go has one looping keyword that is for
package main

import "fmt"

func main() {

	// There is only one keyword for looping in Go is for
	// The Classic C style loop will have 3 parts init, condition and post
	fmt.Println("Classic for loop in go")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	//  for can also act like while loop
	// just leave out the init and post statement
	fmt.Println("While style loop")
	j := 0
	for j < 5 {
		fmt.Println("J =", j)
		j++
	}

	// You can also write infinite loop with break
	fmt.Println("Infinite loop with break")
	k := 0
	for { // an empty for is an infinite loop
		fmt.Println("K =", k)
		if k == 2 {
			fmt.Println("Breaking point")
			break
		}
		k++
	}
}
