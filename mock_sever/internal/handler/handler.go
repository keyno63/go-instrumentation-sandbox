package handler

import (
	"encoding/json"
	"net/http"
)

// Response is a struct for the mock response
type Response struct {
	Message string `json:"message"`
}

// MockHandler handles the mock server requests
func MockHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, this is a mock server!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
