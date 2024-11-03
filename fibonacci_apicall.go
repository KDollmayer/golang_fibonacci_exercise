package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func CallAPIWithRetries(maxRetries int) {
	attempt := 0
	for attempt = 1; attempt <= maxRetries; attempt++ {
		success := makeAPICall(attempt)
		if success {
			fmt.Println("API call successful!")
			return
		}
		delay := fibonacci(attempt) * 100
		fmt.Printf("Retry attemp %d failed. Retrying in %d ms \n", attempt, delay)
		time.Sleep(time.Duration(delay) * time.Millisecond)

	}

	fmt.Println("All retries failed")
}

func makeAPICall(n int) bool {
	return n == 5
}
