---
name: go-http
description: Conventions for Go HTTP API projects - preferred packages, layers and boundaries, repositories and migrations, service errors, REST API design, pagination, health checks, commands, configuration, tests, and CI. Use when writing, changing, or reviewing Go code, SQL migrations, or CI in a Go HTTP API project or Go library.
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

# Go HTTP projects

General conventions for Go HTTP API projects. Where this project's own code or
instructions differ, follow the project. Read the reference for what you're
about to change, from this skill's `references/` directory:

| Reference                                         | Read when                                                                     |
| ------------------------------------------------- | ----------------------------------------------------------------------------- |
| [packages.md](references/packages.md)             | Choosing a package or adding a dependency                                     |
| [layers.md](references/layers.md)                 | Writing any Go code: handlers, services, boundary interfaces, domain models   |
| [repositories.md](references/repositories.md)     | Writing repositories, SQL, or migrations, or changing the schema              |
| [errors.md](references/errors.md)                 | Returning errors from services or mapping them in handlers                    |
| [rest-api.md](references/rest-api.md)             | Designing or changing endpoints: paths, methods, bodies, statuses             |
| [pagination.md](references/pagination.md)         | Adding or changing any endpoint that returns a list                           |
| [health.md](references/health.md)                 | Adding health endpoints, probes, or a provider's `Ping`                       |
| [commands.md](references/commands.md)             | Adding or changing the binary's commands, subcommands, or flags               |
| [config.md](references/config.md)                 | Adding or reading configuration, or wiring commands                           |
| [testing.md](references/testing.md)               | Writing or running tests, including acceptance tests and libraries            |
| [ci.md](references/ci.md)                         | Adding or changing the CI workflow                                            |

These say how to build. How much a capability needs, such as which tests,
health endpoints, or CI jobs, is what its stability level expects (Docket's rubric,
when the repo has it; see forge's `AGENTS.md`, "Where intention lives"). The
`testing` and `observability` skills hold the practice that applies to every
stack.
