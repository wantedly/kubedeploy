package main

import (
	"fmt"
	"os"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func isRunning(pod api.Pod) bool {
	state := pod.Status.ContainerStatuses[0].State
	if state.Running != nil {
		return true
	}
	return false
}

func checkRunning(kubeClient *client.Client, targetPod api.Pod, initialRunning bool) bool {

	count := 60
	bar := pb.StartNew(count)
	for s := 0; s < count; s++ {
		newPod := getTargetPod(kubeClient, targetPod.Name, targetPod.Namespace)
		bar.Increment()
		time.Sleep(time.Second)
		if !isRunning(newPod) && initialRunning {
			return false
		}
	}
	return true
}

func checkHealth(kubeClient *client.Client, targetPod api.Pod) bool {

	// health check
	if isRunning(targetPod) {
		if !checkRunning(kubeClient, targetPod, true) {
			return false
		}
	} else {
		if !checkRunning(kubeClient, targetPod, false) {
			return false
		}
	}

	// last check
	newPod := getTargetPod(kubeClient, targetPod.Name, targetPod.Namespace)
	if !isRunning(newPod) {
		return false
	}

	return true
}

func deploy(kubeClient *client.Client, params map[string]string) {

	// get service
	service := getTargetService(kubeClient, params["service"], params["namespace"])
	if service.Spec.Selector["color"] == "" {
		fmt.Println("blue-green pods don't exist.")
		os.Exit(1)
	}

	// get blue and green pods
	bluePods, _ := getBlueAndGreenPods(kubeClient, service.Name, service.Namespace)

	// get newest master tag
	image := bluePods[0].Spec.Containers[0].Image
	trimedImage := trimImageName(image)
	tagList := getTagList(trimedImage)
	tag := getNewestMasterTag(tagList)
	newImage := QUAYPATH + trimedImage + ":" + tag

	// deploy new image to standby pod
	currentImage := bluePods[0].Spec.Containers[0].Image

	active := service.Spec.Selector["color"]
	if active == "blue" {
		for _, pod := range greenPods {
			replaceImage(pod.Name, currentImage, newImage)
		}
	} else if active == "green" {
		for _, pod := range bluePods {
			replaceImage(pod.Name, currentImage, newImage)
		}
	}

	// health check

	// chenge blue-green
	// replaceColor(service.Name, active)
}
