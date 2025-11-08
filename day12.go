package main

import "fmt"

// Struct is a blueprint that describes the data
type Person struct {
	Name string
	Age  int
}

// Method with value receives (u User).
// Method attaches the function to Struct
// Without (u *User) method operates on the copy of the data, good for reading
func (p Person) Greet() {
	fmt.Printf("Hi, my name is %s and I am %d.\n", p.Name, p.Age)
}

// Method with pointer receiver
// (p *Person) is pointer receiver, operates on a pointer original data.
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func main() {
	person1 := Person{Name: "Alice", Age: 30}
	person1.Greet()

	person1.SetAge(31)

	person1.Greet()
}
