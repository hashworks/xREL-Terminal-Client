package main

import (
	"fmt"
	"github.com/hashworks/xRELTerminalClient/src/xrel"
	"os"
	"regexp"
	"sort"
	"strings"
)

func showCategories(isP2P bool) {
	orderedCategories := map[string][]string{}
	if isP2P {
		p2pCategories, err := xrel.GetP2PCategories()
		ok(err, "Failed to get p2p categories:\n")
		for i := 0; i < len(p2pCategories); i++ {
			metaCat := strings.ToUpper(p2pCategories[i].MetaCat)
			if p2pCategories[i].SubCat != "" {
				orderedCategories[metaCat] = append(orderedCategories[metaCat], p2pCategories[i].SubCat)
			} else {
				if len(orderedCategories[metaCat]) == 0 {
					orderedCategories[metaCat] = []string{}
				}
			}
		}
		fmt.Println("Available p2p categories:")
	} else {
		categories, err := xrel.GetReleaseCategories()
		ok(err, "Failed to get scene categories:\n")
		for i := 0; i < len(categories); i++ {
			category := &categories[i]
			if category.ParentCat != "" && category.Name != "" {
				orderedCategories[category.ParentCat] = append(orderedCategories[category.ParentCat], category.Name)
			} else {
				if len(orderedCategories[category.Name]) == 0 {
					orderedCategories[category.Name] = []string{}
				}
			}
		}
		fmt.Println("Available scene categories:")
	}
	var keys []string
	for k := range orderedCategories {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := orderedCategories[k]
		fmt.Println("\n" + k)
		for i := 0; i < len(v); i++ {
			fmt.Println("\t" + v[i])
		}
	}
}

func browseCategory(categoryName, extInfoType string, isP2P bool, perPage, page int) {
	if isP2P {
		categoryID, err := findP2PCategoryID(categoryName)
		ok(err, "Failed to get category id:\n")
		data, err := xrel.GetP2PReleases(perPage, page, categoryID, "", "")
		ok(err, "Failed to browse p2p category:\n")
		printP2PReleases(data, true, false)
	} else {
		// Currently all categories are upper case. That might change?
		categoryName = strings.ToUpper(categoryName)
		data, err := xrel.BrowseReleaseCategory(categoryName, extInfoType, perPage, page)
		ok(err, "Failed to browse scene category:\n")
		printSceneReleases(data, true, false)
	}
}

func showFilters(isP2PFlag bool) {
	if isP2PFlag {
		fmt.Println("There are no P2P filters available.")
		os.Exit(1)
	}
	filters, err := xrel.GetReleaseFilters()
	ok(err, "Failed to get filters:\n")
	fmt.Println("Available scene filters:\n")
	for i := 0; i < len(filters); i++ {
		fmt.Println(filters[i].Id + ": " + filters[i].Name)
	}
}

func showLatest(filterFlag string, isP2PFlag bool, perPageFlag, pageFlag int) {
	if isP2PFlag {
		data, err := xrel.GetP2PReleases(perPageFlag, pageFlag, "", "", "")
		ok(err, "Failed to get latest p2p releases:\n")
		printP2PReleases(data, false, false)
	} else {
		data, err := xrel.GetLatestReleases(perPageFlag, pageFlag, filterFlag, "")
		ok(err, "Failed to get latest scene releases:\n")
		printSceneReleases(data, false, false)
	}
}

func browseArchive(browseArchiveFlag, filterFlag string, isP2PFlag bool, perPageFlag, pageFlag int) {
	if isP2PFlag {
		fmt.Println("Due to API limitations it is impossible to browse the P2P archive.")
		os.Exit(1)
	} else {
		matched, err := regexp.MatchString("^[0-9]{4}-[1-9]{2}$", browseArchiveFlag)
		if err == nil && matched {
			data, err := xrel.GetLatestReleases(perPageFlag, pageFlag, filterFlag, browseArchiveFlag)
			ok(err, "Failed to browse the scene archive:\n")
			printSceneReleases(data, false, false)
		} else {
			fmt.Println("Please use the following format: --browseArchive=YYYY-MM")
		}
	}
}
