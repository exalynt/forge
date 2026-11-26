// Command forge installs forge into a repository for the agents it uses,
// keeps it up to date, and takes it out again. internal/commands has its
// commands.
package main

import (
	"context"
	"os"

	"github.com/exalynt/forge/internal/commands"
)

func main() {
	os.Exit(commands.Run(context.Background(), os.Args))
}
