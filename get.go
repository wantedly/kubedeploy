package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func getTargetPod(kubeClient *client.Client, podName string, namespace string) api.Pod {
	if namespace == "" {
		namespace = "default"
	}
	targetPod, err := kubeClient.Pods(namespace).Get(podName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return *targetPod
}

func getPods(kubeClient *client.Client, namespace string) []api.Pod {
	if namespace == "" {
		namespace = api.NamespaceAll
	}
	pods, err := kubeClient.Pods(namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return pods.Items
}
