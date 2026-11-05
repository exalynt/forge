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
│   ├── stability.md
│   ├── observability.md         # OTel traces, logs, and metrics in any language
│   ├── testing.md               # unit, integration, e2e, and acceptance tests by level
│   └── subagents.md             # when to parallelize work with subagents
├── skills/
│   ├── stability/               # level definitions and review classification
│   ├── promote/                 # /promote [low|medium|high] <feature> <level>, manual only
│   ├── commit/                  # /commit [session|all], grouped commits, manual only
│   ├── implement/               # /implement [level] [worktree] <task>, on the current branch; worktree adds own branch and worktree, ends in a PR, manual only
│   ├── pr/                      # /pr [level] [draft], commit, push, open or update a labeled PR, manual only
│   ├── qa/                      # /qa [low|medium|high] <feature> [focus], end-to-end QA, manual only
│   ├── ux/                      # /ux [low|medium|high] <feature> [as <user>], main-flow walkthrough, manual only
│   ├── use/                     # /use [low|medium|high] <flow> [as <role>, ...], usability test with one agent per person, manual only
│   └── personas/                # finds or creates reusable user personas in a project's docs/personas, for /ux and /use
├── agents/
│   ├── stability-reviewer.md    # read-only review against a stability level
│   └── use-role.md              # plays one person in a /use usability test
├── settings.json                # permissions, merged into ~/.claude/settings.json
├── install.sh                   # symlinks the above into ~/.claude, merges settings.json
├── scripts/playwright/
│   ├── setup.sh                 # installs and registers Playwright MCP for /qa, /ux, and /use
│   ├── chrome.sh                # opens each browser window in the next free spot on the monitors
│   └── keep-focus.ps1           # under WSLg, keeps the focus off a browser that's opening
├── stacks/
│   ├── go-http/rules/           # general rules for Go HTTP API projects
│   └── react/rules/             # general rules for React + MUI frontends
└── apply.sh                     # copies a stack's rules into a project
```

### Using it in your personal Claude config (optional)

Claude Code reads global configuration from `~/.claude/`. To use this config
across all your projects, symlink it there. Your copy then stays in sync with
`git pull`.

```bash
./claude/install.sh
```

The script creates these links:

```text
~/.claude/CLAUDE.md -> forge/claude/CLAUDE.md
~/.claude/rules     -> forge/claude/rules
~/.claude/agents    -> forge/claude/agents
~/.claude/skills    -> forge/claude/skills
```

It also merges `settings.json` into `~/.claude/settings.json`, or copies it
there if you don't have one. It doesn't link it, because Claude Code saves
your "always allow" approvals to that file, and those shouldn't end up in
forge. The merge adds forge's keys and list entries, such as its permission
rules, and never changes or removes yours. It needs `jq`.

It never overwrites or deletes anything, and it's safe to re-run. It reports
each item as `added`, `updated` (settings merged in), `ok`, `STALE` (a link to
something no longer in forge), or `CONFLICT`, then prints a summary. Edits to linked files take effect
immediately, with no re-run needed. If you
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

### Browser tools for /qa, /ux, and /use (optional)

`/qa` and `/ux` drive a feature through its UI with whichever browser tools
the session has. Without any, they report the UI as not verified. `/use` needs
Playwright MCP, because it gives each person in a test their own browser. To
give them [Playwright MCP](https://github.com/microsoft/playwright-mcp), set it
up once per machine:

```bash
./claude/scripts/playwright/setup.sh
```

It installs `@playwright/mcp` globally with npm, installs Google Chrome if it's
missing (the server's default browser; this step may ask for `sudo`), and
registers the server for all projects with `claude mcp add --scope user`.
`install.sh` can't do this part, because Claude Code keeps MCP servers in
`~/.claude.json`, not in a directory it can link. It finds `claude` on your
PATH, or the CLI bundled with the VS Code extension. Set `CLAUDE_BIN` to use
another one.

- Arguments are passed to the server. The browser is headed by default, which
  works under WSLg, so add `--headless` where there's no display. Pass
  `--browser firefox` to use Firefox instead of Chrome. Playwright's bundled
  Chromium isn't one of the server's options.
- Browser windows arrange themselves across your monitors, so you can follow
  along without them covering your work. `chrome.sh` reads the
  monitors each time a browser opens, with `xrandr` on Linux and WSLg or AppKit
  on macOS, and gives the window an eighth of a monitor, or more where that would
  be smaller than Chrome's minimum window size. Windows fill from the top
  right, right to left and then down, on secondary monitors first and the primary
  last. Once every spot is taken, new ones stack on the same spots in order. A
  spot frees up when its browser closes. Spots are shared across sessions, but
  only count browsers it opened, not your own windows. Where it can't read the
  monitors, the browser opens as it otherwise would, and with `--headless` or
  another browser it isn't used at all. The skills set each test width with the
  browser's resize tool, which doesn't move or resize the window.
- Under WSLg, a browser opening doesn't keep the focus. `keep-focus.ps1` runs on
  the Windows side for the browser's first 10 seconds, and sends focus straight
  back to the window you were in. Clicking into a browser window during those
  seconds sends you back too. On macOS and native Linux, a new browser can still
  take the focus.
- The server always gets `--isolated`, so each session starts with a fresh,
  in-memory profile. A sign-in from an earlier run doesn't carry into
  first-use, signed-out, or other-tenant checks. The cost is signing in again
  each session.
- It registers five servers, each its own browser and sign-in: `playwright`, the
  one used by default, and `playwright-1` through `playwright-4` for when a skill
  calls for more, such as several sign-ins at once or the people in a `/use`
  test. Every Claude session starts its own five, so separate sessions never
  share a browser, and one can't reach another's through its browser tools.
  Subagents share their session's, so each subagent driving a browser needs a
  numbered server of its own. They belong to the main session, so a person's
  sign-in survives their agent waiting between turns. Browsers don't separate
  data: every agent and session uses the same app and database. Killing Chrome
  by name closes every session's browser, so the skills close only their own.
- The server writes screenshots, snapshots, and console logs to
  `.playwright-mcp/` in the project, and refuses paths outside the project or
  that folder. `/qa` and `/ux` keep everything they create there and add
  `/.playwright-mcp/` to the project's `.gitignore` if it's missing.
- It registers Node and the package by absolute path, so the server starts even
  when Claude Code's environment lacks nvm's PATH. Re-run it after switching
  Node versions. Re-running also updates the package and applies any new flags.
- Forge's `settings.json`, merged in by `install.sh`, lets the browser tools run
  without asking each time. It still asks before `browser_run_code_unsafe`,
  which runs arbitrary code on your machine, and before `browser_file_upload`
  and `browser_drop`, which can send local files to a page.

Start a new Claude Code session afterwards. `claude mcp list` should show
`playwright` as connected.

### Project stacks

A stack is a set of rules for one kind of project, applied per repository
rather than globally. `go-http` holds general conventions for Go HTTP API
projects: preferred packages, layers and boundaries, repositories, service
errors, REST API design, configuration, and tests. `react` holds conventions
for React + MUI frontends: page and component layout, reusing MUI and existing
components, mobile responsiveness, light and dark mode, services for external
APIs and SDKs, and CI. Each rule loads only when Claude works on files it
applies to.

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
