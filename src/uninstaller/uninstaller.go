package uninstaller

import (
	"debris/types"
	"fmt"
)

func Uninstaller(app types.Application, config types.Config) error {
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
