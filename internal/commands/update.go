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
		Description: `Makes the install match this binary. It keeps the lock's languages and
stacks, unless --language and --stack name the ones to have instead, or
--no-languages removes them all. A stack brings its language with it.`,
		Flags:  []cli.Flag{jsonFlag, languageFlag, stackFlag, noLanguagesFlag},
		Action: runUpdate,
	}
}

func runUpdate(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	languages, stacks, _, err := languagesWanted(cmd, lock)
	if err != nil {
		return err
	}
	res, err := install.Apply(root, forge.Files, install.Want{Targets: lock.Targets, Languages: languages, Stacks: stacks}, version(), false)
	if err != nil {
		return err
	}
	printResult(cmd, res, nil, "")
	return nil
}
