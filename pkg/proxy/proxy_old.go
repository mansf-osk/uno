package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func StartReverseProxy() {
	originServerURL, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		log.Fatalf("invalid origin server URL: %s", err)
	}

	reverseProxyHandler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Printf("[PROXY] received request at: %s\n", time.Now())
		// fmt.Println("[PROXY] request content:", request)

		// change request data to send to origin server instead
		request.Host = originServerURL.Host
		request.URL.Host = originServerURL.Host
		request.URL.Scheme = originServerURL.Scheme
		request.RequestURI = ""

		// set "X-Forwarded-For"-Header to retain remote address
		remoteHostAddr, _, _ := net.SplitHostPort(request.RemoteAddr)
		responseWriter.Header().Set("X-Forwarded-For", remoteHostAddr)

		// send request to origin server and save response
		originServerResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(responseWriter, err)
			return
		}

		copyOriginHeaders(originServerResponse, responseWriter)

		// return response to client
		responseWriter.WriteHeader(originServerResponse.StatusCode)
		io.Copy(responseWriter, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8081", reverseProxyHandler))
}

// Copy http-headers from origin server response
func copyOriginHeaders(response *http.Response, rw http.ResponseWriter) {
	for key, values := range response.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}
}
