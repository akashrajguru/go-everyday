// --- DEBUG ME ---
package main

import "fmt"

func main() {
	ch := make(chan int)

	// Producer
	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
		}
		// Something is missing here...
		close(ch)
	}()

	// Consumer
	// The loop keeps asking the channel for more data
	for n := range ch {
		fmt.Println(n)
	}

	fmt.Println("All done!")
}
