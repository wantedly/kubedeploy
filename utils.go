package main

import (
	"log"
	"os/exec"
	"strings"
)

func execOutput(app string, commands []string) string {
	out, err := exec.Command(app, commands...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func trimImageName(imageName string) string {
	head := strings.Index(imageName, QUAYPATH)
	tale := strings.Index(imageName, ":")

	if head != -1 {
		if tale != -1 {
			imageName = imageName[len(QUAYPATH):tale]
		} else {
			imageName = imageName[len(QUAYPATH):]
		}
	}
	return imageName
}
