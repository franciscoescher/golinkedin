package golinkedin

import (
	"encoding/json"
)

const EndpointEmailAddress = "https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))"

// GetEmailAddress gets user email address. Please note that email address is only available with the scope r_emailaddress.
func (c *Client) GetEmailAddress() (r EmailAddress, err error) {
	resp, err := c.httpClient.Get(EndpointEmailAddress)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}

type EmailAddress struct {
	ErrorResponse
	Elements []struct {
		Handle      string `json:"handle"`
		HandleTilde struct {
			EmailAddress string `json:"emailAddress"`
		} `json:"handle~"`
	} `json:"elements"`
}
