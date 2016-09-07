package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func printHelp(commands []map[string]string) {

	fmt.Println("=== help === ")

	data := [][]string{}
	for _, command := range commands {
		data = append(data, []string{
			command["command"],
			command["option"],
			command["description"],
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"command",
		"option",
		"description",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	fmt.Println()
}

func help() {
	commands := []map[string]string{
		map[string]string{
			"command":     "kubedeploy get",
			"option":      "[-n namespace]",
			"description": "Get Pods and Services information.",
		},
		map[string]string{
			"command":     "kubedeploy list",
			"option":      "-i image",
			"description": "List image tags",
		},
		map[string]string{
			"command":     "kubedeploy replace",
			"option":      "-p pod -i newImage -n namespace",
			"description": "Replace new image from old one in pod ",
		},
		map[string]string{
			"command":     "kubedeploy deploy",
			"option":      "-s service",
			"description": "Blue-Green Deploy",
		},
		map[string]string{
			"command":     "kubedeploy help",
			"option":      "",
			"description": "Print Usage",
		},
	}

	printHelp(commands)

}
