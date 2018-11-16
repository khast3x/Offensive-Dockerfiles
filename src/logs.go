package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"
)

// ExecPrintCMD : Read logfile and print to screen
func ExecPrintCMD(dpath string, target string) {
	data, err := ioutil.ReadFile(path.Join(dpath, "output", target))
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(tui.Bold(tui.Yellow("Contents of file:\n")), string(data))
	fmt.Println(tui.Yellow("---------------------\n"))
}

// BuildLogList : Returns the log names
func BuildLogList(dpath string) []string {
	logList := make([]string, 0)
	logfiles, err := ioutil.ReadDir(path.Join(dpath, "output"))
	if err != nil {
		log.Println(err)
	}
	for _, f := range logfiles {
		logList = append(logList, f.Name())
	}
	return logList
}

// ToolLogMenu : Function called from menu to select tool logs to print
func ToolLogMenu(dpath string) {
	toolSelection := []string{tui.Dim(">> Back")}

	logList := BuildLogList(dpath)
	// Add Back to menu selection
	logList = append(toolSelection, logList...)

	fmt.Println(logList)
	prompt := promptui.Select{
		Size:      10,
		Templates: selecttemplate,
		Label:     "Choose log to read",
		Items:     logList,
	}
	// Main Loop
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// Parse user selections
		switch result {
		case tui.Dim(">> Back"):
			return

		default:
			// Exec cat/bat on file
			ExecPrintCMD(dpath, result)
		}
	}
}
