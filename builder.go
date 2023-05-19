package golinkedin

import (
	"context"

	"golang.org/x/oauth2"
)

var (
	oauth2endpoint = oauth2.Endpoint{
		AuthURL:   "https://www.linkedin.com/oauth/v2/authorization",
		TokenURL:  "https://www.linkedin.com/oauth/v2/accessToken",
		AuthStyle: oauth2.AuthStyleInParams,
	}
	defaultHeaders = map[string]string{
		"Linkedin-Version": "202305",
	}
)

// builderInterface is the interface that wraps the basic GetAuthURL and GetClient methods.
type BuilderInterface interface {
	GetAuthURL(state string) string
	GetClient(ctx context.Context, code string) (ClientInterface, error)
}

// Newbuilder returns a new linkedin client, not yet authenticated.
func NewBuilder(clientID string, clientSecret string, scopes []string, redirectURL string) BuilderInterface {
	return &builder{
		clientID:     clientID,
		clientSecret: clientSecret,
		scopes:       scopes,
		redirectURL:  redirectURL,
	}
}

// builder is the implementation of the builderInterface
type builder struct {
	// ClientID is the api key client's ID.
	clientID string
	// ClientSecret is the api key client's secret.
	clientSecret string
	// Scopes is the list of scopes that the client will request.
	scopes []string
	// redirectURL is the URL that the user will be redirected to after
	// authenticating with linkedin in the GetAuthURL url response.
	redirectURL string
}

// getOAuth2Config returns the oauth2 config
func (c *builder) getOAuth2Config() *oauth2.Config {
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
func (c *builder) GetAuthURL(state string) string {
	oa2config := c.getOAuth2Config()
	url := oa2config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url
}

// GetClient will exchange the code for an access token and
// use it to create a new client with an authorized http client
func (c *builder) GetClient(ctx context.Context, code string) (ClientInterface, error) {
	oa2config := c.getOAuth2Config()
	token, err := oa2config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	client := oa2config.Client(ctx, token)
	new := NewClient(*client, defaultHeaders)
	return new, nil
}
