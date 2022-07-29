package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var defaultCommitMessage string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit all files in the repo",
	Long:  `Commit all files in the repo with default commit message`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghCommit, defaultCommitMessage)
	},
}

const descMsgFlag string = `message for commit command`

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVar(&defaultCommitMessage, "msg", "automatic full commit", descMsgFlag)
}

