package handler

import (
	"fmt"
	"net/http"
)

type SampleControllerInterface interface {
}

// SampleHandler is a simple HTTP handler function
func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a sample handler!")
}
