package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"

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

func deployBG(kubeClient *client.Client, params map[string]string) {

	pod := getTargetPod(kubeClient, params["pod"], params["namespace"])

	// get svc
	svc := getTargetService(kubeClient, pod.Namespace)
	fmt.Println(svc)
	// check active

	// get the other pod

	// replace standby image

	// health check

	// chenge blue-green

}

func deploy(kubeClient *client.Client, params map[string]string) {

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
