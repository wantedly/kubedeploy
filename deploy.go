package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func bgDeploy(kubeClient *client.Client, params map[string]string) {

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
	active := service.Spec.Selector["color"]
	var standby string
	var replacePods []api.Pod
	if active == "blue" {
		standby = "green"
		replacePods = greenPods
	} else if active == "green" {
		standby = "blue"
		replacePods = bluePods
	}
	for _, pod := range replacePods {
		replaceParams := map[string]string{
			"pod":       pod.Name,
			"image":     newImage,
			"namespace": pod.Namespace,
		}
		replace(kubeClient, replaceParams)
	}

	// chenge blue-green
	replaceColor(kubeClient, service)

	// check color
	service = getTargetService(kubeClient, params["service"], params["namespace"])
	if service.Spec.Selector["color"] == standby {
		color.Green("Blue-Green Deploy Success!!")
	} else {
		color.Red("Blue-Green Deploy Falied!!")
	}

}

func oneDeploy(kubeClient *client.Client, params map[string]string) {

	// get service
	service := getTargetService(kubeClient, params["service"], params["namespace"])
	if service.Spec.Selector["color"] == "" {
		fmt.Println("blue-green pods don't exist.")
		os.Exit(1)
	}

	replacePods := getPodsWithService(kubeClient, service.Name, service.Namespace)

	// get newest master tag
	image := replacePods[0].Spec.Containers[0].Image
	trimedImage := trimImageName(image)
	tagList := getTagList(trimedImage)
	tag := getNewestMasterTag(tagList)
	newImage := QUAYPATH + trimedImage + ":" + tag

	for _, pod := range replacePods {
		replaceParams := map[string]string{
			"pod":       pod.Name,
			"image":     newImage,
			"namespace": pod.Namespace,
		}
		replace(kubeClient, replaceParams)
	}
}
