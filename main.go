package main

import (
	"flag"
	"fmt"
	"os"

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

	var image = flag.String("i", "blank", "string")
	var pod = flag.String("p", "blank", "string")
	flag.Parse()

	if flag.NArg() == 0 || flag.NArg() > 5 {
		help()
		os.Exit(1)
	}

	fmt.Println(*image)
	fmt.Println(*pod)

	kubeClient, err := newKubeClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli(kubeClient)

}
