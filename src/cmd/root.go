/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mulrepo",
	Short: "mulrepo is tool for multi git repo project management",
	Long: `mulrepo commands and sub commands:

			   mulrepo was meant for managing multi git repo project by using more complex git commands
			   mulrepo was meant for managing commit on multi git repo project 

			   mulrepo config:
					mulrepo config   [ import / -i ] <path to config json file>     (import config from json file)
					mulrepo config   [ reload / -r ]     (will work if mulrepo already ran and has configuration file path set as default)
					mulrepo config   [ export / -e ] <path to export to>      (default is currnt path)
					mulrepo config   [ list / -l ]     (shows list of repositories and metadata specified in config)
					mulrepo config   [ seek / -s]     (shows default configuration file location)
			   		mulrepo config   [ interactive / -I ]     (interactive mode of config creation)
					mulrepo config   [ template / -t ]     (show template for repos json file)

				mulrepo repo:
					NOTE: modifying with the repo subcommand will always be interactive!!

					mulrepo repo   [ delete / remove / -rm ] <repo_name>    (remove repo from config by name)
					mulrepo repo   [ update / edit / -e ] <repo_name>    (update repo config by name)
					mulrepo repo   [ add / -a ]     (add new repository to config)

				mulrepo commit  ( [ on / -o ] <repo_name> )   (iterates on repos and commits them, or commit on specified repo)
				mulrepo push ( [ on / -o ]  <repo_name> )   (iterates on included repos and pushes them, or push on specified repo)
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mulrepo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
