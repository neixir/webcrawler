package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("============================================")

	sorted := sortPages(pages)

	for _, url := range sorted {
		fmt.Printf("Found %d internal links to %s\n", url.Value, url.Key)
	}

}
