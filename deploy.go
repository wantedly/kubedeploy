package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func getMatchedPodInfo(pod string, podInfos []string) map[string]string {

	var matchedPodInfo = map[string]string{
		"pod":       "",
		"image":     "",
		"namespace": "",
	}

	for _, record := range podInfos {
		names := strings.Split(record, ",")
		if names[0] == pod {
			matchedPodInfo["pod"] = names[0]
			matchedPodInfo["image"] = names[1]
			matchedPodInfo["namespace"] = names[2]
		}
	}

	return matchedPodInfo

}

func deploy(kubeClient *client.Client, params map[string]string) {

	podInfos := get(kubeClient, "")

	myPodInfo := getMatchedPodInfo(params["pod"], podInfos)
	if myPodInfo["pod"] == "" {
		fmt.Println(params["pod"] + " doesn't exist.")
		os.Exit(1)
	}

	commandOptions := []string{"get", "pod", myPodInfo["pod"], "-o", "yaml"}
	result := execOutput("kubectl", commandOptions)
	result = strings.Replace(result, myPodInfo["image"], params["image"], -1)

	ioutil.WriteFile("tmp.dat", []byte(result), os.ModePerm)
	defer os.Remove("tmp.dat")

	commandOptions = []string{"replace", "-f", "tmp.dat"}
	result = execOutput("kubectl", commandOptions)
	fmt.Println(result)
}
