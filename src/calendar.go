package main

import (
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"os"
	"strings"
)

func showUpcomingTitles(showReleases, showP2P bool) {
	titles, err := xrel.GetUpcomingTitles()
	ok(err, "Failed to get upcoming titles:\n")
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
						fmt.Println("\t" + release.GetPubTime().Format(xRELReleaseTimeFormat) + " - " + release.Dirname)
					}
				} else {
					for i := 0; i < len(title.Releases); i++ {
						release := title.Releases[i]
						fmt.Println("\t" + release.GetTime().Format(xRELReleaseTimeFormat) + " - " + release.Dirname)
					}
				}
			}
		}
	}
}
