# Standards, languages and stacks

How code is built comes in three layers, each one or more skills:

| Layer     | Installed                        | Skills                                  | Covers                                                                                                                                                                                                                                                         |
| --------- | -------------------------------- | --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Universal | Always                           | `standards`, `testing`, `observability` | What holds in any language: layers and boundaries, errors, configuration, dependencies, repositories and schema changes, REST APIs, pagination, health checks, CI, where code lives, and for web UIs responsiveness, light and dark mode, and accessibility. |
| Language  | `--language go`, or with a stack | `go`, `typescript`                      | What a language adds: Go's packages, sentinel errors, config package, goose migrations, urfave/cli commands, tests and CI; TypeScript's services modules and npm scripts.                                                                                     |
| Stack     | `--stack go/http`                | `go-http`, `typescript-react`           | What a kind of project adds: chi, render, turn and the health library for Go HTTP APIs; MUI, pages and components, data, lists and forms for React.                                                                                                           |

A stack is named `language/stack` and brings its language with it, so
`--stack go/http` installs `go` and `go-http`. A repository with a Go API and
a React UI passes `--stack go/http --stack typescript/react`.

## How the layers fit together

- **Each layer only adds to or narrows the one above it**, never restates it.
  The more specific layer wins, and the project's own instructions win over
  all of them. Forge's security and privacy rules apply whatever a layer says.
- **A topic keeps its file name in every layer**: `errors.md` in `standards`,
  `go`, and `go-http` is one topic, read in that order.
- **The most specific skill is the entry point.** Its `SKILL.md` lists, for
  each topic, the file in every layer, linked as `../go/references/errors.md`,
  which resolves because installed skills are siblings. Each `SKILL.md` lists
  the files it applies to, which Claude Code and Cursor use to load it.
- **The block names what's installed**, such as "This repository has
  `standards`, `go` → `go-http`, and `typescript`", so an agent that doesn't
  load skills by itself still knows what to read.

Forge's tests check that every link in the kit resolves to a skill that's
always installed with the one linking to it, and that each language and stack
index links every layer's file for each of its topics.

Project-specific architecture and domain rules don't belong in forge. Keep
them in each project's own instructions.
