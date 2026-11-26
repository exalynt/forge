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
├── AGENTS.md                 # the block written into a repository's instructions
├── skills/                   # each skill forge installs
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
├── forge.go                  # embeds AGENTS.md, skills/, and stacks/ in the binary
├── cmd/forge/                # runs the forge command
├── internal/commands/        # each command in <name>.go, its subcommands in <name>/
├── internal/install/         # where each agent's files go, and how they're written
├── internal/playwright/      # Playwright MCP for qa, ux, and use, for forge browser
└── scripts/release/          # builds the release archives
```

## Installing forge

Forge is one command that carries its skills and stacks with it. Install it
from a release, or with Go.

**From the latest build**, with the GitHub CLI signed in to an account that
can read the repository. The `latest` pre-release is rebuilt from every commit
on `main` that passes CI, for Linux, macOS and Windows, on amd64 and arm64
(Windows on amd64), with the same file names each time:

```bash
gh release download latest --repo exalynt/forge --pattern "forge_latest_linux_amd64.tar.gz" --clobber   # or darwin_arm64, …
tar -xzf forge_latest_linux_amd64.tar.gz
install forge_latest_linux_amd64/forge ~/.local/bin/   # anywhere on your PATH
forge version                                          # the commit it was built from, as a pseudo-version
```

Once there are versioned releases, download one the same way with its
version in place of `latest`.

**With Go.** The repository is private, so tell Go not to look for it through
the public proxy. It needs access to `github.com/exalynt/forge` through git:

```bash
export GOPRIVATE=github.com/exalynt
go install github.com/exalynt/forge/cmd/forge@latest
```

## Installing it in a repository

```bash
cd path/to/repo
forge install                              # for every agent that reads AGENTS.md
forge install --stack go-http              # with a stack
forge install --agent claude --agent codex # for the agents the repository uses
```

Commit what it writes, so everyone working on the repository gets it. In a
repository with Docket, install Docket's kit as well. They sit side by side,
and neither touches the other's block or skill.

### Choosing agents

Each agent reads its own place, so forge installs to **targets**, one per
place, and `--agent` names the agents you use. A first install with no
`--agent` installs the `agents` target. On an existing install, `install
--agent` adds targets. `forge agents` lists every name.

| Target   | `--agent` names                                         | Skills                   | Block                                                                       |
| -------- | ------------------------------------------------------- | ------------------------ | --------------------------------------------------------------------------- |
| `agents` | `codex`, `cursor`, `gemini`, `windsurf`, `copilot`, `agents` | `.agents/skills/<name>/` | `AGENTS.md`                                                                 |
| `claude` | `claude`                                                | `.claude/skills/<name>/` | `CLAUDE.md`, or `AGENTS.md` through an `@AGENTS.md` line with `agents` too |

- With the `agents` target, forge's block goes in `AGENTS.md`. With `claude`
  as well, `CLAUDE.md` gets an `@AGENTS.md` line, and is created if it's
  missing. Claude Code reads `AGENTS.md` itself only in a repository with no
  `CLAUDE.md`, so a repository's own `CLAUDE.md` gets the line with `agents`
  alone too.
- With `claude` alone, the block goes in `CLAUDE.md`, and `AGENTS.md` isn't
  touched.
- The block says where the skills are: `.agents/skills` when the `agents`
  target is installed, otherwise `.claude/skills`.

It writes these, to be committed with the repository:

| Written                                    | What for                                                                       |
| ------------------------------------------ | ------------------------------------------------------------------------------ |
| `AGENTS.md` or `CLAUDE.md`                 | Forge's block, between `<!-- forge:start -->` and `<!-- forge:end -->` markers |
| `.agents/skills/<name>/`                   | A copy of each skill, for the `agents` target                                  |
| `.claude/skills/<name>/`                   | A copy of each skill, for the `claude` target                                  |
| `CLAUDE.md`                                | An `@AGENTS.md` line, with both targets, or with `agents` when it exists       |
| `.agents/forge.lock`, or `.claude/forge.lock` with `claude` alone | Forge's version, the targets, the stacks, and each skill it installed |

An existing `AGENTS.md` or `CLAUDE.md` keeps everything it had: forge's block
is added at the end, or replaced where it already is. If the markers are
unbalanced, forge stops and changes nothing. It changes only what's forge's:
its block, the `@AGENTS.md` line when it added it, and the skill folders it
installed, which it makes match the binary exactly. It never touches another
skill, such as `docket`, or another block, such as Docket's. If a skill folder
with one of forge's names is already in a target and forge didn't install it,
forge stops rather than overwrite it. Copies an earlier forge left in
`.claude/skills/`, from before the `claude` target, are removed, while a
project's own skill of the same name stays.

Don't hand-edit the copies or the block, since the next update overwrites
them. Anything particular to a project goes in its own instructions, outside
forge's block, and wins over forge's.

### Supported agents

Last checked against each agent's docs on 2026-11-25. These change often.
`forge agents` prints the same list from the code.

| Agent                                         | `--agent`  | Reads                                                                 | Target   | What you do                                                                                  | Can't do yet                                                                                   |
| --------------------------------------------- | ---------- | --------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| Claude Code                                   | `claude`   | `CLAUDE.md`, `.claude/skills/`                                        | `claude` | Nothing                                                                                      | Reading `.agents/skills/`, so without the `claude` target it reads skills only through the block |
| Codex                                         | `codex`    | `AGENTS.md`, `.agents/skills/`                                        | `agents` | Nothing                                                                                      |                                                                                                |
| Cursor                                        | `cursor`   | `AGENTS.md`, `.agents/skills/`                                        | `agents` | Nothing                                                                                      |                                                                                                |
| Gemini CLI                                    | `gemini`   | `GEMINI.md` by default, `.agents/skills/`                             | `agents` | Add `"context": {"fileName": ["AGENTS.md", "GEMINI.md"]}` to `.gemini/settings.json`         | Reading `AGENTS.md` without that setting                                                       |
| GitHub Copilot (CLI, cloud agent, VS Code)    | `copilot`  | `AGENTS.md`; `.agents/skills/` in the CLI and cloud agent             | `agents` | Nothing                                                                                      | In VS Code, reading skills anywhere but `.github/skills/`, so there it reads them through the block |
| Windsurf                                      | `windsurf` | `AGENTS.md`, `.agents/skills/`                                        | `agents` | Nothing                                                                                      |                                                                                                |
| Any agent that reads `AGENTS.md`              | `agents`   | `AGENTS.md`, and `.agents/skills/` if it loads skills                 | `agents` | Nothing                                                                                      | Loading skills, for one that doesn't: it reads them through the block                         |
| Aider                                         | use `agents` | `AGENTS.md` only when it's passed in                                | `agents` | Run it with `--read AGENTS.md`, or add that to `.aider.conf.yml`                             | Loading skills: it reads them through the block                                                |
| Zed                                           | use `agents` | The first rules file it finds, `AGENTS.md` among them               | `agents` | Make sure no other rules file, such as a stray `.rules`, shadows `AGENTS.md`                 | Loading skills: it reads them through the block                                                |

## Commands

Every command takes the repository's folder, the current one by default, and
`--json`. Each file is reported as `added`, `updated`, `removed`, or `ok`.

| Command                                                      | Does                                                                                                                                          |
| ------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `forge install [dir] [--agent NAME]... [--stack NAME]... [--no-stacks]` | The first install, `agents` by default. On an existing install, adds the agents' targets and the stacks, and rewrites everything installed. |
| `forge update [dir] [--stack NAME]... [--no-stacks]`        | Rewrites every installed target with this binary's forge, keeping the lock's stacks unless told otherwise.                                     |
| `forge remove [dir] [--agent NAME]... [--stack NAME]...`    | Removes those agents' targets and those stacks, moving the block when its file's target goes. With neither, removes everything forge installed. |
| `forge status [dir]`                                        | The targets installed and the agents they serve, whether they match this binary, the stacks, the lock's version, and agent notes. Exits 0.   |
| `forge check [dir]`                                         | Exits 1 when the install doesn't match this binary, for CI.                                                                                    |
| `forge agents`                                              | The supported agents, their targets, and where their files go.                                                                                 |
| `forge version [--check]`                                   | This binary's version. `--check` compares it with the latest: exit 1 when it's older, 2 when it can't tell.                                   |

Any other failure, such as unbalanced markers or a skill folder forge didn't
install, exits 2 and changes nothing.

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
Docket's kit. Every other skill works either way.

## Stacks

A stack is the conventions for one kind of project, installed as one more
skill, in every target, with `--stack`. `go-http` covers Go HTTP API projects: preferred
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

When forge has a new release, update the binary the way you installed it,
then each repository:

```bash
go install github.com/exalynt/forge/cmd/forge@latest   # or download the latest build again, as above
forge update
```

`forge version --check` says whether the binary is the latest: the newest
release, or until there's one, `main`'s latest commit. It
asks Go, so it needs Go, `GOPRIVATE` set and access to the repository, and
exits 2 when it can't tell.

`forge update` keeps the stacks the lock lists. With any `--stack`, it
installs exactly those, so passing a different set is how you swap one, and
`--no-stacks` removes them all. `forge install --stack` adds a stack, and
`forge remove --stack` removes one. Any skill forge no longer has is removed
too. The lock records the version of forge that last changed what it lists.

## Releasing

Every commit on `main` that passes CI is built and published as the `latest`
pre-release by the **Latest** workflow, so there's nothing to do to get a
build of `main`.

A versioned release is a version tag, such as `v0.2.0`, and the builds for
it. In the repository's Actions tab, run the **Release** workflow with the version: it
tags `main`'s latest commit, builds forge for each platform, and publishes
them as a GitHub release with notes from the commits since the last one.
Pushing a tag yourself does the same:

```bash
git tag v0.2.0 && git push origin v0.2.0
```

The tag is the version `forge version` prints, what `go install …@latest`
installs, and what `forge version --check` compares with. Releases follow
semantic versioning: before 1.0, a breaking change to how forge installs, such
as a lock it can't read, bumps the minor version. CI checks every push and
pull request: formatting, vet, the scripts, and an install into a new
repository that's checked, updated and removed.

## Checking in CI

`forge check` changes nothing, and exits 1 when the repository's copy doesn't
match the binary:

```yaml
- run: go install github.com/exalynt/forge/cmd/forge@<version>
  env:
    GOPRIVATE: github.com/exalynt
