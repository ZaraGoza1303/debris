package uninstaller

import (
	"debris/cleanup"
	"debris/types"
	"fmt"
	"log"

	"charm.land/huh/v2"
)

func uninstaller(app types.Application, config types.Config) error {
	switch app.Source {
	case "pacman":
		return UninstallPacman(app)

	case "apt":
		return UninstallApt(app)

	case "dnf":
		return UninstallDnf(app)

	case "flatpak":
		return UninstallFlatpak(app)

	case "AppImage":
		return UninstallAppImage(app, config)

	default:
		return fmt.Errorf("unsupported application source: %s", app.Source)
	}

}

func beginUninstall(apps []types.Application) (selectedApp types.Application, err error) {
	var options []huh.Option[types.Application]

	for _, app := range apps {
		label := fmt.Sprintf(
			"%s | %s ",
			app.Name,
			app.Source,
		)

		options = append(
			options,
			huh.NewOption(label, app),
		)
	}

	var selected types.Application
	desc := fmt.Sprintf("%d applications found", len(apps))

	km := huh.NewDefaultKeyMap()
	km.Select.Up.SetHelp(fmt.Sprintf("(%s)", km.Select.Up.Help().Key), km.Select.Up.Help().Desc)
	km.Select.Down.SetHelp(fmt.Sprintf("(%s)", km.Select.Down.Help().Key), km.Select.Down.Help().Desc)
	km.Select.Filter.SetHelp(fmt.Sprintf("(%s)", km.Select.Filter.Help().Key), km.Select.Filter.Help().Desc)
	km.Select.Submit.SetHelp(fmt.Sprintf("(%s)", km.Select.Submit.Help().Key), km.Select.Submit.Help().Desc)
	km.Select.SetFilter.SetHelp(fmt.Sprintf("(%s)", km.Select.SetFilter.Help().Key), "exit filter")

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.Application]().
				Title("Select application to uninstall").
				Description(desc).
				Options(options...).
				Value(&selected).
				Height(len(options) + 2),
		),
	).
		WithKeyMap(km).
		Run()

	return selected, err
}

func HandleUninstall(apps []types.Application, config types.Config) {
	selected, err := beginUninstall(apps)
	if err != nil {
		fmt.Println("Failed to select an application: ", err)
		return
	}

	confirmed, err := cleanup.Confirm(selected.Name)
	if err != nil {
		log.Fatal(err)
	}

	if !confirmed {
		HandleUninstall(apps, config)
		return
	}

	if err := uninstaller(selected, config); err != nil {
		fmt.Println("Uninstall Failed: ", err)
		return
	}
	fmt.Printf("\nSuccessfully uninstalled: %s\n\n", selected.Name)

	output, err := cleanup.Cleaner(selected, config)
	if err != nil {
		fmt.Println("Failed to clean residuals files: ", err)
	}

	fmt.Println(output)
}
