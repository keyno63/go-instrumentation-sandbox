package service

import (
	"fmt"
	"net/http"
)

// Service represents the service structure
type Service struct {
	Port string
}

// NewService creates a new service
func NewService(port string) *Service {
	return &Service{Port: port}
}

// Start starts the service
func (s *Service) Start() {
	http.HandleFunc("/", s.handleRequest)
	fmt.Printf("Starting server on port %s\n", s.Port)
	if err := http.ListenAndServe(s.Port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// handleRequest handles incoming HTTP requests
func (s *Service) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
