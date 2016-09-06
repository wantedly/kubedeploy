package main

import client "k8s.io/kubernetes/pkg/client/unversioned"

func cli(kubeClient *client.Client, params map[string]string) {

	switch params["subCommand"] {

	case "get":
		pods := getPods(kubeClient, params["namespace"])
		services := getServices(kubeClient, params["namespace"])
		printPodsTable(pods)
		printServicesTable(services)

	case "replace":
		if params["image"] != "" && params["pod"] != "" {
			replace(kubeClient, params)
		} else {
			help()
		}

	case "deploy":
		if params["image"] != "" && params["service"] != "" {
			deploy(kubeClient, params)
		} else {
			help()
		}

	case "list":
		if params["image"] != "" {
			tagList := getTagList(params["image"])
			printTagList(tagList)
		} else {
			help()
		}

	default:
		help()
	}
}
