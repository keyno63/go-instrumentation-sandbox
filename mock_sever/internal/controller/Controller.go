package controller

import (
	"encoding/json"
	"net/http"

	"tokyo.keyno63.go-instrumentation-sandbox.mock/internal/service"
)

// Controller struct
type Controller struct {
	service service.Service
}

// NewController creates a new Controller
func NewController(c service.Service) *Controller {
	return &Controller{c}
}

// HandleRequest handles incoming HTTP requests
func (c *Controller) HandleRequest(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
