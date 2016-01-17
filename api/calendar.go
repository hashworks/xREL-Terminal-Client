package api

import (
	"io/ioutil"
	"encoding/json"
	"github.com/hashworks/xRELTerminalClient/api/types"
)

/**
	Returns a list upcoming movies and their releases.

	http://www.xrel.to/wiki/1827/api-calendar-upcoming.html
 */
func Calendar_GetUpcomingTitles() ([]types.UpcomingTitle, error) {
	var upcomingTitles []types.UpcomingTitle

	client := getClient()
	response, err := client.Get(apiURL + "calendar/upcoming.json")
	defer response.Body.Close()
	if err == nil {
		err = checkResponseStatusCode(response.StatusCode)
		if err == nil {
			var bytes []byte
			bytes, err = ioutil.ReadAll(response.Body)
			if err == nil {
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &upcomingTitles)
			}
		}
	}

	return upcomingTitles, err
}