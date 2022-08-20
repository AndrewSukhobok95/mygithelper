package repo_manager_test

import (
	"os"
	"path"

	"github.com/AndrewSukhobok95/mygithelper/pkg/helpers"
	"github.com/AndrewSukhobok95/mygithelper/pkg/repo_manager"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var workingDir string

const baseDir = "./../../../tmp/"
const repoDir = "dir-repo"
const nonRepoDir = "dir-non-repo"

var _ = ginkgo.Describe("mghfuncs tests", func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		gomega.Expect(err).Should(gomega.BeNil())
	}

	ginkgo.Context("Tests for success cases", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			repoDirPath, err := helpers.CreateDir(baseDir, repoDir, true)
			gomega.Expect(err).Should(gomega.BeNil())
			helpers.CreateTxtFileWithContent(repoDirPath, "test.txt", "test")
			workingDir, err = os.Getwd()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		ginkgo.AfterEach(func() {
			os.Chdir(workingDir)
			removeAll()
		})

		ginkgo.It("Should correctly check if the folder is git repo", func() {
			dirPath := path.Join(baseDir, repoDir)
			isRepoFlag, err := repo_manager.IsGitRepo(dirPath)
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(isRepoFlag).Should(gomega.BeTrue())
		})

		ginkgo.It("Should correctly add files if this is git repo", func() {
			dirPath := path.Join(baseDir, repoDir)
			os.Chdir(dirPath)
			_, err = repo_manager.MghAddAll()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		ginkgo.It("Should correctly add and commit files if this is git repo", func() {
			dirPath := path.Join(baseDir, repoDir)
			os.Chdir(dirPath)
			_, err = repo_manager.MghAddAll()
			gomega.Expect(err).Should(gomega.BeNil())
			_, err = repo_manager.MghCommit()
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	ginkgo.Context("Tests for failure cases", func() {
		ginkgo.BeforeEach(func() {
			removeAll()
			nonRepoDirPath, err := helpers.CreateDir(baseDir, nonRepoDir, false)
			gomega.Expect(err).Should(gomega.BeNil())
			helpers.CreateTxtFileWithContent(nonRepoDirPath, "test.txt", "test")
			workingDir, err = os.Getwd()
			gomega.Expect(err).Should(gomega.BeNil())
		})

		ginkgo.AfterEach(func() {
			os.Chdir(workingDir)
			removeAll()
		})

		ginkgo.It("Should correctly check if the folder is not a git repo", func() {
			dirPath := path.Join(baseDir, nonRepoDir)
			isRepoFlag, err := repo_manager.IsGitRepo(dirPath)
			gomega.Expect(err).ShouldNot(gomega.BeNil())
			gomega.Expect(isRepoFlag).Should(gomega.BeFalse())
		})

		ginkgo.It("Should fail to add files if this is not a git repo", func() {
			dirPath := path.Join(baseDir, nonRepoDir)
			os.Chdir(dirPath)
			_, err = repo_manager.MghAddAll()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})

		ginkgo.It("Should fail to add and commit files if this is not a git repo", func() {
			dirPath := path.Join(baseDir, repoDir)
			os.Chdir(dirPath)
			_, err = repo_manager.MghCommit()
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})
	})
})
