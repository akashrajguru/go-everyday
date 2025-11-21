package main

import (
	"fmt"
	"time"
)

func main() {

	// Lets create a buffered channel with capacity 2.
	// We can send 2 items without blocking.
	ch := make(chan string, 2)
	ch <- "Message 1"
	ch <- "Message 2"

	fmt.Println("Sent 2 messages without a receiver yet")

	fmt.Println("Received:", <-ch)

	// Closing and ranging
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
			time.Sleep(200 * time.Millisecond)
		}

		// IMP: WE must close the channel when done.
		close(numbers)
		fmt.Println("Channel closed.")
	}()

	fmt.Println("Waiting for numbers...")
	// range loops over the channel.
	// it automatically receives values.
	// automatically stops when channel is empty and close.
	for num := range numbers {
		fmt.Println("Received number is :", num)
	}

	fmt.Println("Done.")
}
