package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit all files in the repo",
	Long:  `Commit all files in the repo with message automaitc_full_commit`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghCommit)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

