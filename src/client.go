package main

import (
	"./xrel"
	"./xrel/types"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	// 2006-01-02 15:04:05.999999999 -0700 MST
	xRELCommentTimeFormat = "02. Jan 2006, 03:04 pm"
	xRELReleaseTimeFormat = "02.01.2006 03:04 pm"
)

var (
	// Set the following uppercase three with -ldflags "-X main.VERSION=v1.2.3 [...]"
	VERSION               string = "unknown"
	OAUTH_CONSUMER_KEY    string
	OAUTH_CONSUMER_SECRET string

	versionFlag        bool
	configFilePathFlag string
	authenticateFlag   bool
	checkRateLimitFlag bool
	isP2PFlag          bool
	perPageFlag        int
	pageFlag           int
	getFiltersFlag     bool
	filterFlag         string
	latestFlag         bool
	browseArchiveFlag  string
	getCategoriesFlag  bool
	extInfoTypeFlag    string
	browseCategoryFlag string
	infoFlag           bool
	imagesFlag         bool
	videosFlag         bool
	addFavEntryFlag    bool
	rateFlag           int
	limitFlag          int
	releasesFlag       bool
	searchExtInfoFlag  string
	searchReleaseFlag  string
	rateVideoFlag      int
	rateAudioFlag      int
	addCommentFlag     string
	releaseFlag        string
	commentsFlag       string
	upcomingTitlesFlag bool
	rmFavEntryFlag     bool
	listFavEntriesFlag bool
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.Usage = Usage

	flagSet.BoolVar(&versionFlag, "version", false, "")

	flagSet.StringVar(&configFilePathFlag, "configFile", "", "")
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

	_ = readConfig(configFilePathFlag)

	xrel.SetOAuthConsumerKeyAndSecret(OAUTH_CONSUMER_KEY, OAUTH_CONSUMER_SECRET)

	switch {
	case versionFlag:
		fmt.Println("xREL Terminal Client")
		fmt.Println("https://github.com/hashworks/xRELTerminalClient")
		fmt.Println("Version: " + VERSION)
		fmt.Println()
		fmt.Println("Published under the GNU General Public License v3.0.")
	case rmFavEntryFlag:
		removeFavEntry()
	case listFavEntriesFlag:
		showFavEntries()
	case upcomingTitlesFlag:
		showUpcomingTitles(releasesFlag, isP2PFlag)
	case releaseFlag != "":
		if addCommentFlag != "" || rateVideoFlag != 0 || rateAudioFlag != 0 {
			if (rateVideoFlag != 0 && rateAudioFlag == 0) || (rateVideoFlag == 0 && rateAudioFlag != 0) {
				fmt.Println("You need to set either both or none of --rateVideo and --rateAudio.")
				os.Exit(2)
			} else {
				addComment(releaseFlag, isP2PFlag, addCommentFlag, rateVideoFlag, rateAudioFlag)
			}
		} else {
			showRelease(releaseFlag, isP2PFlag)
		}
	case searchReleaseFlag != "":
		searchReleases(searchReleaseFlag, isP2PFlag, limitFlag)
	case searchExtInfoFlag != "":
		searchMedia(searchExtInfoFlag, extInfoTypeFlag, perPageFlag, pageFlag, limitFlag, isP2PFlag, infoFlag, releasesFlag, imagesFlag, videosFlag, addFavEntryFlag, rateFlag, browseArchiveFlag)
	case getCategoriesFlag:
		showCategories(isP2PFlag)
	case getFiltersFlag:
		showFilters(isP2PFlag)
	case latestFlag:
		showLatest(filterFlag, isP2PFlag, perPageFlag, perPageFlag)
	case browseArchiveFlag != "":
		browseArchive(filterFlag, browseArchiveFlag, isP2PFlag, perPageFlag, pageFlag)
	case browseCategoryFlag != "":
		browseCategory(browseCategoryFlag, extInfoTypeFlag, isP2PFlag, perPageFlag, pageFlag)
	case commentsFlag != "":
		showComments(commentsFlag, isP2PFlag, perPageFlag, pageFlag)
	case checkRateLimitFlag:
		checkRateLimit()
	case authenticateFlag:
		authenticate()
	default:
		flagSet.Usage()
	}

	ok(writeConfig(), "\nFailed to write the configuration to file system:\n")
}

func ok(err error, prefix string) {
	if err != nil {
		fmt.Println(prefix + err.Error())
		os.Exit(1)
	}
}

func findP2PCategoryID(categoryName string) (string, error) {
	var (
		categoryID string
		err        error
		categories []types.P2PCategory
	)

	categories, err = xrel.GetP2PCategories()
	if err == nil {
		for i := 0; i < len(categories); i++ {
			category := categories[i]
			if category.SubCat != "" {
				if strings.ToLower(category.SubCat) == strings.ToLower(categoryName) {
					categoryID = category.Id
				}
			} else if strings.ToLower(category.MetaCat) == strings.ToLower(categoryName) {
				categoryID = category.Id
			}
			if categoryID != "" {
				break
			}
		}
		if categoryID == "" {
			err = errors.New("Category not found. Please choose one of --categories --p2p.")
		}
	}

	return categoryID, err
}
