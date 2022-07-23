package repo_manager

import (
	"os"
	"os/exec"
	"path"
	"log"
)

func CreateDir(baseDir string, name string, initGit bool) (dirPath string, err error) {
	dirPath = path.Join(baseDir, name)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}

	if !initGit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(currDir)

	os.Chdir(dirPath)
	err = exec.Command("git", "init").Run()
	return
}

func CreateTxtFileWithContent(dirPath, fileName, content string) {
	fullPath := path.Join(dirPath, fileName)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
        file, err := os.Create(fullPath)
        if err != nil {
            log.Fatal(err)
        }
		_, err = file.WriteString(content)
		if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    }
}
