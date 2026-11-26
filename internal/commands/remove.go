package commands

import (
	"context"
	"fmt"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// RemoveCommand removes agents, languages, stacks, or everything forge
// installed.
func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:      "remove",
		Usage:     "Remove agents, languages, stacks, or with no flags everything forge installed",
		ArgsUsage: "[dir]",
		Description: `Removes the targets of the agents named, the languages named with their
stacks, and the stacks named. With none of them, it removes everything forge
installed: its block, its skills, the @AGENTS.md
line if forge added it, its lock, and the folders that leaves empty. When the
target holding the block goes, the block moves to the one that's left.`,
		Flags:  []cli.Flag{jsonFlag, agentFlag, languageFlag, stackFlag},
		Action: runRemove,
	}
}

func runRemove(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	agents, languages, stacks := cmd.StringSlice("agent"), cmd.StringSlice("language"), stackNames(cmd.StringSlice("stack"))
	want := install.Want{}
	if len(agents) > 0 || len(languages) > 0 || len(stacks) > 0 {
		drop, err := install.TargetsFor(agents)
		if err != nil {
			return err
		}
		for _, t := range drop {
			if !slices.Contains(lock.Targets, t) {
				return fmt.Errorf("the %s target isn't installed here", t)
			}
		}
		for _, l := range languages {
			if !slices.Contains(lock.Languages, l) {
				return fmt.Errorf("the %s language isn't installed here", l)
			}
		}
		for _, s := range stacks {
			if !slices.Contains(lock.Stacks, s) {
				return fmt.Errorf("the %s stack isn't installed here", s)
			}
		}
		// A language goes with its stacks.
		for _, s := range lock.Stacks {
			if slices.Contains(languages, install.StackLanguage(s)) {
				stacks = append(stacks, s)
			}
		}
		want.Targets = without(lock.Targets, drop)
		want.Languages = without(lock.Languages, languages)
		want.Stacks = without(lock.Stacks, stacks)
	}
	res, err := install.Apply(root, forge.Files, want, version(), false)
	if err != nil {
		return err
	}
	printResult(cmd, res, nil, "")
	return nil
}

func without(a, drop []string) []string {
	var out []string
	for _, s := range a {
		if !slices.Contains(drop, s) {
			out = append(out, s)
		}
	}
	return out
}
