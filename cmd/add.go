package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var longDescAddCmd string = `Add all files in the repo to staging
Equivalent to command git add .`

var addCmd = &cobra.Command{
	Use:   "addall",
	Short: "Add all files in the repo",
	Long:  longDescAddCmd,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghAddAll)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

