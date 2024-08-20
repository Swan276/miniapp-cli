package usecases

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Swan276/miniapp-cli/utils"
	"github.com/tidwall/gjson"
)

func UpdateVersion() {
	webRoot := "build/web"
	indexPath := fmt.Sprintf("%s/index.html", webRoot)
	jsPath := fmt.Sprintf("%s/flutter.js", webRoot)
	versionPath := fmt.Sprintf("%s/version.json", webRoot)

	// Read version JSON from file
	versionBytes, err := os.ReadFile(versionPath)
	if err != nil {
		utils.Abort(fmt.Sprintf("Error reading version.json: %v", err))
	}

	// Extract version name and build number from JSON
	versionName := gjson.Get(string(versionBytes), "version").String()
	buildNumber := gjson.Get(string(versionBytes), "build_number").String()
	timestamp := fmt.Sprint(time.Now().UnixMilli())
	versionStr := versionName + "b" + buildNumber + "t" + timestamp

	// Read and update flutter.js file
	jsContent, err := os.ReadFile(jsPath)
	if err != nil {
		utils.Abort(fmt.Sprintf("Error reading flutter.js: %v", err))
	}
	jsUpdated := strings.ReplaceAll(string(jsContent), "main.dart.js", fmt.Sprintf("main.dart.js?v=%s", versionStr))

	// Write updated content back to flutter.js file
	fJs, err := os.Create(jsPath)
	if err != nil {
		utils.Abort(fmt.Sprintf("Error writing to flutter.js: %v", err))
	}
	defer fJs.Close()
	fJs.Write([]byte(jsUpdated))
	fJs.Sync()

	// Read and update index.html file
	indexContent, err := os.ReadFile(indexPath)
	if err != nil {
		utils.Abort(fmt.Sprintf("Error reading index.html: %v", err))
	}
	indexUpdated := strings.ReplaceAll(string(indexContent), "flutter.js", fmt.Sprintf("flutter.js?v=%s", versionStr))

	// Write updated content back to index.html file
	fIn, err := os.Create(indexPath)
	if err != nil {
		utils.Abort(fmt.Sprintf("Error writing to index.html: %v", err))
	}
	defer fIn.Close()
	fIn.Write([]byte(indexUpdated))
	fIn.Sync()

	if err != nil {
		utils.Abort(fmt.Sprintf("Error writing to index.html: %v", err))
	}

	utils.Success("Version update completed successfully!")
}
