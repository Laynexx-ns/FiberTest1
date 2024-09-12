package main

import (
	"FiberTest1/DataBase"
	"FiberTest1/jsonEncod"
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

func logsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "there will be logs: \n")
	logmap := DataBase.GetBaseJsonNameText()

	for i, v := range logmap {
		fmt.Fprintf(w, "name : %v | email : %v \n", i, v)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/json", jsonEncod.JsonHandler)
	mux.HandleFunc("/logs", logsHandler)

	wrappedMux := loggingMiddleware(mux)

	http.ListenAndServe(":8080", wrappedMux)
}
