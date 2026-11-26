package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"text/tabwriter"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/playwright"
)

// CheckCommand exits 1 when something setup recorded is gone.
func CheckCommand() *cli.Command {
	return &cli.Command{
		Name:   "check",
		Usage:  "Exit 1 when what setup recorded is gone, such as after a Node version was removed",
		Flags:  []cli.Flag{&cli.BoolFlag{Name: "json", Usage: "print JSON to stdout"}},
		Action: runCheck,
	}
}

func runCheck(ctx context.Context, cmd *cli.Command) error {
	items, err := playwright.Check()
	if err != nil {
		return err
	}
	ok := !slices.ContainsFunc(items, func(it playwright.Item) bool { return it.Status != playwright.OK })
	out := cmd.Root().Writer
	if cmd.Bool("json") {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		_ = enc.Encode(map[string]any{"upToDate": ok, "checks": items})
	} else {
		w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
		_, _ = fmt.Fprintln(w, "CHECK\tSTATUS\tPATH")
		for _, it := range items {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\n", it.Name, it.Status, it.Path)
		}
		_ = w.Flush()
		if !ok {
			_, _ = fmt.Fprintln(out)
		}
		for _, it := range items {
			if it.Fix != "" {
				_, _ = fmt.Fprintf(out, "%s is missing: %s.\n", it.Name, it.Fix)
			}
		}
	}
	if !ok {
		// Exits 1 with nothing more printed, as forge check does.
		return cli.Exit("", 1)
	}
	return nil
}
