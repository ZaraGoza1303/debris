package scanner

import (
	"debris/src/types"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/probonopd/go-appimage/src/goappimage"
)

func ScanAppImage(config types.Config) ([]types.Application, error) {
	var apps []types.Application

	for _, path := range config.Scanner.AppImage.Paths {
		path = expandHome(path)

		err := filepath.WalkDir(path, func(filePath string, data fs.DirEntry, err error) error {
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrPermission) {
					return nil
				}
				return err
			}

			if data.IsDir() {
				return nil
			}

			if goappimage.IsAppImage(filePath) {
				ai, err := goappimage.NewAppImage(filePath)
				if err != nil {

					return nil
				}

				apps = append(apps, types.Application{
					Name:    ai.Name,
					Version: ai.Version,
					Source:  "AppImage",
				})
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return apps, nil
}

func FindFilePathAppImage(name string, config types.Config) (string, error) {
	var appFilePath string

	for _, path := range config.Scanner.AppImage.Paths {
		path = expandHome(path)

		err := filepath.WalkDir(path, func(filePath string, data fs.DirEntry, err error) error {
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrPermission) {
					return nil
				}
				return err
			}

			if data.IsDir() {
				return nil
			}

			if goappimage.IsAppImage(filePath) {
				ai, err := goappimage.NewAppImage(filePath)
				if err != nil {

					return nil
				}

				if name == ai.Name {
					appFilePath = filePath
				}
			}

			return nil
		})

		if err != nil {
			return "", err
		}
	}

	return appFilePath, nil
}

func expandHome(path string) string {
	homeDir, _ := os.UserHomeDir()

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(homeDir, strings.TrimPrefix(path, "~/"))
	}

	return path
}
