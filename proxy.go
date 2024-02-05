package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	go startOriginServer() //start origin server in a new goroutine for testing purpose at localhost:8080
	startReverseProxy()
}

func startOriginServer() {
	originServerHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[ORIGIN] received request at: %s\n", time.Now())
		_, _ = fmt.Fprint(rw, "Response from origin server\n")
	})

	log.Fatal(http.ListenAndServe(":8080", originServerHandler))
}

func startReverseProxy() {
	originServerURL, err := url.Parse("http://jsonplaceholder.typicode.com")
	if err != nil {
		log.Fatalf("invalid origin server URL: %s", err)
	}

	reverseProxyHandler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Printf("[PROXY] received request at: %s\n", time.Now())
		// fmt.Println("[PROXY] request content:", request)

		// fmt.Println("Host:", request.Host)
		// fmt.Println("URLHost:", request.URL.Host)
		// fmt.Println("URLScheme:", request.URL.Scheme)
		// fmt.Println("ReqURI:", request.RequestURI)

		//set request data to forward to origin server
		request.Host = originServerURL.Host
		request.URL.Host = originServerURL.Host
		request.URL.Scheme = originServerURL.Scheme
		request.RequestURI = ""

		// fmt.Println("Host:", request.Host)
		// fmt.Println("URLHost:", request.URL.Host)
		// fmt.Println("URLScheme:", request.URL.Scheme)
		// fmt.Println("ReqURI:", request.RequestURI)

		//save response from origin server
		originServerResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(responseWriter, err)
			return
		}

		//return response to client
		responseWriter.WriteHeader(http.StatusOK)
		io.Copy(responseWriter, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8081", reverseProxyHandler))
}
