package k8s

import (
	"context"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Output:
// map[rr-devops-5f7fbd5bdd-stjr4:Error rr-test-554c9fd599-57zn8:ImagePullBackOff rr-test-85568ccdf7-5ts66:ImagePullBackOff]
func TestAllPodsRunning() map[string]string {
	var podState = make(map[string]string)

	// Create a Kubernetes client from the default config file
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatalf("Error loading Kubernetes config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Set the namespace to check
	namespace := ""
	ctx := context.Background()

	// List all pods in the namespace
	//podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	podList, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	// Check the status of each pod
	for _, pod := range podList.Items {
		if pod.Status.Phase == "Failed" || pod.Status.Phase == "Pending" {
			for _, j := range pod.Status.ContainerStatuses {
				if j.State.Waiting != nil {
					podState[pod.Name] = j.State.Waiting.Reason
				}
				if j.State.Terminated != nil {
					if j.State.Terminated.Reason != "Completed" {
						podState[pod.Name] = j.State.Terminated.Reason
					}
				}
			}
		}
	}
	return podState
}
