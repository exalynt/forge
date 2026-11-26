package install

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

// Where the lock is: in .agents/, or in .claude/ when claude is the only
// target.
const (
	agentsLock = ".agents/forge.lock"
	claudeLock = ".claude/forge.lock"
)

// importFile is the file forge's @AGENTS.md line goes in, and what the lock
// records when forge added it.
const importFile = "CLAUDE.md"

var namePattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// Lock is what forge last installed in a repository.
type Lock struct {
	// Path is where the lock was read from, relative to the repository.
	Path string
	// Version is the forge that last changed the install.
	Version string
	// Targets are the targets installed.
	Targets []string
	// Wrote are the targets forge wrote its skills to. It's Targets, except
	// for older locks: one with "# skills .agents" wrote only to .agents/, and
	// one with neither that nor "# targets" also wrote copies to .claude/.
	Wrote []string
	// Stacks are the stacks installed.
	Stacks []string
	// Skills are the skills installed, stacks included.
	Skills []string
	// Imported says forge added the @AGENTS.md line to CLAUDE.md.
	Imported bool
}

// ReadLock reads the repository's lock, or returns nil when forge isn't
// installed there.
func ReadLock(root string) (*Lock, error) {
	for _, p := range []string{agentsLock, claudeLock} {
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return parseLock(p, string(data)), nil
	}
	return nil, nil
}

func parseLock(path, text string) *Lock {
	l := &Lock{Path: path}
	hasTargets, agentsOnly := false, false
	for line := range strings.SplitSeq(text, "\n") {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "# forge "):
			l.Version = strings.TrimSpace(strings.TrimPrefix(line, "# forge "))
		case line == "# targets" || strings.HasPrefix(line, "# targets "):
			hasTargets = true
			l.Targets = ordered(strings.Fields(strings.TrimPrefix(line, "# targets")))
		case line == "# skills .agents":
			agentsOnly = true
		case line == "# stacks" || strings.HasPrefix(line, "# stacks "):
			l.Stacks = strings.Fields(strings.TrimPrefix(line, "# stacks"))
		case line == "# import "+importFile:
			l.Imported = true
		case strings.HasPrefix(line, "#"), line == "":
		case namePattern.MatchString(line):
			l.Skills = append(l.Skills, line)
		}
	}
	switch {
	case hasTargets:
		l.Wrote = l.Targets
	case agentsOnly:
		l.Targets = []string{TargetAgents}
		l.Wrote = []string{TargetAgents}
	default:
		l.Targets = []string{TargetAgents}
		l.Wrote = []string{TargetAgents, TargetClaude}
	}
	return l
}

// lockPath is where the lock goes for targets.
func lockPath(targets []string) string {
	if slices.Equal(targets, []string{TargetClaude}) {
		return claudeLock
	}
	return agentsLock
}

// lockBody is the lock's text after its "# forge" line.
func lockBody(targets, stacks, skills []string, imported bool) string {
	var b strings.Builder
	b.WriteString("# targets " + strings.Join(targets, " ") + "\n")
	b.WriteString(strings.TrimRight("# stacks "+strings.Join(stacks, " "), " ") + "\n")
	if imported {
		b.WriteString("# import " + importFile + "\n")
	}
	for _, s := range skills {
		b.WriteString(s + "\n")
	}
	return b.String()
}

// withoutVersion drops a lock's "# forge" line, so locks can be compared by
// what they list.
func withoutVersion(text string) string {
	var lines []string
	for line := range strings.SplitSeq(text, "\n") {
		if !strings.HasPrefix(line, "# forge") {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\n")
}
