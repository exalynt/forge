<h1 align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/logo/svg/forge-lockup-dark.svg">
    <img src="assets/logo/svg/forge-lockup-light.svg" alt="forge" height="48">
  </picture>
</h1>

<p align="center">Exalynt's way of working, for every coding agent</p>

Forge gives coding agents Exalynt's way of working: how we engineer, and
skills for the workflows that apply it. You install it into a repository and
commit it, so everyone working there gets the same conventions and skills,
whatever agent they use: Claude Code, Codex, Cursor, Gemini CLI, Copilot,
Windsurf, or anything else that reads `AGENTS.md`.

## What you get

**Workflow skills.** Ask your agent for one by name, such as `/implement` in
Claude Code or "use implement" elsewhere.

| Skill              | What it does                                                                      |
| ------------------ | --------------------------------------------------------------------------------- |
| `implement`        | Build a feature, improvement, or fix; in its own worktree with a PR when asked    |
| `commit`           | Commit changes as small, related commits                                          |
| `pr`               | Commit, push, and open or update a PR labeled with its stability level            |
| `qa`               | QA a capability end to end against its business rules                             |
| `ux`               | Try a main flow as a type of user and report where it's hard to use               |
| `use`              | Usability test a flow, each person played separately in their own browser         |
| `stability-review` | Review work against a capability's stability level (needs Docqet)                 |
| `promote`          | Move a capability up a stability level (needs Docqet)                             |

**Conventions** for how code is built, in layers: `standards`, `testing` and
`observability` in every install, then the languages (`go`, `typescript`) and
stacks (`go/http`, `typescript/react`) you choose. See
[docs/conventions.md](docs/conventions.md).

## Install the command

Download the `latest` build, which is rebuilt from every commit on `main` that
passes CI, for Linux, macOS and Windows:

```bash
curl -fsSLO https://github.com/exalynt/forge/releases/download/latest/forge_latest_linux_amd64.tar.gz   # or darwin_arm64, …
tar -xzf forge_latest_linux_amd64.tar.gz
install forge_latest_linux_amd64/forge ~/.local/bin/   # anywhere on your PATH
forge version
```

Or with Go:

```bash
go install github.com/exalynt/forge/cmd/forge@latest
```

## Add it to a repository

```bash
cd path/to/repo
forge install                                # for agents that read AGENTS.md
forge install --agent claude --agent codex   # for the agents you use
forge install --stack go/http                # with a stack, which brings its language
forge install --language go                  # with a language and no stack
```

Commit what it writes: a block in `AGENTS.md` or `CLAUDE.md`, a copy of each
skill in `.agents/skills/` or `.claude/skills/`, and a lock file. Don't edit
those by hand, since the next update overwrites them. Anything particular to
your project goes in its own instructions, outside forge's block, and wins
over forge's.

Most agents need nothing more. Gemini CLI and Aider need a setting, and a few
agents can't load skills themselves. See
[docs/agents.md](docs/agents.md) for each agent and exactly what forge writes.

## Forge and Docqet

Forge says *how* to do the work. What the software is meant to do, and how
stable each part needs to be, belongs to the project. Where a repository
records that with [Docqet](https://github.com/docqet/docqet) (it has a
`docqet.yaml`), forge's skills read each capability's stability level, rules
and gaps from it, and record what you confirm the Docqet way. Install Docqet's
kit alongside forge's; neither touches the other.

Without Docqet, the skills take the level and rules from you and record
nothing beyond the code and tests you approve. `stability-review` and
`promote` need Docqet and say so without it.

## Browser tools

`qa`, `ux` and `use` drive your app in a browser through
[Playwright MCP](https://github.com/microsoft/playwright-mcp). Set it up once
per machine, then register the servers it prints with your agent:

```bash
forge browser setup                  # headed Chrome
forge browser setup -- --headless    # where there's no display
```

See [docs/browser.md](docs/browser.md) for registering with Claude Code, Codex
and Cursor, permissions, and how the browsers behave.

## Keep it current

Update the binary the way you installed it, then each repository:

```bash
forge update
```

In CI, `forge check` changes nothing and fails when the repository's copy
doesn't match the binary:

```yaml
- run: go install github.com/exalynt/forge/cmd/forge@<version>
- run: forge check
```

Pin the version the repository was last updated with, or use `latest` to be
told whenever forge moves on.

## Commands

| Command         | Does                                                          |
| --------------- | ------------------------------------------------------------- |
| `forge install` | Install forge, or add agents, languages and stacks            |
| `forge update`  | Rewrite the install with this binary's forge                  |
| `forge remove`  | Remove agents, languages or stacks, or everything forge added |
| `forge status`  | Show what's installed and whether it's current                |
| `forge check`   | Fail when the install doesn't match this binary, for CI       |
| `forge agents`  | List the supported agents and where their files go            |
| `forge version` | Print the version; `--check` compares it with the latest      |
| `forge browser` | Set up and run the browser tools                              |

Flags, exit codes and details are in [docs/commands.md](docs/commands.md).

## Uninstall

```bash
forge remove
```

It removes everything forge installed and leaves the rest of your files as
they were.

## More

- [Agents and what forge writes](docs/agents.md)
- [Commands](docs/commands.md)
- [Standards, languages and stacks](docs/conventions.md)
- [Browser tools](docs/browser.md)
- [Developing and releasing forge](docs/development.md)
- [Logo and how to use it](assets/logo/README.md)

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md), and the
Exalynt [Code of Conduct](https://github.com/exalynt/community/blob/main/CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE)
