// In Day 14 we saw Go's famous feature Concurrency using Goroutine.
// Incredibly lightweight thread managed by go runtime.
// previously we use time.sleep() to wait for goroutines.(bad way)

// Today we will use sync.WaitGroup() counter to wait for goroutine to finish.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Create a worker function which accepts a pointer to the sync.WaitGroup()
// so it ca call Done()
func worker(id int, wg *sync.WaitGroup) {
	// We will use special keyword 'defer' which says run this line of code at very end of this function.
	// Right before it returns.
	// This guarantees Done() is called even if the function had an error.
	defer wg.Done()
	fmt.Printf("Worker %d starting......\n", id)

	// Simulate some work
	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %d done. \n", id)
}

func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	newWorkers := 3

	// tell WaitGroup that we are waiting for 3 goroutines
	wg.Add(newWorkers)

	// lets launch 3 goroutines
	for i := 1; i <= newWorkers; i++ {
		go worker(i, &wg)
	}
	fmt.Println("Main is waiting for workers to finish....")
	wg.Wait()
	fmt.Println("All workers finished. Main is exiting.")
}
