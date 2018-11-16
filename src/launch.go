package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"
)

// WriteToolOutput : Outputs a running tool's stdout to a file (look into writer/stream)
func WriteToolOutput(fulltpath string, out bytes.Buffer) {
	file, _ := os.Create(fulltpath)
	fmt.Fprint(file, &out)
	file.Close()
	PrintGood("[*] Finished writing output")
}

// execLaunchCMD : Launchs built tool, asking user for arguments. Output will be saved to output/ dir
func execLaunchCMD(tool string, args string, tpath string) {
	fmt.Println(tui.Dim("[*] Started background launch job for " + tool))

	argsSlice := []string{"run", strings.ToLower(tool)}
	tmpSlice := strings.Split(args, " ")
	// Append is a variadic functions
	// Using "..." unrolls slice as param
	argsSlice = append(argsSlice, tmpSlice...)

	cmd := exec.Command("docker", argsSlice...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	PrintGood("[*] Done with " + tool + " job")

	// Prepare full string to be sent as output destination
	fullpath := path.Join(tpath, "output", time.Now().Format("2006-01-02_150405")+"_"+tool+".log")
	os.MkdirAll(path.Join(tpath, "output"), os.ModePerm)
	fmt.Println("[*] Writing output to ", fullpath)
	WriteToolOutput(fullpath, out)
}

// AskFgBg : Ask user if he wants to run the tool in foreground or background (goroutine)
func AskFgBg() string {
	choiceSelection := []string{"Background", "Foreground"}
	prompt := promptui.Select{
		Size:      10,
		Templates: selecttemplate,
		Label:     "Choose how to run the tool",
		Items:     choiceSelection,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "nil"
	}
	if result == "Foreground" {
		return "fg"
	}
	return "bg"
}

// AskArguments : Ask user for tool arguments with prompt
func AskArguments(tool string) string {
	tool += " args"
	prompt := promptui.Prompt{
		Label: tool,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "nil"
	}
	return result
}

// AskLaunchConfirm : Confirm launch job with user
func AskLaunchConfirm(tool string, args string) bool {
	fmt.Println("\n\n", tui.Bold("=> "), "docker", "run", "-it", strings.ToLower(tool), args)
	prompt := promptui.Select{
		Size:      10,
		Templates: selecttemplate,
		Label:     "Confirm Launch",
		Items:     []string{tui.Bold("Confirm Launch"), "Cancel"},
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

// LaunchToolMenu : Main launch function with prompt
func LaunchToolMenu(tpath string, tools []string) bool {
	buildList := BuildBuiltList()
	builtToolsList := GetBuiltTools(tools, buildList)
	// Build selected dockerfile(s)
	toolSelection := []string{tui.Dim(">> Back")}
	// Add back to menu
	for _, t := range builtToolsList {
		toolSelection = append(toolSelection, t)
	}
	prompt := promptui.Select{
		Label:     "Choose tool to launch",
		Size:      10,
		Templates: selecttemplate,
		Items:     toolSelection,
	}

	// Main Loop
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false
		}
		// Parse user selections
		switch result {
		case tui.Dim(">> Back"):
			return false

		// Launch result
		default:
			args := AskArguments(result)
			if AskLaunchConfirm(result, args) {
				go execLaunchCMD(result, args, tpath)
			}
		}
	}
}
