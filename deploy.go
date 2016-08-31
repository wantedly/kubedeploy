package main

import (
	"fmt"

	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func deploy(kubeClient *client.Client, params map[string]string) {
	podInfos := get(kubeClient)
	fmt.Println(podInfos)

}
