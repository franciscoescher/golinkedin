package golinkedin

import (
	"io"
	"strings"
)

var EndpointPeopleParts = []string{"https://api.linkedin.com/v2/people/(id:", ")?projection=(vanityName)"}

// TODO: parse the response into a struct
func (c *Client) GetPeople(personID string) (r string, err error) {
	sb := new(strings.Builder)
	sb.WriteString(EndpointPeopleParts[0])
	sb.WriteString(personID)
	sb.WriteString(EndpointPeopleParts[1])
	resp, err := c.httpClient.Get(sb.String())
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
