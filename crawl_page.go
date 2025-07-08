package main

import (
	"fmt"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	// Make sure the rawCurrentURL is on the same domain as the rawBaseURL.
	// If it's not, just return. We don't want to crawl the entire internet, just the domain in question.
	if isSameDomain("https://"+rawBaseURL, rawCurrentURL) {
		fmt.Printf("* Crawling %s\n", rawCurrentURL)
	} else {
		fmt.Printf("* Ignoring %s (different domain)\n", rawCurrentURL)
		return
	}

	// Get a normalized version of the rawCurrentURL.
	normalizedCurrentUrl, _ := normalizeURL(rawCurrentURL)

	// If the pages map already has an entry for the normalized version of the current URL,
	// just increment the count and be done, we've already crawled this page.
	_, ok := pages[normalizedCurrentUrl]
	if ok {
		pages[normalizedCurrentUrl]++
	} else {
		// Otherwise, add an entry to the pages map for the normalized version of the current URL,
		// and set the count to 1.
		pages[normalizedCurrentUrl] = 1

		// Get the HTML from the current URL,
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// and add a print statement so you can watch your crawler in real-time.
		fmt.Println(normalizedCurrentUrl)

		// Assuming all went well with the request, get all the URLs from the response body HTML
		moreUrls, err := getURLsFromHTML(html, rawCurrentURL)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Recursively crawl each URL on the page
		for _, link := range moreUrls {
			crawlPage(rawBaseURL, link, pages)
		}
	}
}
