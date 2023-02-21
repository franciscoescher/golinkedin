package golinkedin

import (
	"encoding/json"
	"net/http"
)

const EndpointPrimaryContact = "https://api.linkedin.com/v2/clientAwareMemberHandles?q=members&projection=(elements*(primary,type,handle~))"

// PrimaryContactRequest calls primaryContact api.
// Please note that primary contact is only available with the scope r_liteprofile.
func (c *Client) PrimaryContactRequest() (resp *http.Response, err error) {
	return c.Get(EndpointPrimaryContact)
}

// Same as PrimaryContactRequest but parses the response.
func (c *Client) GetPrimaryContact() (r PrimaryContact, err error) {
	resp, err := c.PrimaryContactRequest()
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}

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
