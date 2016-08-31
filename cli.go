package main

import (
	"flag"
	"fmt"
)

func cli() {
	switch flag.Arg(0) {
	case "get":
		podInfos := get(kubeClient)
		for _, info := range podInfos {
			fmt.Println(info)
		}
	case "deploy":
		// podInfos := getPodInfos(kubeClient)
		deploy(flag.Arg(1))
	default:
		help()
	}
}
