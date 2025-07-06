package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(checkUrl string) (string, error) {
	/*
	   type URL struct {
	   	Scheme      string
	   	Opaque      string    // encoded opaque data
	   	User        *Userinfo // username and password information
	   	Host        string    // host or host:port (see Hostname and Port methods)
	   	Path        string    // path (relative paths may omit leading slash)
	   	RawPath     string    // encoded path hint (see EscapedPath method)
	   	OmitHost    bool      // do not emit empty host (authority)
	   	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	   	RawQuery    string    // encoded query values, without '?'
	   	Fragment    string    // fragment for references, without '#'
	   	RawFragment string    // encoded fragment hint (see EscapedFragment method)
	   }
	*/

	parsed, err := url.Parse(checkUrl)
	if err != nil {
		return "", err
	}

	newUrl := fmt.Sprintf("%s%s", parsed.Host, parsed.Path)
	// fmt.Printf("NewURL: %s\n", newUrl)

	return newUrl, nil
}
