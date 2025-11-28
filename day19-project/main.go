package main

import (
	"day19-project/calculator"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	result := calculator.Add(10, 5)
	fmt.Println("10 + 5 :", result)

	// calculator.subtract(10, 5)

	color.Green("This text printed in Green")
	color.Red("This text is printed in red!")
}
