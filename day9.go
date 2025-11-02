package main

import "fmt"

func main() {

	var primeNumbers [3]int
	primeNumbers[0] = 2
	primeNumbers[1] = 3
	primeNumbers[2] = 5
	fmt.Println("Array :", primeNumbers)

	vowels := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println("Vowels:", vowels)

	var dynamicNumbers []int
	fmt.Println("Empty slice:", dynamicNumbers)

	// we can use apend which is builtin function to add value in slice
	// it returns a new slice so you have to re-assign it
	dynamicNumbers = append(dynamicNumbers, 10)
	dynamicNumbers = append(dynamicNumbers, 20)
	dynamicNumbers = append(dynamicNumbers, 30)

	fmt.Println("Slice :", dynamicNumbers)
	// we can loop over the array and slice using for..range similar to for..each
	for index, val := range vowels {
		fmt.Println("Index : ", index, "Value :", val)
	}

	for _, num := range dynamicNumbers {
		fmt.Println("Just the value: ", num)
	}
}
