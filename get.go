package main

import (
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func getMatchedPodInfo(pod string, podInfos []string) map[string]string {
	var matchedPodInfo = map[string]string{
		"pod":       "",
		"image":     "",
		"namespace": "",
		"status":    "",
	}
	for _, record := range podInfos {
		names := strings.Split(record, ",")
		if names[0] == pod {
			matchedPodInfo["pod"] = names[0]
			matchedPodInfo["image"] = names[1]
			matchedPodInfo["namespace"] = names[2]
		}
	}
	return matchedPodInfo
}

// func getPodStatus(kubeClient *client.Client, pod string, namespace string) string {
// podInfos := get(kubeClient, namespace)
// myPodInfo := getMatchedPodInfo(pod, podInfos)
// return myPodInfo["status"]
// }

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
