package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	startOriginServer()
}

func startOriginServer() {
	originServerHandler := http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Printf("[ORIGIN] received request at: %s\n", time.Now())
		fmt.Printf("[ORIGIN] request TLS:\n%v\n", request.TLS)
		fmt.Printf("[ORIGIN] request Proto:\n%v\n", request.Proto)
		fmt.Printf("[ORIGIN] request-headers:\n%+v\n", request.Header)
		_, _ = fmt.Fprintf(responseWriter, "Response from origin server for remote request from: %s\n", request.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe(":8080", originServerHandler))
}
