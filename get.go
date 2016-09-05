package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

var QUAYIO string = "https://quay.io/api/v1/repository/"

func getTagList(image string) {
	url := QUAYIO + path.Join(image, "tag")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}

func getTargetPod(kubeClient *client.Client, podName string, namespace string) api.Pod {
	if namespace == "" {
		namespace = "default"
	}
	targetPod, err := kubeClient.Pods(namespace).Get(podName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return *targetPod
}

func getPods(kubeClient *client.Client, namespace string) []api.Pod {
	if namespace == "" {
		namespace = api.NamespaceAll
	}
	pods, err := kubeClient.Pods(namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return pods.Items
}
