package main

import "strings"

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
