---
name: go-http
description: Conventions for Go HTTP API projects, building on the go and standards skills - chi, render, validator, turn and health packages, handlers, mapping errors to statuses, REST resources and requests, pagination with turn, health checks, end-to-end tests over HTTP, and CI. Use when writing, changing, or reviewing Go code, SQL migrations, OpenAPI specs, or CI in a Go HTTP API project.
paths:
  - "**/*.go"
  - "**/go.mod"
  - "**/*.sql"
  - "**/repositories/**"
  - "**/repository/**"
  - "**/store/**"
  - "**/migrations/**"
  - "tests/**"
  - "**/Dockerfile*"
  - "**/compose*.y*ml"
  - "**/.kube/**"
  - "**/openapi*.y*ml"
  - "**/static/docs/**"
  - ".github/**"
---

# Go HTTP APIs

Conventions for Go HTTP API projects: a stack layer, building on the `go`
skill, which builds on the `standards` skill.

Each topic is one file name in every layer. Read the rows for what you're
about to change, left to right: the general rule first, then what Go adds,
then what this stack adds. Where they differ, the more specific one wins, and
this project's own code and instructions win over all of them.

| Topic        | Read when                                                              | Read, in order                                                                                                                                  |
| ------------ | ---------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| packages     | Choosing a package or adding a dependency                              | [standards](../standards/references/packages.md), [go](../go/references/packages.md), [go-http](references/packages.md)                          |
| layers       | Writing any Go code: handlers, services, boundary packages, models     | [standards](../standards/references/layers.md), [go](../go/references/layers.md), [go-http](references/layers.md)                                |
| errors       | Returning errors from services or mapping them in handlers             | [standards](../standards/references/errors.md), [go](../go/references/errors.md), [go-http](references/errors.md)                                |
| rest-api     | Designing or changing endpoints: paths, methods, bodies, statuses      | [standards](../standards/references/rest-api.md), [go-http](references/rest-api.md)                                                             |
| pagination   | Adding or changing any endpoint that returns a list                    | [standards](../standards/references/pagination.md), [go-http](references/pagination.md)                                                         |
| health       | Adding health endpoints, probes, or a provider's `Ping`                | [standards](../standards/references/health.md), [go-http](references/health.md)                                                                 |
| config       | Adding or reading configuration, or wiring commands                    | [standards](../standards/references/config.md), [go](../go/references/config.md)                                                                |
| repositories | Writing repositories, SQL, or migrations, or changing the schema       | [standards](../standards/references/repositories.md), [go](../go/references/repositories.md)                                                    |
| commands     | Adding or changing the binary's commands, subcommands, or flags        | [go](../go/references/commands.md)                                                                                                              |
| testing      | Writing or running tests, including acceptance and end-to-end tests    | [testing](../testing/SKILL.md), [go](../go/references/testing.md), [go-http](references/testing.md)                                             |
| ci           | Adding or changing the CI workflow                                     | [standards](../standards/references/ci.md), [go](../go/references/ci.md), [go-http](references/ci.md)                                           |

Topics neither Go nor this stack adds to, such as where code lives, are in
the `standards` skill's index.
