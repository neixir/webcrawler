package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World!\n")
	test()
}

func test() {
	inputURL := "https://blog.boot.dev"
	inputBody := `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`
	// expected := []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"}

	actual, _ := getURLsFromHTML(inputBody, inputURL)
	fmt.Println(actual)
}
