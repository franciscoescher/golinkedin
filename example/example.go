package main

import (
	"context"
	"encoding/json"
	"fmt"
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

func serve(c *golinkedin.Client, port string, route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		// reads query params
		code := r.URL.Query().Get("code")

		err := c.Authorize(context.Background(), code)
		if err != nil {
			log.Fatal(err)
		}

		_, _ = w.Write([]byte("Response for GetProfile:\n\n"))

		profile, err := c.GetProfile()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(profile)
			encodeJson(w, profile)
		}

		fmt.Println("========================================")
		_, _ = w.Write([]byte("\n\n\n\n========================================\nResponse for GetPrimaryContact:\n\n"))

		contact, err := c.GetPrimaryContact()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(contact)
			encodeJson(w, contact)
		}

		fmt.Println("========================================")
		_, _ = w.Write([]byte("\n\n\n\n========================================\nResponse for GetEmailAddress:\n\n"))

		email, err := c.GetEmailAddress()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(email)
			encodeJson(w, email)
		}

		fmt.Println("========================================")
		_, _ = w.Write([]byte("\n\n\n\n========================================\nResponse for GetPeople:\n\n"))

		people, err := c.GetPeople(profile.ID)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(people)
			encodeJson(w, people)
		}
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		authUrl := c.GetAuthURL("state")
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

	c := golinkedin.NewClient(clientId, clientSecret, []string{"r_liteprofile", "r_emailaddress"}, fullRedirectURL)

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
