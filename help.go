package main

import "fmt"

func help() {
	fmt.Println(`
Usage:
  kubedeploy get
  kubedeploy deploy -p pod -i image
	`)
}
