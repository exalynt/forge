# forge

For reusable agent context, prompts, skills, sub agents, etc... across projects

## claude/

Global Claude Code configuration: engineering philosophy, the feature stability
model (Prototype → Alpha → Beta → GA), and the skills and agents that apply it.

```text
claude/
├── CLAUDE.md                    # short global philosophy, always loaded
├── rules/
│   ├── engineering-philosophy.md
│   └── stability.md
├── skills/
│   ├── stability/               # level definitions and review classification
│   └── promote/                 # /promote <feature> <level>, manual only
├── agents/
│   └── stability-reviewer.md    # read-only review against a stability level
├── install.sh                   # symlinks the above into ~/.claude
├── stacks/
│   └── go-http/rules/           # general rules for Go HTTP API projects
└── apply.sh                     # copies a stack's rules into a project
```

### Using it in your personal Claude config (optional)

Claude Code reads global configuration from `~/.claude/`. To use this config
across all your projects, symlink it there. Your copy then stays in sync with
`git pull`.

```bash
./claude/install.sh
```

The script creates these links and nothing else:

```text
~/.claude/CLAUDE.md -> forge/claude/CLAUDE.md
~/.claude/rules     -> forge/claude/rules
~/.claude/agents    -> forge/claude/agents
~/.claude/skills    -> forge/claude/skills
```

It never overwrites or deletes anything, and it's safe to re-run. If you
already have one of these as a real file or directory:

- **Directories** (`rules/`, `skills/`, `agents/`): it links each entry inside
  instead, such as `~/.claude/skills/stability`, so your existing skills stay.
  Re-run it after a new skill, rule, or agent is added here.
- **`CLAUDE.md`**: it reports a conflict and skips it. Either move your
  personal instructions into `rules/` (for example
  `~/.claude/rules/personal.md`) and re-run, or keep your file and add
  `@<path-to-forge>/claude/CLAUDE.md` to it to import this one.

To link by hand instead, use `ln -s "$PWD/claude/<item>" ~/.claude/<item>`
for whichever items you want. If you set `CLAUDE_CONFIG_DIR`, the script
links there instead of `~/.claude`.

Start a new Claude Code session afterwards. To uninstall, remove the links:

```bash
find ~/.claude -maxdepth 2 -type l -lname '*/forge/claude/*' -print -delete
```

### Project stacks

A stack is a set of rules for one kind of project, applied per repository
rather than globally. `go-http` holds general conventions for Go HTTP API
projects: preferred packages, layers and boundaries, repositories, service
errors, REST API design, configuration, and tests. Each rule loads only when
Claude works on files it applies to.

```bash
./claude/apply.sh go-http path/to/repo
```

This copies the rules into `path/to/repo/.claude/rules/go-http/`. Commit them
with the project so everyone working on it gets them. To change a rule, edit it
here and re-apply. Don't hand-edit the copies. Re-running is safe: it reports
each file as `added`, `updated`, or `ok`, and never deletes anything.

Project-specific architecture and domain rules don't belong here. Keep them in
each project's own `CLAUDE.md` or `.claude/`. A project should also say where it
declares feature stability, for example `.exalynt/features/<feature>.yaml`.
