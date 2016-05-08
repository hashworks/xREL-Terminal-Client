package main

import (
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"time"
)

func authenticate() {
	authenticated := false

	data, err := xrel.GetUserInfo()
	if err == nil {
		if data.Name == "" {
			fmt.Println("Received no error, but failed to test authentication.")
		} else {
			fmt.Println("You're already authenticated, " + data.Name + ".")
			authenticated = true
		}
	}
	if !authenticated {
		url := xrel.GetOAuth2AuthURL("")
		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Print("(3) Enter that verification code here: ")
		verificationCode := ""
		fmt.Scanln(&verificationCode)
		err := xrel.PerformOAuth2UserAuthentication(verificationCode)
		ok(err, "Failed to authenticate using oAuth2: ")
		data, err := xrel.GetUserInfo()
		ok(err, "Failed to test authentication: ")
		if data.Name == "" {
			fmt.Println("Received no error, but failed test authentication.")
		} else {
			fmt.Println("Authentication sucessfull, " + data.Name + ".")
		}
	}
}

func checkRateLimit() {
	if types.Config.RateLimitResetUnix == 0 {
		fmt.Println("You need to make a request first before we can display you any rate limit information.")
	} else {
		fmt.Printf("You have %d/%d calls remaining, they will reset in %d seconds.\n",
			types.Config.RateLimitRemaining, types.Config.RateLimitMax, types.Config.RateLimitResetUnix-time.Now().Unix())
	}
}
