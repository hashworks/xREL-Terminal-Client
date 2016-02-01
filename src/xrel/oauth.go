package xrel

import (
	"github.com/mrjones/oauth"
	"net/http"
)

var (
	consumerKey    string
	consumerSecret string
	consumer       *oauth.Consumer
)

func getConsumer() *oauth.Consumer {
	if consumer == nil {
		consumer = oauth.NewConsumer(
			consumerKey,
			consumerSecret,
			oauth.ServiceProvider{
				RequestTokenUrl:   "http://api.xrel.to/api/oauth/temp_token",
				AuthorizeTokenUrl: "http://api.xrel.to/api/oauth/authorize",
				AccessTokenUrl:    "http://api.xrel.to/api/oauth/access_token",
			})
		consumer.Debug(false)
	}
	return consumer
}

/*
SetOAuthConsumerKeyAndSecret sets the OAuth consumer key and secret you received from xREL.
Get them here: http://www.xrel.to/api-apps.html
*/
func SetOAuthConsumerKeyAndSecret(key, secret string) {
	consumerKey = key
	consumerSecret = secret
}

/*
GetOAuthRequestTokenAndURL returns a new request token and an URL where the user can login and get a verification code from.
*/
func GetOAuthRequestTokenAndURL() (*oauth.RequestToken, string, error) {
	return getConsumer().GetRequestTokenAndUrl("oob")
}

/*
GetOAuthAccessToken returns the access token to use in authenticated requests using
the request token and verification code from GetOAuthRequestTokenAndUrl().
Save this in xREL.Config.OAuthAccessToken.
*/
func GetOAuthAccessToken(requestToken *oauth.RequestToken, verificationCode string) (*oauth.AccessToken, error) {
	return getConsumer().AuthorizeToken(requestToken, verificationCode)
}

func makeOAuthClient(accessToken oauth.AccessToken) (*http.Client, error) {
	return getConsumer().MakeHttpClient(&accessToken)
}