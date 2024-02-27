package main

import (
	"github.com/mansf-osk/uno/pkg/proxy"
)

func main() {
	rp := proxy.NewReverseProxy(proxy.ParseURL(proxy.RawOutURL))
	proxy.ServeReverseProxy(rp)
}
