package main

import (
	"github.com/koople/flag-references/src/application"
	"github.com/koople/flag-references/src/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var apiKey string
var projectPath string
var baseUri string
var repository string
var v string

func main() {
	var rootCmd = &cobra.Command{Use: "kpl search references"}
	rootCmd.AddCommand(searchReferencesCmd)
	rootCmd.PersistentFlags().StringVar(&apiKey, "apiKey", "", "Server api key.")
	err := rootCmd.MarkPersistentFlagRequired("apiKey")
	if err != nil {
		panic("Missing api key.")
	}

	rootCmd.PersistentFlags().StringVarP(&repository, "repository", "r", "", "Repository name.")
	err = rootCmd.MarkPersistentFlagRequired("repository")
	if err != nil {
		panic("Missing repository name.")
	}

	rootCmd.PersistentFlags().StringVarP(&projectPath, "projectPath", "p", ".", "Project path.")
	rootCmd.PersistentFlags().StringVar(&baseUri, "baseUri", "https://sdk.koople.io", "Base uri.")
	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", logrus.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic).")

	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var searchReferencesCmd = &cobra.Command{
	Use:   "search-refs",
	Short: "Search flag references in the provided project.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := log.SetUpClientLog(os.Stdout, v); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := application.Run(repository, projectPath, apiKey, baseUri, log.ClientLog)
		if err != nil {
			log.ClientLog.Error(err)
		}

		return nil
	},
}
