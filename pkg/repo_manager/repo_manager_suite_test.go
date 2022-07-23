package repo_manager_test

import (
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "testing"
)

func TestRepoManager(t *testing.T) {
    gomega.RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "RepoManager Suite")
}