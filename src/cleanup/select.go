package cleanup

import (
	"fmt"

	"charm.land/huh/v2"
)

func Select(items []Residual) ([]Residual, error) {
	if len(items) == 0 {
		return nil, nil
	}

	var options []huh.Option[Residual]

	for _, item := range items {
		label := fmt.Sprintf(
			"%s - %s",
			item.Category,
			item.Path,
		)

		options = append(
			options,
			huh.NewOption(label, item),
		)
	}

	height := len(options) + 1
	if height > 12 {
		height = 12
	}

	var selected []Residual

	err := huh.NewMultiSelect[Residual]().
		Title("Select residual files to remove").
		Options(options...).
		Value(&selected).
		Height(height).
		Run()

	if err != nil {
		return nil, err
	}

	return selected, nil
}

func Confirm(count int) (bool, error) {
	var confirmed bool

	err := huh.NewConfirm().
		Title(fmt.Sprintf(
			"Are you sure you want to remove %d items?",
			count,
		)).
		Affirmative("Remove").
		Negative("Cancel").
		Value(&confirmed).
		Run()

	return confirmed, err
}
