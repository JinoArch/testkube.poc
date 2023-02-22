package k8s_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rewards-devops/rewatest/internal/k8s"
)

var podFunc map[string]string

// To run it before suite triggers
var _ = BeforeSuite(func() {
	podFunc = k8s.TestAllPodsRunning()
})

// Container Node
var _ = Describe("TEST: Pods not is Running state", func() {

	// Specs are defined in contexts sections
	Context("In all Namespace", func() {
		It("Should return all pods is Error state", func() {
			By("Get all pods")
			for key0, value0 := range podFunc {
				if value0 == "Error" {
					Expect(key0).To(BeNil())
				}
				//Expect(pods.Items).To(HaveLen(1))
				//Expect(pods.Items[0].Status.Phase).To(Equal("Running"))
			}
		})
	})

	Context("In all Namespace", func() {
		It("Should return all pods in ImagePullBackOff state", func() {
			By("Get all pods")
			for key1, value1 := range podFunc {
				if value1 == "ImagePullBackOff" {
					Expect(key1).To(BeNil())
				}
			}
		})
	})
})

func TestGinkgo(t *testing.T) {
	//  registers a handler, handles Fail function from the ginkgo package with Gomega.
	RegisterFailHandler(Fail)
	// hands over the control to Ginkgo
	RunSpecs(t, "K8s Test Suite")
}
