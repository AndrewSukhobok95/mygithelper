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
	out, err := exec.Command("git", "commit", "-m", "automaitc_full_commit").CombinedOutput()
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
	out, err := exec.Command(
		"git", "branch", "|", 
		"grep", "-v", "master", "|", 
		"xargs", "git", "branch", "-d").CombinedOutput()
	if err != nil {
		return "", err
	}
	output := string(out)
	return output, nil
}


