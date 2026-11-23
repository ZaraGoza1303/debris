package cleanup

import (
	"debris/src/types"
	"fmt"
)

func Cleaner(app types.Application) (string, error) {
	residuals, err := Lookup(app)
	if err != nil {
		return "", err
	}

	if len(residuals) == 0 {
		fmt.Println("No residuals files found")
		return "", nil
	}

	selected, err := Select(residuals)
	if err != nil {
		return "", err
	}

	if len(selected) == 0 {
		fmt.Println("No files selected.")
		return "", nil
	}

	confirmed, err := Confirm(len(selected))
	if err != nil {
		return "", err
	}

	if !confirmed {
		fmt.Println("Cleanup cancelled.")
		return "", nil
	}

	if err := Remove(selected); err != nil {
		return "", err
	}

	return "Cleanup completed.", nil
}
