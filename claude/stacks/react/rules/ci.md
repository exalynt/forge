---
paths:
  - ".github/**"
  - "package.json"
---

# React projects: CI

General conventions for React projects. Where this project's CLAUDE.md
differs, follow it.

## When

CI lives in `.github/workflows/ci.yml`. Add it once any feature in the project
reaches Alpha (see the `stability` skill).

## Scripts

CI runs the project's npm scripts, so the linter and formatter can change
without touching the workflow. Every project defines:

| Script | Does |
| --- | --- |
| `lint` | Runs the linter. |
| `format` | Formats the code in place. |
| `format:check` | Fails if anything isn't formatted. |
| `build` | Type checks (`tsc -b`) and builds the app. |

## Workflow

- Run on pushes to `main` and on pull requests to `main`.
- Set `permissions: contents: read`.
- Set up Node with `actions/setup-node` and `cache: npm`, using the same Node
  version in every workflow, and install with `npm ci`.
- CI only checks. It never commits, pushes, opens pull requests with fixes, or
  deploys. Deployment, if any, is its own workflow.

## Jobs

Separate jobs, so each failure shows on its own:

| Job | Steps |
| --- | --- |
| `lint` | `npm run lint` |
| `format` | `npm run format:check` |
| `build` | `npm run build` |
