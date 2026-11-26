package commands

import (
	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/commands/browser"
)

// BrowserCommand sets up and runs Playwright MCP, once per machine.
func BrowserCommand() *cli.Command {
	return &cli.Command{
		Name:  "browser",
		Usage: "Set up and run Playwright MCP, the browsers qa, ux, and use drive, once per machine",
		Commands: []*cli.Command{
			browser.SetupCommand(),
			browser.ServerCommand(),
			browser.CheckCommand(),
		},
	}
}
