package main

import (
	"debris/cleanup"
	"debris/scanner"
	"debris/types"
	"debris/uninstaller"
	"debris/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	apps, err := scanner.Scan(config)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: debris list/uninstall")
		return
	}

	switch os.Args[1] {
	case "list":
		utils.PrintApplications(apps)

	case "uninstall":
		var selected *types.Application

		if len(os.Args) < 3 {
			fmt.Println("Application name not provided yet")
			break
		}

		appName := os.Args[2]
		for i := range apps {
			if apps[i].Name == appName {
				selected = &apps[i]
			}
		}

		if selected == nil {
			fmt.Println("Application not found\nTry search for an application name using: debris list")
			break
		}

		if err := uninstaller.Uninstaller(*selected, config); err != nil {
			fmt.Println("Uninstall Failed:", err)
			break
		}

		output, err := cleanup.Cleaner(*selected, config)
		if err != nil {
			fmt.Println("Failed to clean residuals files:", err)
		}

		fmt.Println(output)
	}

}
