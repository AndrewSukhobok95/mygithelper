package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all files in the repo",
	Long:  `Pull all files in the repo`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghPull)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

