# Go LinkedIn

This is a Go client library for the LinkedIn API.

## Installation

    go get github.com/franciscoescher/golinkedin

## Usage

Fiest, you need to create a client builder with your api key credentials using `golinkedin.NewBuilder`.

Then, you need to get the auth url to redirect the user to the linkedin login page using `builder.GetAuthURL`.

The user needs to access the url, login and accept the permissions. After that, linkedin will redirect to the redirect url with a code parameter.

Finally, You need to use this code to authenticate and get the client using `builder.GetClient`.

This client will be able to call the linkedin api.

```
package main

import (
	"context"
  "fmt"

  "github.com/franciscoescher/golinkedin"
)

func main() {
  // first, you need to create a builder with your api key credentials
  scopes := []string{"r_liteprofile", "r_emailaddress"}
  builder := golinkedin.NewBuilder("CLIENT_ID", "CLIENT_SECRET", scopes, "REDIRECT_URL")

  // then, you need to get the auth url to redirect the user to the linkedin login page
  url := builder.GetAuthURL("state")
  fmt.Println(url)

  // after user accepts, linkedin will redirect to REDIRECT_URL with a code parameter
  // you need to use this code to authenticate and get the client
  code := "CODE"
  client, err := builder.GetClient(context.Background(), code)
  if err != nil {
    log.Fatal(err)
  }

  // now you can use the client to make requests
  profile, err := client.GetProfile()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(profile)
}
```

Please, check the example folder for a working example that implements a server that handles the linkedin callback.
