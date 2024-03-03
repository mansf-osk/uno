package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	serveRequestLogger()
}

func serveRequestLogger() {
	originServerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[ORIGIN] received request at: %s\n", time.Now())
		_, _ = fmt.Fprintf(w, "Response from origin server for remote request from: %s\n", r.RemoteAddr)
		_, _ = fmt.Fprintf(w, "[Request-Protocol]\n%s\n", r.Proto)
		_, _ = fmt.Fprintf(w, "[Request-Headers]\n%s\n", HeaderToString(&r.Header))
	})

	log.Fatal(http.ListenAndServe(":8080", originServerHandler))
}

func HeaderToString(h *http.Header) string {
	headerString := ""
	for header, values := range *h {
		headerValues := ""
		for i, value := range values {
			if i == len(values)-1 {
				headerValues += value
			} else {
				headerValues += fmt.Sprintf("%s, ", value)
			}
		}
		headerString += fmt.Sprintf("%s: %s\n", header, headerValues)
	}
	return headerString
}
