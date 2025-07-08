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
	// startUrl, err := normalizeURL(arguments[0])
	startUrl := arguments[0]
	fmt.Printf("starting crawl of: %s\n", startUrl)

	html, err := getHTML(startUrl)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(html)
}
