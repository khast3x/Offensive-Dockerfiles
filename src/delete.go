package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"
)

func GetBuiltTools(tools []string, built string) []string {

	builtTools := make([]string, 0)

	for _, t := range tools {
		if strings.Contains(built, strings.ToLower(t)) {
			builtTools = append(builtTools, t)
		}
	}
	return builtTools
}

// ExecBuildCmd : Launches the docker build command
func ExecDeleteCmd(selected []string) {

	argsSlice := []string{"rmi"}
	argsSlice = append(argsSlice, selected...)
	cmd := exec.Command("docker", argsSlice...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(&out)
		log.Print(err, "- Docker Delete")
	}
	fmt.Println(&out)
	PrintGood("[*] Deleted ")
	fmt.Println(selected)
}

// ConfirmDelete : Ask User to confirm delete actions
func ConfirmDelete(selected []string) bool {

	prompt := promptui.Select{
		Label:     "Confirm Delete",
		Size:      10,
		Templates: selecttemplate,
		Items:     []string{"Confirm Delete", "Cancel"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	if result == "Cancel" {
		return false
	}
	return true
}

// DeleteTool : Function called from menu to select tools to build
func DeleteTool(dpath string, tools []string) []string {

	buildList := BuildBuiltList()
	builtToolsList := GetBuiltTools(tools, buildList)
	// Delete selected dockerfile(s)
	toolSelection := []string{tui.Yellow(tui.Bold(">> Delete")), tui.Dim(">> Back")}

	selected := make([]string, 0)
	// Add done & back to menu
	for _, t := range builtToolsList {
		toolSelection = append(toolSelection, t)
	}
	prompt := promptui.Select{
		Label: "Choose tools to delete",
		Items: toolSelection,
	}

	// Main Loop
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}
		// Parse user selections
		switch result {
		case tui.Dim(">> Back"):
			return selected
		case tui.Yellow(tui.Bold(">> Delete")):
			// If no selected images
			if len(selected) == 0 {
				fmt.Println(tui.Red("[!] Please select at least one tool"))
				break
			}
			// Print tools to build
			fmt.Println(tui.Bold("\n\tSelected:"))
			for _, sel := range selected {
				fmt.Println("\t+ ", tui.Yellow(sel))
			}
			// Launch delete CMD
			if ConfirmDelete(selected) {
				ExecDeleteCmd(selected)
			}

			return selected
		default:
			// Check if duplicate
			if CheckDuplicate(selected, result) {
				selected = append(selected, strings.ToLower(result))
			} else {
				fmt.Println(tui.Red("[!] " + result + ": Already added to queue, skipping.."))
			}
		}
	}
}
