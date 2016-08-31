package main

import (
	"fmt"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func cli(kubeClient *client.Client, params map[string]string) {

	switch params["subCommand"] {

	case "get":
		podInfos := get(kubeClient)
		for _, info := range podInfos {
			fmt.Println(info)
		}

	case "deploy":
		if params["image"] != "" && params["pod"] != "" {
			deploy(kubeClient, params)
		} else {
			help()
		}

	default:
		help()
	}

}