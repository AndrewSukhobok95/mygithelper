package cmd

import (
	"os"
	"fmt"
    "github.com/spf13/cobra"
)

var version = "1.2"
var multiRun bool

var longDescRootCmd string = `mgh stands for MyGitHelper - wrapper around git.
It replicates and extends several git commands.
Also it allows to recursively apply git commands
to several repositories.`

var rootCmd = &cobra.Command{
	Use:   "mgh",
	Version: version,
	Short: "mgh stands for MyGitHelper - wrapper around git",
	Long: longDescRootCmd,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
	},
}

var descMultiRunFlag string = `mgh will go through all repositories in the current directory
and apply the command there`

func init() {
	rootCmd.PersistentFlags().BoolVarP(&multiRun, "multi-run", "m", false, descMultiRunFlag)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}