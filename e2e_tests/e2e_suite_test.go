package e2e_tests

import (
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "testing"
)

func TestRepoManager(t *testing.T) {
    gomega.RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "MyGitHelper End to End Test Suite")
}