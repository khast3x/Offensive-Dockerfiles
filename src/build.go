package main

// This file contains functions to build the tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"
)

// BuildBuiltList : Returns list of built docker images
func BuildBuiltList() string {
	cmd := exec.Command("docker", "images", "--format", "\"{{.Repository}}\"")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("cmd.Run() BuildBuiltList failed with %s\n", err)
	}
	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr
}

// BuildRunningList : Returns list of running docker images
func BuildRunningList() string {
	cmd := exec.Command("docker", "ps", "--format", "\"{{.Image}}\"")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}
	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr
}

// BuildToolList : Folder parsing. Builds a list of repos that contain a Dockerfile
func BuildToolList(dpath string) []string {
	files, err := ioutil.ReadDir(dpath)
	if err != nil {
		log.Println(err)
	}
	tools := make([]string, 0)

	// Find folders with a Dockerfile
	for _, f := range files {
		if _, err := os.Stat(path.Join(dpath, f.Name(), "Dockerfile")); !os.IsNotExist(err) {
			if f.IsDir() {
				// fmt.Println(tui.Dim(">> adding " + f.Name()))
				tools = append(tools, f.Name())
			}
		}
	}
	fmt.Printf("[*] Found a total of %v available tools\n\n", len(tools))
	return tools
}

// ExecBuildCmd : Launches the docker build command
func ExecBuildCmd(tpath string, tool string) {
	fmt.Println(tui.Dim("[*] Started background building job for " + tool))
	cmd := exec.Command("docker", "build", path.Join(tpath, tool), "-t", strings.ToLower(tool))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(&out)
		log.Println(err, "- Docker")
	}
	PrintGood("[*] Done with " + tool)

}

// BuildToolDockerfile : Function called once tools to build have been chosen
// Launches building of queued tools, should be async or goroutine
func BuildToolDockerfile(tpath string, selected []string) {
	// Look into interfaces for here
	for _, tool := range selected {
		go ExecBuildCmd(tpath, tool)
	}
}

// BuildTool : Function called from menu to select tools to build
func BuildTool(dpath string, tools []string) []string {

	// Build selected dockerfile(s)
	toolSelection := []string{tui.Bold(">> Build"), tui.Dim(">> Back")}
	selected := make([]string, 0)
	// Add done & back to menu
	for _, t := range tools {
		toolSelection = append(toolSelection, t)
	}
	prompt := promptui.Select{
		Label:     "Choose tools to build",
		Size:      10,
		Templates: selecttemplate,
		Items:     toolSelection,
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
		case tui.Bold(">> Build"):
			// If no selected images
			if len(selected) == 0 {
				fmt.Println(tui.Red("[!] Please select at least one tool"))
				break
			}
			// Print tools to build
			fmt.Println(tui.Bold("\nSelected:"))
			for _, sel := range selected {
				fmt.Println(tui.Yellow(sel))
			}

			BuildToolDockerfile(dpath, selected)

			return selected
		default:
			// Check if duplicate
			if CheckDuplicate(selected, result) {
				selected = append(selected, result)
			} else {
				fmt.Println(tui.Red("[!] " + result + ": Already added to queue, skipping.."))
			}
		}
	}
}
