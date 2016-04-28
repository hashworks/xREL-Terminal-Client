package main

import (
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"os"
	"strings"
)

func showUpcomingTitles(country string, showReleases, showP2P bool) {
	titles, err := xrel.GetUpcomingTitles(country)
	ok(err, "Failed to get upcoming titles: ")
	titleCount := len(titles)
	if titleCount == 0 {
		fmt.Println("No upcoming titles found.")
		os.Exit(1)
	} else {
		for i := 0; i < titleCount; i++ {
			title := titles[i]
			fmt.Println("[" + strings.ToUpper(title.Type) + "] " + title.Title + " [" + title.Genre + "]")
			fmt.Println("\t" + title.LinkHref)
			if showReleases {
				if showP2P {
					for i := 0; i < len(title.P2PReleases); i++ {
						release := title.P2PReleases[i]
						fmt.Println("\t" + release.GetPubTime().Format(types.TIME_FORMAT_RELEASE) + " - " + release.Dirname)
					}
				} else {
					for i := 0; i < len(title.Releases); i++ {
						release := title.Releases[i]
						fmt.Println("\t" + release.GetTime().Format(types.TIME_FORMAT_RELEASE) + " - " + release.Dirname)
					}
				}
			}
		}
	}
}
