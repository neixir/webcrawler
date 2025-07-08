package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int  // Keep track of the pages we've crawled
	baseURL            *url.URL        // Keep track of the original base URL
	mu                 *sync.Mutex     // Ensure the pages map is thread-safe
	concurrencyControl chan struct{}   // Ensure we don't spawn too many goroutines at once
	wg                 *sync.WaitGroup // Ensure the main function waits until all in-flight goroutines (HTTP requests) are done before exiting the program
}

func main() {

	arguments := os.Args[1:]

	// If the number of CLI arguments is less than 1, print "no website provided" and exit with code 1.
	if len(arguments) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	// If the number of CLI arguments is more than 1, print "too many arguments provided" and exit with code 1.
	if len(arguments) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// If we have exactly one CLI argument, it's the BASE_URL,
	// so print a message letting the user know the crawler is starting at that URL:
	// starting crawl of: BASE_URL
	// where BASE_URL is the URL provided as the CLI argument.
	startUrl := arguments[0]
	fmt.Printf("starting crawl of: %s\n", startUrl)

	rawBaseURL, err := normalizeURL(arguments[0])
	fmt.Printf("rawBaseURL = %s\n", rawBaseURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	baseURL, _ := url.Parse(startUrl)

	maxConcurrency := 5
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	// Call Add(1) before starting the goroutine so the counter always correctly matches the number of goroutines running.
	cfg.wg.Add(1)
	go cfg.crawlPage(startUrl)
	cfg.wg.Wait()

	fmt.Println("*** PAGES ***")
	for url, count := range cfg.pages {
		fmt.Printf("URL: %s, Count: %d\n", url, count)
	}
}
