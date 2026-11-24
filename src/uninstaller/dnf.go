package uninstaller

import (
	"debris/types"
	"os"
	"os/exec"
)

func UninstallDnf(app types.Application) error {
	cmd := exec.Command("sudo", "dnf", "remove", app.Name)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
