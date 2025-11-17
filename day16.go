// If Goroutines are the workers, Channels are the pipes connecting them.
// Channels allow goroutines to communicate and synchronize execution.
// They are typed conduits through which you can send and receive values with the channel operator, <-.

// The key concept: Unbuffered channels block.
// Sending (ch <- v) halts execution until someone is ready to receive.
// Receiving (v := <-ch) halts execution until someone sends data.
// This blocking behavior creates automatic synchronization without needing locks.

package main

import (
	"fmt"
	"time"
)

// Lets write a function that will run as goroutine.
// it accepts a channel of type string
func sendMessage(ch chan string) {
	fmt.Println("Goroutine: thinking about what you say...")
	time.Sleep(2 * time.Second)

	// send data to channel.
	// The arrow points to the channel variable.
	ch <- "Hello from the other side." // This line blocks until main is ready to receive.
	fmt.Println("Goroutine: Message sent.")
}

func main() {
	// 1. Create a channel using make.
	// also specify the type of data it carries
	messageChannel := make(chan string)

	// 2. start the goroutine passing the channel
	go sendMessage(messageChannel)

	fmt.Println("Main: Waiting for data...")

	// 3. Receive data from channel.
	//  the arrow points out of the channel variable.
	msg := <-messageChannel // This line blocks the main function until data arrives
	fmt.Println("Main: Received message: ", msg)

}
