package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"k8s.io/kubernetes/pkg/api"
)

func printTagList(tagList []string) {
	for i, tag := range tagList {
		fmt.Print(i + 1)
		fmt.Print(" ")
		fmt.Print(tag)
		if i == 0 {
			color.Green(" [NEW]")
		} else {
			fmt.Println()
		}
	}
}

func printPodsTable(pods []api.Pod) {
	color.Red("=== Pod ===")
	data := [][]string{}
	for _, pod := range pods {
		data = append(data, []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace,
			pod.CreationTimestamp.String(),
		})

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name",
		"Image",
		"Namespace",
		"Created at",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	fmt.Println()
}

func printServicesTable(services []api.Service) {
	color.Blue("=== Service ===")
	data := [][]string{}
	for _, service := range services {
		data = append(data, []string{
			service.Name,
			service.Namespace,
			service.Spec.Selector["color"],
			service.CreationTimestamp.String(),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name",
		"Namespace",
		"Color",
		"Created at",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	fmt.Println()

}

func printReplace(old, new string) {
	fmt.Println("Replace: " + old + " => " + new)
}
