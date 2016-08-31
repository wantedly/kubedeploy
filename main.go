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

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		pod       = fs.String("p", "", "Pod name")
		image     = fs.String("i", "", "Image name")
		namespace = fs.String("n", "", "Namespace name")
	)
	fs.Parse(os.Args[2:])
	var params = map[string]string{
		"subCommand": os.Args[1],
		"image":      *image,
		"pod":        *pod,
		"namespace":  *namespace,
	}

	kubeClient, err := newKubeClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli(kubeClient, params)

}
