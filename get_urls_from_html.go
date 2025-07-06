package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	allUrls := []string{}

	// strings.NewReader(htmlBody) creates a io.Reader
	htmlReader := strings.NewReader(htmlBody)

	// https://pkg.go.dev/golang.org/x/net/html#example-Parse
	// https://pkg.go.dev/golang.org/x/net/html#Node
	// html.Parse(htmlReader) creates a tree of html.Nodes
	htmlNodes, _ := html.Parse(htmlReader)
	//fmt.Printf("htmlNodes: %v", htmlNodes)
	for _, node := range htmlNodes {
		fmt.Printf("Type: %v / Data: %v\n", node.Type, node.Data)
	}

	// Use recursion to traverse the node tree and find the <a> tag "anchor" elements
	// In HTML, "anchor" elements are links. e.g:
	//	<a href="https://www.boot.dev">Learn Backend Development</a>

	// z := html.NewTokenizer(htmlReader)

	// traverseNode(z)

	return allUrls, nil
}

// https://pkg.go.dev/golang.org/x/net/html
// Copiat tal qual, a veure
/*
func traverseNode(z *Tokenizer) error { //}([]string, error) {
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				emitBytes(z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
}
*/
