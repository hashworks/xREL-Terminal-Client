# xREL API Package

A golang package to authorize with and access the complete xREL.to API.

If you use the OAuth authentication make sure to save the Config variable somewhere and set it again on your next run.
Here is an example how to use the OAuth authentication:

```golang
xREL.SetOAuthConsumerKeyAndSecret("CONSUMER_KEY", "CONSUMER_SECRET")
requestToken, url, err := xREL.GetOAuthRequestTokenAndUrl()
ok(err)

// get verificationCode from the provided URL

accessToken, err := xREL.GetOAuthAccessToken(requestToken, verificationCode)
ok(err)
xREL.Config.OAuthAccessToken = *accessToken
```