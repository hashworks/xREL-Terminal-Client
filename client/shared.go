package client

import (
	"os"
	"fmt"
	"strings"
	"errors"
	"github.com/hashworks/xRELTerminalClient/api/types"
	"github.com/hashworks/xRELTerminalClient/api"
)

// 2006-01-02 15:04:05.999999999 -0700 MST
const xRELCommentTimeFormat = "02. Jan 2006, 03:04 pm"
const xRELReleaseTimeFormat = "02.01.2006 03:04 pm"

func OK(err error, prefix string) {
	if err != nil {
		fmt.Println(prefix + err.Error())
		os.Exit(1)
	}
}

func findP2PCategoryID(categoryName string) (string, error) {
	var categoryID	string
	var err			error

	var categories []types.P2PCategory
	categories, err = api.P2P_GetCategories()
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
				break;
			}
		}
		if categoryID == "" {
			err = errors.New("Category not found. Please choose one of --categories --p2p.")
		}
	}

	return categoryID, err
}