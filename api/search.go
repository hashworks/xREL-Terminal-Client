package api

import (
	"errors"
	"strconv"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/hashworks/xRELTerminalClient/api/types"
)

/**
	Searches for Scene and P2P releases. Please note that additional search rate limiting applies.
	See http://www.xrel.to/wiki/2727/api-rate-limiting.html

	query						Search keyword.	(required)
	includeScene	:= false	If true, Scene releases will be included in the search results.
	includeP2P		:= false	If true, P2P releases will be included in the search results.
	limit			:= 25		Number of returned search results. Maximum and default 25.


	http://www.xrel.to/wiki/6320/api-search-releases.html
 */
func Search_Releases(query string, includeScene, includeP2P bool, limit int) (types.ReleaseSearchResult, error) {
	var searchResult	types.ReleaseSearchResult
	var err 			error

	if query == "" {
		err = errors.New("Please provide a query string.")
	} else {
		parameters := make(map[string]string)
		parameters["q"] = url.QueryEscape(query)
		if (includeScene) {
			parameters["scene"]	= "1"
		} else {
			parameters["scene"]	= "0"
		}
		if (includeP2P) {
			parameters["p2p"]	= "1"
		} else {
			parameters["p2p"]	= "0"
		}
		if limit != 0 {
			if limit < 1 { limit = 1 }
			if limit > 25 { limit = 25 }
			parameters["limit"] = strconv.Itoa(limit)
		}
		query = generateGetParametersString(parameters)
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "search/releases.json" + query)
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &searchResult)
				}
			}
		}
	}

	return searchResult, err
}

/**
	Searches for Ext Infos. Please note that additional search rate limiting applies.
	See http://www.xrel.to/wiki/2727/api-rate-limiting.html

	query				Search keyword.
	extInfoType	:= ""	One of: movie|tv|game|console|software|xxx - or leave empty to search Ext Infos of all types.
	limit		:= 25	Number of returned search results. Maximum and default 25.

	http://www.xrel.to/wiki/6319/api-search-ext-info.html
 */
func Search_ExtInfo(query, extInfoType string, limit int) (types.ExtInfoSearchResult, error) {
	var searchResult	types.ExtInfoSearchResult
	var err 			error

	if query == "" {
		err = errors.New("Please provide a query string.")
	} else {
		query = "?q=" + url.QueryEscape(query)
		if limit != 0 {
			if limit < 1 { limit = 1 }
			if limit > 25 { limit = 25 }
			query += "&limit=" + strconv.Itoa(limit)
		}
		switch extInfoType {
		case "":
		case "movie", "tv", "game", "console", "software", "xxx":
			query += "&type=" + extInfoType
		default:
			err = errors.New("Wrong media type - Use one of: movie|tv|game|console|software|xxx" +
								" - or leave empty to search media infos of all types.")
		}
		if err == nil {
			client := getClient()
			var response *http.Response
			response, err = client.Get(apiURL + "search/ext_info.json" + query)
			if err == nil {
				defer response.Body.Close()
				err = checkResponseStatusCode(response.StatusCode)
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						bytes = stripeJSON(bytes)
						err = json.Unmarshal(bytes, &searchResult)
					}
				}
			}
		}
	}

	return searchResult, err
}
