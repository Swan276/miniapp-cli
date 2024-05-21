/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Swan276/miniapp-cli/usecases"
	"github.com/Swan276/miniapp-cli/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter Web App",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var config utils.Config
		var err error
		config, err = utils.GetConfig()
		if err != nil {
			utils.ErrorMsg(err.Error())
			if utils.PromptInitConfig() {
				config, err = utils.InitConfig()
				if err != nil {
					utils.Abort(fmt.Sprintf("Error Initiating config %v", err))
				}
			} else {
				os.Exit(1)
			}
		}

		environment, err := utils.PromptEnv(config.Env)
		if err != nil {
			return
		}

		buildMode, err := utils.PromptBuildUrl(config.Url)
		if err != nil {
			return
		}

		webRenderer, err := utils.PromptWebRenderer(config.Renderer)
		if err != nil {
			return
		}

		chosenEnvVariables := []string{}
		for _, envVariable := range config.Variables {
			if utils.PromptEnvVariable(envVariable) {
				chosenEnvVariables = append(chosenEnvVariables, fmt.Sprintf("%s=%s", envVariable.Key, envVariable.Value))
			}
		}

		useFVM := utils.PromptFVM()

		usecases.BuildAndCheckFiles(environment, buildMode, webRenderer, chosenEnvVariables, useFVM)

		if utils.PromptVersionUpdate() {
			usecases.UpdateVersion()
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
