// Go handles Object-Oriented Programming behavior through Method.

// Pointer receiver example method
package main

import "fmt"

type Counter struct {
	value int
}

// Define method on counter struct
// (c *Counter) is the pointer receiver.
// This method 'Increment' is now "attached" to Counter.
func (c *Counter) Increment() {
	c.value++
	fmt.Println("Incrementing Value is now: ", c.value)
}

func main() {
	myCounter := Counter{value: 0}
	myCounter.Increment()
	myCounter.Increment()
	fmt.Println("Final value: ", myCounter.value)
}
