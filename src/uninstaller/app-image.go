package uninstaller

import (
	"debris/scanner"
	"debris/types"
	"os"
	"os/exec"
)

func UninstallAppImage(app types.Application, config types.Config) error {
	filePath, err := scanner.FindFilePathAppImage(app.Name, config)
	if err != nil {
		return err
	}

	cmd := exec.Command("rm", filePath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
