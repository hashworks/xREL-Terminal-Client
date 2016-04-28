package main

import (
	"errors"
	"fmt"
	"github.com/hashworks/go-xREL-API/xrel"
	"os"
	"strconv"
	"strings"
)

func selectFavList(selectPrefix string) (int, error) {
	var (
		id  int
		err error
	)

	if selectPrefix == "" {
		selectPrefix = "Please choose one: "
	}

	favLists, err := xrel.GetFavsLists()
	if err == nil {
		favListCount := len(favLists)
		if favListCount == 0 {
			err = errors.New("You have no favorites lists.")
		} else if favListCount > 1 {
			for i := 0; i < favListCount; i++ {
				favList := favLists[i]
				if favList.Description != "" {
					fmt.Printf("(%d) %s (%s)\n", i+1, favList.Name, favList.Description)
				} else {
					fmt.Printf("(%d) %s\n", i+1, favList.Name)
				}
			}
			selection := -1
			fmt.Print(selectPrefix)
			for selection < 1 || selection > favListCount {
				fmt.Scanln(&selection)
			}
			fmt.Println()
			id = favLists[selection-1].ID
		} else {
			id = favLists[0].ID
		}
	}
	return id, err
}

func addEntryToFavList(extInfoId string) {
	id, err := selectFavList("Please choose the list you want to add an entry to: ")
	ok(err, "Failed to get your favorites lists: ")
	result, err := xrel.AddFavsListEntry(strconv.Itoa(id), extInfoId)
	ok(err, "Failed to add entry: ")
	fmt.Println("Sucessfully added \"" + result.ExtInfo.Title + "\".")
}

func removeFavEntry() {
	id, err := selectFavList("Please choose the list you want to remove an entry from: ")
	ok(err, "Failed to get your favorites lists: ")
	favListEntries, err := xrel.GetFavsListEntries(strconv.Itoa(id), false)
	ok(err, "Failed to get favorites list entries: ")
	favListEntriesCount := len(favListEntries)
	if favListEntriesCount == 0 {
		fmt.Println("You have no favorites list entries on this list.")
		os.Exit(1)
	} else {
		var entryID string
		if favListEntriesCount == 1 {
			var selection string
			fmt.Print("Do you really want to remove \"" + favListEntries[0].Title + "\"? (y/N) ")
			fmt.Scanln(&selection)
			if selection == "y" {
				entryID = favListEntries[0].ID
			} else {
				os.Exit(0)
			}
		} else {
			for i := 0; i < favListEntriesCount; i++ {
				favListEntry := favListEntries[i]
				fmt.Printf("(%d) [%s] %s\n", i+1, strings.ToUpper(favListEntry.Type), favListEntry.Title)
			}
			selection := -1
			fmt.Print("Please choose the entry you want to remove: ")
			for selection < 1 || selection > favListEntriesCount {
				fmt.Scanln(&selection)
			}
			entryID = favListEntries[selection-1].ID
		}
		fmt.Println()
		result, err := xrel.RemoveFavsListEntry(strconv.Itoa(id), entryID)
		ok(err, "Failed to remove entry: ")
		fmt.Println("Sucessfully removed \"" + result.ExtInfo.Title + "\".")
	}
}

func showFavEntries(markAsRead bool) {
	id, err := selectFavList("")
	ok(err, "Failed to get your favorites lists: ")
	idStr := strconv.Itoa(id)
	favListEntries, err := xrel.GetFavsListEntries(idStr, true)
	ok(err, "Failed to get favorites list entries: ")
	favListEntriesCount := len(favListEntries)
	if favListEntriesCount == 0 {
		fmt.Println("You have no favorites list entries on this list.")
		os.Exit(1)
	} else {
		for i := 0; i < favListEntriesCount; i++ {
			entry := favListEntries[i]
			fmt.Println(entry.Title + " [" + strings.ToUpper(entry.Type) + "]")
			releaseCount := len(entry.Releases)
			p2pReleaseCount := len(entry.P2PReleases)
			if releaseCount == 0 && p2pReleaseCount == 0 {
				fmt.Println("\tNo new releases.")
			} else {
				for i := 0; i < releaseCount; i++ {
					release := entry.Releases[i]
					if markAsRead {
						go xrel.MarkFavsListEntryAsRead(idStr, release.ID, true)
					}
					fmt.Println("\t[SCENE] " + release.Dirname + " (" + release.LinkURL + ")")
				}
				for i := 0; i < p2pReleaseCount; i++ {
					release := entry.P2PReleases[i]
					if markAsRead {
						go xrel.MarkFavsListEntryAsRead(idStr, release.ID, false)
					}
					fmt.Println("\t[P2P]   " + release.Dirname + " (" + release.LinkURL + ")")
				}
			}
			fmt.Println()
		}
	}
}
