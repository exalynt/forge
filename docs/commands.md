# Commands

Every command that works on a repository takes its folder, the current one by
default, and `--json`. Each file is reported as `added`, `updated`, `removed`,
or `ok`.

| Command                                                                                              | Does                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `forge install [dir] [--agent NAME]... [--language NAME]... [--stack LANG/NAME]... [--no-languages]` | The first install, `agents` by default. On an existing install, adds the agents' targets, the languages and the stacks, and rewrites everything installed.                       |
| `forge update [dir] [--language NAME]... [--stack LANG/NAME]... [--no-languages]`                    | Rewrites every installed target with this binary's forge, keeping the lock's languages and stacks unless told otherwise.                                                         |
| `forge remove [dir] [--agent NAME]... [--language NAME]... [--stack LANG/NAME]...`                   | Removes those agents' targets, those languages with their stacks, and those stacks, moving the block when its file's target goes. With none, removes everything forge installed. |
| `forge status [dir]`                                                                                 | The targets installed and the agents they serve, whether they match this binary, the languages and stacks, the lock's version, and agent notes. Exits 0.                         |
| `forge check [dir]`                                                                                  | Exits 1 when the install doesn't match this binary, for CI.                                                                                                                       |
| `forge agents`                                                                                       | The supported agents, their targets, and where their files go.                                                                                                                    |
| `forge version [--check]`                                                                            | This binary's version. `--check` compares it with the latest: exit 1 when it's older, 2 when it can't tell.                                                                      |
| `forge browser setup [-- FLAGS]`                                                                     | Sets up Playwright MCP for `qa`, `ux` and `use`. See [browser.md](browser.md).                                                                                                     |
| `forge browser server`                                                                               | Runs the Playwright MCP server that setup recorded. Your agent runs this, not you.                                                                                                |
| `forge browser check`                                                                                | Exits 1 when something setup recorded is gone, and says what to run.                                                                                                              |

Any other failure, such as unbalanced markers or a skill folder forge didn't
install, exits 2 and changes nothing.

## Languages and stacks on update

`forge update` keeps the languages and stacks the lock lists. With any
`--language` or `--stack`, it installs exactly those, so passing a different
set is how you swap one, and `--no-languages` removes them all.
`forge install --language` or `--stack` adds one, and `forge remove` with
either removes it; removing a language removes its stacks. Any skill forge no
longer has is removed too. The lock records the version of forge that last
changed what it lists.

An install from before languages, with the `go-http` or `react` stack, updates
to `go/http` or `typescript/react` and their languages: `forge update` adds
`standards`, `go`, `typescript`, and `typescript-react`, and removes `react`.
The old names still work with `--stack`.

## Version checks

`forge version --check` says whether the binary is the latest: the newest
release, or until there's one, `main`'s latest commit. It asks Go, so it needs
Go and a network connection, and exits 2 when it can't tell.

## Removing

`forge remove` with no flags removes everything forge installed: its block,
its skills in every target, the `@AGENTS.md` line if forge added it, its lock,
and the folders that leaves empty. A file left with nothing in it, such as a
`CLAUDE.md` forge created, is removed too.
