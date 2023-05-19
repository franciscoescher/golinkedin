package golinkedin

import (
	"encoding/json"
	"net/http"
)

// EndpointPrimaryContact is the endpoint for primary contact api.
const EndpointPrimaryContact = "https://api.linkedin.com/rest/clientAwareMemberHandles?q=members&projection=(elements*(primary,type,handle~))"

// PrimaryContact is the response from primary contact api.
type PrimaryContact struct {
	ErrorResponse
	Elements []struct {
		Handle      string `json:"handle"`
		HandleTilde struct {
			EmailAddress string `json:"emailAddress"`
			PhoneNumber  struct {
				Number string `json:"number"`
			} `json:"phoneNumber"`
		} `json:"handle~"`
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
	} `json:"elements"`
}

// PrimaryContactRequest calls primaryContact api.
// Please note that primary contact is only available with the scope r_liteprofile.
func (c *client) PrimaryContactRequest() (resp *http.Response, err error) {
	return c.Get(EndpointPrimaryContact)
}

// Same as PrimaryContactRequest but parses the response.
func (c *client) GetPrimaryContact() (r PrimaryContact, err error) {
	resp, err := c.PrimaryContactRequest()
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}
