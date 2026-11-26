package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintHelp() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "usage: debris <command> [options]")
	fmt.Fprintln(w, "commands:")
	fmt.Fprintln(w, "  list\tShow all applications detected on this system")
	fmt.Fprintln(w, "  uninstall\tUninstall an app interactively, then remove leftovers")
	w.Flush()

	// Separate writer so the options column doesn't follow the command width.
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "options:")
	fmt.Fprintln(w, "  -h, --help\tShow this help message and exit")
	w.Flush()
}
