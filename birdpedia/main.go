package main

import (
	"fmt"
	"net/http"
)

func main() {
	// We map the 'handler' function for every call onto '/'
	http.HandleFunc("/", handler)

	// As the serving is defined, we start listening on port 12111
	http.ListenAndServe(":12111", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pipeing Hello, World into the response (w <- "Hello, World!")
	fmt.Fprintf(w, "Hello, World!")
}

