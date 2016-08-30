package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
)

func newKubeClient() (*client.Client, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = clientcmd.RecommendedHomeFile

	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

	clientConfig, err := loader.ClientConfig()

	if err != nil {
		return nil, err
	}

	kubeClient, err := client.New(clientConfig)

	if err != nil {
		return nil, err
	}

	return kubeClient, nil
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: ./kubedeploy get")
		os.Exit(1)
	}

	kubeClient, err := newKubeClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pods, err := kubeClient.Pods(api.NamespaceAll).List(api.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	podInfos := []string{}
	for _, pod := range pods.Items {
		podInfoList := []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace}
		podInfos = append(podInfos, strings.Join(podInfoList, ","))
	}

	for _, info := range podInfos {
		fmt.Println(info)
	}
}
