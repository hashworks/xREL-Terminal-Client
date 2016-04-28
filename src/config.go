package main

import (
	"encoding/json"
	"github.com/hashworks/go-xREL-API/xrel/types"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var configFilePath string

func readConfig(filePath string) error {
	var (
		configData []byte
		err        error
	)

	if filePath == "" {
		configFilePath = getDefaultConfigPath()
	} else {
		configFilePath = filePath
	}
	configData, err = ioutil.ReadFile(configFilePath)
	if err == nil {
		err = json.Unmarshal(configData, &types.Config)
	}

	return err
}

func writeConfig() error {
	err := os.MkdirAll(filepath.Dir(configFilePath), 0700)
	if err == nil {
		var jsonString []byte
		jsonString, err = json.Marshal(types.Config)
		if err == nil {
			err = ioutil.WriteFile(configFilePath, jsonString, 0700)
		}
	}
	return err
}

func getDefaultConfigPath() string {
	var (
		defaultPath string
		separator   string = string(filepath.Separator)
	)

	usr, err := user.Current()
	if err != nil {
		defaultPath = "."
	} else {
		defaultPath = usr.HomeDir + separator + ".config" + separator + "xREL"
	}
	defaultPath += separator + "config.json"

	return defaultPath
}
