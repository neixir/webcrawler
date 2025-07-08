package main

import (
	"fmt"
	"os"
)

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

	pages := make(map[string]int)
	crawlPage(rawBaseURL, startUrl, pages)

	fmt.Println("*** PAGES ***")
	for url, count := range pages {
		fmt.Printf("URL: %s, Count: %d\n", url, count)
	}
}
