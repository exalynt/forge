// Package install writes forge into a repository for the agents it uses, and
// takes it out again: forge's block in the agents' instructions, its skills
// in each target's skill folder, and the lock that records what was
// installed.
package install

import (
	"fmt"
	"slices"
	"strings"
)

// Target names, as the lock records them.
const (
	TargetAgents = "agents"
	TargetClaude = "claude"
)

// Agent is a name a user can pass to --agent, and what that agent reads.
type Agent struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Reads string `json:"reads"`
	Note  string `json:"note,omitempty"`
}

// Target is a place agents read: a skill folder, and the file forge's block
// goes in.
type Target struct {
	Name   string  `json:"name"`
	Skills string  `json:"skills"`
	Block  string  `json:"block"`
	Agents []Agent `json:"agents"`
}

// Targets is every target forge installs to, with the agents each one
// serves. A new agent that reads one of these places is a row here; one that
// reads somewhere else is a new target.
var Targets = []Target{
	{
		Name:   TargetAgents,
		Skills: ".agents/skills",
		Block:  "AGENTS.md",
		Agents: []Agent{
			{Name: "codex", Label: "Codex", Reads: "AGENTS.md, .agents/skills/"},
			{Name: "cursor", Label: "Cursor", Reads: "AGENTS.md, .agents/skills/"},
			{
				Name: "gemini", Label: "Gemini CLI", Reads: "GEMINI.md, or AGENTS.md when configured; .agents/skills/",
				Note: `Gemini CLI reads AGENTS.md only when its context.fileName setting includes it: add "context": {"fileName": ["AGENTS.md", "GEMINI.md"]} to .gemini/settings.json.`,
			},
			{Name: "windsurf", Label: "Windsurf", Reads: "AGENTS.md, .agents/skills/"},
			{
				Name: "copilot", Label: "GitHub Copilot", Reads: "AGENTS.md; .agents/skills/ in the CLI and cloud agent",
				Note: "Copilot in VS Code reads skills only from .github/skills/, so there it reads a skill only because forge's block says where it is.",
			},
			{Name: "agents", Label: "Any agent that reads AGENTS.md", Reads: "AGENTS.md, .agents/skills/"},
		},
	},
	{
		Name:   TargetClaude,
		Skills: ".claude/skills",
		Block:  "CLAUDE.md, or AGENTS.md through an @AGENTS.md line when agents is installed too",
		Agents: []Agent{
			{Name: "claude", Label: "Claude Code", Reads: "CLAUDE.md, .claude/skills/"},
		},
	},
}

// TargetNamed returns the target called name.
func TargetNamed(name string) (Target, bool) {
	i := slices.IndexFunc(Targets, func(t Target) bool { return t.Name == name })
	if i < 0 {
		return Target{}, false
	}
	return Targets[i], true
}

// AgentNames lists every name --agent takes.
func AgentNames() []string {
	var names []string
	for _, t := range Targets {
		for _, a := range t.Agents {
			names = append(names, a.Name)
		}
	}
	return names
}

// FindAgent returns the agent called name and its target, or an error listing
// the names forge supports.
func FindAgent(name string) (Agent, Target, error) {
	for _, t := range Targets {
		for _, a := range t.Agents {
			if a.Name == name {
				return a, t, nil
			}
		}
	}
	return Agent{}, Target{}, fmt.Errorf("unknown agent %q: forge supports %s (see forge agents)", name, strings.Join(AgentNames(), ", "))
}

// TargetsFor returns the targets the named agents read, in the table's order.
func TargetsFor(agents []string) ([]string, error) {
	var targets []string
	for _, name := range agents {
		_, t, err := FindAgent(name)
		if err != nil {
			return nil, err
		}
		if !slices.Contains(targets, t.Name) {
			targets = append(targets, t.Name)
		}
	}
	return ordered(targets), nil
}

// NotesFor returns the notes for the agents a target serves, or for the named
// agents when there are any.
func NotesFor(targets, agents []string) []string {
	var notes []string
	for _, t := range Targets {
		if !slices.Contains(targets, t.Name) {
			continue
		}
		for _, a := range t.Agents {
			if a.Note == "" {
				continue
			}
			if len(agents) > 0 && !slices.Contains(agents, a.Name) && !slices.Contains(agents, TargetAgents) {
				continue
			}
			notes = append(notes, a.Name+": "+a.Note)
		}
	}
	return notes
}

// ordered sorts target names in the table's order, dropping any it doesn't
// know.
func ordered(names []string) []string {
	var out []string
	for _, t := range Targets {
		if slices.Contains(names, t.Name) {
			out = append(out, t.Name)
		}
	}
	return out
}
