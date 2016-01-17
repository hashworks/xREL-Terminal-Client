package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/hashworks/xRELTerminalClient/configHandler"
	"github.com/hashworks/xRELTerminalClient/client"
)

var versionFlag			bool

var configFilePath		string
var authenticateFlag	bool
var checkRateLimitFlag	bool

var isP2PFlag			bool
var perPageFlag			int
var pageFlag			int

var getFiltersFlag		bool
var filterFlag			string
var latestFlag			bool
var browseArchiveFlag	string

var getCategoriesFlag	bool
var extInfoTypeFlag		string
var browseCategoryFlag	string

var infoFlag			bool
var imagesFlag			bool
var videosFlag			bool
var addFavEntryFlag		bool
var rateFlag			int
var limitFlag			int

var releasesFlag		bool
var searchExtInfoFlag	string

var searchReleaseFlag	string

var rateVideoFlag		int
var rateAudioFlag		int
var addCommentFlag		string
var releaseFlag			string

var commentsFlag		string

var upcomingTitlesFlag	bool

var rmFavEntryFlag		bool
var listFavEntriesFlag	bool

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.Usage = client.Usage

	flagSet.BoolVar(&versionFlag, "version", false, "")

	flagSet.StringVar(&configFilePath, "configFile", "", "")
	flagSet.BoolVar(&authenticateFlag, "authenticate", false, "")
	flagSet.BoolVar(&checkRateLimitFlag, "rateLimit", false, "")

	flagSet.BoolVar(&isP2PFlag, "p2p", false, "")

	flagSet.IntVar(&perPageFlag, "perPage", 5, "")
	flagSet.IntVar(&pageFlag, "page", 0, "")

	flagSet.BoolVar(&getFiltersFlag, "filters", false, "")
	flagSet.StringVar(&filterFlag, "filter", "", "")
	flagSet.BoolVar(&latestFlag, "latest", false, "")
	flagSet.StringVar(&browseArchiveFlag, "browseArchive", "", "")

	flagSet.StringVar(&extInfoTypeFlag, "mediaType", "", "")

	flagSet.BoolVar(&getCategoriesFlag, "categories", false, "")
	flagSet.StringVar(&browseCategoryFlag, "browseCategory", "", "")

	flagSet.BoolVar(&imagesFlag, "images", false, "")
	flagSet.BoolVar(&videosFlag, "videos", false, "")
	flagSet.BoolVar(&infoFlag, "info", false, "")
	flagSet.IntVar(&rateFlag, "rate", 0, "")
	flagSet.BoolVar(&addFavEntryFlag, "addToFavorites", false, "")
	flagSet.IntVar(&limitFlag, "limit", 0, "")

	flagSet.BoolVar(&releasesFlag, "releases", false, "")
	flagSet.StringVar(&searchExtInfoFlag, "searchMedia", "", "")

	flagSet.StringVar(&searchReleaseFlag, "searchRelease", "", "")

	flagSet.IntVar(&rateVideoFlag, "rateVideo", 0, "")
	flagSet.IntVar(&rateAudioFlag, "rateAudio", 0, "")
	flagSet.StringVar(&addCommentFlag, "addComment", "", "")
	flagSet.StringVar(&releaseFlag, "release", "", "")
	flagSet.StringVar(&commentsFlag, "comments", "", "")

	flagSet.BoolVar(&upcomingTitlesFlag, "upcomingTitles", false, "")

	flagSet.BoolVar(&listFavEntriesFlag, "showUnreadFavorites", false, "")
	flagSet.BoolVar(&rmFavEntryFlag, "removeFavoriteEntry", false, "")

	flagSet.Parse(os.Args[1:])

	if limitFlag == 0 && perPageFlag != 0 {
		limitFlag = perPageFlag
	}

	config, _ := configHandler.GetConfig(configFilePath)

	if versionFlag {
		fmt.Println("xREL Terminal Client v1.0.1")
		fmt.Println("https://github.com/hashworks/xRELTerminalClient")
		fmt.Println()
		fmt.Println("Published under the GNU General Public License v3.0.")
	} else if rmFavEntryFlag {
		client.RemoveFavEntry()
	} else if listFavEntriesFlag {
		client.ShowFavEntries()
	} else if upcomingTitlesFlag {
		client.ShowUpcomingTitles(releasesFlag, isP2PFlag)
	} else if releaseFlag != "" {
		if addCommentFlag != "" || rateVideoFlag != 0 || rateAudioFlag != 0 {
			if (rateVideoFlag != 0 && rateAudioFlag == 0) || (rateVideoFlag == 0 && rateAudioFlag != 0) {
				fmt.Println("You need to set either both or none of --rateVideo and --rateAudio.")
				os.Exit(2)
			} else {
				client.AddComment(releaseFlag, isP2PFlag, addCommentFlag, rateVideoFlag, rateAudioFlag)
			}
		} else {
			client.ShowRelease(releaseFlag, isP2PFlag)
		}
	} else if searchReleaseFlag != "" {
		client.SearchReleases(searchReleaseFlag, isP2PFlag, limitFlag)
	} else if searchExtInfoFlag != "" {
		client.SearchMedia(searchExtInfoFlag, extInfoTypeFlag, perPageFlag, pageFlag, limitFlag, isP2PFlag, infoFlag, releasesFlag, imagesFlag, videosFlag, addFavEntryFlag, rateFlag, browseArchiveFlag)
	} else if getCategoriesFlag {
		client.ShowCategories(isP2PFlag)
	} else if getFiltersFlag {
		client.ShowFilters(isP2PFlag)
	} else if latestFlag {
		client.ShowLatest(filterFlag, isP2PFlag, perPageFlag, perPageFlag)
	} else if browseArchiveFlag != "" {
		client.BrowseArchive(filterFlag, browseArchiveFlag, isP2PFlag, perPageFlag, pageFlag)
	} else if browseCategoryFlag != "" {
		client.BrowseCategory(browseCategoryFlag, extInfoTypeFlag, isP2PFlag, perPageFlag, pageFlag)
	} else if commentsFlag != "" {
		client.ShowComments(commentsFlag, isP2PFlag, perPageFlag, pageFlag)
	} else if checkRateLimitFlag {
		client.CheckRateLimit()
	} else if authenticateFlag {
		client.Authenticate()
	} else {
		flagSet.Usage()
	}

	err := config.WriteConfig()
	client.OK(err, "\nFailed to write the configuration to file system:\n")
}