package golinkedin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EndpointSnapshot is the endpoint for email address api.
const EndpointSnapshot = "https://api.linkedin.com/rest/memberSnapshotData"

type SnapshotRequest struct {
	Domain string `json:"domain,omitempty"`
}

// Snapshot is the response from email address api.
type Snapshot struct {
	ErrorResponse
	Paging struct {
		Start int `json:"start"`
		Count int `json:"count"`
		Links []struct {
			Type string `json:"type"`
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Total int `json:"total"`
	} `json:"paging"`
	Elements []struct {
		SnapshotData   []byte `json:"snapshotData"`
		SnapshotDomain string `json:"snapshotDomain"`
	} `json:"elements"`
}

// SnapshotRequest calls Snapshot api.
// Please note that email address is only available with the scope r_Snapshot.
func (c *client) SnapshotRequest(req SnapshotRequest) (resp *http.Response, err error) {
	endpoint := EndpointSnapshot
	if req.Domain != "" {
		endpoint += "?domain=" + req.Domain
	}
	return c.Get(endpoint)
}

// Same as SnapshotRequest but parses the response.
func (c *client) GetSnapshot(req SnapshotRequest) (r Snapshot, err error) {
	resp, err := c.SnapshotRequest(req)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	// print raw response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	fmt.Println(string(body))
	resp.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}
