# TypeScript: CI

For TypeScript projects, adding to the `standards` skill's
[ci.md](../../standards/references/ci.md). Where this project's own code or
instructions differ, follow the project.

## Scripts

CI runs the project's npm scripts, so the linter and formatter can change
without touching the workflow. Every project defines:

| Script | Does |
| --- | --- |
| `lint` | Runs the linter. |
| `format` | Formats the code in place. |
| `format:check` | Fails if anything isn't formatted. |
| `build` | Type checks (`tsc -b`) and builds the app. |

## Setup

Set up Node with `actions/setup-node` and `cache: npm`, using the same Node
version in every workflow, and install with `npm ci`.

## Jobs

| Job | Steps |
| --- | --- |
| `lint` | `npm run lint` |
| `format` | `npm run format:check` |
| `build` | `npm run build` |
