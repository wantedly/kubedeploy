package main

import (
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func printInfos(podInfos []string) {

}

func get(kubeClient *client.Client, namespace string) []string {

	if namespace == "" {
		namespace = api.NamespaceAll
	}

	pods, err := kubeClient.Pods(namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	podInfos := []string{}
	for _, pod := range pods.Items {
		podInfoList := []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace}
		podInfos = append(podInfos, strings.Join(podInfoList, ","))
	}

	return podInfos
}
