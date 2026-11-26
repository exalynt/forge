package commands

import (
	"context"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// CheckCommand exits 1 when the install doesn't match this binary, for CI.
func CheckCommand() *cli.Command {
	return &cli.Command{
		Name:      "check",
		Usage:     "Exit 1 when the install doesn't match this binary, for CI",
		ArgsUsage: "[dir]",
		Flags:     []cli.Flag{jsonFlag},
		Action:    runCheck,
	}
}

func runCheck(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: lock.Targets, Stacks: lock.Stacks}, version(), true)
	if err != nil {
		return err
	}
	printResult(cmd, res, nil, "check")
	if res.Stale() {
		return errStale
	}
	return nil
}
