package utils

import (
	"debris/types"
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintApplications(apps []types.Application) {
	if len(apps) == 0 {
		fmt.Println("No applications found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

	fmt.Fprintln(w, "NAME\tSOURCE\tVERSION")

	for _, a := range apps {
		version := a.Version
		if version == "" {
			version = "-"
		}

		fmt.Fprintf(
			w,
			"%s\t%s\t%s\n",
			a.Name,
			a.Source,
			version,
		)
	}

	w.Flush()

	fmt.Printf("\nTotal: %d applications found\n", len(apps))
}
