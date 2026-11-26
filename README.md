# forge

Exalynt's way of working, for coding agents: how we engineer, and skills for
the workflows that apply it. It's installed into each repository and committed
with it, so everyone working on the repository gets it, whatever agent they
use.

Forge says *how* to do the work. What the software is meant to do, and how
much each part of it needs, is the project's intention. Where a repository
records that with [Docket](https://github.com/exalynt/docket), forge depends on
it: Docket holds each capability's stability level, its business rules, the
personas it serves, its gaps and decisions, and what each level expects.
Forge never recreates any of that. Without Docket, forge's skills still work,
taking the level and the rules from the user and recording nothing beyond the
code and tests the user approves.

```text
forge/
├── AGENTS.md                 # the block written into a repository's AGENTS.md
├── install.sh                # installs forge into a repository
├── skills/
│   ├── implement/            # build a feature, improvement, or fix; a worktree and PR when asked
│   ├── commit/               # small, related commits
│   ├── pr/                   # commit, push, and open or update a PR labeled with its level
│   ├── qa/                   # QA a capability end to end against its rules
│   ├── ux/                   # walk a main flow as a type of user
│   ├── use/                  # usability test, one player per person (role.md is their brief)
│   ├── stability-review/     # review work against a level; needs Docket
│   ├── promote/              # move a capability up a level; needs Docket
│   ├── testing/              # how we write tests
│   └── observability/        # how we add traces, logs, and metrics
├── stacks/
│   ├── go-http/              # Go HTTP APIs: SKILL.md and references/
│   └── react/                # React + MUI frontends: SKILL.md and references/
└── scripts/playwright/       # browser tools for qa, ux, and use
```

## Installing it in a repository

```bash
./install.sh path/to/repo
./install.sh path/to/repo --stack go-http
./install.sh path/to/repo --stack go-http --stack react
```

In a repository with Docket, install Docket's kit as well. They sit side by
side, and neither touches the other's files:

```bash
cd path/to/repo
docket kit
path/to/forge/install.sh . --stack go-http
```

It writes these into the repository, to be committed with it:

| Written                        | What for                                                                       |
| ------------------------------ | ------------------------------------------------------------------------------ |
| `AGENTS.md`                    | Forge's block, between `<!-- forge:start -->` and `<!-- forge:end -->` markers |
| `.agents/skills/<name>/`       | A copy of each skill, where Codex, Cursor, Gemini CLI, and others look         |
| `.agents/forge.lock`           | Forge's version, the stacks, and each skill it installed                       |
| `CLAUDE.md`                    | An `@AGENTS.md` line, if the file exists without one                           |

Claude Code finds the skills through forge's block in `AGENTS.md` until it
reads `.agents/skills/`. It reads `AGENTS.md` itself only when a repository has
no `CLAUDE.md`, so where there is one, the import line makes it read
`AGENTS.md` too. The script never creates `CLAUDE.md`. Copies an earlier forge
put in `.claude/skills/` are removed.

An existing `AGENTS.md` keeps everything it had: forge's block is added at the
end, or replaced where it already is. If the markers are unbalanced, the script
stops and changes nothing. It changes only what's forge's: its block, and the
skill folders it installed, which it makes match forge exactly. It never
touches another skill, such as `docket`, or another block, such as Docket's.
If a skill folder with one of forge's names is already there and forge didn't
install it, it stops rather than overwrite it.

It reports each file as `added`, `updated`, `removed`, or `ok`, then a summary.

Don't hand-edit the copies or the block, since the next install overwrites
them. Anything particular to a project goes in its own instructions, outside
forge's block, and wins over forge's.

## With and without Docket

A repository has Docket when there's a `docket.yaml` at the working directory
or above it, within the repository. Forge's `AGENTS.md`, under "Where
intention lives", is where the two modes are defined, and every skill follows
it.

- **With Docket**, skills read each capability's level, rules, personas, gaps,
  and decisions with `docket touches`, `docket show`, and `docket list`, and
  what a level expects with `docket rubric`. A business rule's source of truth
  is its Docket rule doc, and the Gherkin `Rule:` that tests it carries a
  `docket:verifies` annotation. Anything a skill records, it records the
  Docket kit's way, after a person confirms it.
- **Without Docket**, a capability's level is whatever the user says, and
  acceptance tests are the specification. Nothing is recorded beyond the code
  and tests the user approves.

`stability-review` and `promote` are about levels, so they need Docket.
Without it they stop and say so: a repository adopts it with `docket init` and
`docket kit`. Every other skill works either way.

## Stacks

A stack is the conventions for one kind of project, installed as one more
skill with `--stack`. `go-http` covers Go HTTP API projects: preferred
packages, layers and boundaries, repositories, service errors, REST API
design, pagination, health checks, configuration, tests, and CI. `react` covers React + MUI frontends:
page and component layout, reusing MUI and existing components, mobile
responsiveness, light and dark mode, accessibility, services for external
APIs and SDKs, loading states and fetching data, lists with search, filters,
sorting, and pagination, forms, and CI. Each stack's `SKILL.md` indexes its `references/`, and lists the files it
applies to, which Claude Code and Cursor use to load it when you work on them.

Project-specific architecture and domain rules don't belong here. Keep them in
each project's own instructions.

## Updating

Re-run the script. Without `--stack`, it keeps the stacks the lock lists. With
any `--stack`, it installs exactly those, so passing a different set is how
you drop one, and `--no-stacks` removes them all. Any skill forge no longer
has is removed too. The lock records forge's version, so you can see which one
a repository has.

```bash
path/to/forge/install.sh .                 # update, keeping the stacks
path/to/forge/install.sh . --stack react   # swap the stacks for react
path/to/forge/install.sh . --no-stacks     # remove every stack
```

To check in CI that a repository's copy matches forge, run it with `--check`.
It changes nothing, and exits 1 if anything would change. Like an update, it
checks the stacks the lock lists unless given others.

```bash
path/to/forge/install.sh . --check
```

## Browser tools for qa, ux, and use

`qa` and `ux` drive a capability through its UI with whichever browser tool
the agent has. Without one, they report the UI as not verified. `use` needs a
browser for each person in a test.
[Playwright MCP](https://github.com/microsoft/playwright-mcp) gives them that.
Set it up once per machine:

```bash
./scripts/playwright/setup.sh
```

It installs `@playwright/mcp` globally with npm, installs Google Chrome if it's
missing (the server's default browser; this step may ask for `sudo`), and
prints the command and arguments for five servers: `playwright`, the one used
by default, and `playwright-1` through `playwright-4` for when a skill calls
for more, such as several sign-ins at once or the people in a `use` test. Each
is its own browser and sign-in. Register all five with your agent, with the
same command and arguments:

- **Claude Code**: for each name,
  `claude mcp add --scope user <name> -- <command> <arguments>`. Re-running
  needs `claude mcp remove --scope user <name>` first.
- **Codex**: a table for each name in `~/.codex/config.toml`:

  ```toml
  [mcp_servers.playwright]
  command = "/path/to/node"
  args = ["/path/to/@playwright/mcp/cli.js", "--isolated", "--executable-path", "/path/to/forge/scripts/playwright/chrome.sh"]
  ```

- **Cursor**: an entry for each name under `mcpServers` in `.cursor/mcp.json`
  for one project, or `~/.cursor/mcp.json` for all of them, with the same
  `command` and `args`.

Let the browser tools run without asking each time, but keep asking before
`browser_run_code_unsafe`, which runs arbitrary code on your machine, and
before `browser_file_upload` and `browser_drop`, which can send local files to
a page. Set that in your agent's permissions.

- Arguments to `setup.sh` are passed to the server. The browser is headed by
  default, which works under WSLg, so add `--headless` where there's no
  display. Pass `--browser firefox` to use Firefox instead of Chrome.
  Playwright's bundled Chromium isn't one of the server's options.
- Browser windows arrange themselves across your monitors, so you can follow
  along without them covering your work. `chrome.sh` reads the monitors each
  time a browser opens, with `xrandr` on Linux and WSLg or AppKit on macOS,
  and gives the window an eighth of a monitor, or more where that would be
  smaller than Chrome's minimum window size. Windows fill from the top right,
  right to left and then down, on secondary monitors first and the primary
  last. Once every spot is taken, new ones stack on the same spots in order. A
  spot frees up when its browser closes. Spots are shared across sessions, but
  only count browsers it opened, not your own windows. Where it can't read the
  monitors, the browser opens as it otherwise would, and with `--headless` or
  another browser it isn't used at all. The skills set each test width with
  the browser's resize tool, which doesn't move or resize the window.
- Under WSLg, a browser opening doesn't keep the focus. `keep-focus.ps1` runs
  on the Windows side for the browser's first 10 seconds, and sends focus
  straight back to the window you were in. Clicking into a browser window
  during those seconds sends you back too. On macOS and native Linux, a new
  browser can still take the focus.
- The server always gets `--isolated`, so each session starts with a fresh,
  in-memory profile. A sign-in from an earlier run doesn't carry into
  first-use, signed-out, or other-tenant checks. The cost is signing in again
  each session.
- Agents that start MCP servers per session, as Claude Code does, give every
  session its own five browsers, so separate sessions never share one.
  Subagents share their session's, so each subagent driving a browser needs a
  numbered server of its own. Browsers don't separate data: every agent and
  session uses the same app and database. Killing Chrome by name closes every
  session's browser, so the skills close only their own.
- The server writes screenshots, snapshots, and console logs to
  `.playwright-mcp/` in the project, and refuses paths outside the project or
  that folder. `qa`, `ux`, and `use` keep everything they create there and add
  `/.playwright-mcp/` to the project's `.gitignore` if it's missing.
- It registers Node and the package by absolute path, so the server starts
  even when the agent's environment lacks nvm's PATH. Re-run it, and
  re-register, after switching Node versions. Re-running also updates the
  package.

## Uninstalling

Delete forge's block from `AGENTS.md`, from `<!-- forge:start -->` to
`<!-- forge:end -->`; each skill folder `.agents/forge.lock` lists, from
`.agents/skills/`; and `.agents/forge.lock` itself.
Remove the `@AGENTS.md` line from `CLAUDE.md` if nothing else in `AGENTS.md`
needs it.

## Known gaps

- **Claude Code doesn't load forge's skills by itself.** It doesn't read
  `.agents/skills/` yet, so it reads a skill only because forge's block in
  `AGENTS.md` says where it is, and `/implement`, `/qa` and the rest aren't in
  its slash menu: ask for one by name instead, such as "use implement". Once
  Claude Code reads `.agents/skills/`, this closes with no change here.
