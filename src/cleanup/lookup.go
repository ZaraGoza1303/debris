package cleanup

import (
	"debris/types"
	"debris/utils"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Residual struct {
	Path     string
	Category string
	IsDir    bool
}

func Lookup(app types.Application, config types.Config) ([]Residual, error) {
	roots := lookupRoots(app, config)

	targets := []string{app.Name}
	if app.Source == "flatpak" && app.ID != "" {
		targets = append(targets, app.ID)
	}

	var results []Residual
	seen := make(map[string]bool)

	for _, root := range roots {
		err := filepath.WalkDir(root.Path, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrPermission) {
					return nil
				}
				return err
			}

			// Root itu sendiri tidak perlu dimasukkan.
			if path == root.Path {
				return nil
			}

			// Symlink tidak diikuti.
			if entry.Type()&os.ModeSymlink != 0 {
				return nil
			}

			if !matchesAny(entry.Name(), targets) {
				return nil
			}

			if seen[path] {
				return nil
			}
			seen[path] = true

			results = append(results, Residual{
				Path:     path,
				Category: root.Category,
				IsDir:    entry.IsDir(),
			})

			// Kalau direktori cocok, isinya tidak perlu didaftarkan satu per satu.
			if entry.IsDir() {
				return filepath.SkipDir
			}

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

func lookupRoots(app types.Application, config types.Config) []types.SearchRoot {
	paths := make([]types.SearchRoot, 0, len(config.Lookup.Paths))

	for _, path := range config.Lookup.Paths {
		paths = append(paths, path)
	}

	switch app.Source {
	case "flatpak":
		paths = append(paths, config.Lookup.Flatpak.Paths...)

	case "AppImage":
		paths = append(paths, config.Lookup.AppImage.Paths...)
	}

	seen := make(map[string]bool, len(paths))
	roots := make([]types.SearchRoot, 0, len(paths))

	for _, path := range paths {
		expanded := utils.ExpandHome(path.Path)
		if expanded == "" || seen[expanded] {
			continue
		}

		seen[expanded] = true
		roots = append(roots, types.SearchRoot{
			Path:     expanded,
			Category: path.Category,
		})
	}

	return roots
}

func normalize(name string) string {
	return strings.Map(func(r rune) rune {
		// Cek apakah unicode termasuk letter dan digit, kalau iya maka buat lowercase
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, name)
}

func matchesAny(entry string, targets []string) bool {
	for _, target := range targets {
		if matches(entry, target) {
			return true
		}
	}

	return false
}

func matches(entry string, targetApp string) bool {
	normalized := normalize(entry)

	// Ambil nama file/folder tanpa extension
	withoutExt := normalize(strings.TrimSuffix(
		entry,
		filepath.Ext(entry),
	))

	target := normalize(targetApp)

	if normalized == target || withoutExt == target {
		return true
	}

	return false
}
