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

func PromptEnv(envValues []Env) (env string, err error) {
	envNames := []string{}
	for _, env := range envValues {
		envNames = append(envNames, env.Flavor)
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
		if v.Flavor == envName {
			return v.Entry, nil
		}
	}
	Abort(fmt.Sprintf("Select Environment failed %v\n", err))
	return
}

func PromptBuildUrl(buildUrlValues map[string]string) (buildUrl string, err error) {
	buildUrlNames := []string{}
	for key := range buildUrlValues {
		buildUrlNames = append(buildUrlNames, key)
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
	return buildUrlValues[buildModeName], nil
}

func PromptWebRenderer(rendererValues map[string]string) (webRenderer string, err error) {
	rendererNames := []string{}
	for key := range rendererValues {
		rendererNames = append(rendererNames, key)
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
	return rendererValues[webRendererName], nil
}

func PromptCleanAndRebuild() bool {
	cleanAndRebuildPrompt := promptui.Prompt{
		Label:     "Would you like to clean and rebuild",
		IsConfirm: true,
	}

	_, err := cleanAndRebuildPrompt.Run()

	return err == nil
}
