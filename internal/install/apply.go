package install

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// Change is what happened, or with a dry run what would happen, to one path.
type Change struct {
	Path   string `json:"path"`
	Status string `json:"status"`
}

// Statuses a change can have.
const (
	Added   = "added"
	Updated = "updated"
	Removed = "removed"
	OK      = "ok"
)

// Want is the install a repository should end up with. No targets means
// nothing of forge at all.
type Want struct {
	Targets   []string
	Languages []string
	Stacks    []string
}

// Result is what Apply did, or would do.
type Result struct {
	Changes   []Change `json:"files"`
	Targets   []string `json:"targets"`
	Languages []string `json:"languages"`
	Stacks    []string `json:"stacks"`
	Lock      string   `json:"lock,omitempty"`
}

// Stale reports whether anything changed, or would.
func (r Result) Stale() bool {
	return slices.ContainsFunc(r.Changes, func(c Change) bool { return c.Status != OK })
}

// Apply makes the repository at root hold want, from src, and returns what it
// changed. With dryRun it changes nothing and returns what it would change.
// version is written to the lock when what it lists changes, so the lock
// names the forge that last changed the install's shape.
func Apply(root string, src fs.FS, want Want, version string, dryRun bool) (Result, error) {
	want.Targets = ordered(want.Targets)
	old, err := ReadLock(root)
	if err != nil {
		return Result{}, err
	}
	if old == nil {
		old = &Lock{}
	}
	a := &applier{root: root, src: src, old: old, dryRun: dryRun}
	return a.apply(want, version)
}

type applier struct {
	root    string
	src     fs.FS
	old     *Lock
	dryRun  bool
	changes []Change
	removed bool
}

func (a *applier) apply(want Want, version string) (Result, error) {
	// What to install.
	want, skills, err := resolve(a.src, want)
	if err != nil {
		return Result{}, err
	}
	names := make([]string, len(skills))
	for i, s := range skills {
		names[i] = s.name
	}

	// Everything that would stop the install, before anything changes.
	docs := map[string]*doc{}
	for _, f := range []string{"AGENTS.md", importFile} {
		d, err := a.readDoc(f)
		if err != nil {
			return Result{}, err
		}
		docs[f] = d
	}
	for _, name := range want.Targets {
		t, _ := TargetNamed(name)
		for _, s := range skills {
			if err := a.checkForeign(t, s); err != nil {
				return Result{}, err
			}
		}
	}

	// Skills, in each target.
	for _, name := range want.Targets {
		t, _ := TargetNamed(name)
		for _, s := range skills {
			if err := a.sync(path.Join(t.Skills, s.name), s); err != nil {
				return Result{}, err
			}
		}
		// Skills forge installed here before and no longer does.
		if slices.Contains(a.old.Wrote, t.Name) {
			for _, n := range a.old.Skills {
				if n != "docket" && !slices.Contains(names, n) {
					if err := a.removeAll(path.Join(t.Skills, n)); err != nil {
						return Result{}, err
					}
				}
			}
		}
	}

	// Targets no longer installed, and copies an earlier forge left.
	for _, t := range Targets {
		if slices.Contains(want.Targets, t.Name) {
			continue
		}
		candidates := slices.Clone(a.old.Skills)
		for _, s := range skills {
			if !slices.Contains(candidates, s.name) {
				candidates = append(candidates, s.name)
			}
		}
		for _, n := range candidates {
			if n == "docket" {
				continue
			}
			rel := path.Join(t.Skills, n)
			if !a.exists(rel) {
				continue
			}
			ours := slices.Contains(a.old.Wrote, t.Name) && slices.Contains(a.old.Skills, n)
			if !ours && t.Name == TargetClaude {
				// A link, or an exact copy of forge's skill, is a copy an
				// earlier forge left. Anything else is the project's own.
				ours = a.isLink(rel) || a.sameAs(rel, skillNamed(skills, n))
			}
			if ours {
				if err := a.removeAll(rel); err != nil {
					return Result{}, err
				}
			}
		}
	}

	// The block, in AGENTS.md with the agents target, otherwise in CLAUDE.md.
	agents, claude := docs["AGENTS.md"], docs[importFile]
	hasAgents := slices.Contains(want.Targets, TargetAgents)
	hasClaude := slices.Contains(want.Targets, TargetClaude)
	block, err := a.block(want)
	if err != nil {
		return Result{}, err
	}
	switch {
	case hasAgents:
		agents.set(setBlock(agents.text, block))
		claude.set(dropBlock(claude.text))
	case hasClaude:
		claude.set(setBlock(claude.text, block))
		agents.set(dropBlock(agents.text))
	default:
		agents.set(dropBlock(agents.text))
		claude.set(dropBlock(claude.text))
	}

	// The @AGENTS.md line, so Claude Code reads AGENTS.md beside CLAUDE.md:
	// with both targets, and with agents when the project has a CLAUDE.md of
	// its own, since Claude Code reads AGENTS.md itself only without one.
	imported := a.old.Imported
	ownClaude := claude.exists && strings.TrimSpace(claude.text) != ""
	switch {
	case hasAgents && (hasClaude || ownClaude):
		if !hasLine(claude.text, "@AGENTS.md") {
			claude.set(addLine(claude.text, "@AGENTS.md"))
			imported = true
		}
	case imported:
		// Kept while Claude is still served and AGENTS.md has more in it that
		// Claude would otherwise stop reading.
		keep := hasClaude && strings.TrimSpace(agents.text) != ""
		if !keep {
			claude.set(dropLine(claude.text, "@AGENTS.md"))
			imported = false
		}
	}
	claude.ours = hasAgents && hasClaude
	for _, f := range []string{"AGENTS.md", importFile} {
		if err := a.writeDoc(docs[f]); err != nil {
			return Result{}, err
		}
	}

	// The lock.
	lock := ""
	if len(want.Targets) > 0 {
		lock = lockPath(want.Targets)
		body := lockBody(want, names, imported)
		have, err := os.ReadFile(a.abs(lock))
		switch {
		case err == nil && withoutVersion(string(have)) == body:
			a.report(lock, OK)
		case err == nil || errors.Is(err, os.ErrNotExist):
			status := Added
			if err == nil {
				status = Updated
			}
			if err := a.write(lock, []byte("# forge "+version+"\n"+body)); err != nil {
				return Result{}, err
			}
			a.report(lock, status)
		default:
			return Result{}, err
		}
	}
	for _, p := range []string{agentsLock, claudeLock} {
		if p != lock && a.exists(p) {
			if err := a.removeAll(p); err != nil {
				return Result{}, err
			}
		}
	}

	// Folders forge's removals left empty, and nothing else.
	if a.removed && !a.dryRun {
		for _, dir := range []string{".agents/skills", ".agents", ".claude/skills", ".claude"} {
			// Remove fails on a folder that isn't empty, which stays.
			_ = os.Remove(a.abs(dir))
		}
	}
	return Result{Changes: a.changes, Targets: want.Targets, Languages: want.Languages, Stacks: want.Stacks, Lock: lock}, nil
}

