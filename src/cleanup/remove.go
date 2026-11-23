package cleanup

import (
	"fmt"
	"os"
	"path/filepath"
)

func Remove(selected []Residual) error {
	for _, item := range selected {
		path, err := filepath.Abs(item.Path)
		if err != nil {
			return err
		}

		info, err := os.Lstat(path)
		if err != nil {
			return err
		}

		// Jangan mengikuti symlink
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("refusing to remove symlink: %s", path)
		}

		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
