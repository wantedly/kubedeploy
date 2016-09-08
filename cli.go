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

	case "deploy-bg":
		if params["service"] != "" {
			bgDeploy(kubeClient, params)
		} else {
			help()
		}

	case "deploy-one":
		if params["service"] != "" {
			oneDeploy(kubeClient, params)
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
