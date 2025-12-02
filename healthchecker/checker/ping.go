package checker

import (
	"net/http"
	"time"
)

type Result struct {
	URL     string
	Status  string
	Latency time.Duration
}

func Ping(url string) Result {
	start := time.Now()

	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)

	duration := time.Since(start)

	res := Result{
		URL:     url,
		Latency: duration,
	}

	if err != nil {
		res.Status = "FAILED"

	} else {
		res.Status = resp.Status

		resp.Body.Close()
	}
	return res
}
