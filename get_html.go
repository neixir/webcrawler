package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Use http.Get to fetch the webpage of the rawURL
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Return an error if the HTTP status code is an error-level code (400+)
	if res.StatusCode > 400 {
		return "", fmt.Errorf("error-level code (400+)")
	}

	// Return an error if the response content-type header is not text/html
	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("content-type is not text/html")
	}

	// Return any other possible errors

	// Return the webpage's HTML if successful
	return string(body), nil
}
