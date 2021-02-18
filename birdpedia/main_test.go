package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Here, we form a new HTTP request. This is the request that's going to be passed to our handler.
	// The first argument is the method, the second argument is the route (which we leave blank for now,
	// and will get back to soon), and the third is the request body, which we don't have in this case.
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Go's httptest library to create a recorder
	// This recorder will act as the target of our http request
	// This is somewhat of a minibrowser, which will accept the result of the http request 
	// that we make
	recorder := httptest.NewRecorder()

	// This is the function we defined as the handler in the main.go file
	hf := http.HandlerFunc(handler)

	// Pass the request for the recorder
	hf.ServeHTTP(recorder, req)

	// Check the http status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expected
	expected := `Hello, World!`
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
