package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) string {
	homeDir, _ := os.UserHomeDir()

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, strings.TrimPrefix(path, "~/"))
	}

	if strings.HasPrefix(path, "$HOME") {
		path = os.ExpandEnv(path)
	}

	return path
}
