package commands

import (
	"context"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// InstallCommand installs forge into a repository, or adds agents, languages
// and stacks to an install.
func InstallCommand() *cli.Command {
	return &cli.Command{
		Name:      "install",
		Usage:     "Install forge, or add agents, languages and stacks to an install",
		ArgsUsage: "[dir]",
		Description: `Writes forge's block, its skills, its standards, and its lock into the
repository at dir, the current folder by default. A first install with no
--agent installs for every agent that reads AGENTS.md and .agents/skills/.
--language adds a language's conventions, and --stack a stack's, with its
language: --stack go/http installs go and go-http. On an existing install, it
adds the targets of the agents named and the languages and stacks named, and
rewrites everything installed. --no-languages removes every language and
stack.`,
		Flags:  []cli.Flag{jsonFlag, agentFlag, languageFlag, stackFlag, noLanguagesFlag},
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
	languages, stacks, given, err := languagesWanted(cmd, lock)
	if err != nil {
		return err
	}
	if given && !cmd.Bool("no-languages") && lock != nil {
		languages = union(lock.Languages, languages)
		stacks = union(lock.Stacks, stacks)
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: targets, Languages: languages, Stacks: stacks}, version(), false)
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
