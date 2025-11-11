// An interface is a contract for behavior.
// It defines a set of methods that a type must have to be considered "of" that interface type.
// Go's interfaces are satisfied implicitly.
// You never have to write type MyStruct implements MyInterface.
// If your struct simply has all the methods defined in the interface, it automatically satisfies that interface.
package main

import (
	"fmt"
	"math"
)

// Define shape interface.
// it must have method like Area() that return a float64
type Shape interface {
	Area() float64
}

// Define struct Rectangle
type Rectangle struct {
	Width  float64
	Health float64
}

// Implement Area method on Rectangle.
// Rectangle has Area() float64 it implicitly satisfies the shape interface.
func (r Rectangle) Area() float64 {
	return r.Width * r.Health
}

// Another struct circle
type Circle struct {
	Radius float64
}

// Implement area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// The magic of Go
// We can write function that takes the interface as a parameter.
func PrintShapeArea(s Shape) {
	fmt.Println("The area of this shape is: ", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Health: 5}
	circ := Circle{Radius: 7}

	PrintShapeArea(rect)
	PrintShapeArea(circ)
}
