package main

import (
	"fmt"
	"os"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func deploy(kubeClient *client.Client, params map[string]string) {

	// get service
	service := getTargetService(kubeClient, params["service"], params["namespace"])
	if service.Spec.Selector["color"] == "" {
		fmt.Println("blue-green pods don't exist.")
		os.Exit(1)
	}

	// get blue and green pods
	bluePods, greenPods := getBlueAndGreenPods(kubeClient, service.Name, service.Namespace)

	// get newest master tag
	image := bluePods[0].Spec.Containers[0].Image
	trimedImage := trimImageName(image)
	tagList := getTagList(trimedImage)
	tag := getNewestMasterTag(tagList)
	newImage := QUAYPATH + trimedImage + ":" + tag

	// deploy new image to standby pods
	currentImage := bluePods[0].Spec.Containers[0].Image
	active := service.Spec.Selector["color"]
	if active == "blue" {
		for _, pod := range greenPods {
			replaceImage(pod, currentImage, newImage)
		}
	} else if active == "green" {
		for _, pod := range bluePods {
			replaceImage(pod, currentImage, newImage)
		}
	}

	// chenge blue-green
	// replaceColor(service.Name, active)
}
