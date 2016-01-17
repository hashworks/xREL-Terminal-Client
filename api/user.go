package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/hashworks/xRELTerminalClient/api/types"
)

/**
	Returns information about the currently active user.
	Requires oAuth authentication.

	http://www.xrel.to/wiki/1718/api-user-get-authd-user.html
 */
func User_GetAuthdUser() (types.User, error) {
	var user types.User

	client, err := getOAuthClient()
	if err == nil {
		var response *http.Response
		response, err = client.Get(apiURL + "user/get_authd_user.json")
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &user)
				}
			}
		}
	}

	return user, err
}

/**
	Shows how many calls the user (if an OAuth session is present)
	or the IP address (otherwise) has left before none will be answered.

	http://www.xrel.to/wiki/1795/api-user-rate-limit-status.html
 */
func User_RateLimitStatus() (types.RateLimitStatus, error) {
	var rateLimitStatus types.RateLimitStatus

	client := getClient()
	response, err := client.Get(apiURL + "user/rate_limit_status.json")
	defer response.Body.Close()
	if err == nil {
		err = checkResponseStatusCode(response.StatusCode)
		if err == nil {
			var bytes []byte
			bytes, err = ioutil.ReadAll(response.Body)
			if err == nil {
				var rateLimitStatus types.RateLimitStatus
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &rateLimitStatus)
			}
		}
	}

	return rateLimitStatus, err
}