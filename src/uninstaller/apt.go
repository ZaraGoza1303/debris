package uninstaller

import (
	"debris/types"
	"os"
	"os/exec"
)

func UninstallApt(app types.Application) error {
	cmd := exec.Command("sudo", "apt", "purge", "-y", app.Name)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
