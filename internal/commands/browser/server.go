package browser

import (
	"context"

	"github.com/urfave/cli/v3"

	"github.com/exalynt/forge/internal/playwright"
)

// ServerCommand runs Playwright MCP as setup recorded it. It's the command
// agents register.
func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run Playwright MCP as setup recorded it: the command agents register",
		Description: `Runs @playwright/mcp with the Node and flags in setup.json, speaking MCP over
stdin and stdout. Each server writes the Chrome launcher first, so it's always
the one this binary carries.`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return playwright.Serve()
		},
	}
}
