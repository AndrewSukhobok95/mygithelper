package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var longDescCleanCmd string = `Delete something from the repository
The target of the cleaning is defined by required sub commands`

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean repository from something described by given parameter",
	Long:  longDescCleanCmd,
}

var cleanBranchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "Delete all the branches in the repo besides the master",
	Long:  `Delete all the branches in the repo besides the master`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo_manager.Run(wd, multiRun, repo_manager.MghCleanBranches)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.AddCommand(cleanBranchesCmd)
}

