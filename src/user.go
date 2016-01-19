package main

import (
	"fmt"
	"github.com/hashworks/xRELTerminalClient/src/xrel"
	"time"
)

func authenticate() {
	authenticated := false

	if xrel.Config.OAuthAccessToken.Token != "" && xrel.Config.OAuthAccessToken.Secret != "" {
		data, err := xrel.GetAuthdUser()
		if err == nil {
			fmt.Println("You're already authenticated, " + data.Name + ".")
			authenticated = true
		}
	}
	if !authenticated {
		requestToken, url, err := xrel.GetOAuthRequestTokenAndUrl()
		ok(err, "Failed to authenticate using oAuth:\n")
		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Print("(3) Enter that verification code here: ")
		verificationCode := ""
		fmt.Scanln(&verificationCode)
		accessToken, err := xrel.GetOAuthAccessToken(requestToken, verificationCode)
		ok(err, "Failed to authenticate using oAuth:\n")
		xrel.Config.OAuthAccessToken = *accessToken
		data, err := xrel.GetAuthdUser()
		if err == nil {
			fmt.Println("Authentication sucessfull, " + data.Name + ".")
		} else {
			fmt.Println("Authentication sucessfull, but we failed to test it:\n" + err.Error())
		}
	}
}

func checkRateLimit() {
	data, err := xrel.GetRateLimitStatus()
	ok(err, "Failed to check rate limit:\n")
	fmt.Printf("You have %d calls remaining, they will reset in %d seconds.\n",
		data.RemainingCalls, data.GetResetTime().Unix()-time.Now().Unix())
}
