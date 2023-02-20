package k8s_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" // test framework
	. "github.com/onsi/gomega"    // assertion library
)

// function must to run tests
func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8s Test Suite")
}
