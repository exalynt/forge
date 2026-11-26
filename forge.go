// Package forge holds what forge installs into a repository, its kit: the
// block for its agents' instructions, its skills, its standards, and its
// languages and stacks. cmd/forge installs them, and internal/install says
// how.
package forge

import (
	"embed"
	"io/fs"
)

//go:embed kit
var kit embed.FS

// Files is the kit: AGENTS.md, the block written into a repository's
// instructions; workflows/, the workflow skills; universal/, the conventions
// every repository gets; and languages/, each language a repository can
// choose, with the stacks under it. Every install has the workflows and the
// universal skills.
var Files fs.FS = mustSub(kit, "kit")

func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
