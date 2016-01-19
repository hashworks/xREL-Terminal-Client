package main

import (
	"./xREL"
	"fmt"
	"time"
)

func authenticate() {
	authenticated := false

	if xREL.Config.OAuthAccessToken.Token != "" && xREL.Config.OAuthAccessToken.Secret != "" {
		data, err := xREL.GetAuthdUser()
		if err == nil {
			fmt.Println("You're already authenticated, " + data.Name + ".")
			authenticated = true
		}
	}
	if !authenticated {
		requestToken, url, err := xREL.GetOAuthRequestTokenAndUrl()
		ok(err, "Failed to authenticate using oAuth:\n")
		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Print("(3) Enter that verification code here: ")
		verificationCode := ""
		fmt.Scanln(&verificationCode)
		accessToken, err := xREL.GetOAuthAccessToken(requestToken, verificationCode)
		ok(err, "Failed to authenticate using oAuth:\n")
		xREL.Config.OAuthAccessToken = *accessToken
		data, err := xREL.GetAuthdUser()
		if err == nil {
			fmt.Println("Authentication sucessfull, " + data.Name + ".")
		} else {
			fmt.Println("Authentication sucessfull, but we failed to test it:\n" + err.Error())
		}
	}
}

func checkRateLimit() {
	data, err := xREL.GetRateLimitStatus()
	ok(err, "Failed to check rate limit:\n")
	fmt.Printf("You have %d calls remaining, they will reset in %d seconds.",
		data.RemainingCalls, data.GetResetTime().Unix()-time.Now().Unix())
}