// block is forge's block, naming the skill folder agents should read
// (.agents/skills when that target is installed, otherwise .claude/skills)
// and the conventions installed.
func (a *applier) block(want Want) (string, error) {
	data, err := fs.ReadFile(a.src, "AGENTS.md")
	if err != nil {
		return "", err
	}
	dir := ".agents/skills"
	if !slices.Contains(want.Targets, TargetAgents) {
		dir = ".claude/skills"
	}
	text := strings.ReplaceAll(string(data), "{{skills}}", dir)
	return strings.ReplaceAll(text, "{{conventions}}", conventions(want)), nil
}

// checkForeign refuses a skill folder at a target that forge didn't install,
// unless it's already exactly forge's.
func (a *applier) checkForeign(t Target, s skill) error {
	rel := path.Join(t.Skills, s.name)
	if !a.exists(rel) {
		return nil
	}
	if slices.Contains(a.old.Wrote, t.Name) && slices.Contains(a.old.Skills, s.name) {
		return nil
	}
	if !a.isLink(rel) && a.sameAs(rel, &s) {
		return nil
	}
	return fmt.Errorf("%s exists and wasn't installed by forge; move or remove it first", rel)
}

// sync makes the folder rel hold exactly the skill's files.
func (a *applier) sync(rel string, s skill) error {
	if a.isLink(rel) {
		if err := a.removeAll(rel); err != nil {
			return err
		}
	}
	want := map[string]bool{}
	err := s.walk(a.src, func(sub string, data []byte) error {
		want[sub] = true
		dest := path.Join(rel, sub)
		have, err := os.ReadFile(a.abs(dest))
		status := OK
		switch {
		case errors.Is(err, os.ErrNotExist) || (a.dryRun && a.isLink(rel)):
			status = Added
		case err != nil:
			return err
		case !bytes.Equal(have, data):
			status = Updated
		}
		if status != OK {
			if err := a.write(dest, data); err != nil {
				return err
			}
		}
		a.report(dest, status)
		return nil
	})
	if err != nil {
		return err
	}
	if !a.exists(rel) || a.isLink(rel) {
		return nil
	}
	var stray []string
	err = filepath.WalkDir(a.abs(rel), func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		sub, err := filepath.Rel(a.abs(rel), p)
		if err != nil {
			return err
		}
		if !want[filepath.ToSlash(sub)] {
			stray = append(stray, filepath.ToSlash(sub))
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, sub := range stray {
		if err := a.removeAll(path.Join(rel, sub)); err != nil {
			return err
		}
	}
	if len(stray) > 0 && !a.dryRun {
		removeEmptyDirs(a.abs(rel))
	}
	return nil
}

// doc is AGENTS.md or CLAUDE.md, read once and written once.
type doc struct {
	path    string
	exists  bool
	orig    string
	text    string
	changed bool
	// ours says forge has something in it besides its block: the
	// @AGENTS.md line.
	ours bool
}

func (d *doc) set(text string) {
	if text != d.text {
		d.text = text
		d.changed = true
	}
}

func (a *applier) readDoc(file string) (*doc, error) {
	d := &doc{path: file}
	data, err := os.ReadFile(a.abs(file))
	switch {
	case errors.Is(err, os.ErrNotExist):
		return d, nil
	case err != nil:
		return nil, err
	}
	d.exists, d.orig, d.text = true, string(data), string(data)
	if _, _, _, err := findBlock(file, d.text); err != nil {
		return nil, err
	}
	return d, nil
}

// writeDoc writes a doc that changed, or removes it when nothing but
// whitespace is left.
func (a *applier) writeDoc(d *doc) error {
	if !d.changed && !d.exists {
		return nil
	}
	if !d.changed || d.text == d.orig {
		if d.exists && (d.ours || strings.Contains(d.text, blockStart)) {
			a.report(d.path, OK)
		}
		return nil
	}
	if strings.TrimSpace(d.text) == "" {
		if d.exists {
			return a.removeAll(d.path)
		}
		return nil
	}
	status := Updated
	if !d.exists {
		status = Added
	}
	if err := a.write(d.path, []byte(d.text)); err != nil {
		return err
	}
	a.report(d.path, status)
	return nil
}

func (a *applier) abs(rel string) string {
	return filepath.Join(a.root, filepath.FromSlash(rel))
}

func (a *applier) exists(rel string) bool {
	_, err := os.Lstat(a.abs(rel))
	return err == nil
}

func (a *applier) isLink(rel string) bool {
	info, err := os.Lstat(a.abs(rel))
	return err == nil && info.Mode()&fs.ModeSymlink != 0
}

// sameAs reports whether the folder rel holds exactly the skill's files.
func (a *applier) sameAs(rel string, s *skill) bool {
	if s == nil {
		return false
	}
	want := map[string][]byte{}
	err := s.walk(a.src, func(sub string, data []byte) error {
		want[sub] = data
		return nil
	})
	if err != nil {
		return false
	}
	have := 0
	err = filepath.WalkDir(a.abs(rel), func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		sub, err := filepath.Rel(a.abs(rel), p)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if w, ok := want[filepath.ToSlash(sub)]; !ok || !bytes.Equal(w, data) {
			return errors.New("differs")
		}
		have++
		return nil
	})
	return err == nil && have == len(want)
}

func (a *applier) write(rel string, data []byte) error {
	if a.dryRun {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(a.abs(rel)), 0o755); err != nil {
		return err
	}
	return os.WriteFile(a.abs(rel), data, 0o644)
}

func (a *applier) removeAll(rel string) error {
	a.report(rel, Removed)
	a.removed = true
	if a.dryRun {
		return nil
	}
	return os.RemoveAll(a.abs(rel))
}

func (a *applier) report(rel, status string) {
	a.changes = append(a.changes, Change{Path: rel, Status: status})
}

// skillNamed is the skill called name, or nil when forge doesn't install it
// now.
func skillNamed(skills []skill, name string) *skill {
	for i := range skills {
		if skills[i].name == name {
			return &skills[i]
		}
	}
	return nil
}

// removeEmptyDirs removes the empty folders under dir, deepest first.
func removeEmptyDirs(dir string) {
	var dirs []string
	_ = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err == nil && d.IsDir() && p != dir {
			dirs = append(dirs, p)
		}
		return nil
	})
	sort.Sort(sort.Reverse(sort.StringSlice(dirs)))
	for _, d := range dirs {
		_ = os.Remove(d)
	}
}
