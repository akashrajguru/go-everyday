package main

import (
	"fmt"
	"time"
)

// A simple function we want to run concurrently
func say(message string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(message)
	}
}

func main() {
	say("Hello", 3)
	// run function as goroutine
	go say("world", 5)
	// run another function as goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("I am an anonymous goroutine...")

	fmt.Println("Main is waiting...")
	time.Sleep(1 * time.Second)
	fmt.Println("Main is done.")
}
