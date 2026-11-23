package main

import (
	"debris/src/cleanup"
	"debris/src/scanner"
	"debris/src/types"
	"debris/src/uninstaller"
	"debris/src/utils"
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
			fmt.Println("Application not found")
		}

		if err := uninstaller.Uninstaller(*selected, config); err != nil {
			fmt.Println("Uninstall Failed:", err)
		}

		output, err := cleanup.Cleaner(*selected)
		if err != nil {
			fmt.Println("Failed to clean residuals files:", err)
		}

		fmt.Println(output)
	}

}
