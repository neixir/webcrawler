package main

import (
	"fmt"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	// Block if too many goroutines are running:
	// Send an empty struct into the channel.
	cfg.concurrencyControl <- struct{}{}

	// Guarantee cleanup:
	// Use defer to make sure two things always happen, even if your function returns early:
	// Remove this goroutine from the WaitGroup counter:
	defer cfg.wg.Done()

	// Free up a spot in the concurrency channel:
	defer func() { <-cfg.concurrencyControl }()
	// This ensures that your program won’t hang, and future crawlers can run when there’s room.

	// Make sure the rawCurrentURL is on the same domain as the rawBaseURL.
	// If it's not, just return. We don't want to crawl the entire internet, just the domain in question.
	rawBaseUrl := fmt.Sprintf("%s://%s", cfg.baseURL.Scheme, cfg.baseURL.Host)
	if isSameDomain(rawBaseUrl, rawCurrentURL) {
		fmt.Printf("* Crawling %s\n", rawCurrentURL)
	} else {
		fmt.Printf("* Ignoring %s (different domain)\n", rawCurrentURL)
		// fmt.Printf("  rawBaseUrl = %s\n", rawBaseUrl)
		return
	}

	// Get a normalized version of the rawCurrentURL.
	normalizedURL, _ := normalizeURL(rawCurrentURL)

	if cfg.addPageVisit(normalizedURL) {
		// Get the HTML from the current URL,
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// and add a print statement so you can watch your crawler in real-time.
		fmt.Println(normalizedURL)

		// Assuming all went well with the request, get all the URLs from the response body HTML
		moreUrls, err := getURLsFromHTML(html, rawCurrentURL)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Recursively crawl each URL on the page
		for _, link := range moreUrls {
			cfg.wg.Add(1)
			go cfg.crawlPage(link)
		}
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return false
	} else {
		cfg.pages[normalizedURL] = 1
		return true
	}
}
