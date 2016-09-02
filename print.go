package main

import (
	"fmt"
	"strings"

	"k8s.io/kubernetes/pkg/api"
)

func printPodCSV(pods []api.Pod) {
	podInfos := []string{}
	for _, pod := range pods {
		podInfoList := []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace,
		}
		podInfos = append(podInfos, strings.Join(podInfoList, ","))
	}
	for _, info := range podInfos {
		fmt.Println(info)
	}
}

func printDeploy(oldImage, newImage string) {
	fmt.Println("Deploy: " + oldImage + " => " + newImage)
}
