---
name: typescript
description: Conventions for every TypeScript project, building on the standards skill - services for external APIs and SDKs, one module per resource with its types, and the npm scripts and jobs CI runs. Use when writing, changing, or reviewing TypeScript code, package.json, or the CI of a TypeScript project.
paths:
  - "**/*.ts"
  - "**/*.tsx"
  - "**/package.json"
  - "**/tsconfig*.json"
  - ".github/**"
---

# TypeScript

Conventions for every TypeScript project: a language layer, building on the
`standards` skill. A stack skill under it, such as `typescript-react`, builds
on this in turn; where one is installed for the code you're changing, start
from its index instead.

Each topic is one file name in every layer. Read the rows for what you're
about to change, left to right: the general rule first, then what TypeScript
adds. Where they differ, the more specific one wins, and this project's own
code and instructions win over both.

| Topic    | Read when                                                      | Read, in order                                                                       |
| -------- | -------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| services | Calling an external API or SDK                                 | [standards](../standards/references/layers.md), [typescript](references/services.md)  |
| ci       | Adding or changing the CI workflow or the npm scripts it runs  | [standards](../standards/references/ci.md), [typescript](references/ci.md)            |
| testing  | Writing or running tests                                       | [testing](../testing/SKILL.md)                                                       |

Topics TypeScript adds nothing to are in the `standards` skill's index.
