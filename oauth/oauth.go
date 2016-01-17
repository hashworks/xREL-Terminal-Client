package oauth

import (
	"github.com/mrjones/oauth"
	"fmt"
	"net/http"
)

// Set these with  -ldflags "-X github.com/hashworks/xRELTerminalClient/oauth.CONSUMER_KEY=foo [...]"
var CONSUMER_KEY	string
var CONSUMER_SECRET	string

var consumer *oauth.Consumer = nil

func GetConsumer() *oauth.Consumer {
	if consumer == nil {
		consumer = oauth.NewConsumer(
			CONSUMER_KEY,
			CONSUMER_SECRET,
			oauth.ServiceProvider{
				RequestTokenUrl:   "http://api.xrel.to/api/oauth/temp_token",
				AuthorizeTokenUrl: "http://api.xrel.to/api/oauth/authorize",
				AccessTokenUrl:    "http://api.xrel.to/api/oauth/access_token",
			})
		consumer.Debug(false)
	}
	return consumer
}

func GetNewAuthorizeToken() (*oauth.AccessToken, error) {
	fmt.Println("Authenticating with xREL.to using oAuth...")

	requestToken, u, err := GetConsumer().GetRequestTokenAndUrl("oob")
	if err != nil {
		return &oauth.AccessToken{}, err
	}

	fmt.Println("(1) Go to: " + u)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Print("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := consumer.AuthorizeToken(requestToken, verificationCode)
	return accessToken, err
}

func GetOAuthClient(accessToken oauth.AccessToken) (*http.Client, error) {
	client, err := GetConsumer().MakeHttpClient(&accessToken)
	return client, err;
}