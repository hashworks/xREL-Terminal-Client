package main

import (
	"./xREL"
	"./xREL/types"
	"fmt"
	"os"
)

func showRelease(dirname string, isP2P bool) {
	if isP2P {
		release, err := xREL.GetP2PReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the p2p release:\n")
		fmt.Println(release.Dirname)
		fmt.Println("Link: " + release.LinkHref)
		fmt.Println("Media: " + release.ExtInfo.LinkHref)
		fmt.Println()
		fmt.Printf("Size: %d MB\n", release.SizeInMB)
		pubTime := release.GetPubTime()
		postTime := release.GetPostTime()
		if postTime == pubTime {
			fmt.Println("Pub- & PostTime: " + pubTime.Format(xRELReleaseTimeFormat))
		} else {
			fmt.Println("PubTime: " + pubTime.Format(xRELReleaseTimeFormat))
			fmt.Println("PostTime: " + postTime.Format(xRELReleaseTimeFormat))
		}
		fmt.Println()
		if release.VideoRating != 0 {
			fmt.Printf("Video: %2.1f★\n", release.VideoRating)
			fmt.Printf("Audio: %2.1f★\n", release.AudioRating)
		} else {
			fmt.Println("No rating yet.")
		}
		fmt.Printf("Release has %d comments.\n", release.Comments)
	} else {
		release, err := xREL.GetReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the scene release:\n")
		fmt.Println(release.Dirname)
		fmt.Println("Link: " + release.LinkHref)
		fmt.Println("Media: " + release.ExtInfo.LinkHref)
		fmt.Println()
		fmt.Printf("Size: %d %s\n", release.Size.Number, release.Size.Unit)
		if release.ProofUrl != "" {
			fmt.Println("Proof: " + release.ProofUrl)
		}
		fmt.Println("PreTime: " + release.GetTime().Format(xRELReleaseTimeFormat))
		fmt.Println()
		if release.VideoRating != 0 {
			fmt.Printf("Video: %2.1f★\n", release.VideoRating)
			fmt.Printf("Audio: %2.1f★\n", release.AudioRating)
		} else {
			fmt.Println("No rating yet.")
		}
		fmt.Printf("Release has %d comments.\n", release.Comments)
		fmt.Println()
		if release.NukeReason == "" {
			fmt.Println("Release isn't nuked.")
		} else {
			fmt.Println("Release is nuked for \"" + release.NukeReason + "\"!")
		}
	}
}

func addComment(dirname string, isP2P bool, addComment string, rateVideo, rateAudio int) {
	if (rateVideo != 0 && rateAudio == 0) || (rateVideo == 0 && rateAudio != 0) {
		fmt.Println("You need to set either both or none of --rateVideo and --rateAudio.")
		os.Exit(1)
	}
	var id string
	if isP2P {
		release, err := xREL.GetP2PReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the p2p release:\n")
		id = release.Id
	} else {
		release, err := xREL.GetReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the scene release:\n")
		id = release.Id
	}
	comment, err := xREL.AddComment(id, isP2P, addComment, rateVideo, rateAudio)
	ok(err, "Failed to add comment:\n")
	fmt.Println("Sucessfully added comment:")
	printComment(comment)
}

func showComments(query string, isP2P bool, perPage, page int) {
	var (
		id  string
		err error
	)

	if isP2P {
		var p2pRelease types.P2PRelease
		p2pRelease, err = xREL.GetP2PReleaseInfo(query, false)
		if err == nil {
			id = p2pRelease.Id
		}
	} else {
		var release types.Release
		release, err := xREL.GetReleaseInfo(query, false)
		if err == nil {
			id = release.Id
		}
	}
	ok(err, "Failed to get release:\n")
	data, err := xREL.GetComments(id, isP2P, perPage, page)
	ok(err, "Failed to get comments:\n")
	commentCount := len(data.List)
	if commentCount > 0 {
		pagination := data.Pagination
		if pagination.TotalPages > 1 {
			fmt.Printf("Comments %d of %s (Page %d of %d):\n\n", commentCount, data.TotalCount, pagination.CurrentPage, pagination.TotalPages)
		} else {
			fmt.Println("Comments:\n")
		}
		for i := 0; i < commentCount; i++ {
			if i > 0 {
				fmt.Println("----------------------------------------------------------------\n")
			}
			printComment(data.List[i])
			fmt.Print("\n")
		}
	} else {
		fmt.Println("Release has no comments.")
	}
}
