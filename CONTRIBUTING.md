# Contributing to forge

Contributions are welcome: bug reports, ideas, documentation, and pull
requests of any size.

forge follows the Exalynt
[contributing guide](https://github.com/exalynt/community/blob/main/CONTRIBUTING.md)
and [Code of Conduct](https://github.com/exalynt/community/blob/main/CODE_OF_CONDUCT.md).
Read those first. They cover issues, pull requests, code review, AI-assisted
contributions, security, and licensing. This document adds only what's
specific to forge.

## How to contribute

1. For a bug, open an issue with what you ran, what happened, and what you
   expected. Include `forge version`, `forge status`, and the agent you use.
2. For anything bigger than a small fix, such as a new skill, a new language
   or stack, or support for another agent, open an issue or discussion first
   so we can agree on the direction before you build it.
3. Fork the repository, make your change on a branch, and run the checks
   below.
4. Open a pull request against `main` that says what changed, why, and how
   you tested it.

## Development setup

You need:

- Go at the version in `go.mod` or newer.
- [ShellCheck](https://www.shellcheck.net/), for the scripts.

There's nothing else to install. To try your build in a repository:

```bash
go build -o /tmp/forge ./cmd/forge
/tmp/forge install path/to/some/repo --agent claude
```

[docs/development.md](docs/development.md) shows how the repository is laid
out.

## Checks

Before opening a pull request, run what CI runs:

```bash
test -z "$(gofmt -l .)"
go vet ./...
go build ./...
go test ./...
shellcheck internal/playwright/chrome.sh scripts/release/*.sh
```

CI also installs forge into a new repository for every target and stack, then
checks, updates, and removes it, and expects nothing to be left behind.

## Changing the kit

`kit/` is what forge installs into other people's repositories, so a change
there reaches every project on its next `forge update`.

- **Skills are instructions for agents.** Write them as plainly as you would
  for a new engineer: say what to do and when, not just what's true.
- **Conventions are layered.** A language or stack only adds to or narrows
  the layer above it, never restates it, and a topic keeps the same file name
  in every layer. A new topic file needs a link from each language and stack
  index that covers it. Forge's tests check that every link resolves and that
  every index is complete. See [docs/conventions.md](docs/conventions.md).
- **Keep it general.** Conventions in forge apply to any project that uses
  that language or stack. Rules for one project's architecture or domain
  belong in that project's own instructions.
- **Respect the project's intention.** Skills never decide what the software
  should do, or how stable it should be. With Docqet they read it from
  Docqet's docs, and without it they ask the user. `kit/AGENTS.md`, under
  "Where intention lives", defines both modes, and every skill follows it.

## Changing the command

- Each command lives in `internal/commands/<name>.go`, and its subcommands in
  `internal/commands/<name>/`.
- Supported agents are rows in `Targets` in `internal/install/targets.go`.
  When you add or change one, update the table in
  [docs/agents.md](docs/agents.md) and its "last checked" date in the same
  pull request.
- Forge changes only the files it owns in a repository. Keep it that way: a
  change that could touch a user's own skills or instructions needs a test
  showing it doesn't.
- A change to a command's flags or output updates
  [docs/commands.md](docs/commands.md) and, where it matters to users, the
  README.

## License

By contributing, you agree that your contributions are licensed under the
[MIT License](LICENSE).
