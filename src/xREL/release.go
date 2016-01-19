package xREL

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"errors"
	"./types"
)

/**
	Returns information about a single release, specified by the complete dirname or an API release id.

	http://www.xrel.to/wiki/1680/api-release-info.html
 */
func GetReleaseInfo(query string, isID bool) (types.Release, error) {
	var release types.Release

	if isID {
		query = "?id=" + query
	} else {
		query = "?dirname=" + query
	}
	client := getClient()
	response, err := client.Get(apiURL + "release/info.json" + query)
	defer response.Body.Close()
	if err == nil {
		err = checkResponseStatusCode(response.StatusCode)
		if err == nil {
			var bytes []byte
			bytes, err = ioutil.ReadAll(response.Body)
			if err == nil {
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &release)
			}
		}
	}

	return release, err
}

func getReleases(url string) (types.Releases, error) {
	var releases types.Releases

	client := getClient()
	response, err := client.Get(url)
	defer response.Body.Close()
	if err == nil {
		err = checkResponseStatusCode(response.StatusCode)
		if err == nil {
			var bytes []byte
			bytes, err = ioutil.ReadAll(response.Body)
			if err == nil {
				bytes = stripeJSON(bytes)
				err = json.Unmarshal(bytes, &releases)
			}
		}
	}

	return releases, err
}

/**
	Returns the latest releases. Also allows to browse the archive by month.
	perPage	:= 25	Number of releases per page. Min. 5, max. 100.
	page	:= 1	Page number (1 to N).
	filter	:= ""	Filter ID (from Release_GetFilters()) or "overview" to use the currently logged in user's overview filter.
	archive	:= ""	Empty = current releases, YYYY-MM for archive.

	Note that this won't include any release rating information.

	http://www.xrel.to/wiki/2994/api-release-latest.html
 */
func GetLatestReleases(perPage, page int, filter, archive string) (types.Releases, error) {
	parameters := make(map[string]string)

	if perPage != 0 {
		if perPage < 5 { perPage = 5 }
		if perPage > 100 { perPage = 100 }
		parameters["per_page"] = strconv.Itoa(perPage)
	}
	if page > 0 {
		parameters["page"] = strconv.Itoa(page)
	}
	if filter != "" {
		parameters["filter"] = filter
	}
	if archive != "" {
		parameters["archive"] = archive
	}
	query := generateGetParametersString(parameters)

	return getReleases(apiURL + "release/latest.json" + query)
}

/**
	Returns a list of public, predefined release filters. You can use the filter ID in Release_GetLatest().

	http://www.xrel.to/wiki/2996/api-release-filters.html
 */
func GetReleaseFilters() ([]types.Filter, error) {
	var err error

	currentUnix := time.Now().Unix()
	if Config.LastFilterRequest == 0 || currentUnix - Config.LastFilterRequest > 86400 || len(Config.Filters) == 0 {
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "release/filters.json")
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &Config.Filters)
					if err == nil {
						Config.LastFilterRequest = currentUnix
					}
				}
			}
		}
	}

	return Config.Filters, err
}

/**
	Returns scene releases from the given category.
	categoryName		Category name from Release_GetCategories()
	extInfoType := ""	Use one of: movie|tv|game|console|software|xxx - or leave empty to browse releases of all types

	http://www.xrel.to/wiki/3751/api-release-browse-category.html
 */
func BrowseReleaseCategory(categoryName, extInfoType string, perPage, page int) (types.Releases, error) {
	var releasesStruct	types.Releases
	var err				error

	if categoryName == "" {
		err = errors.New("Please specifiy a category name.");
	} else {
		query := "?category_name=" + categoryName
		if perPage != 0 {
			if perPage < 5 { perPage = 5 }
			if perPage > 100 { perPage = 100 }
			query += "&per_page=" + strconv.Itoa(perPage)
		}
		if page > 0 {
			query += "&page=" + strconv.Itoa(page)
		}
		switch extInfoType {
		case "":
		case "movie", "tv", "game", "console", "software", "xxx":
			query += "&ext_info_type=" + extInfoType
		default:
			err = errors.New("Wrong extInfoType - Use one of: movie|tv|game|console|software|xxx" +
								" - or leave empty to browse releases of all types.")
		}
		if err == nil {
			releasesStruct, err = getReleases(apiURL + "release/browse_category.json" + query)
		}
	}

	return releasesStruct, err
}

/**
	Returns a list of available release categories. You can use the category name in Release_BrowseCategory().

	http://www.xrel.to/wiki/6318/api-release-categories.html
 */
func GetReleaseCategories() ([]types.Category, error) {
	var err error

	currentUnix := time.Now().Unix()
	if Config.LastCategoryRequest == 0 || currentUnix - Config.LastCategoryRequest > 86400 || len(Config.Categories) == 0 {
		client := getClient()
		var response *http.Response
		response, err = client.Get(apiURL + "release/categories.json")
		defer response.Body.Close()
		if err == nil {
			err = checkResponseStatusCode(response.StatusCode)
			if err == nil {
				var bytes []byte
				bytes, err = ioutil.ReadAll(response.Body)
				if err == nil {
					bytes = stripeJSON(bytes)
					err = json.Unmarshal(bytes, &Config.Categories)
					if err == nil {
						Config.LastCategoryRequest = currentUnix
					}
				}
			}
		}
	}

	return Config.Categories, err
}

/**
	Returns all releases associated with a given Ext Info.
	id				Ext info ID.
	perPage	:= 25	Number of releases per page. Min. 5, max. 100.
	page	:= 1	Page number (1 to N).

	http://www.xrel.to/wiki/2822/api-release-ext-info.html
 */
func GetReleaseByExtInfoId(id string, perPage, page int) (types.Releases, error) {
	query := "?id=" + id
	if perPage != 0 {
		if perPage < 5 { perPage = 5 }
		if perPage > 100 { perPage = 100 }
		query += "&per_page=" + strconv.Itoa(perPage)
	}
	if page > 0 {
		query += "&page=" + strconv.Itoa(page)
	}

	return getReleases(apiURL + "release/ext_info.json" + query)
}

