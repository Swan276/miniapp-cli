package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env       []KeyValuePair `json:"env"`
	Url       []KeyValuePair `json:"url"`
	Renderer  []KeyValuePair `json:"renderer"`
	Variables []KeyValuePair `json:"variables"`
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
	fmt.Println(config)
	return
}

func InitConfig() (config Config, err error) {
	json := `{
		"env": [
			{
				"key": "tnd",
				"value": "main_tnd.dart"
			},
			{
				"key": "sit",
				"value": "main_sit.dart"
			},
			{
				"key": "preprod",
				"value": "main_preprod.dart"
			},
			{
				"key": "prod",
				"value": "main.dart"
			}
		],
		"url": [
			{
				"key": "empty",
				"value": "/"
			}
		],
		"renderer": [
			{
				"key": "canvaskit",
				"value": "canvaskit"
			},
			{
				"key": "html",
				"value": "html"
			}
		],
		"env_variables": [
			{
				"key": "BROWSER_IMAGE_DECODING_ENABLED",
				"value": "false"
			}
		]
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
