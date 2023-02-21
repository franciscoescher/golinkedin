package golinkedin

import (
	"io"
	"net/http"
	"strings"
)

var EndpointPeopleListParts = []string{"https://api.linkedin.com/v2/people?ids=List(", ")"}

// PeopleRequest calls people api with multiple person ids.
func (c *Client) PeopleListRequest(personIDs []string) (resp *http.Response, err error) {
	// (id:{Person ID1}),(id:{Person ID2}),(id:{Person ID3})
	sb := new(strings.Builder)
	sb.WriteString(EndpointPeopleListParts[0])
	for i, id := range personIDs {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString("(id:")
		sb.WriteString(id)
		sb.WriteString(")")
	}
	sb.WriteString(EndpointPeopleListParts[1])
	return c.Get(sb.String())
}

// Same as PeopleListRequest but parses the response.
// This API will only return data for members who haven't limited their Off-LinkedIn Visibility
// TODO: parse the response into a struct
func (c *Client) GetPeopleList(personIDs []string) (r string, err error) {
	resp, err := c.PeopleListRequest(personIDs)
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
