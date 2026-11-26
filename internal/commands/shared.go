package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/install"
)

// repoDir is the command's [dir], the current folder by default.
func repoDir(cmd *cli.Command) (string, error) {
	if cmd.Args().Len() > 1 {
		return "", errors.New("expected at most one folder")
	}
	dir := cmd.Args().First()
	if dir == "" {
		dir = "."
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return "", errors.New("not a folder: " + dir)
	}
	return abs, nil
}

// installed is the repository and its lock, or an error when forge isn't
// installed there.
func installed(cmd *cli.Command) (string, *install.Lock, error) {
	root, err := repoDir(cmd)
	if err != nil {
		return "", nil, err
	}
	lock, err := install.ReadLock(root)
	if err != nil {
		return "", nil, err
	}
	if lock == nil {
		return "", nil, errors.New("forge isn't installed here: run forge install")
	}
	return root, lock, nil
}

// languagesWanted works out the languages and stacks from --language,
// --stack and --no-languages, or the lock's when none is given. It reports
// whether any was given.
func languagesWanted(cmd *cli.Command, lock *install.Lock) (languages, stacks []string, given bool, err error) {
	languages, stacks = cmd.StringSlice("language"), stackNames(cmd.StringSlice("stack"))
	if cmd.Bool("no-languages") {
		if len(languages) > 0 || len(stacks) > 0 {
			return nil, nil, false, errors.New("--no-languages can't be used with --language or --stack")
		}
		return nil, nil, true, nil
	}
	if len(languages) > 0 || len(stacks) > 0 {
		return languages, stacks, true, nil
	}
	if lock != nil {
		return lock.Languages, lock.Stacks, false, nil
	}
	return nil, nil, false, nil
}

// stackNames is each stack's name as forge writes it.
func stackNames(stacks []string) []string {
	var out []string
	for _, s := range stacks {
		out = append(out, install.StackName(s))
	}
	return out
}

// printResult prints each file's change and a summary, or with --json the
// whole result.
func printResult(cmd *cli.Command, res install.Result, notes []string, mode string) {
	out := cmd.Root().Writer
	if cmd.Bool("json") {
		writeJSON(out, map[string]any{
			"files": res.Changes, "targets": res.Targets, "languages": res.Languages, "stacks": res.Stacks,
			"upToDate": !res.Stale(), "notes": notes,
		})
		return
	}
	counts := map[string]int{}
	for _, c := range res.Changes {
		_, _ = fmt.Fprintf(out, "%-9s %s\n", c.Status, c.Path)
		counts[c.Status]++
	}
	_, _ = fmt.Fprintln(out)
	if mode == "check" {
		_, _ = fmt.Fprintf(out, "check: %d to add, %d to update, %d to remove, %d ok\n", counts[install.Added], counts[install.Updated], counts[install.Removed], counts[install.OK])
		if res.Stale() {
			_, _ = fmt.Fprintln(out, "Not up to date: run forge update.")
		}
	} else {
		_, _ = fmt.Fprintf(out, "%d added, %d updated, %d removed, %d ok\n", counts[install.Added], counts[install.Updated], counts[install.Removed], counts[install.OK])
	}
	printNotes(out, notes)
}

func printNotes(out io.Writer, notes []string) {
	if len(notes) == 0 {
		return
	}
	_, _ = fmt.Fprintln(out, "\nNotes:")
	for _, n := range notes {
		_, _ = fmt.Fprintf(out, "- %s\n", n)
	}
}

func writeJSON(w io.Writer, v any) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}
