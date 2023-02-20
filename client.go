package golinkedin

import (
	"net/http"
)

type Client struct {
	http.Client
}

// Common struct for error response in linkedin API
type ErrorResponse struct {
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Message          string `json:"message"`
	Status           int    `json:"status"`
}
