package xREL

import (
	"net/http"
	"errors"
	"strconv"
	"./types"
	"github.com/mrjones/oauth"
)

const apiURL = "http://api.xrel.to/api/"

var Config = struct {
	OAuthAccessToken		oauth.AccessToken

	// 24h caching http://www.xrel.to/wiki/6318/api-release-categories.html
	LastCategoryRequest		int64
	Categories				[]types.Category

	// 24h caching http://www.xrel.to/wiki/2996/api-release-filters.html
	LastFilterRequest		int64
	Filters					[]types.Filter

	// 24h caching http://www.xrel.to/wiki/3698/api-p2p-categories.html
	LastP2PCategoryRequest	int64
	P2PCategories			[]types.P2PCategory
}{}

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
	var client	*http.Client
	var err		error

	if err == nil && Config.OAuthAccessToken.Token != "" && Config.OAuthAccessToken.Secret != "" {
		client, err = GetOAuthClient(Config.OAuthAccessToken)
	} else {
		err = errors.New("You're not authenticated.")
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