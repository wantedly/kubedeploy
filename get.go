package main

import (
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// func getPodStatus(kubeClient *client.Client, pod string, namespace string) string {
// 	pod := getPods(kubeClient, namespace)
//
// }

func getTargetPod(pods []api.Pod, podName string) api.Pod {
	var ret api.Pod
	f := false
	for _, pod := range pods {
		if pod.Name == podName {
			ret = pod
			f = true
		}
	}
	if !f {
		fmt.Println(podName + " doesn't exist.")
		os.Exit(1)
	}
	return ret
}

func getPodInfos(pods []api.Pod) []string {
	podInfos := []string{}
	for _, pod := range pods {
		podInfoList := []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace,
		}
		podInfos = append(podInfos, strings.Join(podInfoList, ","))
	}
	return podInfos
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
