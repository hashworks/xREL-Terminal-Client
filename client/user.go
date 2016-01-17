package client

import (
	"fmt"
	"time"
	"github.com/hashworks/xRELTerminalClient/oauth"
	"github.com/hashworks/xRELTerminalClient/configHandler"
	"github.com/hashworks/xRELTerminalClient/api"
)

func Authenticate() {
	authenticated := false
	config, _ := configHandler.GetConfig("")

	if (config.OAuthAccessToken.Token != "" && config.OAuthAccessToken.Secret != "") {
		data, err := api.User_GetAuthdUser()
		if err == nil {
			fmt.Println("You're already authenticated, " + data.Name + ".")
			authenticated = true
		}
	}
	if !authenticated {
		accessToken, err := oauth.GetNewAuthorizeToken()
		OK(err, "Failed to authenticate using oAuth: \n")
		config.OAuthAccessToken = *accessToken
		data, err := api.User_GetAuthdUser()
		if err == nil {
			fmt.Println("Authentication sucessfull, " + data.Name + ".")
		} else {
			fmt.Println("Authentication sucessfull, but we failed to test it:\n" + err.Error())
		}
	}
}

func CheckRateLimit() {
	data, err := api.User_RateLimitStatus()
	OK(err, "Failed to check rate limit:\n")
	fmt.Printf("You have %d calls remaining, they will reset in %d seconds.",
		data.RemainingCalls, data.GetResetTime().Unix() - time.Now().Unix())
}
