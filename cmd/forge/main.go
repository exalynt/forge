// Command forge installs forge into a repository for the agents it uses,
// keeps it up to date, and takes it out again.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli/v3"
	"golang.org/x/mod/semver"

	"github.com/exalynt/forge"
	"github.com/exalynt/forge/internal/install"
)

// module is forge's module, for go install and forge version --check.
const module = "github.com/exalynt/forge"

// errStale is returned once a command has printed what it found stale, or
// that the binary is older than the latest. It exits 1 with nothing more
// printed. Any other error means the command couldn't run, and exits 2.
var errStale = errors.New("stale")

// couldNot is an error a command returns when it couldn't run.
type couldNot struct{ msg string }

func (e couldNot) Error() string { return e.msg }

func main() {
	root := rootCommand()
	err := root.Run(context.Background(), os.Args)
	switch {
	case err == nil:
		os.Exit(0)
	case errors.Is(err, errStale):
		os.Exit(1)
	}
	if slices.Contains(os.Args[1:], "--json") {
		writeJSON(os.Stdout, map[string]string{"error": err.Error()})
	} else {
		fmt.Fprintf(os.Stderr, "forge: %s\n", err)
	}
	os.Exit(2)
}

// version is the binary's version from its build info, or "dev" when it has
// none, as when it's built from a checkout without version control.
func version() string {
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

func rootCommand() *cli.Command {
	jsonFlag := &cli.BoolFlag{Name: "json", Usage: "print JSON to stdout"}
	agentFlag := &cli.StringSliceFlag{Name: "agent", Usage: "an agent the repository uses, such as claude or codex; see forge agents"}
	stackFlag := &cli.StringSliceFlag{Name: "stack", Usage: "a stack to install, such as go-http or react"}
	noStacksFlag := &cli.BoolFlag{Name: "no-stacks", Usage: "install no stacks, removing any installed"}
	return &cli.Command{
		Name:    "forge",
		Usage:   "Install Exalynt's way of working into a repository, for the agents it uses",
		Version: version(),
		Flags:   []cli.Flag{jsonFlag},
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "Install forge, or add agents and stacks to an install",
				ArgsUsage: "[dir]",
				Description: `Writes forge's block, its skills, and its lock into the repository at dir,
the current folder by default. A first install with no --agent installs for
every agent that reads AGENTS.md and .agents/skills/. On an existing install,
it adds the targets of the agents named and the stacks named, and rewrites
everything installed. --no-stacks removes every stack.`,
				Flags:  []cli.Flag{jsonFlag, agentFlag, stackFlag, noStacksFlag},
				Action: installAction,
			},
			{
				Name:      "update",
				Usage:     "Rewrite every installed target with this binary's forge",
				ArgsUsage: "[dir]",
				Description: `Makes the install match this binary. It keeps the lock's stacks, unless
--stack names the stacks to have instead, or --no-stacks removes them all.`,
				Flags:  []cli.Flag{jsonFlag, stackFlag, noStacksFlag},
				Action: updateAction,
			},
			{
				Name:      "remove",
				Usage:     "Remove agents, stacks, or with no flags everything forge installed",
				ArgsUsage: "[dir]",
				Description: `Removes the targets of the agents named, and the stacks named. With neither,
it removes everything forge installed: its block, its skills, the @AGENTS.md
line if forge added it, its lock, and the folders that leaves empty. When the
target holding the block goes, the block moves to the one that's left.`,
				Flags:  []cli.Flag{jsonFlag, agentFlag, stackFlag},
				Action: removeAction,
			},
			{
				Name:      "status",
				Usage:     "Show what's installed and whether it matches this binary",
				ArgsUsage: "[dir]",
				Flags:     []cli.Flag{jsonFlag},
				Action:    statusAction,
			},
			{
				Name:      "check",
				Usage:     "Exit 1 when the install doesn't match this binary, for CI",
				ArgsUsage: "[dir]",
				Flags:     []cli.Flag{jsonFlag},
				Action:    checkAction,
			},
			{
				Name:   "agents",
				Usage:  "List the agents forge supports, and where their files go",
				Flags:  []cli.Flag{jsonFlag},
				Action: agentsAction,
			},
			{
				Name:  "version",
				Usage: "Print this binary's version",
				Description: `--check asks Go for the latest forge, and exits 1 when this binary is older,
or 2 when it can't tell, such as without Go or access to the repository.`,
				Flags:  []cli.Flag{jsonFlag, &cli.BoolFlag{Name: "check", Usage: "compare with the latest forge"}},
				Action: versionAction,
			},
		},
	}
}

// repoDir is the command's [dir], the current folder by default.
func repoDir(cmd *cli.Command) (string, error) {
	if cmd.Args().Len() > 1 {
		return "", couldNot{"expected at most one folder"}
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
		return "", couldNot{"not a folder: " + dir}
	}
	return abs, nil
}