- run: forge check
```

Pin the version the repository was last updated with, or `latest` to be told
whenever forge moves on.

## Browser tools for qa, ux, and use

`qa` and `ux` drive a capability through its UI with whichever browser tool
the agent has. Without one, they report the UI as not verified. `use` needs a
browser for each person in a test.
[Playwright MCP](https://github.com/microsoft/playwright-mcp) gives them that.
Set it up once per machine:

```bash
forge browser setup                  # headed Chrome
forge browser setup -- --headless    # where there's no display
```

It installs `@playwright/mcp` globally with npm, installs Google Chrome if it's
missing (the server's default browser; this step may ask for `sudo`), and
records where Node and the package are in `setup.json`, in
`~/.local/share/forge/playwright/` (`$XDG_DATA_HOME/forge/playwright/` where
that's set). Then it prints the command to register, `forge browser server`
by forge's absolute path, under five names: `playwright`, the one used by
default, and `playwright-1` through `playwright-4` for when a skill calls for
more, such as several sign-ins at once or the people in a `use` test. Each is
its own browser and sign-in. Register all five with your agent, with the same
command, once:

- **Claude Code**: setup prints a loop that runs
  `claude mcp add --scope user <name> -- /path/to/forge browser server` for
  each name. Replacing one needs `claude mcp remove --scope user <name>` first.
- **Codex**: a table for each name in `~/.codex/config.toml`:

  ```toml
  [mcp_servers.playwright]
  command = "/path/to/forge"
  args = ["browser", "server"]
  ```

- **Cursor**: an entry for each name under `mcpServers` in `.cursor/mcp.json`
  for one project, or `~/.cursor/mcp.json` for all of them, with the same
  `command` and `args`.

`forge browser server` reads `setup.json` and runs the package with Node in
its place. So after switching Node versions, or to change the server's flags,
run `forge browser setup` again: the servers use the new setup the next time
they start, without registering them again. Re-registering is needed only if
forge's binary moves. `forge browser check` exits 1 when something setup
recorded is gone, such as a Node version that was uninstalled, and says what
to run.

Let the browser tools run without asking each time, but keep asking before
`browser_run_code_unsafe`, which runs arbitrary code on your machine, and
before `browser_file_upload` and `browser_drop`, which can send local files to
a page. Set that in your agent's permissions.

- Flags after `--` are passed to the server. The browser is headed by
  default, which works under WSLg, so add `--headless` where there's no
  display. Pass `--browser firefox` to use Firefox instead of Chrome.
  Playwright's bundled Chromium isn't one of the server's options.
- Browser windows arrange themselves across your monitors, so you can follow
  along without them covering your work. Playwright opens Chrome through
  `chrome.sh`, which setup writes beside `setup.json` and each server rewrites
  as it starts, so it's always the one forge carries. It reads the monitors
  each time a browser opens, with `xrandr` on Linux and WSLg or AppKit on
  macOS, and gives the window an eighth of a monitor, or more where that would be
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
- Setup records Node and the package by absolute path, so the server starts
  even when the agent's environment lacks nvm's PATH. Re-running it also
  updates the package.
- On Windows, Chrome's windows don't arrange themselves, since `chrome.sh`
  needs bash: Chrome opens as the server would open it anyway.

## Uninstalling

```bash
forge remove
```

It removes everything forge installed: its block, its skills in every target,
the `@AGENTS.md` line if forge added it, its lock, and the folders that leaves
empty. A file left with nothing in it, such as a `CLAUDE.md` forge created, is
removed too. `forge remove --agent NAME` removes one target instead.

## Known gaps

- **Agents without their own target read skills through the block.** Claude
  Code doesn't read `.agents/skills/`, so where the `claude` target isn't
  installed, it reads a skill only because forge's block says where it is, and
  `/implement`, `/qa` and the rest aren't in its slash menu: ask for one by
  name, such as "use implement". Install `--agent claude` to give it its own
  copies. The same holds for Aider, Zed, and others that don't load skills.
- **Gemini CLI reads `AGENTS.md` only with its `context.fileName` setting.**
  Forge doesn't write that setting: `forge status` and `forge agents` say what
  to add.
- **Copilot in VS Code reads skills only from `.github/skills/`.** There's no
  target for it yet, so in VS Code it reads skills through the block.
