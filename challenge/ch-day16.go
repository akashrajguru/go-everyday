package main

import "fmt"

func main() {
	// Create a channel
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	// We try to send a message to the channel

	// Then we try to read it
	msg := <-ch
	fmt.Println(msg)
}
