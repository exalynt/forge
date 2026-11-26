# Go: preferred packages

For Go projects, adding to the `standards` skill's
[packages.md](../../standards/references/packages.md). Where this project's
own code or instructions differ, follow the project.

- `github.com/urfave/cli/v3`: the binary's commands (see
  [commands.md](commands.md))
- `github.com/gofrs/uuid/v5`: IDs (UUIDv7)
- `github.com/pressly/goose/v3`: SQL migrations, embedded in the binary
- `github.com/matryer/moq`: generated mocks for tests

No ORMs.
