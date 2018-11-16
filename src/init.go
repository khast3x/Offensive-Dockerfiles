package main

// This file contains helper functions called before starting working with the tools

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"
)

func CheckDuplicate(src []string, target string) bool {
	for _, s := range src {
		if s == target {
			return false
		}
	}
	return true
}

func PrintBanner() {
	banner := tui.Bold("\t---------------------\n")
	println("\n", banner, tui.Yellow("\tOffensive Dockerfiles\n"), tui.Yellow("\tDeploy Tool -khast3x\n"), banner)
}

func PrintGood(str string) {
	fmt.Println(tui.Green(str))
}
func GetRepoFiles() string {
	items := []string{"Perform git clone"}
	var result string
	var err error

	prompt := promptui.SelectWithAdd{
		Label:    "Perform a \"git pull Offensive-Dockerfiles\"?",
		Items:    items,
		AddLabel: "Provide Path (Default: ./Offensive-Dockerfiles)",
	}
	fmt.Println(tui.Red("[!] No local Dockerfiles found. Provide path or clone repository\n"))
	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "err"
	}
	if len(result) == 0 {
		result = "./Offensive-Dockerfiles"

	}

	if result == "Perform git clone" {
		// Check if file already exist. If yes, err
		if _, err := os.Stat("./Offensive-Dockerfiles"); os.IsNotExist(err) {
			fmt.Println("[*] Performing git clone")
			cmd := exec.Command("git", "clone", "https://github.com/khast3x/Offensive-Dockerfiles.git")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err, "- Is git installed?")
			}
			PrintGood("[*] Pulled remote Offensive Dockerfile repository")
			result = "./Offensive-Dockerfiles"
		} else {
			log.Println("Offensive-Dockerfiles directory already exists. Please remove and relaunch")
		}
	}
	fmt.Printf("[*] Setting Dockerfiles location to %v\n", result)

	if _, err := os.Stat(result); !os.IsNotExist(err) {
		PrintGood("[*] Directory exists - Found " + result)
	}

	return result
}
