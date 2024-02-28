package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mansf-osk/uno/pkg/proxy"
)

const (
	rawTargetURL = "http://127.0.0.0:8080"
	rawProxyAddr = ":8081"
)

func main() {
	fmt.Printf("Uno listening on %s and redirecting to origin server at %s\n", rawProxyAddr, rawTargetURL)
	proxy.ServeReverseProxy(rawProxyAddr, proxy.NewReverseProxy(parseURL(rawTargetURL)))
}

// Parses raw URL as string into url.URL and checks for errors
func parseURL(target string) *url.URL {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}
	return url
}
