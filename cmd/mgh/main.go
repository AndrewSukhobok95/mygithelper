package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

func checkCommands(commands []string, maxNumCmd int) {
	if len(commands) > maxNumCmd {
		runCommands := strings.Join(commands[:maxNumCmd], " ")
		fmt.Printf("Warning: redundant commands - only %s will be executed.\n", runCommands)
	}
}

func chooseGitCommand(commands []string) (mghFunc func (...string) (string, error)) {
	switch commands[0] {
	case "addall":
		checkCommands(commands, 1)
		mghFunc = repo_manager.MghAddAll
	case "commit":
		checkCommands(commands, 1)
		mghFunc = repo_manager.MghCommit
	case "push":
		checkCommands(commands, 1)
		mghFunc = repo_manager.MghPush
	case "pull":
		checkCommands(commands, 1)
		mghFunc = repo_manager.MghPull
	case "clean":
		checkCommands(commands, 2)
		if commands[1] == "branches" {
			mghFunc = repo_manager.MghCleanBranches
		} else {
			fmt.Printf("Error: uknown sub command %s.\n", commands[1])
        	os.Exit(1)
		}
	default:
		fmt.Printf("Error: uknown command %s.\n", commands[0])
        os.Exit(1)
	}
	return mghFunc
}

func main() {
	wd, _ := os.Getwd()

	multiRun := flag.Bool("m", false, "Go through all 1st level directories and apply the command there.")
	flag.Parse()
	commands := flag.Args()

	if len(commands) == 0 {
        fmt.Println("Usage: mgh [-m] COMMANDS")
        flag.PrintDefaults()
        os.Exit(1)
    }

	DirManager, err := repo_manager.NewDirManager(wd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	mghFunc := chooseGitCommand(commands)
	output, err := DirManager.RunMghFunc(mghFunc, *multiRun)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for repo, message := range output {
		fmt.Printf("[%s output message]:\n", repo)
		fmt.Println(message)
	}
}