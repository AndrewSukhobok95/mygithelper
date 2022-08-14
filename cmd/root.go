package cmd

import (
	"os"
	"path"

	"github.com/AndrewSukhobok95/mygithelper/pkg/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var multiRun bool
var configPath string

const longDescRootCmd string = `mgh stands for MyGitHelper - wrapper around git.
It replicates and extends several git commands.
Also it allows to recursively apply git commands
to several repositories.`

const descMultiRunFlag string = `mgh will go through all repositories in the current directory
and apply the command there`

var rootCmd = &cobra.Command{
	Use:   "mgh",
	Short: "mgh stands for MyGitHelper - wrapper around git",
	Long:  longDescRootCmd,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	err := viper.BindEnv("home")
	common.Check(err)
	home := viper.GetString("home")
	configPath = path.Join(home, ".mgh/config.yaml")

	rootCmd.PersistentFlags().BoolVarP(&multiRun, "multi-run", "m", false, descMultiRunFlag)
}

func initConfig() {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		common.Check(err)
	}
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	common.Check(err)
}

func Execute() {
	initConfig()
	err := rootCmd.Execute()
	common.Check(err)
}
