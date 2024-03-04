package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/echo", http.HandlerFunc(requestEcho))

	err := http.ListenAndServe("127.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ORIGIN] received request at %s from %s\n", time.Now(), r.RemoteAddr)
	_, _ = fmt.Fprintf(w, "Response from origin server for remote request from: %s\n", r.RemoteAddr)
}

func requestEcho(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ORIGIN] received request at: %s\n", time.Now())
	_, _ = fmt.Fprintf(w, "Response from origin server for remote request from: %s\n", r.RemoteAddr)
	_, _ = fmt.Fprintf(w, "[Request-Protocol]\n%s\n", r.Proto)
	_, _ = fmt.Fprintf(w, "[Request-Headers]\n%s\n", headerToString(&r.Header))
}

func headerToString(h *http.Header) string {
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
