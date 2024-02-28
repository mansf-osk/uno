package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	rp := httputil.ReverseProxy{
		Rewrite: buildRewriteFunc(target),
	}
	return &rp
}

func buildRewriteFunc(target *url.URL) func(pr *httputil.ProxyRequest) {
	return func(pr *httputil.ProxyRequest) {
		pr.SetXForwarded()
		pr.SetURL(target)
		pr.Out.Header.Set("X-Additional-Header", "this header was added by the proxy")
	}
}

func ServeReverseProxy(addr string, rp *httputil.ReverseProxy) {
	log.Fatal(http.ListenAndServe(addr, rp))
}

func ServeTLSProxy(addr string, certFile string, keyFile string, rp *httputil.ReverseProxy) {
	log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, rp))
}
