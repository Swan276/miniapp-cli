package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptInitConfig() bool {
	initConfigPrompt := promptui.Prompt{
		Label:     "Would you like to initialize config file",
		IsConfirm: true,
	}

	_, err := initConfigPrompt.Run()

	return err == nil
}

func PromptEnv(envValues []KeyValuePair) (env string, err error) {
	envNames := []string{}
	for _, env := range envValues {
		envNames = append(envNames, env.Key)
	}

	if len(envNames) == 0 {
		Abort("There's no Environment to select from")
	}

	envPrompt := promptui.Select{
		Label: "Environment",
		Items: envNames,
	}
	_, envName, err := envPrompt.Run()

	if err != nil {
		Abort(fmt.Sprintf("Select Environment failed %v\n", err))
		return "", err
	}

	for _, v := range envValues {
		if v.Key == envName {
			return v.Value, nil
		}
	}
	Abort(fmt.Sprintf("Select Environment failed %v\n", err))
	return
}

func PromptBuildUrl(buildUrlValues []KeyValuePair) (buildUrl string, err error) {
	buildUrlNames := []string{}
	for _, url := range buildUrlValues {
		buildUrlNames = append(buildUrlNames, url.Key)
	}

	if len(buildUrlNames) == 0 {
		Abort("There's no Url to select from")
	}

	buildModePrompt := promptui.Select{
		Label: "Url Build Mode",
		Items: buildUrlNames,
	}
	_, buildModeName, err := buildModePrompt.Run()

	if err != nil {
		Abort(fmt.Sprintf("Select Url Build Mode failed %v\n", err))
		return "", err
	}

	for _, v := range buildUrlValues {
		if v.Key == buildModeName {
			return v.Value, nil
		}
	}
	Abort(fmt.Sprintf("Select Environment failed %v\n", err))
	return
}

func PromptWebRenderer(rendererValues []KeyValuePair) (webRenderer string, err error) {
	rendererNames := []string{}
	for _, renderer := range rendererValues {
		rendererNames = append(rendererNames, renderer.Key)
	}

	if len(rendererNames) == 0 {
		Abort("There's no Renderer to select from")
	}

	webRendererPrompt := promptui.Select{
		Label: "Web Renderer",
		Items: rendererNames,
	}
	_, webRendererName, err := webRendererPrompt.Run()

	if err != nil {
		Abort(fmt.Sprintf("Select Web Renderer failed %v\n", err))
		return "", err
	}

	for _, v := range rendererValues {
		if v.Key == webRendererName {
			return v.Value, nil
		}
	}
	Abort(fmt.Sprintf("Select Environment failed %v\n", err))
	return
}

func PromptEnvVariable(envVariable KeyValuePair) bool {
	useEnvVariablePrompt := promptui.Prompt{
		Label:     fmt.Sprintf("Use Env Variable: %s=%s", envVariable.Key, envVariable.Value),
		IsConfirm: true,
	}
	_, err := useEnvVariablePrompt.Run()

	return err == nil
}

func PromptFVM() bool {
	useFvmPrompt := promptui.Prompt{
		Label:     "Use FVM: ",
		IsConfirm: true,
	}
	_, err := useFvmPrompt.Run()

	return err == nil
}

func PromptCleanAndRebuild() bool {
	cleanAndRebuildPrompt := promptui.Prompt{
		Label:     "Would you like to clean and rebuild",
		IsConfirm: true,
	}

	_, err := cleanAndRebuildPrompt.Run()

	return err == nil
}

func PromptVersionUpdate() bool {
	versionUpdatePrompt := promptui.Prompt{
		Label:     "Would you update version for flutter.js",
		IsConfirm: true,
	}

	_, err := versionUpdatePrompt.Run()

	return err == nil
}
