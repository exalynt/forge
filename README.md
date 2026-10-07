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
└── install.sh                   # symlinks the above into ~/.claude
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

Project-specific architecture and domain rules don't belong here. Keep them in
each project's own `CLAUDE.md` or `.claude/`. A project should also say where it
declares feature stability, for example `.exalynt/features/<feature>.yaml`.