// stacksWanted works out the stacks from --stack and --no-stacks, or the
// lock's when neither is given.
func stacksWanted(cmd *cli.Command, lock *install.Lock) ([]string, bool, error) {
	given := cmd.StringSlice("stack")
	if cmd.Bool("no-stacks") {
		if len(given) > 0 {
			return nil, false, couldNot{"--no-stacks and --stack can't be used together"}
		}
		return nil, true, nil
	}
	if len(given) > 0 {
		return given, true, nil
	}
	if lock != nil {
		return lock.Stacks, false, nil
	}
	return nil, false, nil
}

func installAction(ctx context.Context, cmd *cli.Command) error {
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
		return couldNot{err.Error()}
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

func updateAction(ctx context.Context, cmd *cli.Command) error {
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

func removeAction(ctx context.Context, cmd *cli.Command) error {
	root, lock, err := installed(cmd)
	if err != nil {
		return err
	}
	agents, stacks := cmd.StringSlice("agent"), cmd.StringSlice("stack")
	want := install.Want{}
	if len(agents) > 0 || len(stacks) > 0 {
		drop, err := install.TargetsFor(agents)
		if err != nil {
			return couldNot{err.Error()}
		}
		for _, t := range drop {
			if !slices.Contains(lock.Targets, t) {
				return couldNot{fmt.Sprintf("the %s target isn't installed here", t)}
			}
		}
		for _, s := range stacks {
			if !slices.Contains(lock.Stacks, s) {
				return couldNot{fmt.Sprintf("the %s stack isn't installed here", s)}
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

func checkAction(ctx context.Context, cmd *cli.Command) error {
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

func statusAction(ctx context.Context, cmd *cli.Command) error {
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
	res, err := install.Apply(root, forge.Files, install.Want{Targets: lock.Targets, Stacks: lock.Stacks}, version(), true)
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
			"targets": targets, "stacks": lock.Stacks, "upToDate": !res.Stale(), "files": res.Changes, "notes": notes,
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
	stacks := strings.Join(lock.Stacks, ", ")
	if stacks == "" {
		stacks = "none"
	}
	_, _ = fmt.Fprintf(out, "\nStacks: %s\n", stacks)
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

func agentsAction(ctx context.Context, cmd *cli.Command) error {
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

func versionAction(ctx context.Context, cmd *cli.Command) error {
	out := cmd.Root().Writer
	v := version()
	if !cmd.Bool("check") {
		if cmd.Bool("json") {
			writeJSON(out, map[string]string{"version": v})
		} else {
			_, _ = fmt.Fprintln(out, v)
		}
		return nil
	}
	latest, err := latestVersion(ctx)
	if err != nil {
		return err
	}
	older := false
	switch {
	case v == "dev" || !semver.IsValid(v):
		return couldNot{fmt.Sprintf("this binary is %s, so it can't be compared with the latest, %s: install it with go install %s/cmd/forge@latest", v, latest, module)}
	case semver.Compare(v, latest) < 0:
		older = true
	}
	if cmd.Bool("json") {
		writeJSON(out, map[string]any{"version": v, "latest": latest, "upToDate": !older})
	} else if older {
		_, _ = fmt.Fprintf(out, "forge %s is older than the latest, %s. Update it with:\n\n  go install %s/cmd/forge@latest\n\nthen run forge update in each repository.\n", v, latest, module)
	} else {
		_, _ = fmt.Fprintf(out, "forge %s is the latest.\n", v)
	}
	if older {
		return errStale
	}
	return nil
}

// latestVersion asks Go for forge's latest version.
func latestVersion(ctx context.Context) (string, error) {
	goBin, err := exec.LookPath("go")
	if err != nil {
		return "", couldNot{"can't check: Go isn't installed, or isn't on PATH"}
	}
	c := exec.CommandContext(ctx, goBin, "list", "-m", "-json", module+"@latest")
	// Outside any module, so the current folder's go.mod can't get in the way.
	c.Dir = os.TempDir()
	var stderr strings.Builder
	c.Stderr = &stderr
	data, err := c.Output()
	if err != nil {
		return "", couldNot{fmt.Sprintf("can't check: go list couldn't reach %s (%s). It's private, so set GOPRIVATE=github.com/exalynt and make sure git can fetch from GitHub", module, strings.TrimSpace(stderr.String()))}
	}
	var m struct{ Version string }
	if err := json.Unmarshal(data, &m); err != nil || m.Version == "" {
		return "", couldNot{"can't check: go list didn't say which version is the latest"}
	}
	return m.Version, nil
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
		return "", nil, couldNot{"forge isn't installed here: run forge install"}
	}
	return root, lock, nil
}

// printResult prints each file's change and a summary, or with --json the
// whole result.
func printResult(cmd *cli.Command, res install.Result, notes []string, mode string) {
	out := cmd.Root().Writer
	if cmd.Bool("json") {
		writeJSON(out, map[string]any{
			"files": res.Changes, "targets": res.Targets, "stacks": res.Stacks,
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

func union(a, b []string) []string {
	out := slices.Clone(a)
	for _, s := range b {
		if !slices.Contains(out, s) {
			out = append(out, s)
		}
	}
	return out
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

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
