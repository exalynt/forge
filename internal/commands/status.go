package commands

import (
	"context"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// StatusCommand shows what's installed and whether it matches this binary.
func StatusCommand() *cli.Command {
	return &cli.Command{
		Name:      "status",
		Usage:     "Show what's installed and whether it matches this binary",
		ArgsUsage: "[dir]",
		Flags:     []cli.Flag{jsonFlag},
		Action:    runStatus,
	}
}

func runStatus(ctx context.Context, cmd *cli.Command) error {
	root, err := repoDir(cmd)
	if err != nil {
		return err
	}
	lock, err := install.ReadLock(root)
	if err != nil {
		return err
	}
	out := cmd.Root().Writer
	if lock == nil {
		if cmd.Bool("json") {
			writeJSON(out, map[string]any{"installed": false, "binary": version()})
		} else {
			_, _ = fmt.Fprintln(out, "forge isn't installed here: run forge install.")
		}
		return nil
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: lock.Targets, Languages: lock.Languages, Stacks: lock.Stacks}, version(), true)
	if err != nil {
		return err
	}
	type targetStatus struct {
		Name     string   `json:"name"`
		Skills   string   `json:"skills"`
		Agents   []string `json:"agents"`
		UpToDate bool     `json:"upToDate"`
	}
	var targets []targetStatus
	for _, name := range lock.Targets {
		t, _ := install.TargetNamed(name)
		ts := targetStatus{Name: t.Name, Skills: t.Skills, UpToDate: true}
		for _, a := range t.Agents {
			ts.Agents = append(ts.Agents, a.Name)
		}
		for _, c := range res.Changes {
			if c.Status != install.OK && strings.HasPrefix(c.Path, t.Skills+"/") {
				ts.UpToDate = false
			}
		}
		targets = append(targets, ts)
	}
	notes := install.NotesFor(lock.Targets, nil)
	if cmd.Bool("json") {
		writeJSON(out, map[string]any{
			"installed": true, "binary": version(), "lockVersion": lock.Version, "lock": lock.Path,
			"targets": targets, "languages": lock.Languages, "stacks": lock.Stacks, "upToDate": !res.Stale(), "files": res.Changes, "notes": notes,
		})
		return nil
	}
	_, _ = fmt.Fprintf(out, "forge %s installed (%s); this binary is %s\n\n", lock.Version, lock.Path, version())
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "TARGET\tSKILLS\tAGENTS\tMATCHES THIS BINARY")
	for _, t := range targets {
		_, _ = fmt.Fprintf(w, "%s\t%s/\t%s\t%s\n", t.Name, t.Skills, strings.Join(t.Agents, ", "), yesNo(t.UpToDate))
	}
	_ = w.Flush()
	_, _ = fmt.Fprintf(out, "\nLanguages: %s\nStacks: %s\n", listOrNone(lock.Languages), listOrNone(lock.Stacks))
	if res.Stale() {
		_, _ = fmt.Fprintln(out, "\nNot up to date: forge update would change these:")
		for _, c := range res.Changes {
			if c.Status != install.OK {
				_, _ = fmt.Fprintf(out, "  %-9s %s\n", c.Status, c.Path)
			}
		}
	} else {
		_, _ = fmt.Fprintln(out, "Everything matches this binary.")
	}
	printNotes(out, notes)
	return nil
}

func listOrNone(list []string) string {
	if len(list) == 0 {
		return "none"
	}
	return strings.Join(list, ", ")
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
