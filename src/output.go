package main

import (
	"fmt"
	"time"
	"strings"
	"regexp"
	"html"
	"./xREL"
	"./xREL/types"
)

// Note that this should only contain outputs that are used multiple times.

func printP2PRelease(release types.P2PRelease) {
	printRelease(release.Dirname, release.ExtInfo.Title, release.LinkHref, release.ExtInfo.Type, "", "MB",
		release.SizeInMB, release.TVSeason, release.TVEpisode, release.Comments, release.ExtInfo.Rating,
		release.VideoRating, release.AudioRating, release.GetPubTime())
}

func printSceneRelease(release types.Release) {
	printRelease(release.Dirname, release.ExtInfo.Title, release.LinkHref, release.ExtInfo.Type, release.NukeReason, release.Size.Unit,
		release.Size.Number, release.TVSeason, release.TVEpisode, release.Comments, release.ExtInfo.Rating,
		release.VideoRating, release.AudioRating, release.GetTime())
}

func printRelease(dirname, title, link, releaseType, nukeReason, sizeUnit string, sizeNumber, tvSeason, tvEpisode, commentCount int, rating , videoRating, audioRating float32, preTime time.Time) {
	if preTime != (time.Time{}) {
		fmt.Print(preTime.Format(xRELReleaseTimeFormat) + " - ")
	}

	if title != "" {
		if releaseType != "" {
			fmt.Print("[" + strings.ToUpper(releaseType) + "] ")
		}
		fmt.Print(title)
		if tvEpisode != 0 {
			fmt.Printf(" [S%02dE%02d]", tvSeason, tvEpisode)
		}
		if rating != 0 {
			fmt.Printf(" [%2.1f★]", rating);
		}
		fmt.Print("\n\t" + dirname)
	} else {
		fmt.Print(dirname)
	}

	fmt.Println("\n\t" + link)
	fmt.Printf("\t%d %s - %d comments", sizeNumber, sizeUnit, commentCount)
	if videoRating != 0 && audioRating != 0 {
		fmt.Printf(" - Video: %2.1f★, Audio: %2.1f★", videoRating, audioRating)
	}
	if nukeReason != "" {
		fmt.Print(" - nuked for \"" + nukeReason + "\"")
	}
	fmt.Println()
}

func printP2PReleases(releases types.P2PReleases, printPagination, noTitle bool) {
	pagination := releases.Pagination
	if (printPagination && pagination.TotalPages > 1) {
		fmt.Printf("P2P releases %d/%d (Page %d of %d):\n", len(releases.List), releases.TotalCount, pagination.CurrentPage, pagination.TotalPages);
	} else {
		fmt.Println("P2P releases:")
	}
	for i := 0; i < len(releases.List); i++ {
		p2pRelease := releases.List[i]
		if noTitle {
			p2pRelease.ExtInfo.Title = ""
		}
		printP2PRelease(p2pRelease)
	}
}

func printSceneReleases(releases types.Releases, printPagination, noTitle bool) {
	pagination := releases.Pagination
	if (printPagination && pagination.TotalPages > 1) {
		fmt.Printf("Scene releases %d/%d (Page %d of %d):\n", len(releases.List), releases.TotalCount, pagination.CurrentPage, pagination.TotalPages);
	} else {
		fmt.Println("Scene releases:")
	}
	for i := 0; i < len(releases.List); i++ {
		release := releases.List[i]
		if noTitle {
			release.ExtInfo.Title = ""
		}
		printSceneRelease(release)
	}
}

func printExtendedExtInfo(extInfo types.ExtendedExtInfo) {
	fmt.Printf("%s [%s]", extInfo.Title, strings.ToUpper(extInfo.Type))
	if extInfo.Rating != 0 {
		fmt.Printf(" [%2.1f★]", extInfo.Rating)
	}
	fmt.Print(" " + extInfo.LinkHref)
	fmt.Println()
	if extInfo.AltTitle != "" {
		fmt.Println("Alternative title: " + extInfo.AltTitle)
	}
	if extInfo.Genre != "" {
		fmt.Println("Genre: " + extInfo.Genre)
	}
	if len(extInfo.ReleaseDates) > 0 {
		fmt.Println("\nRelease dates: ")
		for i := 0; i < len(extInfo.ReleaseDates); i++ {
			fmt.Println(extInfo.ReleaseDates[i].Date + ": " + strings.ToUpper(extInfo.ReleaseDates[i].Type))
		}
	}
	if len(extInfo.URIs) > 0 {
		fmt.Println("\nURIs: ")
		for i := 0; i < len(extInfo.URIs); i++ {
			fmt.Println(extInfo.URIs[i])
		}
		fmt.Println()
	}

	if len(extInfo.Externals) > 0 {
		for i := 0; i < len(extInfo.Externals); i++ {
			external := extInfo.Externals[i]
			if (external.Plot != "") {
				fmt.Println("Plot laut " + external.Source.Name + " (" + external.LinkUrl + "):")
				plot := regexp.MustCompile("<(.+?)[\\s]*\\/?[\\s]*>").ReplaceAllString(external.Plot, "")
				fmt.Println(html.UnescapeString(plot))
				break;
			}
		}
	}
}

