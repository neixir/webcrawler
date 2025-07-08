package main

import (
	"fmt"
	"net/url"
	"strings"
)

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
