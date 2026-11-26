// Package browser holds forge browser's subcommands.
package browser

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/playwright"
)

// SetupCommand installs Playwright MCP and Chrome, and prints what to
// register with an agent.
func SetupCommand() *cli.Command {
	return &cli.Command{
		Name:      "setup",
		Usage:     "Install Playwright MCP and Chrome, and print what to register with your agent",
		ArgsUsage: "[-- server flags...]",
		Description: `Installs @playwright/mcp globally with npm, and Google Chrome if it's the
browser and missing, then records where they are in setup.json. Flags after
-- are passed to the server, such as --headless where there's no display, or
--browser firefox. Register forge browser server with your agent once, under
the five names it prints. Run setup again after switching Node versions or to
change the flags. The servers use it the next time they start, without
registering them again.`,
		Action: runSetup,
	}
}

func runSetup(ctx context.Context, cmd *cli.Command) error {
	out := cmd.Root().Writer
	errOut := cmd.Root().ErrWriter
	if errOut == nil {
		errOut = os.Stderr
	}
	s, path, err := playwright.Install(ctx, out, errOut, cmd.Args().Slice())
	if err != nil {
		return err
	}
	command := shellQuote(s.Forge) + " browser server"
	_, _ = fmt.Fprintf(out, "\nSet up Playwright MCP, recorded in %s:\n\n", path)
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(w, "  node\t%s\n  package\t%s\n  browser\t%s\n", s.Node, s.Package, s.Browser)
	if s.Launcher != "" {
		_, _ = fmt.Fprintf(w, "  launcher\t%s\n", s.Launcher)
	}
	if len(s.Flags) > 0 {
		_, _ = fmt.Fprintf(w, "  flags\t%s\n", strings.Join(s.Flags, " "))
	}
	_ = w.Flush()
	names := strings.Join(playwright.Servers, " ")
	_, _ = fmt.Fprintf(out, `
Servers already registered with this command use it the next time they start.
Otherwise, register these five MCP servers with your agent, each with the same
command (docs/browser.md in forge shows how for Codex and Cursor):

  names:    %s
  command:  %s

With Claude Code:

  for name in %s; do
    claude mcp add --scope user "$name" -- %s
  done
`, names, command, names, command)
	return nil
}

// shellQuote quotes s for a POSIX shell where it needs it.
func shellQuote(s string) string {
	if s != "" && !strings.ContainsFunc(s, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || strings.ContainsRune("/._-+=:,@", r))
	}) {
		return s
	}
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}
