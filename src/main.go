package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/evilsocket/islazy/tui"

	"github.com/manifoldco/promptui"
)

var selecttemplate = &promptui.SelectTemplates{
	Label:    "{{ . |  bold | white }}",
	Active:   "\U0001F449 {{ . | cyan | bold }}",
	Inactive: "  {{ . | cyan }} ",
	Selected: "\U0001F449 {{ . | green | white }}",
	Help:     "   ",
}

// PrintToolList : Fetchs list of docker images, compares and prints tool list + build stat
func PrintToolList(tools []string) {
	buildList := BuildBuiltList()
	runList := BuildRunningList()
	columns := []string{
		"Available Tools",
		"Status",
	}
	// Check if tool is built or running
	rows := make([][]string, 0)
	for _, t := range tools {
		if len(t) != 0 {
			// 'RUNNING' has priority over 'BUILD OK'
			buildStat := ""
			if strings.Contains(runList, strings.ToLower(t)) {
				buildStat = tui.Bold(tui.Green("RUNNING"))
				t = tui.Bold(tui.Green(t))
			} else if strings.Contains(buildList, strings.ToLower(t)) {
				buildStat = tui.Green("BUILD OK")
				t = tui.Green(t)
			}
			rows = append(rows, []string{t, buildStat})
		}
	}
	tui.Table(os.Stdout, columns, rows)
}

// StartPrompt : Launches the main loop, shows initial menu selection
func StartPrompt(dpath string, tools []string) {

	prompt := promptui.Select{
		Label:     tui.Bold("Main Menu"),
		Size:      10,
		Templates: selecttemplate,
		Items:     []string{"List", "History", "Build", "Delete", "Launch", "Quit"},
	}

	for {
		PrintBanner()

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// Parse user selection
		switch result {
		case "Quit":
			fmt.Println("\n\nBye\n\n")
			return
		case "Build":
			BuildTool(dpath, tools)

		case "List":
			PrintToolList(tools)
		case "Delete":
			DeleteTool(dpath, tools)
		case "Launch":
			LaunchToolMenu(dpath, tools)
		case "History":
			ToolLogMenu(dpath)
		}
	}
}

func main() {
	dpath := "./Offensive-Dockerfiles"

	if _, err := os.Stat(dpath); os.IsNotExist(err) {
		dpath = GetRepoFiles()
	} else {
		PrintGood("[*] Directory exists - Found " + dpath)
	}
	tools := BuildToolList(dpath)

	StartPrompt(dpath, tools)
}
