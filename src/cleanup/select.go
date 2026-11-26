package cleanup

import (
	"fmt"

	"charm.land/huh/v2"
)

func Select(items []Residual) ([]Residual, error) {
	if len(items) == 0 {
		return nil, nil
	}

	// Deklarasi options dengan type Residual pada library huh
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

	var selected []Residual

	km := huh.NewDefaultKeyMap()
	km.MultiSelect.Toggle.SetHelp("(space)", "select")
	km.MultiSelect.Up.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.Up.Help().Key), km.MultiSelect.Up.Help().Desc)
	km.MultiSelect.Down.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.Down.Help().Key), km.MultiSelect.Down.Help().Desc)
	km.MultiSelect.Filter.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.Filter.Help().Key), km.MultiSelect.Filter.Help().Desc)
	km.MultiSelect.Submit.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.Submit.Help().Key), km.MultiSelect.Submit.Help().Desc)
	km.MultiSelect.SelectAll.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.SelectAll.Help().Key), km.MultiSelect.SelectAll.Help().Desc)
	km.MultiSelect.SelectNone.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.SelectNone.Help().Key), "unselect all")
	km.MultiSelect.SetFilter.SetHelp(fmt.Sprintf("(%s)", km.MultiSelect.SetFilter.Help().Key), "exit filter")

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[Residual]().
				Title("Select residual files to remove").
				Description("Residual files found").
				Options(options...).
				Value(&selected).
				Height(len(options) + 2),
		),
	).
		WithKeyMap(km).
		Run()

	if err != nil {
		return nil, err
	}

	return selected, nil
}

func Confirm(name string) (bool, error) {
	var confirmed bool

	err := huh.NewConfirm().
		Title(fmt.Sprintf("Are you sure you want to uninstall %s ?", name)).
		Affirmative("Uninstall").
		Negative("Cancel").
		Value(&confirmed).
		Run()

	return confirmed, err
}

func ConfirmResiduals(count int) (bool, error) {
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
