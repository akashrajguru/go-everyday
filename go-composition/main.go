package main

import "fmt"

// Basic types
type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) start() {
	fmt.Printf("Strating %s engine with %d HP\n", e.Type, e.Horsepower)
}

type Wheels struct {
	Count int
}

func (w Wheels) Roll() {
	fmt.Printf("Rolling on %d wheels\n", w.Count)

}

// Understanding Conposition
// Compose type  Car "has-a" Engine and "has-a" Wheels
type Car struct {
	Engine //embedded type
	Wheels
	Model string
}

// Car has its own methods
func (c Car) Drive() {
	fmt.Printf("Driving the %s\n", c.Model)
	c.start()
	c.Roll()
}

func main() {
	car := Car{
		Engine: Engine{
			Horsepower: 300,
			Type:       "V8",
		},
		Wheels: Wheels{
			Count: 4,
		},
		Model: "Mustang",
	}

	car.Drive()
	// You can also access embedded fields and methods directly
	fmt.Printf("\nDirect access : %d HP\n", car.Horsepower)
	car.start()
}
