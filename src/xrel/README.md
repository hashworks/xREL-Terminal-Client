# xREL API Package

[![GoDoc](https://godoc.org/github.com/hashworks/xRELTerminalClient/src/xrel?status.svg)](https://godoc.org/github.com/hashworks/xRELTerminalClient/src/xrel)

A golang package to authorize with and access the complete xrel.to API.

If you use the OAuth authentication make sure to save the Config variable somewhere and set it again on your next run.
Here is an example how to use the OAuth authentication:

```go
xrel.SetOAuthConsumerKeyAndSecret("CONSUMER_KEY", "CONSUMER_SECRET")
requestToken, url, err := xrel.GetOAuthRequestTokenAndUrl()
ok(err)

// get verificationCode from the provided URL

accessToken, err := xrel.GetOAuthAccessToken(requestToken, verificationCode)
ok(err)
xrel.Config.OAuthAccessToken = *accessToken
```