package main

import (
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"os"
)

func showRelease(dirname string, isP2P bool) {
	if isP2P {
		release, err := xrel.GetP2PReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the p2p release: ")
		fmt.Println(release.Dirname)
		fmt.Println("Link: " + release.LinkURL)
		fmt.Println("Media: " + release.ExtInfo.LinkURL)
		fmt.Println()
		fmt.Printf("Size: %d MB\n", release.SizeInMB)
		pubTime := release.GetPubTime()
		postTime := release.GetPostTime()
		if postTime == pubTime {
			fmt.Println("Pub- & PostTime: " + pubTime.Format(types.TIME_FORMAT_RELEASE))
		} else {
			fmt.Println("PubTime: " + pubTime.Format(types.TIME_FORMAT_RELEASE))
			fmt.Println("PostTime: " + postTime.Format(types.TIME_FORMAT_RELEASE))
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
		release, err := xrel.GetReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the scene release: ")
		fmt.Println(release.Dirname)
		fmt.Println("Link: " + release.LinkURL)
		fmt.Println("Media: " + release.ExtInfo.LinkURL)
		fmt.Println()
		fmt.Printf("Size: %d %s\n", release.Size.Number, release.Size.Unit)
		if release.ProofURL != "" {
			fmt.Println("Proof: " + release.ProofURL)
		}
		fmt.Println("PreTime: " + release.GetTime().Format(types.TIME_FORMAT_RELEASE))
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
		release, err := xrel.GetP2PReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the p2p release: ")
		id = release.ID
	} else {
		release, err := xrel.GetReleaseInfo(dirname, false)
		ok(err, "Failed to get information about the scene release: ")
		id = release.ID
	}
	comment, err := xrel.AddComment(id, isP2P, addComment, rateVideo, rateAudio)
	ok(err, "Failed to add comment: ")
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
		p2pRelease, err = xrel.GetP2PReleaseInfo(query, false)
		if err == nil {
			id = p2pRelease.ID
		}
	} else {
		var release types.Release
		release, err := xrel.GetReleaseInfo(query, false)
		if err == nil {
			id = release.ID
		}
	}
	ok(err, "Failed to get release: ")
	data, err := xrel.GetComments(id, isP2P, perPage, page)
	ok(err, "Failed to get comments: ")
	commentCount := len(data.List)
	if commentCount > 0 {
		pagination := data.Pagination
		if pagination.TotalPages > 1 {
			fmt.Printf("Comments %d of %d (Page %d of %d):\n\n", commentCount, data.TotalCount, pagination.CurrentPage, pagination.TotalPages)
		} else {
			fmt.Println("Comments: ")
		}
		for i := 0; i < commentCount; i++ {
			if i > 0 {
				fmt.Print("----------------------------------------------------------------\n\n")
			}
			printComment(data.List[i])
			fmt.Print("\n")
		}
	} else {
		fmt.Println("Release has no comments.")
	}
}

func getNFOImage(dirname string, isP2P bool) {
	var id string
	if isP2P {
		release, err := xrel.GetP2PReleaseInfo(dirname, false)
		ok(err, "Failed to get p2p release id: ")
		id = release.ID
	} else {
		release, err := xrel.GetReleaseInfo(dirname, false)
		ok(err, "Failed to get release id: ")
		id = release.ID
	}
	imageData, err := xrel.GetNFOByID(id, isP2P)
	ok(err, "Failed to get NFO image: ")
	err = ioutil.WriteFile(dirname+".nfo.png", imageData, 0644)
	ok(err, "Failed to save NFO image: ")
	fmt.Println("Saved image '" + dirname + ".nfo.png'.")
}

func addProof(filepath string, isP2P bool, dirnames []string) {
	if isP2P {
		fmt.Println("Proof pictures aren't supported for P2P releases yet.")
		os.Exit(2)
	}
	dirnamesCount := len(dirnames)
	if dirnamesCount == 0 {
		fmt.Println("Please provide at least one release dirname.")
		os.Exit(2)
	}
	var ids []string
	for i := 0; i < dirnamesCount; i++ {
		release, err := xrel.GetReleaseInfo(dirnames[i], false)
		if err != nil {
			fmt.Println("Failed to get release id of '" + dirnames[i] + "': " + err.Error())
		} else {
			ids = append(ids, release.ID)
		}
	}
	if len(ids) == 0 {
		os.Exit(1)
	}
	results, err := xrel.AddReleaseProofImageByPath(ids, filepath)
	ok(err, "Failed to add proof picture: ")
	count := len(results.ReleaseIDs)
	fmt.Printf("Sucessfully added %s to %d release", results.ProofURL, count)
	if count > 1 {
		fmt.Print("s")
	}
	fmt.Print(".\n")
}
