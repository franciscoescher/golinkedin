package golinkedin

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

var (
	oauth2endpoint = oauth2.Endpoint{
		AuthURL:   "https://www.linkedin.com/oauth/v2/authorization",
		TokenURL:  "https://www.linkedin.com/oauth/v2/accessToken",
		AuthStyle: oauth2.AuthStyleInParams,
	}
)

// NewClient returns a new linkedin client, not yet authenticated.
func NewClient(clientID string, clientSecret string, scopes []string, redirectURL string) *Client {
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		scopes:       scopes,
		redirectURL:  redirectURL,
	}
}

type Client struct {
	// ClientID is the api key client's ID.
	clientID string
	// ClientSecret is the api key client's secret.
	clientSecret string
	// Scopes is the list of scopes that the client will request.
	scopes []string
	// httpClient is the http client that will be used to make requests.
	// it must be authenticated with oauth2 before making requests
	// using the Authorize method.
	httpClient *http.Client
	// redirectURL is the URL that the user will be redirected to after
	// authenticating with linkedin in the GetAuthURL url response.
	redirectURL string
}

// getOAuth2Config returns the oauth2 config
func (c *Client) getOAuth2Config() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Scopes:       c.scopes,
		Endpoint:     oauth2endpoint,
		RedirectURL:  c.redirectURL,
	}
	return conf
}

// GetAuthURL returns the URL to the linkedin login page
// The state is a string that will be returned to the redirect URL
// so it can be used to prevent CSRF attacks
func (c *Client) GetAuthURL(state string) string {
	oa2config := c.getOAuth2Config()
	url := oa2config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url
}

// Authorize will exchange the code for an access token and
// use it to create a new client with an authorized http client
func (c *Client) Authorize(ctx context.Context, code string) (*Client, error) {
	oa2config := c.getOAuth2Config()
	fmt.Println(oa2config)
	fmt.Println(code)
	token, err := oa2config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	new := &Client{
		clientID:     c.clientID,
		clientSecret: c.clientSecret,
		scopes:       c.scopes,
		redirectURL:  c.redirectURL,
		httpClient:   oa2config.Client(ctx, token),
	}
	return new, nil
}

// Common struct for error response in linkedin API
type ErrorResponse struct {
	ServiceErrorCode int    `json:"serviceErrorCode"`
	Message          string `json:"message"`
	Status           int    `json:"status"`
}
