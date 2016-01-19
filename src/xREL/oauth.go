package xREL

import (
	"github.com/mrjones/oauth"
	"net/http"
)

var consumerKey		string
var consumerSecret	string

var consumer *oauth.Consumer = nil

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

func SetOAuthConsumerKeyAndSecret(key, secret string) {
	consumerKey		= key
	consumerSecret	= secret
}

func GetOAuthRequestTokenAndUrl() (*oauth.RequestToken, string, error) {
	return getConsumer().GetRequestTokenAndUrl("oob")
}

func GetOAuthAccessToken(requestToken *oauth.RequestToken, verificationCode string) (*oauth.AccessToken, error) {
	return getConsumer().AuthorizeToken(requestToken, verificationCode)
}

func GetOAuthClient(accessToken oauth.AccessToken) (*http.Client, error) {
	return getConsumer().MakeHttpClient(&accessToken);
}