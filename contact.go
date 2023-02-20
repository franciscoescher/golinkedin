package golinkedin

import (
	"encoding/json"
)

const EndpointPrimaryContact = "https://api.linkedin.com/v2/clientAwareMemberHandles?q=members&projection=(elements*(primary,type,handle~))"

// GetPrimaryContact gets user primary contact. Please note that primary contact is only available with the scope r_liteprofile.
func (c *Client) GetPrimaryContact() (r PrimaryContact, err error) {
	resp, err := c.httpClient.Get(EndpointPrimaryContact)
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
		} `json:"handle~"`
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
	} `json:"elements"`
}
