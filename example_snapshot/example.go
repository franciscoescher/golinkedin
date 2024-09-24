package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/franciscoescher/golinkedin"
)

func encodeJson(w http.ResponseWriter, data interface{}) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	_, _ = w.Write(val)
}

func serve(b golinkedin.BuilderInterface, port string, route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		// reads query params
		errparam := r.URL.Query().Get("error")
		if errparam != "" {
			desc := html.UnescapeString(r.URL.Query().Get("error_description"))
			msg := fmt.Sprintf("Error: %s\nError Description: %s", errparam, desc)
			log.Fatal(msg)
		}

		code := r.URL.Query().Get("code")

		c, err := b.GetClient(context.Background(), code)
		if err != nil {
			log.Fatal(err)
		}

		_, _ = w.Write([]byte("Response for Domain Profile:\n\n"))

		snapshot1, err := c.GetSnapshot(golinkedin.SnapshotRequest{Domain: "PROFILE"})
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(snapshot1)
			encodeJson(w, snapshot1)
		}

	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		authUrl := b.GetAuthURL("state")
		_, _ = w.Write([]byte(authUrl))
	})

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// To run this test, you need to setup a callback endpoint and place it at LINKEDIN_REDIRECT_HOST env.
// This endpoint should receive the callback from LinkedIn and receive the access code,
// which you can then use to call InitHTTPClientWithUserToken.
func main() {
	clientId := os.Getenv("LINKEDIN_CLIENT_ID")
	clientSecret := os.Getenv("LINKEDIN_CLIENT_SECRET")
	redirect_host := os.Getenv("LINKEDIN_REDIRECT_HOST")
	redirect_callback_url := os.Getenv("LINKEDIN_REDIRECT_ROUTE")
	redirect_port := os.Getenv("LINKEDIN_REDIRECT_PORT")

	fullRedirectURL := fmt.Sprintf("%s:%s%s", redirect_host, redirect_port, redirect_callback_url)
	fmt.Println("fullRedirectURL", fullRedirectURL)

	// add r_basicprofile to enable GetProfile (if you have the permission)
	c := golinkedin.NewBuilder(golinkedin.NewBuilderParams{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"r_dma_portability_3rd_party"},
		RedirectURL:  fullRedirectURL,
	})

	// to run this test, you need to uncomment the following line and
	// visit the URL in your browser. You will be redirected to the
	// redirect URL with a code. Copy the code and set it as an
	// environment variable called LINKEDIN_CODE
	//
	authUrl := c.GetAuthURL("state")
	fmt.Println("========================================")
	fmt.Println("Visit the following URL in your browser:")
	fmt.Println(authUrl)
	fmt.Println("========================================")

	serve(c, redirect_port, redirect_callback_url)
}
