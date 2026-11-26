// Package forge holds what forge installs into a repository: the block for
// its agents' instructions, its skills, and its stacks. cmd/forge installs
// them, and internal/install says how.
package forge

import "embed"

// Files holds AGENTS.md, the block written into a repository's instructions;
// skills/, each skill forge installs; and stacks/, each stack a repository can
// choose.
//
//go:embed AGENTS.md skills stacks
var Files embed.FS
