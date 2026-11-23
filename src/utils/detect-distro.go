package utils

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func GetDistro() string {
	if runtime.GOOS != "linux" {
		return runtime.GOOS
	}

	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "linux (unknown distro)"
	}
	defer file.Close()

	var name string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "ID=") {
			name = strings.TrimPrefix(line, "ID=")
			name = strings.Trim(name, " \" ")
			break
		}
	}

	return name
}
