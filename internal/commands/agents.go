package commands

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/install"
)

// AgentsCommand lists the agents forge supports, and where their files go.
func AgentsCommand() *cli.Command {
	return &cli.Command{
		Name:   "agents",
		Usage:  "List the agents forge supports, and where their files go",
		Flags:  []cli.Flag{jsonFlag},
		Action: runAgents,
	}
}

func runAgents(ctx context.Context, cmd *cli.Command) error {
	out := cmd.Root().Writer
	if cmd.Bool("json") {
		writeJSON(out, install.Targets)
		return nil
	}
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "--agent\tAGENT\tTARGET\tSKILLS\tREADS")
	var notes []string
	for _, t := range install.Targets {
		for _, a := range t.Agents {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s/\t%s\n", a.Name, a.Label, t.Name, t.Skills, a.Reads)
			if a.Note != "" {
				notes = append(notes, a.Name+": "+a.Note)
			}
		}
	}
	_ = w.Flush()
	_, _ = fmt.Fprintln(out, "\nThe block goes in AGENTS.md when the agents target is installed. With only")
	_, _ = fmt.Fprintln(out, "claude, it goes in CLAUDE.md. With both, CLAUDE.md gets an @AGENTS.md line.")
	printNotes(out, notes)
	return nil
}
