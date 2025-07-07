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
		expected  []string
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
		{
			name:     "no anchor tags",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
			<span>Boot.dev</span>
			<p>Nothing to see here</p>
			</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name:     "only absolute urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<a href="https://www.3cat.cat/3cat/els-conclaves-de-1978/audio/1249437/">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
			</html>
			`,
			expected: []string{"https://www.3cat.cat/3cat/els-conclaves-de-1978/audio/1249437/", "https://other.com/path/one"},
		},
		{
			name:     "different types of relative URLs ",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<a href="../els-conclaves-de-1978/audio/1249437/">
					<span>Boot.dev</span>
				</a>
				<a href="./nose">
					<span>Boot.dev</span>
				</a>
				<a href="nose.html">
					<span>Boot.dev</span>
				</a>
			</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/els-conclaves-de-1978/audio/1249437/", "https://blog.boot.dev/nose", "https://blog.boot.dev/nose.html"},
		},
		// add more test cases here
	}

	// Adaptat de normalize_url_test.go
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			// You may find reflect.DeepEqual to be particularly useful for testing.
			// https://pkg.go.dev/reflect#DeepEqual
			// reflect.DeepEqual(x, y any) bool
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
