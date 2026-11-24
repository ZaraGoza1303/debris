package scanner

import (
	"debris/types"
	"os/exec"
	"strings"
)

func ScanDnf() ([]types.Application, error) {
	cmd := exec.Command("dnf", "repoquery", "--installed", "--queryformat", "%{name} %{version}\n")

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
			Source:  "dnf",
		})
	}

	return apps, nil
}
