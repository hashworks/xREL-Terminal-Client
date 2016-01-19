/*
Package xrel contains functions to authorize with and access the complete xREL.to API.

Here is an example on how to use the OAuth authentication:

	xREL.SetOAuthConsumerKeyAndSecret("CONSUMER_KEY", "CONSUMER_SECRET")
	requestToken, url, err := xREL.GetOAuthRequestTokenAndUrl()
	ok(err)
	// get verificationCode from the provided URL
	accessToken, err := xREL.GetOAuthAccessToken(requestToken, verificationCode)
	ok(err)
	xREL.Config.OAuthAccessToken = *accessToken

*/
package xrel

import (
	"errors"
	"github.com/hashworks/xRELTerminalClient/src/xrel/types"
	"github.com/mrjones/oauth"
	"net/http"
	"strconv"
)

const apiURL = "http://api.xrel.to/api/"

/*
Config contains the OAuth Token and cached results. Save this somewhere and restore it on every run.
 */
var Config = struct {
	OAuthAccessToken oauth.AccessToken

	// 24h caching http://www.xrel.to/wiki/6318/api-release-categories.html
	LastCategoryRequest int64
	Categories          []types.Category

	// 24h caching http://www.xrel.to/wiki/2996/api-release-filters.html
	LastFilterRequest int64
	Filters           []types.Filter

	// 24h caching http://www.xrel.to/wiki/3698/api-p2p-categories.html
	LastP2PCategoryRequest int64
	P2PCategories          []types.P2PCategory
}{}

/*
stripeJSON removes /*-secure-\n{"payload":\n and their closings surrounding xREL.to JSON responds.
Follow the xREL API changelog, we might need to remove this partly in future releases.
*/
func stripeJSON(json []byte) []byte {
	return json[22 : len(json)-4]
}

/*
getClient returns an OAuth client if authenticated and a normal client otherwise.
*/
func getClient() *http.Client {
	client, err := getOAuthClient()
	if err != nil {
		client = http.DefaultClient
	}
	return client
}

func getOAuthClient() (*http.Client, error) {
	var (
		client *http.Client
		err    error
	)

	if err == nil && Config.OAuthAccessToken.Token != "" && Config.OAuthAccessToken.Secret != "" {
		client, err = makeOAuthClient(Config.OAuthAccessToken)
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
	// TODO: Find out what happens if we send wrong or expired OAuth data
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
