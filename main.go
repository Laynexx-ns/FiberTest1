package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var logcount int = 0

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logcount++
		log.Printf("| log: %v | %s %s %v ", logcount, r.Method, r.URL.Path, time.Since(start))
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	wrappedMux := loggingMiddleware(mux)

	http.ListenAndServe(":8080", wrappedMux)
}
