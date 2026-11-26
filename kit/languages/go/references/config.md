# Go: configuration

For Go projects, adding to the `standards` skill's
[config.md](../../standards/references/config.md). Where this project's own
code or instructions differ, follow the project.

- A `config` package holds a single `Config` struct, built by `config.New()`
  from environment variables. Nothing else calls `os.Getenv`.
- Related settings are sub-config structs (`HTTP`, `Database`, `Email`,
  `Payments`), with env var prefixes to match.
- Each field's comment says what it controls and what happens when it's unset.
- Defaults are named constants in the `config` package, each with a comment
  saying why.
- Commands build the config once (see [commands.md](commands.md)) and pass
  each component the sub-config it needs.
