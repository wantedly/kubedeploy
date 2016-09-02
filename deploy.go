package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func isRunning(kubeClient *client.Client, pod string, namespace string) {
	// get info
	// status := getPodStatus(kubeClient, pod, namespace)
	// check status

	// return
}

func healthCheck(kubeClient *client.Client, targetPod api.Pod) bool {
	count := 60
	bar := pb.StartNew(count)
	for s := 0; s < count; s++ {
		bar.Increment()
		time.Sleep(time.Second)
		if s%10 == 0 {
			newTargetPod := getTargetPod(kubeClient, targetPod.Name, targetPod.Namespace)
			state := newTargetPod.Status.ContainerStatuses[0].State
			if state.Running == nil {
				return false
			}
		}
	}
	return true
}

func replaceImage(podName, oldImage, newImage string) {

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

func deploy(kubeClient *client.Client, params map[string]string) {

	targetPod := getTargetPod(kubeClient, params["pod"], params["namespace"])
	printDeploy(targetPod.Spec.Containers[0].Image, params["image"])
	replaceImage(targetPod.Name, targetPod.Spec.Containers[0].Image, params["image"])

	if healthCheck(kubeClient, targetPod) {
		fmt.Println("Success!!")
	} else {
		fmt.Println("Failed!!")
		fmt.Println("Revert Start")
		printDeploy(params["image"], targetPod.Spec.Containers[0].Image)
		replaceImage(targetPod.Name, params["image"], targetPod.Spec.Containers[0].Image)
		fmt.Println("Revert End")
	}
}
