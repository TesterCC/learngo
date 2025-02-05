package main

import (
	"net/http"
	"sync"
)

// https://pkg.go.dev/sync#example-WaitGroup
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com/",
		"http://www.douban.com/",
		"http://www.163.com/",
	}

	for _,url := range urls {
		//  Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Fetch the URL.
			http.Get(url)

		}(url)
	}

	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
