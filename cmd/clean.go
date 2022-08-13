package cmd

import (
	"os"

	"github.com/AndrewSukhobok95/mygithelper/pkg/common"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const longDescCleanCmd string = `Delete something from the repository
The target of the cleaning is defined by required sub commands`
const longDescCleanBranchesCmd string = `Delete all the branches in the repo besides the main one
The name of the main branch can be defined in config 
file ($HOME/.mgh/config.yaml) or passin an argument --main-branch`

const descMainBranchFlag string = `name of the branch, that should not be deleted`

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean repository from something described by given parameter",
	Long:  longDescCleanCmd,
}

var cleanBranchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "Delete all the branches in the repo besides the main one",
	Long:  longDescCleanBranchesCmd,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		var mainBranchName string
		mainBranchNameCmdArg := viper.GetString("main-branch")
		mainBranchNameConfig := viper.GetString("main-branch-name")
		if mainBranchNameCmdArg == "" {
			mainBranchName = mainBranchNameConfig
		} else {
			mainBranchName = mainBranchNameCmdArg
		}
		repo_manager.Run(wd, multiRun, repo_manager.MghCleanBranches, mainBranchName)
	},
}

func init() {
	cleanBranchesCmd.Flags().String("main-branch", "", descMainBranchFlag)
	err := viper.BindPFlag("main-branch", cleanBranchesCmd.Flags().Lookup("main-branch"))
	common.Check(err)

	rootCmd.AddCommand(cleanCmd)
	cleanCmd.AddCommand(cleanBranchesCmd)
}
