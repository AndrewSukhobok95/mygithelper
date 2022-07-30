package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
	"github.com/AndrewSukhobok95/mygithelper/pkg/common"
)

const descMsgFlag string = `message for commit command`

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit all files in the repo",
	Long:  `Commit all files in the repo with default commit message`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghCommit, viper.GetString("msg"))
	},
}

func init() {
	commitCmd.Flags().String(
		"msg", 
		viper.GetString("default-commit-message"), 
		descMsgFlag)
	err := viper.BindPFlag("msg", commitCmd.Flags().Lookup("msg"))
	common.Check(err)
	rootCmd.AddCommand(commitCmd)
}

