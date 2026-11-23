package cleanup

import (
	"debris/src/types"
	"errors"
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

type searchRoot struct {
	Path     string
	Category string
}

func Lookup(app types.Application) ([]Residual, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	config, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	cache, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	data := os.Getenv("XDG_DATA_HOME")
	if data == "" {
		data = filepath.Join(home, ".local", "share")
	}

	state := os.Getenv("XDG_STATE_HOME")
	if state == "" {
		state = filepath.Join(home, ".local", "state")
	}

	roots := []searchRoot{
		{config, "Configuration"},
		{data, "Application Data"},
		{cache, "Cache"},
		{state, "Application State"},
	}

	var results []Residual

	for _, root := range roots {
		// Ambil data isi direktori
		entries, err := os.ReadDir(root.Path)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}

		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			// Cek apakah symlink atau bukan
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}

			// Cek apakah entry yang masuk itu sama dengan target app yang dicari
			if !matches(entry.Name(), app.Name) {
				continue
			}

			results = append(results, Residual{
				Path:     filepath.Join(root.Path, entry.Name()),
				Category: root.Category,
				IsDir:    entry.IsDir(),
			})
		}
	}

	return results, nil
}

func normalize(name string) string {
	return strings.Map(func(r rune) rune {
		// Cek apakah uncicode termasuk letter dan digit, kalau iya maka buat lowercase
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, name)
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
