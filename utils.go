package main

import (
	"log"
	"os/exec"
)

func ExecOutput(app string, commands []string) string {
	out, err := exec.Command(app, commands...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
