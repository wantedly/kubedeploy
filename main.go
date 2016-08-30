package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var KUBECTL = "kubectl"

func getExecOutput(app string, commands []string) string {
	out, err := exec.Command(app, commands...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: ./kubedeploy get")
		os.Exit(1)
	}

	sub_commands := []string{"get", "po"}
	out := getExecOutput(KUBECTL, sub_commands)
	fmt.Println(out)

	// get namespace

	// get describe

	// get images

}
