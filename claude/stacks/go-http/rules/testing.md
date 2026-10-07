---
paths:
  - "**/*.go"
---

# Go HTTP projects: tests

General conventions for Go HTTP API projects. Where this project's own code or
CLAUDE.md differs, follow the project.

## When to write tests

Write tests only when the user explicitly asks, or when the feature's
stability level requires them (see the `stability` skill). Then add only the
kinds and depth that the request or level calls for. For example, Alpha asks
for main-path tests, not exhaustive error paths. Otherwise, don't add tests on
your own initiative. Existing tests must still pass after any change.

## Three kinds, separated by build tags

Build tags let each suite run on its own.

| Kind | Build tag | Run with | Covers |
| --- | --- | --- | --- |
| Unit | none | `go test ./...` | Table-driven tests (`tests := []struct{…}` with `t.Run`), dependencies mocked with moq: service logic, validation, error mapping. |
| Integration | `//go:build integration` | `go test -tags integration ./...` | A specific piece of happy-path functionality against real dependencies such as Postgres, for example a repository against a real database. |
| Feature / end-to-end | `//go:build e2e` | `go test -tags e2e ./...` | Black box: calls the API endpoints over HTTP against the running app and its real dependencies, asserting only on responses. Happy paths. |

## Mocks

- Generate mocks with moq from the interface, through a `//go:generate moq …`
  directive, and commit them. Don't hand-write mocks.
- Mock the boundary interfaces (repositories, providers), not the code under
  test.

## Real dependencies

Integration and end-to-end tests get their dependencies from the project's
local stack (such as Docker Compose) through environment variables. Each test
isolates its own data so the suite can run repeatedly.
