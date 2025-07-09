package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func isSameDomain(url1, url2 string) bool {

	domain1, _ := getDomain(url1)
	domain2, _ := getDomain(url2)

	return strings.EqualFold(domain1, domain2)
}

func getDomain(link string) (string, error) {
	parsed, err := url.Parse(link)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return parsed.Host, nil

}

// https://claude.ai/chat/530556b2-1565-42f8-9e2c-ef60535a29d4
func sortPages(pages map[string]int) []kv {

	// Convertim el map a un slice
	var sorted []kv
	for k, v := range pages {
		sorted = append(sorted, kv{k, v})
	}

	// Ordenem per valor (descending)
	// i despres per clau (alfabeticament)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}

		return sorted[i].Value > sorted[j].Value
	})

	return sorted

}
