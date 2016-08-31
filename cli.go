package main

import (
	"flag"
	"fmt"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func cli(kubeClient *client.Client, params []string) {
	switch flag.Arg(0) {
	case "get":
		podInfos := get(kubeClient)
		for _, info := range podInfos {
			fmt.Println(info)
		}
	case "deploy":
		if params[0] != "no" && params[1] != "no" {
			deploy(params)
		}
	default:
		help()
	}
}
