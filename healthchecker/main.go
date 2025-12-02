package main

import (
	"fmt"
	"healthchecker/checker"
)

func main() {
	// list of websites
	website := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://go.dev",
		"https://this-site-does-not-exist-123.com",
	}

	// create buffer channel to hold result
	resultChannel := make(chan checker.Result, len(website))
	fmt.Println("--- Starting Concurrent Health check")
	for _, url := range website {
		go func(u string) {
			result := checker.Ping((u))
			resultChannel <- result
		}(url)
	}

	// FAN-In collect the results
	for i := 0; i < len(website); i++ {
		result := <-resultChannel
		if result.Status == "FAILED" {
			fmt.Printf("[DOWN] %s (Time: %s)\n", result.URL, result.Latency)
		} else {
			fmt.Printf("[UP]  %s - %s (Time: %s)\n", result.URL, result.Latency)
		}

	}

	fmt.Println("--- All Checks Complete ---")

}
