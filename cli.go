package main

import client "k8s.io/kubernetes/pkg/client/unversioned"

func cli(kubeClient *client.Client, params map[string]string) {

	switch params["subCommand"] {

	case "get":
		pods := getPods(kubeClient, params["namespace"])
		printPodCSV(pods)

	case "deploy":
		if params["image"] != "" && params["pod"] != "" {
			deploy(kubeClient, params)
		} else {
			help()
		}

	case "list":
		getTagList(params["image"])

	default:
		help()
	}
}
