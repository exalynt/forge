package commands

import (
	"context"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// InstallCommand installs forge into a repository, or adds agents and stacks
// to an install.
func InstallCommand() *cli.Command {
	return &cli.Command{
		Name:      "install",
		Usage:     "Install forge, or add agents and stacks to an install",
		ArgsUsage: "[dir]",
		Description: `Writes forge's block, its skills, and its lock into the repository at dir,
the current folder by default. A first install with no --agent installs for
every agent that reads AGENTS.md and .agents/skills/. On an existing install,
it adds the targets of the agents named and the stacks named, and rewrites
everything installed. --no-stacks removes every stack.`,
		Flags:  []cli.Flag{jsonFlag, agentFlag, stackFlag, noStacksFlag},
		Action: runInstall,
	}
}

func runInstall(ctx context.Context, cmd *cli.Command) error {
	root, err := repoDir(cmd)
	if err != nil {
		return err
	}
	lock, err := install.ReadLock(root)
	if err != nil {
		return err
	}
	agents := cmd.StringSlice("agent")
	targets, err := install.TargetsFor(agents)
	if err != nil {
		return err
	}
	if lock != nil {
		targets = union(lock.Targets, targets)
	} else if len(targets) == 0 {
		agents = []string{install.TargetAgents}
		targets = []string{install.TargetAgents}
	}
	stacks, given, err := stacksWanted(cmd, lock)
	if err != nil {
		return err
	}
	if given && !cmd.Bool("no-stacks") && lock != nil {
		stacks = union(lock.Stacks, stacks)
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: targets, Stacks: stacks}, version(), false)
	if err != nil {
		return err
	}
	var notes []string
	if len(agents) > 0 {
		notes = install.NotesFor(res.Targets, agents)
	}
	printResult(cmd, res, notes, "")
	return nil
}

func union(a, b []string) []string {
	out := slices.Clone(a)
	for _, s := range b {
		if !slices.Contains(out, s) {
			out = append(out, s)
		}
	}
	return out
}
