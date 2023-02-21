package golinkedin

import (
	"io"
	"net/http"
	"strings"
)

var EndpointPeopleParts = []string{"https://api.linkedin.com/v2/people/(id:", ")?projection=(vanityName)"}

// PeopleRequest calls people api.
func (c *Client) PeopleRequest(personID string) (resp *http.Response, err error) {
	sb := new(strings.Builder)
	sb.WriteString(EndpointPeopleParts[0])
	sb.WriteString(personID)
	sb.WriteString(EndpointPeopleParts[1])
	return c.Get(sb.String())
}

// Same as PeopleRequest but parses the response.
// TODO: parse the response into a struct
func (c *Client) GetPeople(personID string) (r string, err error) {
	resp, err := c.PeopleRequest(personID)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, nil
	}
	return string(b), nil
}
