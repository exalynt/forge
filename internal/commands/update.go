package commands

import (
	"context"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// UpdateCommand rewrites every installed target with this binary's forge.
func UpdateCommand() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "Rewrite every installed target with this binary's forge",
		ArgsUsage: "[dir]",
		Description: `Makes the install match this binary. It keeps the lock's stacks, unless
--stack names the stacks to have instead, or --no-stacks removes them all.`,
		Flags:  []cli.Flag{jsonFlag, stackFlag, noStacksFlag},
		Action: runUpdate,
	}
}

func runUpdate(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	stacks, _, err := stacksWanted(cmd, lock)
	if err != nil {
		return err
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: lock.Targets, Stacks: stacks}, version(), false)
	if err != nil {
		return err
	}
	printResult(cmd, res, nil, "")
	return nil
}
