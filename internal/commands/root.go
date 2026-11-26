// Package commands is the forge command: each top-level command in a file of
// its own, and a command's subcommands in a folder named for it.
package commands

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"slices"

	"github.com/urfave/cli/v3"
)

// module is forge's module, for go install and forge version --check.
const module = "github.com/exalynt/forge"

// errStale is returned once a command has printed what it found stale, or
// that the binary is older than the latest. It exits 1 with nothing more
// printed. Any other error means the command couldn't run, and exits 2.
var errStale = cli.Exit("", 1)

var (
	jsonFlag     = &cli.BoolFlag{Name: "json", Usage: "print JSON to stdout"}
	agentFlag    = &cli.StringSliceFlag{Name: "agent", Usage: "an agent the repository uses, such as claude or codex; see forge agents"}
	stackFlag    = &cli.StringSliceFlag{Name: "stack", Usage: "a stack to install, such as go-http or react"}
	noStacksFlag = &cli.BoolFlag{Name: "no-stacks", Usage: "install no stacks, removing any installed"}
)

// Root is the forge command.
func Root() *cli.Command {
	return &cli.Command{
		Name:    "forge",
		Usage:   "Install Exalynt's way of working into a repository, for the agents it uses",
		Version: version(),
		Flags:   []cli.Flag{jsonFlag},
		Commands: []*cli.Command{
			InstallCommand(),
			UpdateCommand(),
			RemoveCommand(),
			StatusCommand(),
			CheckCommand(),
			AgentsCommand(),
			BrowserCommand(),
			VersionCommand(),
		},
	}
}

// Run runs forge with args and returns its exit code. A command that finds
// something stale exits 1 itself. Any other error is printed, as JSON with
// --json, and exits 2.
func Run(ctx context.Context, args []string) int {
	err := Root().Run(ctx, args)
	if err == nil {
		return 0
	}
	if slices.Contains(args[1:], "--json") {
		writeJSON(os.Stdout, map[string]string{"error": err.Error()})
	} else {
		fmt.Fprintf(os.Stderr, "forge: %s\n", err)
	}
	return 2
}

// version is the binary's version from its build info, or "dev" when it has
// none, as when it's built from a checkout without version control.
func version() string {
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}
