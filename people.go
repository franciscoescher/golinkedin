package golinkedin

import (
	"io"
	"net/http"
	"strings"
)

// EndpointPeopleParts are the parts to build the endpoint for people api.
var EndpointPeopleParts = []string{"https://api.linkedin.com/rest/people/(id:", ")"}

// PeopleRequest calls people api.
func (c *client) PeopleRequest(personID string) (resp *http.Response, err error) {
	sb := new(strings.Builder)
	sb.WriteString(EndpointPeopleParts[0])
	sb.WriteString(personID)
	sb.WriteString(EndpointPeopleParts[1])
	return c.Get(sb.String())
}

// Same as PeopleRequest but parses the response.
// This API will only return data for members who haven't limited their Off-LinkedIn Visibility
// TODO: parse the response into a struct
func (c *client) GetPeople(personID string) (r string, err error) {
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
