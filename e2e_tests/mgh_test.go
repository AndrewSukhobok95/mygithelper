package e2e_tests

import (
	"os"
	"os/exec"
	"log"
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "github.com/AndrewSukhobok95/mygithelper/pkg/helpers"
)

var workingDir string
const baseDir = "./../../tmp/"

var _ = ginkgo.Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
        err = os.RemoveAll(baseDir)
        gomega.Expect(err).Should(gomega.BeNil())
    }

	ginkgo.Context("Tests for success cases with one branch", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			repoDirList := []string{"dir-repo-1", "dir-repo-2"}
			for _, dir := range repoDirList {
				repoDirPath, err := helpers.CreateDir(baseDir, dir, true)
				gomega.Expect(err).Should(gomega.BeNil())
				helpers.CreateTxtFileWithContent(repoDirPath, "test.txt", "test")
			}
			workingDir, err = os.Getwd()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		ginkgo.AfterEach(func() {
			os.Chdir(workingDir)
			//removeAll()
		})

		ginkgo.It("Should correctly recursively add files in git repos", func() {
			os.Chdir(baseDir)
			commands := []string{"-m", "addall"}
			_, err := helpers.RunMyGitHelper(commands)
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	ginkgo.Context("Tests for failure cases with one branch", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			repoDirList := []string{"dir-repo-1", "dir-repo-2"}
			for _, dir := range repoDirList {
				repoDirPath, err := helpers.CreateDir(baseDir, dir, false)
				gomega.Expect(err).Should(gomega.BeNil())
				helpers.CreateTxtFileWithContent(repoDirPath, "test.txt", "test")
			}
			workingDir, err = os.Getwd()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		ginkgo.AfterEach(func() {
			os.Chdir(workingDir)
			removeAll()
		})

		ginkgo.It("Should correctly recursively add files in git repos", func() {
			os.Chdir(baseDir)
			commands := []string{"-m", "addall"}
			_, err := helpers.RunMyGitHelper(commands)
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})
	})

	ginkgo.Context("Tests for success cases with many branches", func() {
		ginkgo.BeforeEach(func() {
			var err error
			var commands []string
			//removeAll()
			workingDir, err = os.Getwd()
			gomega.Expect(err).Should(gomega.BeNil())
			repoDirList := []string{"dir-repo-1", "dir-repo-2"}
			branchesList := []string{"master", "b1", "b2"}
			for _, dir := range repoDirList {
				repoDirPath, err := helpers.CreateDir(baseDir, dir, true)
				gomega.Expect(err).Should(gomega.BeNil())
				for _, b := range branchesList {
					os.Chdir(repoDirPath)
					_, err := exec.Command("git", "checkout", "-b", b).CombinedOutput()
					if err != nil {
						log.Fatal(err)
					}
					fileToAdd := "test" + b + ".txt"
					helpers.CreateTxtFileWithContent(repoDirPath, fileToAdd, "test")
					commands = []string{"addall"}
					_, err = helpers.RunMyGitHelper(commands)
					gomega.Expect(err).Should(gomega.BeNil())
					commands = []string{"commit"}
					_, err = helpers.RunMyGitHelper(commands)
					gomega.Expect(err).Should(gomega.BeNil())
					os.Chdir(workingDir)
				}
			}
		})

		ginkgo.AfterEach(func() {
			os.Chdir(workingDir)
			removeAll()
		})

		ginkgo.It("Should correctly clean branches in the repo", func() {
			os.Chdir(baseDir)
			commands := []string{"-m", "clean", "branches"}
			_, err := helpers.RunMyGitHelper(commands)
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})
})

