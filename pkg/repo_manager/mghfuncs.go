package repo_manager

import (
	"os"
	"os/exec"
)

func IsGitRepo(path string) (bool, error) {
	_, err := os.Stat(path + "/.git")
	if err != nil {
		return false, err
	}
	return true, nil
}

func gitInit(args ...string) (string, error) {
	out, err := exec.Command("git", "init").CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}

func MghAddAll(args ...string) (string, error) {
	out, err := exec.Command("git", "add", ".").CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}

func MghCommit(args ...string) (string, error) {
	var commitMsg string
	if len(args)==0 {
		commitMsg = "-"
	} else {
		commitMsg = "\"" + args[0] + "\""
	}
	out, err := exec.Command("git", "commit", "-m", commitMsg).CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}

func MghPush(args ...string) (string, error) {
	out, err := exec.Command("git", "push").CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}

func MghPull(args ...string) (string, error) {
	out, err := exec.Command("git", "pull").CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}

func MghCleanBranches(args ...string) (string, error) {
	var mainBranchName string
	if len(args)==0 {
		mainBranchName = "master"
	} else {
		mainBranchName = args[0]
	}
	_, err := exec.Command("git", "checkout", mainBranchName).CombinedOutput()
	if err != nil {
		return "", err
	}
	bashCmd := "git branch | grep -v " + mainBranchName + " | xargs git branch -D"
	out, err := exec.Command("bash", "-c", bashCmd).Output()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}


