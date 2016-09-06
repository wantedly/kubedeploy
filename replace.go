package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func replaceImage(podName, oldImage, newImage string) {
	printDeploy(oldImage, newImage)
	commandOptions := []string{"get", "pod", podName, "-o", "yaml"}
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

func replaceColor(service string) {
	commandOptions := []string{"get", "svc", service, "-o", "yaml"}
	result := execOutput("kubectl", commandOptions)

	if service.Spec.Selector["color"] == "blue" {
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

func replace(kubeClient *client.Client, params map[string]string) {
	targetPod := getTargetPod(kubeClient, params["pod"], params["namespace"])
	replaceImage(targetPod.Name, targetPod.Spec.Containers[0].Image, params["image"])

	if checkHealth(kubeClient, targetPod) {
		color.Green("Success!!")
	} else {
		color.Red("Failed!!")
		fmt.Println("Revert Start")
		replaceImage(targetPod.Name, params["image"], targetPod.Spec.Containers[0].Image)
		fmt.Println("Revert End")
	}
}
