---
name: standards
description: Conventions for every project, in any language or stack - layers and boundaries, errors, configuration, dependencies, repositories and schema changes, REST APIs, pagination, health checks, CI, where code lives, and for web UIs mobile responsiveness, light and dark mode, and accessibility. Use when writing, changing, or reviewing code, APIs, schemas, CI, or UI. Language and stack skills, such as go or go-http, build on it.
paths:
  - "**/*.{go,ts,tsx,js,jsx,mjs,cjs,py,rb,java,kt,rs,php,cs,swift,sql}"
  - "**/openapi*.y*ml"
  - ".github/**"
---

# Standards

The conventions every project follows, whatever its language or stack. They
are the first of three layers:

1. **Standards**, this skill, with the `testing` and `observability` skills.
2. **A language**, such as `go` or `typescript`, where the repository has it.
3. **A stack under that language**, such as `go-http` or `typescript-react`.

Each layer only adds to or narrows the one above it. A topic keeps its file
name in every layer, so `errors.md` here, in `go`, and in `go-http` are one
topic, read in that order. Where two layers differ, the more specific one
wins, and this project's own code and instructions win over all of them.
Forge's security and privacy rules apply whatever any layer says.

When a language or stack skill is installed for the code you're changing,
start from its index instead: it lists this skill's files for each topic it
covers. Read the reference for what you're about to change, from this skill's
`references/` directory:

| Reference                                          | Read when                                                                      |
| -------------------------------------------------- | ------------------------------------------------------------------------------ |
| [layers.md](references/layers.md)                  | Writing any code: entry points, services, boundary interfaces, domain models   |
| [errors.md](references/errors.md)                  | Returning errors from services or mapping them at an entry point               |
| [config.md](references/config.md)                  | Adding or reading configuration                                                |
| [packages.md](references/packages.md)              | Choosing a package or adding a dependency                                      |
| [repositories.md](references/repositories.md)      | Writing repositories, SQL, or migrations, or changing the schema               |
| [rest-api.md](references/rest-api.md)              | Designing or changing endpoints: paths, methods, bodies, statuses              |
| [pagination.md](references/pagination.md)          | Adding or changing any endpoint that returns a list                            |
| [health.md](references/health.md)                  | Adding health endpoints, probes, or a provider's ping                          |
| [ci.md](references/ci.md)                          | Adding or changing the CI workflow                                             |
| [structure.md](references/structure.md)            | Adding files or folders, or deciding where code lives                          |
| [responsive.md](references/responsive.md)          | Building or changing any page or component of a web UI                         |
| [color-modes.md](references/color-modes.md)        | Styling anything in a web UI, or setting up its theme                          |
| [accessibility.md](references/accessibility.md)    | Building or changing any page or component of a web UI                         |
| [testing](../testing/SKILL.md)                     | Writing, changing, or placing tests                                            |
| [observability](../observability/SKILL.md)         | Adding traces, logs, or metrics                                                |

These say how to build. How much a capability needs, such as which tests,
health endpoints, or CI jobs, is what its stability level expects (Docqet's
rubric, when the repo has it; see forge's block, "Where intention lives").
