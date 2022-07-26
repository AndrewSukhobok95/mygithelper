package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push all files in the repo",
	Long:  `Push all files in the repo`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghPush)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

