package uninstaller

import (
	"debris/types"
	"os"
	"os/exec"
)

func UninstallFlatpak(app types.Application) error {
	target := app.Name
	if app.ID != "" {
		target = app.ID
	}

	cmd := exec.Command("flatpak", "uninstall", "--delete-data", target)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
