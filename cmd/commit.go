package cmd

import (
	"os"

	"github.com/AndrewSukhobok95/mygithelper/pkg/common"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const descMsgFlag string = `message for commit command`

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit all files in the repo",
	Long:  `Commit all files in the repo with default commit message`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		var commitMessage string
		commitMessageCmdArg := viper.GetString("msg")
		commitMessageConfig := viper.GetString("default-commit-message")
		if commitMessageCmdArg == "" {
			commitMessage = commitMessageConfig
		} else {
			commitMessage = commitMessageCmdArg
		}
		repo_manager.Run(wd, multiRun, repo_manager.MghCommit, commitMessage)
	},
}

func init() {
	commitCmd.Flags().String("msg", "", descMsgFlag)
	err := viper.BindPFlag("msg", commitCmd.Flags().Lookup("msg"))
	common.Check(err)

	rootCmd.AddCommand(commitCmd)
}
