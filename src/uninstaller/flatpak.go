package uninstaller

import (
	"debris/src/types"
	"os"
	"os/exec"
)

func UninstallFlatpak(app types.Application) error {
	cmd := exec.Command("flatpak", "uninstall", "--delete-data", app.Name)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
