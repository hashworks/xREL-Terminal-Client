package configHandler

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"github.com/mrjones/oauth"
	"path/filepath"
	"github.com/hashworks/xRELTerminalClient/api/types"
)

type config struct {
	ConfigFilePath			string
	OAuthAccessToken		oauth.AccessToken

	// 24h caching http://www.xrel.to/wiki/6318/api-release-categories.html
	LastCategoryRequest		int64
	Categories				[]types.Category

	// 24h caching http://www.xrel.to/wiki/2996/api-release-filters.html
	LastFilterRequest		int64
	Filters					[]types.Filter

	// 24h caching http://www.xrel.to/wiki/3698/api-p2p-categories.html
	LastP2PCategoryRequest	int64
	P2PCategories			[]types.P2PCategory
}

var instantiatedConfig *config;

func GetConfig(configFilePath string) (*config, error) {
	var err error
	if (instantiatedConfig == nil) {
		if (configFilePath == "") {
			configFilePath = GetDefaultConfigPath();
		}
		var configData []byte
		configData, err = ioutil.ReadFile(configFilePath);
		if err == nil {
			err = json.Unmarshal(configData, &instantiatedConfig)
		} else {
			instantiatedConfig = &config{}
		}
		instantiatedConfig.ConfigFilePath = configFilePath
	}
	return instantiatedConfig, err;
}

func (c config) WriteConfig() error {
	err := os.MkdirAll(filepath.Dir(c.ConfigFilePath), 0700);
	if err == nil {
		var jsonString []byte
		jsonString, err = json.Marshal(c)
		if err == nil {
			err = ioutil.WriteFile(c.ConfigFilePath, jsonString, 0700);
		}
	}
	return err
}

func GetDefaultConfigPath() string {
	var defaultPath string
	seperator := string(filepath.Separator)

	usr, err := user.Current()
	if (err != nil) {
		defaultPath = "."
	} else {
		defaultPath = usr.HomeDir + seperator + ".config" + seperator + "xrel"
	}
	defaultPath += seperator + "config.json"

	return defaultPath;
}