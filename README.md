# Go LinkedIn

This is a Go client library for the LinkedIn API.

It authenticates with Oauth2.

## Installation

    go get github.com/franciscoescher/go-linkedin

## Usage

```
package main

import (
	"context"
  "fmt"

  "github.com/franciscoescher/go-linkedin"
)

func main() {
  // first, you need to create a client with your api key credentials
  scopes := []string{"r_liteprofile", "r_emailaddress"}
  client := linkedin.NewClient("CLIENT_ID", "CLIENT_SECRET", scopes, "REDIRECT_URL")

  // then, you need to get the auth url to redirect the user to the linkedin login page
  url := client.GetAuthURL("state")
  fmt.Println(url)

  // after user accepts, linkedin will redirect to REDIRECT_URL with a code parameter
  // you need to use this code to authenticate the client
  code := "CODE"
  err := client.Authenticate(context.Background(), code)
  if err != nil {
    log.Fatal(err)
  }

  // now you can use the client to make requests
  profile, err := c.GetProfile()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(profile)
}
```

Please, check the example folder for a working example that implements a server that handles the linkedin callback.
