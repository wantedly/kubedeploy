package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func replaceImage(kubeClient *client.Client, pod api.Pod, newImage string) {
	pod.Spec.Containers[0].Image = newImage
	_, err := kubeClient.Pods(pod.Namespace).Update(&pod)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func replaceColor(kubeClient *client.Client, service api.Service) {

	var newColor = ""
	if service.Spec.Selector["color"] == "blue" {
		fmt.Println("Change: blue => green")
		newColor = "green"
	} else if service.Spec.Selector["color"] == "green" {
		fmt.Println("Change: green => blue")
		newColor = "blue"
	}

	service.Spec.Selector["color"] = newColor
	_, err := kubeClient.Services(service.Namespace).Update(&service)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func replace(kubeClient *client.Client, params map[string]string) {

	targetPod := getTargetPod(kubeClient, params["pod"], params["namespace"])
	oldImage := targetPod.Spec.Containers[0].Image
	newImage := params["image"]
	success := false

	// replace image
	printReplace(oldImage, newImage)
	replaceImage(kubeClient, targetPod, newImage)

	// health check
	if check(kubeClient, targetPod) {
		color.Green("Deploy Success!!")
		success = true
	} else {
		color.Red("Deploy Failed!!")
	}

	// revert image
	if !success {
		color.Red("Revert!!")

		targetPod = getTargetPod(kubeClient, params["pod"], params["namespace"])
		printReplace(newImage, oldImage)
		replaceImage(kubeClient, targetPod, oldImage)

		if check(kubeClient, targetPod) {
			color.Green("Revert Success!!")
		} else {
			color.Red("Revert Failed!!")
			fmt.Println("Check " + targetPod.Name)
			os.Exit(1)
		}
	}

}
