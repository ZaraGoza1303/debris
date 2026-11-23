package uninstaller

import (
	"debris/src/types"
	"os"
	"os/exec"
)

func UninstallPacman(app types.Application) error {
	cmd := exec.Command("sudo", "pacman", "-Rns", app.Name)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
