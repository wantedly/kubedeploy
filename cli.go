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
			deploy(kubeClient, params)
		} else {
			help()
		}

	case "deploy":
		if params["image"] != "" && params["pod"] != "" {
			deployBG(kubeClient, params)
		} else {
			help()
		}

	case "list":
		if params["image"] != "" {
			getTagList(params["image"])
		} else {
			help()
		}

	default:
		help()
	}
}
