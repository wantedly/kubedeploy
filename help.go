package main

import "fmt"

func help() {
	fmt.Println(`
Usage:
$ kubedeploy get [-n namespace]
$ kubedeploy deploy -p pod -i image
$ kubedeploy list -i image
	`)
}
