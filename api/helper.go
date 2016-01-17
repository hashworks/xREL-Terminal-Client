package api

import (
	"net/http"
	"errors"
	"strconv"
	"github.com/hashworks/xRELTerminalClient/oauth"
	"github.com/hashworks/xRELTerminalClient/configHandler"
)

const apiURL = "http://api.xrel.to/api/"

/**
	xREL JSON responses are surrounded /*-secure-\n{"payload":\n and their closings.
	The following removes this. Follow the xREL API changelog,
	we might need to remove this partly in future releases.
 */
func stripeJSON(json []byte) []byte {
	return json[22:len(json) - 4]
}

/**
	Returns an OAuth client if authenticated and a normal client otherwise.
 */
func getClient() *http.Client {
	client, err := getOAuthClient()
	if (err != nil) {
		client = http.DefaultClient
	}
	return client
}

/**
	Returns an OAuth client
 */
func getOAuthClient() (*http.Client, error) {
	var client *http.Client

	config, err := configHandler.GetConfig("")
	if err == nil && config.OAuthAccessToken.Token != "" && config.OAuthAccessToken.Secret != "" {
		client, err = oauth.GetOAuthClient(config.OAuthAccessToken)
	} else {
		err = errors.New("You're not authenticated, please do so by executing with --authenticate.")
	}

	return client, err
}

func checkResponseStatusCode(statusCode int) error {
	var err error

	switch statusCode {
	case 200:
		return err
	case 429:
		err = errors.New("Rate limit reached (http://www.xrel.to/wiki/2727/api-rate-limiting.html). Please try again later.")
	case 404:
		err = errors.New("Not found.")
	// TODO: Find out what happens if we send wrong or expired oAuth data
	default:
		err = errors.New("xREL returned unexpected HTTP status code " + strconv.Itoa(statusCode) + ".")
	}

	return err
}

func generateGetParametersString(parameters map[string]string) string {
	var query string

	for k, v := range parameters {
		if query == "" {
			query = "?"
		} else {
			query += "&"
		}
		query += k + "=" + v
	}

	return query
}