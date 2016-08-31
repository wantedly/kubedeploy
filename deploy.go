package main

import (
	"fmt"
	"os"
	"strings"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func getMatchedPodInfo(pod string, podInfos []string) map[string]string {

	var matchedPodInfo = map[string]string{
		"pod":       "",
		"image":     "",
		"namespace": "",
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

func deploy(kubeClient *client.Client, params map[string]string) {

	podInfos := get(kubeClient)

	infoMap := getMatchedPodInfo(params["pod"], podInfos)
	if infoMap["pod"] == "" {
		fmt.Println(params["pod"] + " doesn't exist.")
		os.Exit(1)
	}

}
