package install

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"strings"
)

// Where forge's skills are in its kit. Every install has the workflow skills
// in workflows/ and the universal conventions in universal/: standards,
// testing, and observability. A language is languages/<language>/, and a
// stack under it languages/<language>/stacks/<stack>/, chosen per
// repository.
const (
	workflowsDir = "workflows"
	universalDir = "universal"
	languagesDir = "languages"
	stacksDir    = "stacks"
	// standards is the universal skill that languages and stacks build on.
	standardsDir = "standards"
)

// renamed maps the stack names an earlier forge used to their names now, so
// its locks and flags keep working.
var renamed = map[string]string{
	"go-http": "go/http",
	"react":   "typescript/react",
}

// skill is one skill to install: its name, its folder in src, and a folder
// inside it that isn't part of it (a language's stacks).
type skill struct {
	name string
	dir  string
	skip string
}

// walk calls fn with each of the skill's files, by its path inside the skill.
func (s skill) walk(src fs.FS, fn func(sub string, data []byte) error) error {
	return fs.WalkDir(src, s.dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if s.skip != "" && p == path.Join(s.dir, s.skip) {
				return fs.SkipDir
			}
			return nil
		}
		data, err := fs.ReadFile(src, p)
		if err != nil {
			return err
		}
		return fn(strings.TrimPrefix(p, s.dir+"/"), data)
	})
}

// StackName is a stack's name as forge writes it, language/stack, from the
// name given, which may be one an earlier forge used.
func StackName(name string) string {
	if n, ok := renamed[name]; ok {
		return n
	}
	return name
}

// StackLanguage is the language a stack is under.
func StackLanguage(stack string) string {
	lang, _, _ := strings.Cut(stack, "/")
	return lang
}

// stackSkill is the name a stack is installed under: go/http is go-http.
func stackSkill(stack string) string {
	return strings.ReplaceAll(stack, "/", "-")
}

// Languages lists the languages src has, each with its stacks as
// language/stack.
func Languages(src fs.FS) (map[string][]string, error) {
	entries, err := fs.ReadDir(src, languagesDir)
	if err != nil {
		return nil, err
	}
	langs := map[string][]string{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		langs[e.Name()] = nil
		stacks, err := fs.ReadDir(src, path.Join(languagesDir, e.Name(), stacksDir))
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return nil, err
		}
		for _, s := range stacks {
			if s.IsDir() {
				langs[e.Name()] = append(langs[e.Name()], e.Name()+"/"+s.Name())
			}
		}
	}
	return langs, nil
}

// resolve checks want's languages and stacks against src, adds the language
// of each stack, and returns them sorted with every skill to install: the
// workflow skills, the universal ones, each language, then each stack.
func resolve(src fs.FS, want Want) (Want, []skill, error) {
	if len(want.Targets) == 0 {
		want.Languages, want.Stacks = nil, nil
		return want, nil, nil
	}
	known, err := Languages(src)
	if err != nil {
		return Want{}, nil, err
	}
	var langs, stacks []string
	for _, l := range want.Languages {
		if _, ok := known[l]; !ok {
			return Want{}, nil, fmt.Errorf("unknown language %q: forge has %s", l, strings.Join(sortedKeys(known), ", "))
		}
		langs = appendNew(langs, l)
	}
	for _, s := range want.Stacks {
		s = StackName(s)
		if !slices.Contains(known[StackLanguage(s)], s) {
			return Want{}, nil, fmt.Errorf("unknown stack %q: forge has %s", s, strings.Join(allStacks(known), ", "))
		}
		stacks = appendNew(stacks, s)
		langs = appendNew(langs, StackLanguage(s))
	}
	slices.Sort(langs)
	slices.Sort(stacks)
	want.Languages, want.Stacks = langs, stacks

	var skills []skill
	for _, dir := range []string{workflowsDir, universalDir} {
		entries, err := fs.ReadDir(src, dir)
		if err != nil {
			return Want{}, nil, err
		}
		for _, e := range entries {
			if e.IsDir() {
				skills = append(skills, skill{name: e.Name(), dir: path.Join(dir, e.Name())})
			}
		}
	}
	for _, l := range langs {
		skills = append(skills, skill{name: l, dir: path.Join(languagesDir, l), skip: stacksDir})
	}
	for _, s := range stacks {
		lang, name, _ := strings.Cut(s, "/")
		skills = append(skills, skill{name: stackSkill(s), dir: path.Join(languagesDir, lang, stacksDir, name)})
	}
	return want, skills, nil
}

// conventions is the line the block gives for what this repository has
// installed, such as "`standards`, `go` → `go-http`, and `typescript`".
func conventions(want Want) string {
	parts := []string{"`" + standardsDir + "`"}
	for _, l := range want.Languages {
		part := "`" + l + "`"
		var under []string
		for _, s := range want.Stacks {
			if StackLanguage(s) == l {
				under = append(under, "`"+stackSkill(s)+"`")
			}
		}
		if len(under) > 0 {
			part += " → " + strings.Join(under, ", ")
		}
		parts = append(parts, part)
	}
	if len(parts) == 1 {
		return "This repository has `standards` only."
	}
	last := len(parts) - 1
	list := strings.Join(parts[:last], ", ")
	if last > 1 {
		list += ","
	}
	return "This repository has " + list + " and " + parts[last] + "."
}

func appendNew(list []string, s string) []string {
	if slices.Contains(list, s) {
		return list
	}
	return append(list, s)
}

func sortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func allStacks(m map[string][]string) []string {
	var all []string
	for _, l := range sortedKeys(m) {
		all = append(all, m[l]...)
	}
	return all
}
