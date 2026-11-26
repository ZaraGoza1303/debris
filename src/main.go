package main

import (
	"debris/scanner"
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

	// Scan all apps
	apps, err := scanner.Scan(config)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("help: debris -h for help")
		return
	}

	switch os.Args[1] {
	case "list":
		utils.PrintApplications(apps)

	case "-h", "--help":
		utils.PrintHelp()

	case "uninstall":
		uninstaller.HandleUninstall(apps, config)
	}
}