func outputExtInfoData(id string, perPageFlag, pageFlag int, isP2PFlag, infoFlag, releasesFlag, imagesFlag, videosFlag bool, rateFlag int, browseCategoryFlag string) {
	multipleItems := false

	if infoFlag || (!releasesFlag && !imagesFlag && !videosFlag && rateFlag == 0) {
		extInfo, err := xREL.GetExtInfo(id)
		ok(err, "Failed to get media information:\n")
		printExtendedExtInfo(extInfo)
		multipleItems = true
	}

	if imagesFlag || videosFlag {
		if (multipleItems) {
			fmt.Println()
		}
		items, err := xREL.GetExtInfoMedia(id);
		itemCount := len(items)
		if err == nil && itemCount > 0 {
			if (imagesFlag) {
				fmt.Println("Images:")
				for i := 0; i < itemCount; i++ {
					if (items[i].IsImage()) {
						fmt.Println(items[i].Description + " - " + items[i].UrlFull)
					}
				}
			}
			if (videosFlag) {
				fmt.Println("Videos:")
				for i := 0; i < itemCount; i++ {
					if (items[i].IsVideo()) {
						fmt.Println(items[i].Description + " - " + items[i].VideoURL)
					}
				}
			}
		} else {
			fmt.Println("No images or videos found.")
		}
		multipleItems = true
	}

	if (rateFlag > 0) {
		if (multipleItems) {
			fmt.Println()
		}
		extInfo, err := xREL.RateExtInfo(id, rateFlag)
		ok(err, "Failed to rate media:\n")
		if (infoFlag) {
			fmt.Print("R")
		} else {
			fmt.Printf("%s [%s] r", extInfo.Title, strings.ToUpper(extInfo.Type))
		}
		fmt.Printf("ated with %s★, overall rating is %2.1f★.\n", extInfo.OwnRating, extInfo.Rating)
		multipleItems = true
	}

	if (releasesFlag) {
		if (multipleItems) {
			fmt.Println()
		}
		if (isP2PFlag) {
			var categoryID	string
			var err			error
			if browseCategoryFlag != "" {
				categoryID, err = findP2PCategoryID(browseCategoryFlag)
				ok(err, "Failed to get category id:\n")
			}
			p2pReleases, err := xREL.GetP2PReleases(perPageFlag, pageFlag, categoryID, "", id)
			ok(err, "Failed to load p2p releases by media:\n")
			printP2PReleases(p2pReleases, false, true)
		} else {
			releases, err := xREL.GetReleaseByExtInfoId(id, perPageFlag, pageFlag)
			ok(err, "Failed to load scene releases by media:\n")
			printSceneReleases(releases, false, true)
		}
	}
}

func printComment(comment types.Comment) {
	fmt.Print("[ " + comment.Author.Name)
	postTime, timeErr := comment.GetTime()
	if timeErr == nil && postTime != (time.Time{}) {
		fmt.Print(" - " + postTime.Format(xRELCommentTimeFormat))
	}
	fmt.Print(" - ")
	if comment.Rating.Video != "" {
		fmt.Print("Video: " + comment.Rating.Video + "★ | Audio: " + comment.Rating.Audio + "★")
	} else {
		fmt.Print("No rating by user")
	}
	fmt.Print(" ] \n")

	if comment.Text != "" {
		fmt.Print("\n" + comment.Text)
		fmt.Print("\n\n")
		if (comment.Edits.Count != 0) {
			lastEditTime, timeErr := comment.Edits.GetLast()
			if timeErr == nil {
				fmt.Printf("[ Edited %d times, last on %s ]\n", comment.Edits.Count, lastEditTime.Format(xRELCommentTimeFormat))
			}
		}
		fmt.Printf("[ %d+ vs %d- | %s ]\n", comment.Votes.Positive, comment.Votes.Negative, comment.LinkHref)
	}
}