package scanner

import (
	"debris/types"
	"os/exec"
	"strings"
)

func ScanPacman() ([]types.Application, error) {
	cmd := exec.Command("pacman", "-Qe")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var apps []types.Application
	for _, app := range strings.Split(string(output), "\n") {
		fields := strings.Fields(app)

		if len(fields) < 2 {
			continue
		}

		apps = append(apps, types.Application{
			Name:    fields[0],
			Version: fields[1],
			Source:  "pacman",
		})
	}

	return apps, nil
}
