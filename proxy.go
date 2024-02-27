package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	go startOriginServer() // start origin server in a new goroutine for testing purpose at localhost:8080
	startReverseProxy()
}

func startOriginServer() {
	originServerHandler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Printf("[ORIGIN] received request at: %s\n", time.Now())
		_, _ = fmt.Fprintf(responseWriter, "Response from origin server for remote request from: %s\n", request.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe(":8080", originServerHandler))
}

func startReverseProxy() {
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

		// copy http-headers from origin server response
		for key, values := range originServerResponse.Header {
			for _, value := range values {
				responseWriter.Header().Set(key, value)
			}
		}

		// return response to client
		responseWriter.WriteHeader(originServerResponse.StatusCode)
		io.Copy(responseWriter, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8081", reverseProxyHandler))
}
