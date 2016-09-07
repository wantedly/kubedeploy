package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func getNewestMasterTag(tagList []string) string {

	for _, tag := range tagList {
		if strings.Index(tag, "master") != -1 {
			return tag
		}
	}
	return ""
}

func getTagList(image string) []string {

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

	var tagList = []string{}
	m := f.(map[string]interface{})
	for _, v := range m {
		switch vv := v.(type) {
		case string:
		case int:
		case bool:
		case []interface{}:
			for _, u := range vv {
				_, ok := u.(map[string]interface{})["name"]
				if ok {
					tagList = append(tagList, u.(map[string]interface{})["name"].(string))
				}
			}
		default:
		}
	}
	return tagList
}

func getBlueAndGreenPods(kubeClient *client.Client, service, namespace string) ([]api.Pod, []api.Pod) {
	pods := getPods(kubeClient, namespace)
	var bluePods = []api.Pod{}
	var greenPods = []api.Pod{}

	for _, pod := range pods {
		if pod.Labels["name"] == service {
			if pod.Labels["color"] == "blue" {
				bluePods = append(bluePods, pod)
			} else if pod.Labels["color"] == "green" {
				greenPods = append(greenPods, pod)
			}
		}
	}
	if len(bluePods) == 0 || len(greenPods) == 0 {
		fmt.Println("blue-green pods don't exist.")
		os.Exit(1)
	}
	return bluePods, greenPods
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

func getTargetService(kubeClient *client.Client, serviceName string, namespace string) api.Service {
	if namespace == "" {
		namespace = "default"
	}
	targetService, err := kubeClient.Services(namespace).Get(serviceName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return *targetService
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
