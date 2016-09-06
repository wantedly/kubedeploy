package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func replaceImage(pod api.Pod, oldImage, newImage string) {

	commandOptions := []string{"get", "pod", pod.Name, "--namespace=" + pod.Namespace, "-o", "yaml"}
	result := execOutput("kubectl", commandOptions)
	result = strings.Replace(result, oldImage, newImage, -1)

	ioutil.WriteFile("tmp.dat", []byte(result), os.ModePerm)
	defer os.Remove("tmp.dat")

	commandOptions = []string{"replace", "-f", "tmp.dat"}
	_, err := exec.Command("kubectl", commandOptions...).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func replaceColor(serviceName, color string) {
	commandOptions := []string{"get", "svc", serviceName, "-o", "yaml"}
	result := execOutput("kubectl", commandOptions)

	if color == "blue" {
		result = strings.Replace(result, "blue", "green", -1)
	} else {
		result = strings.Replace(result, "green", "blue", -1)
	}

	ioutil.WriteFile("tmp.dat", []byte(result), os.ModePerm)
	defer os.Remove("tmp.dat")

	commandOptions = []string{"replace", "-f", "tmp.dat"}
	_, err := exec.Command("kubectl", commandOptions...).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func replaceBlueGreen(kubeClient *client.Client, targetPod api.Pod, oldImage string, newImage string) {

	printReplace(oldImage, newImage)
	replaceImage(targetPod, oldImage, newImage)

	if check(kubeClient, targetPod) {
		color.Green("Deploy Success!!")
	} else {

		color.Red("Deploy Failed!!")
		color.Red("Revert!!")

		targetPod = getTargetPod(kubeClient, targetPod.Name, targetPod.Namespace)
		printReplace(newImage, oldImage)
		replaceImage(targetPod, newImage, oldImage)

		if check(kubeClient, targetPod) {
			color.Green("Revert Success!!")
		} else {
			color.Red("Revert Failed!!")
			fmt.Println("Check " + targetPod.Name)
			os.Exit(1)
		}
	}

}

func replace(kubeClient *client.Client, params map[string]string) {

	targetPod := getTargetPod(kubeClient, params["pod"], params["namespace"])
	oldImage := targetPod.Spec.Containers[0].Image
	newImage := params["image"]

	printReplace(oldImage, newImage)
	replaceImage(targetPod, oldImage, newImage)

	if check(kubeClient, targetPod) {
		color.Green("Deploy Success!!")
	} else {

		color.Red("Deploy Failed!!")
		color.Red("Revert!!")

		targetPod = getTargetPod(kubeClient, params["pod"], params["namespace"])
		printReplace(newImage, oldImage)
		replaceImage(targetPod, newImage, oldImage)

		if check(kubeClient, targetPod) {
			color.Green("Revert Success!!")
		} else {
			color.Red("Revert Failed!!")
			fmt.Println("Check " + targetPod.Name)
			os.Exit(1)
		}
	}

}
