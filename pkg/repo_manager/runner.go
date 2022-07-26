package repo_manager

import (
	"fmt"
	"os"
)

func Run(
	baseDir string, 
	multiRun bool, 
	mghFunc func (...string) (string, error),
	mghFuncArgs ...string) {
	DirManager, err := NewDirManager(baseDir)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	output, err := DirManager.RunMghFunc(mghFunc, multiRun)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for repo, message := range output {
		fmt.Printf("[%s output message]:\n", repo)
		fmt.Println(message)
	}
}

