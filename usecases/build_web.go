package usecases

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Swan276/miniapp-cli/utils"
	"github.com/briandowns/spinner"
)

func BuildAndCheckFiles(env string, buildMode string, renderer string, envVariables []string, useFVM bool) {
	buildWeb(env, buildMode, renderer, envVariables, useFVM)
	if fileMissingError := checkFiles(); fileMissingError != nil {
		utils.ErrorMsg(fileMissingError.Error())

		cleanAndRebuild := utils.PromptCleanAndRebuild()

		if cleanAndRebuild {
			if err := cleanWorkspace(); err != nil {
				utils.Abort(err.Error())
			}
			BuildAndCheckFiles(env, buildMode, renderer, envVariables, useFVM)
		} else {
			utils.Info("You can run `flutter clean && flutter pub get` to clean the workspace")
			utils.Abort("Exiting...")
		}
	}
}

func buildWeb(env string, buildMode string, renderer string, envVariables []string, useFVM bool) {
	var definedEnvVariables = ""
	for _, v := range envVariables {
		definedEnvVariables += fmt.Sprintf("--dart-define %s ", v)
	}
	definedEnvVariables = strings.Trim(definedEnvVariables, " ")
	buildCommand := fmt.Sprintf("%sflutter build web -t lib/%s --web-renderer %s --base-href %s %s --release", _fvm(useFVM), strings.ToLower(env), renderer, buildMode, definedEnvVariables)
	buildCommandArgs := utils.ParseCommand(buildCommand)
	fmt.Println(buildCommand)

	if err := utils.RunCommand(buildCommandArgs); err != nil {
		utils.Abort(err.Error())
	}
}

func cleanWorkspace() error {
	cleanCommand := "flutter clean && flutter pub get"
	cleanCommandArgs := utils.ParseCommand(cleanCommand)

	if err := utils.RunCommand(cleanCommandArgs); err != nil {
		return err
	}
	return nil
}

func checkFiles() error {
	utils.Info("Checking Files ...")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	errFlutterJs := checkFlutterJs()
	if errFlutterJs != nil {
		return errFlutterJs
	}

	errIndexHtml := checkIndexHtml()
	if errIndexHtml != nil {
		return errIndexHtml
	}

	s.Stop()

	utils.Success("No Error Found")

	return nil
}

func checkFlutterJs() error {
	path := "build/web/flutter.js"
	if errFlutterJs := checkFile(path); errFlutterJs != nil {
		if errFlutterJsUpper := checkFile(path); errFlutterJsUpper != nil {
			return fmt.Errorf("flutter.js file is missing")
		}
	}
	return nil
}

func checkIndexHtml() error {
	path := "build/web/index.html"
	if errIndexHtml := checkFile(path); errIndexHtml != nil {
		if errIndexHtmlUpper := checkFile(path); errIndexHtmlUpper != nil {
			return fmt.Errorf("index.html file is missing")
		}
	}
	return nil
}

func checkFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func _fvm(useFVM bool) string {
	if useFVM {
		return "fvm "
	}
	return ""
}
