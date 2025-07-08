package main

import (
	"net/url"
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
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return allUrls, err
	}

	parsedBaseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return allUrls, err
	}

	traverseNode(htmlNodes, &allUrls, *parsedBaseUrl)

	return allUrls, nil
}

func traverseNode(node *html.Node, links *[]string, parsedBaseUrl url.URL) {
	if node == nil {
		return
	}

	switch node.Type {
	// html.ElementNode - for HTML elements like <a>, <div>, <span>, etc.
	case html.ElementNode:
		// node.Data - contains the tag name for element nodes (like "a", "div", etc.)
		if node.Data == "a" {
			href := findAttr("href", node)
			if href != "" {
				parsedHref, err := url.Parse(href)
				if err != nil {
					return
				}

				absoluteUrl := parsedBaseUrl.ResolveReference(parsedHref)

				// Debug print
				// fmt.Printf("Original href: %s, Absolute URL: %s\n", href, absoluteUrl.String())

				*links = append(*links, absoluteUrl.String())
			}
		}

		// html.TextNode - for text content between tags
		// html.CommentNode - for HTML comments <!-- like this -->
		// html.DoctypeNode - for the <!DOCTYPE html> declaration
		// html.DocumentNode - for the root document node
	}

	// Recursively process children
	// node.FirstChild - pointer to the first child node
	traverseNode(node.FirstChild, links, parsedBaseUrl)

	// Recursively process siblings
	// node.NextSibling - pointer to the next sibling node
	traverseNode(node.NextSibling, links, parsedBaseUrl)

}

func findAttr(attribute string, node *html.Node) string {
	// node.Attr - a slice of attributes for the node
	for _, attr := range node.Attr {
		if attr.Key == attribute {
			return attr.Val
		}
	}
	return ""
}
