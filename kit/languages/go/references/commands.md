# Go: commands

For every Go binary. Where this project's own code or instructions differ,
follow the project.

## urfave/cli

- Build every binary's command line with `github.com/urfave/cli/v3`: its
  commands, subcommands, flags and help. Don't use the standard library's
  `flag` package or another CLI framework, even for a binary with one command.
- Each command has a `Usage` line, and a `Description` where its behavior
  needs more than a line: what it changes, its defaults, and its exit codes
  when they mean something.

## Layout

```
cmd/<binary>/main.go     runs the root command, and nothing else
internal/commands/
  <name>.go              one top-level command: <Name>Command() *cli.Command
  <name>/                that command's subcommands, if it has any
    <sub>.go             one subcommand: <Sub>Command() *cli.Command
  shared.go              helpers several commands use
```

- The root command lists the top-level commands by their constructors, in
  `main.go` or in `internal/commands/root.go`.
- Each top-level command is its own file, named for the command, with a
  constructor that returns it: `serve.go` has `ServeCommand()`.
- A command with subcommands keeps them in a folder named for it, one file
  each, in a package of the same name. Its `<name>.go` names the group and
  lists them, such as `migrate.go` listing `migrate.UpCommand()`,
  `migrate.DownCommand()` and `migrate.StatusCommand()`.
- A helper one command uses stays in that command's file. Helpers several use
  go in `shared.go`.
- Commands parse flags and arguments, build the config once (see
  [config.md](config.md)), call the packages that do the work, and print the
  result. The work itself lives in its own packages, such as services or
  `internal/<thing>/`, so it can be tested without the command line.
- Don't give a package that does the work the same name as a subcommand
  folder, since one would have to import the other under an alias.
