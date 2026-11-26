package install

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/exalynt/forge"
)

// everything is the kit's every skill: each language with all its stacks.
func everything(t *testing.T) []skill {
	t.Helper()
	langs, err := Languages(forge.Files)
	if err != nil {
		t.Fatal(err)
	}
	want := Want{Targets: []string{TargetAgents}, Languages: sortedKeys(langs), Stacks: allStacks(langs)}
	_, skills, err := resolve(forge.Files, want)
	if err != nil {
		t.Fatal(err)
	}
	return skills
}

// isStack reports whether s is a stack, and if so its language's skill name.
func isStack(s skill) (string, bool) {
	parts := strings.Split(s.dir, "/")
	if len(parts) == 4 && parts[0] == languagesDir && parts[2] == stacksDir {
		return parts[1], true
	}
	return "", false
}

var linkPattern = regexp.MustCompile(`\]\(([^)\s#]+)(#[^)]*)?\)`)

// TestKitLinks checks that every relative link in the kit resolves to a file
// that's installed whenever the linking skill is: the workflow skills and
// universal skills always, a language with itself, and a stack with its
// language.
func TestKitLinks(t *testing.T) {
	skills := everything(t)
	installed := map[string]bool{}
	for _, s := range skills {
		if err := s.walk(forge.Files, func(sub string, _ []byte) error {
			installed[path.Join(s.name, sub)] = true
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range skills {
		guaranteed := []string{s.name, standardsDir}
		for _, o := range skills {
			if strings.HasPrefix(o.dir, workflowsDir+"/") || strings.HasPrefix(o.dir, universalDir+"/") {
				guaranteed = append(guaranteed, o.name)
			}
		}
		if lang, ok := isStack(s); ok {
			guaranteed = append(guaranteed, lang)
		}
		err := s.walk(forge.Files, func(sub string, data []byte) error {
			if !strings.HasSuffix(sub, ".md") {
				return nil
			}
			for _, m := range linkPattern.FindAllStringSubmatch(string(data), -1) {
				link := m[1]
				if strings.Contains(link, ":") {
					continue
				}
				target := path.Clean(path.Join(s.name, path.Dir(sub), link))
				owner, _, _ := strings.Cut(target, "/")
				switch {
				case !slices.Contains(guaranteed, owner):
					t.Errorf("%s/%s links to %s, in %s, which isn't always installed with it", s.name, sub, link, owner)
				case !installed[target]:
					t.Errorf("%s/%s links to %s, which doesn't exist", s.name, sub, link)
				}
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestIndexes checks that each skill's index links every layer's file for
// each topic it or its language has, so an agent finds a topic's whole chain
// from one place.
func TestIndexes(t *testing.T) {
	skills := everything(t)
	byName := map[string]skill{}
	for _, s := range skills {
		byName[s.name] = s
	}
	refs := func(s skill) []string {
		entries, err := readDirNames(path.Join(s.dir, "references"))
		if err != nil {
			t.Fatal(err)
		}
		return entries
	}
	for _, s := range skills {
		var chain []skill
		switch lang, stack := isStack(s); {
		case s.name == standardsDir:
			chain = []skill{s}
		case stack:
			chain = []skill{byName[standardsDir], byName[lang], s}
		case strings.HasPrefix(s.dir, languagesDir+"/"):
			chain = []skill{byName[standardsDir], s}
		default:
			continue
		}
		index, err := readFile(path.Join(s.dir, "SKILL.md"))
		if err != nil {
			t.Fatal(err)
		}
		var topics []string
		for _, layer := range chain {
			if layer.name != standardsDir || s.name == standardsDir {
				for _, topic := range refs(layer) {
					topics = appendNew(topics, topic)
				}
			}
		}
		for _, topic := range topics {
			for _, layer := range chain {
				if !slices.Contains(refs(layer), topic) {
					continue
				}
				link := "references/" + topic
				if layer.name != s.name {
					link = "../" + layer.name + "/" + link
				}
				if !strings.Contains(index, "]("+link+")") {
					t.Errorf("%s's index doesn't link %s for the %s topic", s.name, link, strings.TrimSuffix(topic, ".md"))
				}
			}
		}
	}
}

func readDirNames(dir string) ([]string, error) {
	entries, err := readDir(dir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if !e.IsDir() {
			names = append(names, e.Name())
		}
	}
	return names, nil
}

// TestInstall installs a stack for both targets, checks it's up to date, and
// removes it again.
func TestInstall(t *testing.T) {
	root := t.TempDir()
	want := Want{Targets: []string{TargetAgents, TargetClaude}, Stacks: []string{"go/http"}}
	res, err := Apply(root, forge.Files, want, "test", false)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(res.Languages, []string{"go"}) || !slices.Equal(res.Stacks, []string{"go/http"}) {
		t.Errorf("installed languages %v and stacks %v, want [go] and [go/http]", res.Languages, res.Stacks)
	}
	for _, dir := range []string{".agents/skills", ".claude/skills"} {
		for _, f := range []string{"standards/SKILL.md", "testing/SKILL.md", "go/SKILL.md", "go-http/references/pagination.md"} {
			if _, err := os.Stat(filepath.Join(root, dir, f)); err != nil {
				t.Errorf("%s/%s wasn't installed", dir, f)
			}
		}
		if _, err := os.Stat(filepath.Join(root, dir, "go", "stacks")); err == nil {
			t.Errorf("%s/go has its stacks in it", dir)
		}
	}
	agents, _ := os.ReadFile(filepath.Join(root, "AGENTS.md"))
	if !strings.Contains(string(agents), "This repository has `standards` and `go` → `go-http`.") {
		t.Error("the block doesn't name the conventions installed")
	}
	lock, _ := os.ReadFile(filepath.Join(root, agentsLock))
	if !strings.Contains(string(lock), "# languages go\n# stacks go/http\n") {
		t.Errorf("the lock doesn't list the language and stack:\n%s", lock)
	}

	l, err := ReadLock(root)
	if err != nil {
		t.Fatal(err)
	}
	res, err = Apply(root, forge.Files, Want{Targets: l.Targets, Languages: l.Languages, Stacks: l.Stacks}, "test", true)
	if err != nil {
		t.Fatal(err)
	}
	if res.Stale() {
		t.Errorf("a fresh install isn't up to date: %v", res.Changes)
	}

	if _, err := Apply(root, forge.Files, Want{}, "test", false); err != nil {
		t.Fatal(err)
	}
	left, _ := os.ReadDir(root)
	if len(left) > 0 {
		t.Errorf("removing everything left %v", left)
	}
}

// TestOldLock updates an install from before languages, whose stacks were
// go-http and react.
func TestOldLock(t *testing.T) {
	root := t.TempDir()
	old := "# forge v0.1.0\n# targets agents\n# stacks go-http react\nimplement\ngo-http\nreact\n"
	for f, text := range map[string]string{
		agentsLock:                          old,
		".agents/skills/go-http/SKILL.md":   "old",
		".agents/skills/react/SKILL.md":     "old",
		".agents/skills/implement/SKILL.md": "old",
	} {
		p := filepath.Join(root, filepath.FromSlash(f))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	l, err := ReadLock(root)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(l.Stacks, []string{"go/http", "typescript/react"}) || !slices.Equal(l.Languages, []string{"go", "typescript"}) {
		t.Fatalf("read stacks %v and languages %v", l.Stacks, l.Languages)
	}
	if _, err := Apply(root, forge.Files, Want{Targets: l.Targets, Languages: l.Languages, Stacks: l.Stacks}, "test", false); err != nil {
		t.Fatal(err)
	}
	for f, want := range map[string]bool{
		"react":                     false,
		"go-http/references/ci.md":  true,
		"go/SKILL.md":               true,
		"typescript/SKILL.md":       true,
		"typescript-react/SKILL.md": true,
	} {
		_, err := os.Stat(filepath.Join(root, ".agents/skills", f))
		if (err == nil) != want {
			t.Errorf(".agents/skills/%s: exists %v, want %v", f, err == nil, want)
		}
	}
}

func TestStackNames(t *testing.T) {
	_, _, err := resolve(forge.Files, Want{Targets: []string{TargetAgents}, Stacks: []string{"go/nope"}})
	if err == nil || !strings.Contains(err.Error(), "go/http") {
		t.Errorf("an unknown stack's error should list the stacks: %v", err)
	}
	_, _, err = resolve(forge.Files, Want{Targets: []string{TargetAgents}, Languages: []string{"cobol"}})
	if err == nil || !strings.Contains(err.Error(), "typescript") {
		t.Errorf("an unknown language's error should list the languages: %v", err)
	}
	got, _, err := resolve(forge.Files, Want{Targets: []string{TargetAgents}, Stacks: []string{"react"}})
	if err != nil || !slices.Equal(got.Stacks, []string{"typescript/react"}) || !slices.Equal(got.Languages, []string{"typescript"}) {
		t.Errorf("react resolved to %v and %v (%v)", got.Stacks, got.Languages, err)
	}
}

func TestConventions(t *testing.T) {
	tests := []struct {
		want Want
		line string
	}{
		{Want{}, "This repository has `standards` only."},
		{Want{Languages: []string{"go"}}, "This repository has `standards` and `go`."},
		{
			Want{Languages: []string{"go", "typescript"}, Stacks: []string{"go/http", "typescript/react"}},
			"This repository has `standards`, `go` → `go-http`, and `typescript` → `typescript-react`.",
		},
	}
	for _, tt := range tests {
		if got := conventions(tt.want); got != tt.line {
			t.Errorf("conventions(%v) = %q, want %q", tt.want, got, tt.line)
		}
	}
}

func readFile(p string) (string, error) {
	data, err := fs.ReadFile(forge.Files, p)
	return string(data), err
}

// readDir lists a folder of the kit, or nothing when it doesn't exist.
func readDir(p string) ([]fs.DirEntry, error) {
	entries, err := fs.ReadDir(forge.Files, p)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	return entries, err
}
