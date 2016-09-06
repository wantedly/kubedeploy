package main

import (
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

	// get svc
	service := getTargetService(kubeClient, params["service"], params["namespace"])

	// check active
	bluePods, greenPods := getBlueAndGreenPods(kubeClient, service.Name, service.Namespace)
	active := service.Spec.Selector["color"]

	// get new image
	// newImage := get

	// replace standby image
	// if active == "blue" {
	// 	for _, pod := range bluePods {
	// 		replaceImage(pod.Name, pod., newImage)
	// 	}
	// } else if active == "green" {
	// 	printPodsTable(greenPods)
	// }

	// health check

	// chenge blue-green
	// replaceColor(service.Name, active)
}
