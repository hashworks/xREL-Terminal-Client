package xrel

import (
	"encoding/json"
	"errors"
	"github.com/hashworks/xRELTerminalClient/src/xrel/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
Returns a list of all the current user's favorite lists.
Requires oAuth authentication.

http://www.xrel.to/wiki/1754/api-favs-lists.html
*/
func GetFavsLists() ([]types.FavList, error) {
	var (
		favLists []types.FavList
		err      error
	)

	var client *http.Client
	client, err = getOAuthClient()
	if err == nil {
		var response *http.Response
		response, err = client.Get(apiURL + "favs/lists.json")
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &favLists)
				}
			}
		}
	}

	return favLists, err
}

/*
Returns entries of a favorite list.
Requires oAuth authentication.

id						The favorite list ID, as obtained through Favs_GetLists().
getReleases	:= false	If true, an inline list of unread(!) releases will be returned with each ext_info entry.

http://www.xrel.to/wiki/1823/api-favs-list-entries.html
*/
func GetFavsListEntries(id string, getReleases bool) ([]types.ExtendedExtInfo, error) {
	var (
		extendedExtInfos []types.ExtendedExtInfo
		err              error
	)

	if id == "" {
		err = errors.New("Please specify a favs list id.")
	} else {
		var client *http.Client
		client, err = getOAuthClient()
		if err == nil {
			parameters := url.Values{}
			parameters.Add("id", id)
			if getReleases {
				parameters.Add("get_releases", "true")
			}
			var response *http.Response
			response, err = client.PostForm(apiURL+"favs/list_entries.json", parameters)
			defer response.Body.Close()
			if err == nil {
				err = checkResponseStatusCode(response.StatusCode)
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						bytes = stripeJSON(bytes)
						err = json.Unmarshal(bytes, &extendedExtInfos)
					}
				}
			}
		}
	}

	return extendedExtInfos, err
}

/*
Add an Ext Info to a favorite list.
Requires oAuth authentication.

id			The favorite list ID, as obtained through Favs_GetLists().
extInfoId	The Ext Info ID, as obtained through other API calls.

http://www.xrel.to/wiki/6316/api-favs-list-addentry.html
*/
func AddFavsListEntry(id, extInfoId string) (types.FavListEntryModificationResult, error) {
	var (
		favListAddEntryResult types.FavListEntryModificationResult
		err                   error
	)

	if id == "" {
		err = errors.New("Please specify a favs list id.")
	} else if extInfoId == "" {
		err = errors.New("Please specify an ext info id to add.")
	} else {
		var client *http.Client
		client, err = getOAuthClient()
		if err == nil {
			parameters := url.Values{}
			parameters.Add("id", id)
			parameters.Add("ext_info_id", extInfoId)
			var response *http.Response
			response, err = client.PostForm(apiURL+"favs/list_addentry.json", parameters)
			defer response.Body.Close()
			if err == nil {
				switch response.StatusCode {
				case 404:
					err = errors.New("There was an error with the favorite list. Does it exist?")
				case 400:
					err = errors.New("There was an error with the ExtInfo. Maybe it doesn't exist or it is already on the list?")
				default:
					err = checkResponseStatusCode(response.StatusCode)
				}
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						bytes = stripeJSON(bytes)
						err = json.Unmarshal(bytes, &favListAddEntryResult)
					}
				}
			}
		}
	}

	return favListAddEntryResult, err
}

/*
Removes an Ext Info from a favorite list.
Requires oAuth authentication.

id			The favorite list ID, as obtained through Favs_GetLists().
extInfoId	The Ext Info ID, as obtained through other API calls.

http://www.xrel.to/wiki/6317/api-favs-list-delentry.html
*/
func RemoveFavsListEntry(id, extInfoId string) (types.FavListEntryModificationResult, error) {
	var (
		favListRemoveEntryResult types.FavListEntryModificationResult
		err                      error
	)

	if id == "" {
		err = errors.New("Please specify a favs list id.")
	} else if extInfoId == "" {
		err = errors.New("Please specify an ext info id to remove.")
	} else {
		var client *http.Client
		client, err = getOAuthClient()
		if err == nil {
			parameters := url.Values{}
			parameters.Add("id", id)
			parameters.Add("ext_info_id", extInfoId)
			var response *http.Response
			response, err = client.PostForm(apiURL+"favs/list_delentry.json", parameters)
			defer response.Body.Close()
			if err == nil {
				switch response.StatusCode {
				case 404:
					err = errors.New("There was an error with the favorite list. Does it exist?")
				case 400:
					err = errors.New("There was an error with the ExtInfo. Maybe it doesn't exist or it is not on the list?")
				default:
					err = checkResponseStatusCode(response.StatusCode)
				}
				if err == nil {
					var bytes []byte
					bytes, err = ioutil.ReadAll(response.Body)
					if err == nil {
						bytes = stripeJSON(bytes)
						err = json.Unmarshal(bytes, &favListRemoveEntryResult)
					}
				}
			}
		}
	}

	return favListRemoveEntryResult, err
}
