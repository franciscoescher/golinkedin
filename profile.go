package golinkedin

import (
	"encoding/json"
	"net/http"
)

// EndpointProfile is the endpoint for profile api.
const EndpointProfile = "https://api.linkedin.com/v2/me?projection=(id,firstName,lastName,vanityName,localizedHeadline,localizedFirstName,localizedLastName,localizedHeadline,headline,profilePicture(displayImage~:playableStreams))"

// Profile is the response from profile api.
type Profile struct {
	ErrorResponse
	ID                 string `json:"id"`
	LocalizedFirstName string `json:"localizedFirstName"`
	LocalizedLastName  string `json:"localizedLastName"`
	FirstName          struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"firstName"`
	LastName struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"lastName"`
	ProfilePicture struct {
		DisplayImage     string `json:"displayImage"`
		DisplayImageFull struct {
			Paging struct {
				Count int   `json:"count"`
				Start int   `json:"start"`
				Links []any `json:"links"`
			} `json:"paging"`
			Elements []struct {
				Artifact            string `json:"artifact"`
				AuthorizationMethod string `json:"authorizationMethod"`
				Data                struct {
					ComLinkedinDigitalmediaMediaartifactStillImage struct {
						MediaType    string `json:"mediaType"`
						RawCodecSpec struct {
							Name string `json:"name"`
							Type string `json:"type"`
						} `json:"rawCodecSpec"`
						DisplaySize struct {
							Width  float64 `json:"width"`
							Uom    string  `json:"uom"`
							Height float64 `json:"height"`
						} `json:"displaySize"`
						StorageSize struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"storageSize"`
						StorageAspectRatio struct {
							WidthAspect  float64 `json:"widthAspect"`
							HeightAspect float64 `json:"heightAspect"`
							Formatted    string  `json:"formatted"`
						} `json:"storageAspectRatio"`
						DisplayAspectRatio struct {
							WidthAspect  float64 `json:"widthAspect"`
							HeightAspect float64 `json:"heightAspect"`
							Formatted    string  `json:"formatted"`
						} `json:"displayAspectRatio"`
					} `json:"com.linkedin.digitalmedia.mediaartifact.StillImage"`
				} `json:"data"`
				Identifiers []struct {
					Identifier                 string `json:"identifier"`
					Index                      int    `json:"index"`
					MediaType                  string `json:"mediaType"`
					File                       string `json:"file"`
					IdentifierType             string `json:"identifierType"`
					IdentifierExpiresInSeconds int    `json:"identifierExpiresInSeconds"`
				} `json:"identifiers"`
			} `json:"elements"`
		} `json:"displayImage~"`
	} `json:"profilePicture"`
	Headline struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"headline"`
	LocalizedHeadline string `json:"localizedHeadline"`
	VanityName        string `json:"vanityName"`
}

// ProfileRequest calls profile api.
// Please note that this is only available with the scope r_liteprofile.
// Also, vanity name and headline are only available with the scope r_basicprofile.
func (c *client) ProfileRequest() (resp *http.Response, err error) {
	return c.Get(EndpointProfile)
}

// Same as ProfileRequest but parses the response.
func (c *client) GetProfile() (r Profile, err error) {
	resp, err := c.ProfileRequest()
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}
