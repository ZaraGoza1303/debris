package scanner

import (
	"debris/types"
	"os/exec"
	"strings"
)

func ScanFlatpak() ([]types.Application, error) {
	// Use explicit columns so display names containing spaces are not split.
	cmd := exec.Command("flatpak", "list", "--app", "--columns=name,application,version")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var apps []types.Application
	for _, app := range strings.Split(string(output), "\n") {
		fields := strings.Split(app, "\t")

		if len(fields) < 2 {
			continue
		}

		name := strings.TrimSpace(fields[0])
		applicationID := strings.TrimSpace(fields[1])

		if name == "" || applicationID == "" {
			continue
		}

		version := ""
		if len(fields) > 2 {
			version = strings.TrimSpace(fields[2])
		}

		apps = append(apps, types.Application{
			Name:    name,
			Version: version,
			Source:  "flatpak",
			ID:      applicationID,
		})
	}

	return apps, nil
}
