package xrel

import (
	"encoding/json"
	"github.com/hashworks/xRELTerminalClient/src/xrel/types"
	"io/ioutil"
	"net/http"
)

/*
GetAuthdUser returns information about the currently active user.
Requires OAuth authentication.

http://www.xrel.to/wiki/1718/api-user-get-authd-user.html
*/
func GetAuthdUser() (types.User, error) {
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

/*
GetRateLimitStatus shows how many calls the user (if an OAuth session is present)
or the IP address (otherwise) has left before none will be answered.

http://www.xrel.to/wiki/1795/api-user-rate-limit-status.html
*/
func GetRateLimitStatus() (types.RateLimitStatus, error) {
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
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &rateLimitStatus)
			}
		}
	}

	return rateLimitStatus, err
}
