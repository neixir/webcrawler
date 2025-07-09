package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		pages    map[string]int
		expected []kv
	}{
		{
			name: "same count",
			pages: map[string]int{
				"https://example.com/page1": 45,
				"https://example.com/page2": 123,
				"https://example.com/page3": 45, // same count as page1
				"https://example.com/page4": 89,
				"https://example.com/page5": 156,
				"https://example.com/about": 45, // same count as page1 and page3
			},
			expected: []kv{
				{"https://example.com/page5", 156},
				{"https://example.com/page2", 123},
				{"https://example.com/page4", 89},
				{"https://example.com/about", 45},
				{"https://example.com/page1", 45},
				{"https://example.com/page3", 45},
			},
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.pages)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
