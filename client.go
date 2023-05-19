package golinkedin

import (
	"net/http"
)

type ClientInterface interface {
	Headers() map[string]string
	SetHeader(key string, value string)
	UnsetHeader(key string)
	Get(url string) (resp *http.Response, err error)
	PrimaryContactRequest() (resp *http.Response, err error)
	GetPrimaryContact() (r PrimaryContact, err error)
	GetEmailAddress() (r EmailAddress, err error)
	EmailAddressRequest() (resp *http.Response, err error)
	PeopleRequest(personID string) (resp *http.Response, err error)
	GetPeople(personID string) (r string, err error)
	PeopleListRequest(personIDs []string) (resp *http.Response, err error)
	GetPeopleList(personIDs []string) (r string, err error)
	ProfileRequest() (resp *http.Response, err error)
	GetProfile() (r Profile, err error)
}

type client struct {
	httpClient http.Client
	headers    map[string]string
}

// Common struct for error response in linkedin API
type ErrorResponse struct {
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Message          string `json:"message"`
	Status           int    `json:"status"`
}

func NewClient(httpClient http.Client, initialHeaders map[string]string) ClientInterface {
	return &client{
		httpClient: httpClient,
		headers:    initialHeaders,
	}
}

// Get overwrites get method in http.Client
func (c *client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for h, v := range c.headers {
		req.Header.Add(h, v)
	}
	return c.httpClient.Do(req)
}

// Headers returns the current headers
func (c *client) Headers() map[string]string {
	return c.headers
}

// SetHeader sets a header
func (c *client) SetHeader(key, value string) {
	c.headers[key] = value
}

// UnsetHeader unsets a header
func (c *client) UnsetHeader(key string) {
	delete(c.headers, key)
}
