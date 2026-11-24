package scanner

import (
	"debris/types"
	"debris/utils"
)

func Scan(config types.Config) ([]types.Application, error) {
	var apps []types.Application
	var err error

	distro := utils.GetDistro()

	switch distro {
	case "ubuntu", "debian", "linuxmint", "pop":
		apps, err = ScanApt()

	case "arch":
		apps, err = ScanPacman()

	case "fedora":
		apps, err = ScanDnf()
	}

	if config.Scanner.Flatpak.Enabled {
		flatpakApp, err := ScanFlatpak()
		if err != nil {
			return nil, err
		}
		apps = append(apps, flatpakApp...)
	}

	if config.Scanner.AppImage.Enabled {
		appImages, err := ScanAppImage(config)
		if err != nil {
			return nil, err
		}
		apps = append(apps, appImages...)
	}

	return apps, err
}
