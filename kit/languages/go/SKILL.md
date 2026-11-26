---
name: go
description: Conventions for every Go project, building on the standards skill - preferred packages, layers and boundary packages, sentinel errors, the config package, repositories and goose migrations, urfave/cli commands, tests with build tags, moq and godog, libraries, and CI. Use when writing, changing, or reviewing Go code, SQL migrations, or the CI of a Go project or library.
paths:
  - "**/*.go"
  - "**/go.mod"
  - "**/*.sql"
  - "**/migrations/**"
  - "tests/**"
  - "**/Dockerfile*"
  - ".github/**"
---

# Go

Conventions for every Go project: a language layer, building on the
`standards` skill. A stack skill under it, such as `go-http`, builds on this
in turn; where one is installed for the code you're changing, start from its
index instead.

Each topic is one file name in every layer. Read the rows for what you're
about to change, left to right: the general rule first, then what Go adds.
Where they differ, the more specific one wins, and this project's own code and
instructions win over both.

| Topic        | Read when                                                              | Read, in order                                                                                                    |
| ------------ | ---------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| packages     | Choosing a package or adding a dependency                              | [standards](../standards/references/packages.md), [go](references/packages.md)                                     |
| layers       | Writing any Go code: services, boundary packages, domain models        | [standards](../standards/references/layers.md), [go](references/layers.md)                                         |
| errors       | Returning errors from services or mapping them                         | [standards](../standards/references/errors.md), [go](references/errors.md)                                         |
| config       | Adding or reading configuration, or wiring commands                    | [standards](../standards/references/config.md), [go](references/config.md)                                         |
| repositories | Writing repositories, SQL, or migrations, or changing the schema       | [standards](../standards/references/repositories.md), [go](references/repositories.md)                             |
| commands     | Adding or changing the binary's commands, subcommands, or flags        | [go](references/commands.md)                                                                                      |
| testing      | Writing or running tests, including acceptance tests and libraries     | [testing](../testing/SKILL.md), [go](references/testing.md)                                                       |
| ci           | Adding or changing the CI workflow                                     | [standards](../standards/references/ci.md), [go](references/ci.md)                                                 |

Topics Go adds nothing to, such as where code lives or REST API design, are in
the `standards` skill's index.
