package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var KUBECTL = "kubectl"

func execOutput(app string, commands []string) string {
	out, err := exec.Command(app, commands...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func getNamespaces() []string {
	out := execOutput(KUBECTL, []string{"get", "namespace"})

	records := strings.Split(out, "\n")[1:]
	records = records[:len(records)-1]

	namespaces := []string{}
	for key := range records {
		namespaces = append(namespaces, strings.Split(records[key], " ")[0])
	}

	return namespaces
}

func getDescribePods() {

}

func getImages() {

}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: ./kubedeploy get")
		os.Exit(1)
	}

	// get namespace
	namespaces := getNamespaces()
	fmt.Println(namespaces)

	// get describe

	// get images

}
