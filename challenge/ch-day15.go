// --- DEBUG ME ---
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d running\n", id)
	time.Sleep(500 * time.Millisecond)

	// Oops, something is wrong here!
	if id == 2 {
		fmt.Println("Worker 2 forgot something!")
		return // Worker 2 exits early!
	}

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	// This program has two bugs!

	// Bug 1 is here
	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)
	// Bug 2 is related to worker 3

	fmt.Println("Waiting for workers...")
	wg.Wait()
	fmt.Println("Done.")
}
