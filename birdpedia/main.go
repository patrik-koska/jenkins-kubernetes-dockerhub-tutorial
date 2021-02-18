package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()

	http.ListenAndServe(":12111", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pipeing Hello, World into the response (w <- "Hello, World!")
	fmt.Fprintf(w, "Hello, World!")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}
