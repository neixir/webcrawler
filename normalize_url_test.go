package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove user/pass",
			inputURL: "https://username:password@blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove query params",
			inputURL: "https://blog.boot.dev/path?s=1",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove fragments",
			inputURL: "https://blog.boot.dev/path#troubleshooting",
			expected: "blog.boot.dev/path",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
