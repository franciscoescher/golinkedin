package golinkedin

import (
	"net/http"
)

type Client struct {
	http.Client
	Headers map[string]string
}

// Common struct for error response in linkedin API
type ErrorResponse struct {
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Message          string `json:"message"`
	Status           int    `json:"status"`
}

// Get overwrites get method in http.Client
func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for h, v := range c.Headers {
		req.Header.Add(h, v)
	}
	return c.Do(req)
}
