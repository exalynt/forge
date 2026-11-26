package commands

import (
	"context"
	"fmt"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// RemoveCommand removes agents, stacks, or everything forge installed.
func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:      "remove",
		Usage:     "Remove agents, stacks, or with no flags everything forge installed",
		ArgsUsage: "[dir]",
		Description: `Removes the targets of the agents named, and the stacks named. With neither,
it removes everything forge installed: its block, its skills, the @AGENTS.md
line if forge added it, its lock, and the folders that leaves empty. When the
target holding the block goes, the block moves to the one that's left.`,
		Flags:  []cli.Flag{jsonFlag, agentFlag, stackFlag},
		Action: runRemove,
	}
}

func runRemove(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	agents, stacks := cmd.StringSlice("agent"), cmd.StringSlice("stack")
	want := install.Want{}
	if len(agents) > 0 || len(stacks) > 0 {
		drop, err := install.TargetsFor(agents)
		if err != nil {
			return err
		}
		for _, t := range drop {
			if !slices.Contains(lock.Targets, t) {
				return fmt.Errorf("the %s target isn't installed here", t)
			}
		}
		for _, s := range stacks {
			if !slices.Contains(lock.Stacks, s) {
				return fmt.Errorf("the %s stack isn't installed here", s)
			}
		}
		want.Targets = without(lock.Targets, drop)
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
