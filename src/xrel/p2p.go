package xrel

import (
	"encoding/json"
	"errors"
	"github.com/hashworks/xRELTerminalClient/src/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
GetP2PReleaseInfo returns information about a single release, specified by the complete dirname or an API release id.

http://www.xrel.to/wiki/3697/api-p2p-rls-info.html
*/
func GetP2PReleaseInfo(query string, isID bool) (types.P2PRelease, error) {
	var (
		p2pReleaseStruct types.P2PRelease
		err              error
	)

	if query == "" {
		err = errors.New("Please provide a dirname or an id as query.")
	} else {
		if isID {
			query = "?id=" + query
		} else {
			query = "?dirname=" + query
		}
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "p2p/rls_info.json" + query)
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &p2pReleaseStruct)
				}
			}
		}
	}

	return p2pReleaseStruct, err
}

/*
GetP2PReleases allows to browse P2P/non-scene releases.

	perPage		:= 25	Number of releases per page. Min. 5, max. 100.
	page     	:= 1	Page number (1 to N).

	Set only one of the following:
	categoryID	:= ""	P2P category ID from GetP2PCategories()
	groupID		:= ""	P2P release group ID
	extInfoID	:= ""	Ext Info ID


http://www.xrel.to/wiki/3699/api-p2p-releases.html
*/
func GetP2PReleases(perPage, page int, categoryID, groupID, extInfoID string) (types.P2PReleases, error) {
	var (
		p2pReleasesStruct types.P2PReleases
		parameters        = url.Values{}
	)

	if perPage != 0 {
		if perPage < 5 {
			perPage = 5
		}
		if perPage > 100 {
			perPage = 100
		}
		parameters.Add("per_page", strconv.Itoa(perPage))
	}
	if page > 0 {
		parameters.Add("page", strconv.Itoa(page))
	}
	if categoryID != "" {
		parameters.Add("category_id", categoryID)
	} else if groupID != "" {
		parameters.Add("group_id", groupID)
	} else if extInfoID != "" {
		parameters.Add("ext_info_id", extInfoID)
	}

	client := getClient()
	response, err := client.PostForm(apiURL+"p2p/releases.json", parameters)
	defer response.Body.Close()
	if err == nil {
		err = checkResponseStatusCode(response.StatusCode)
		if err == nil {
			var bytes []byte
			bytes, err = ioutil.ReadAll(response.Body)
			if err == nil {
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &p2pReleasesStruct)
			}
		}
	}

	return p2pReleasesStruct, err
}

/*
GetP2PCategories returns a list of available P2P release categories and their IDs. You can use the category ID in GetP2PReleases().

http://www.xrel.to/wiki/3698/api-p2p-categories.html
*/
func GetP2PCategories() ([]types.P2PCategory, error) {
	var err error

	// According to xREL we should cache the results for 24h
	currentUnix := time.Now().Unix()
	if Config.LastP2PCategoryRequest == 0 || currentUnix-Config.LastP2PCategoryRequest > 86400 || len(Config.P2PCategories) == 0 {
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "p2p/categories.json")
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &Config.P2PCategories)
					if err == nil {
						Config.LastP2PCategoryRequest = currentUnix
					}
				}
			}
		}
	}

	return Config.P2PCategories, err
}
