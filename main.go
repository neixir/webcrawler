package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int  // Keep track of the pages we've crawled
	baseURL            *url.URL        // Keep track of the original base URL
	mu                 *sync.Mutex     // Ensure the pages map is thread-safe
	concurrencyControl chan struct{}   // Ensure we don't spawn too many goroutines at once
	wg                 *sync.WaitGroup // Ensure the main function waits until all in-flight goroutines (HTTP requests) are done before exiting the program
	maxPages           int             // maximum number of pages to crawl.
}

func main() {

	arguments := os.Args[1:]

	if len(arguments) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	if len(arguments) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	startUrl := arguments[0]

	maxConcurrency, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Second argument \"maxConcurrency\" must be a number")
	}

	maxPages, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Thred argument \"maxPages\" must be a number")
	}

	fmt.Printf("starting crawl of: %s\n", startUrl)

	rawBaseURL, err := normalizeURL(arguments[0])
	fmt.Printf("rawBaseURL = %s\n", rawBaseURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	baseURL, _ := url.Parse(startUrl)

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	// Call Add(1) before starting the goroutine so the counter always correctly matches the number of goroutines running.
	cfg.wg.Add(1)
	go cfg.crawlPage(startUrl)
	cfg.wg.Wait()

	fmt.Println("*** PAGES ***")
	i := 1
	for url, count := range cfg.pages {
		fmt.Printf("[%02d] URL: %s, Count: %d\n", i, url, count)
		i++
	}
}
