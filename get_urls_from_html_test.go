package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	//
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
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
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		// add more test cases here
	}

	// copiat de normalize_url_test.go, canviant nomes la funcio
	// TOO posar la resta de parametres, i a veure si la comprovacio funciona,
	// pq getURLsFromHTML retorna un slice...
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// You may find reflect.DeepEqual to be particularly useful for testing.
			// https://pkg.go.dev/reflect#DeepEqual
			// reflect.DeepEqual(x, y any) bool
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			/*
				if err != nil {
					t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
					return
				}
				if actual != tc.expected {
					t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				}
			*/
		})
	}

}
