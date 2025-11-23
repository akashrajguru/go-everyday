package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Data from Server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Data from Server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	for i := 0; i < 2; i++ {
		fmt.Println("Waiting....")
		select {
		case msg1 := <-output1:
			fmt.Println("Received:", msg1)
		case msg2 := <-output2:
			fmt.Println("Received:", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout one server is too slow.")
		}
	}
}
