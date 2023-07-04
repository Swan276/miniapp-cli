package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env      []Env             `json:"env"`
	Url      map[string]string `json:"url"`
	Renderer map[string]string `json:"renderer"`
}

type Env struct {
	Flavor string `json:"flavor"`
	Entry  string `json:"entry"`
}

func GetConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
	return
}

func InitConfig() (config Config, err error) {
	json := `{
		"env": [
			{
				"flavor": "tnd",
				"entry": "main_tnd.dart"
			},
			{
				"flavor": "sit",
				"entry": "main_sit.dart"
			},
			{
				"flavor": "preprod",
				"entry": "main_preprod.dart"
			},
			{
				"flavor": "prod",
				"entry": "main.dart"
			}
		],
		"url": {
			"empty": "/"
		},
		"renderer": {
			"canvaskit": "canvaskit",
			"html": "html"
		}
	}`

	f, err := os.Create(".miniapp-cli.json")
	if err != nil {
		Abort(fmt.Sprintf("Error creating config file %v", err))
	}
	defer f.Close()
	f.Write([]byte(json))
	f.Sync()

	viper.ReadInConfig()

	return
}
