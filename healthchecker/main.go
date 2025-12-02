package main

import (
	"encoding/csv"
	"fmt"
	"healthchecker/checker"
	"os"
)

func main() {
	file2, err := os.Create("data1.csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
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

	writer := csv.NewWriter(file2)
	defer writer.Flush()

	writer.Write([]string{"URL", "Status", "Latency"})
	// FAN-In collect the results
	for i := 0; i < len(website); i++ {
		result := <-resultChannel
		var row []string

		if result.Status == "FAILED" {
			fmt.Printf("[DOWN] %s (Time: %s)\n", result.URL, result.Latency)
			row = []string{result.URL, "DOWN", result.Latency.String()}
		} else {
			fmt.Printf("[UP]  %s - %s (Time: %s)\n", result.URL, result.Status, result.Latency)
			row = []string{result.URL, result.Status, result.Latency.String()}
		}

		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error writing to CSV", err)
		}

	}

	fmt.Println("--- All Checks Complete ---")

}
