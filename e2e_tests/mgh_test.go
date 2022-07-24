package e2e_tests

import (
	"fmt"
	"os"
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
)

var workingDir string
const baseDir = "./../tmp/"

var _ = ginkgo.Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
        err = os.RemoveAll(baseDir)
        gomega.Expect(err).Should(gomega.BeNil())
    }

	ginkgo.Context("Tests for success cases", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			repoDirList := []string{"dir-repo-1", "dir-repo-2"}
			for _, dir := range repoDirList {
				repoDirPath, err := repo_manager.CreateDir(baseDir, dir, true)
				gomega.Expect(err).Should(gomega.BeNil())
				repo_manager.CreateTxtFileWithContent(repoDirPath, "test.txt", "test")
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
			_, err := repo_manager.RunMyGitHelper(commands)
			fmt.Println(err)
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	ginkgo.Context("Tests for failure cases", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			repoDirList := []string{"dir-repo-1", "dir-repo-2"}
			for _, dir := range repoDirList {
				repoDirPath, err := repo_manager.CreateDir(baseDir, dir, false)
				gomega.Expect(err).Should(gomega.BeNil())
				repo_manager.CreateTxtFileWithContent(repoDirPath, "test.txt", "test")
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
			_, err := repo_manager.RunMyGitHelper(commands)
			fmt.Println(err)
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})
	})
})

