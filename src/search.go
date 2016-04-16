package main

import (
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"os"
	"strings"
)

func searchMedia(query, extInfoType string, perPage, page, limit int, isP2P, showInfo, showReleases, showImages, showVideos, addFavEntry bool, rateMedia int, category string) {
	extInfoType = strings.ToLower(extInfoType)
	results, err := xrel.SearchExtInfos(query, extInfoType, limit)
	ok(err, "Failed to search for media:\n")
	if results.Total == 0 {
		fmt.Println("Nothing found.")
		os.Exit(1)
	}
	var id string
	if results.Total > 1 {
		fmt.Printf("Found %d results:\n", results.Total)
		for i := 0; i < len(results.Results); i++ {
			result := results.Results[i]
			var title string
			if extInfoType == "" {
				title = "[" + strings.ToUpper(result.Type) + "] "
			} else {
			}
			title += result.Title
			if result.NumRatings > 0 {
				fmt.Printf("(%d) %s [%2.1f★]\n", i+1, title, result.Rating)
			} else {
				fmt.Printf("(%d) %s\n", i+1, title)
			}
		}
		selection := -1
		fmt.Print("Please choose one: ")
		for selection < 1 || selection > results.Total {
			fmt.Scanln(&selection)
		}
		fmt.Println()
		id = results.Results[selection-1].Id
	} else {
		id = results.Results[0].Id
	}
	if addFavEntry {
		addEntryToFavList(id)
	} else {
		outputExtInfoData(id, perPage, page, isP2P, showInfo, showReleases, showImages, showVideos, rateMedia, category)
	}
}

func searchReleases(query string, isP2P bool, limit int) {
	var (
		results types.ReleaseSearchResult
		err     error
	)
	if isP2P {
		results, err = xrel.SearchReleases(query, false, true, limit)
	} else {
		results, err = xrel.SearchReleases(query, true, false, limit)
	}
	ok(err, "Failed to search for releases:\n")
	if results.Total == 0 {
		fmt.Println("Nothing found.")
		os.Exit(1)
	}
	if results.Total > 1 {
		fmt.Printf("Found %d results:\n", results.Total)
	}
	if isP2P {
		for i := 0; i < len(results.P2PResults); i++ {
			result := results.P2PResults[i]
			printP2PRelease(result)
		}
	} else {
		for i := 0; i < len(results.SceneResults); i++ {
			result := results.SceneResults[i]
			printSceneRelease(result)
		}
	}
}
