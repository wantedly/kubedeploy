package main

import (
	"encoding/json"
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

	var f interface{}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteArray, &f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	for _, v := range m {
		switch vv := v.(type) {
		case string:
		case int:
		case bool:
		case []interface{}:
			for i, u := range vv {
				id := u.(map[string]interface{})["docker_image_id"].(string)
				fmt.Println(i+1, u.(map[string]interface{})["name"].(string)+"_"+id)
			}
		default:
		}
	}

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

func getTargetService(kubeClient *client.Client, namespace string) []api.Service {
	services, err := kubeClient.Services(namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return services.Items
}

func getServices(kubeClient *client.Client, namespace string) []api.Service {
	if namespace == "" {
		namespace = api.NamespaceAll
	}
	services, err := kubeClient.Services(namespace).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return services.Items
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
