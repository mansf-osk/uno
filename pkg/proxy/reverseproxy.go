package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const RawOutURL = "http://127.0.0.0:8080"

func NewReverseProxy(outURL *url.URL) *httputil.ReverseProxy {
	rp := httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetXForwarded()
			pr.SetURL(outURL)
			pr.Out.Header.Set("X-Additional-Header", "this header was added by the proxy")
		},
	}
	return &rp
}

func NewReverseProxyHandler(rp *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		rp.ServeHTTP(rw, req)
	}
}

func ServeReverseProxy(rp *httputil.ReverseProxy) {
	log.Fatal(http.ListenAndServe(":8081", rp))
}

func Rewrite(pr *httputil.ProxyRequest) {
	pr.SetXForwarded()
	pr.SetURL(ParseURL(RawOutURL))
}

// Parses raw URL as string into url.URL and checks for errors
func ParseURL(target string) *url.URL {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}
	return url
}
