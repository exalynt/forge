---
paths:
  - "**/*.go"
  - "tests/**"
---

# Go HTTP projects: tests

General conventions for Go HTTP API projects. Where this project's own code or
CLAUDE.md differs, follow the project.

## When to write tests

Write tests only when the user explicitly asks, or when the feature's
stability level requires them (see the `stability` skill). Then add only the
kinds and depth that the request or level calls for. For example, Alpha asks
only for feature / end-to-end tests of common paths. Unit and integration
tests start at Beta. Otherwise, don't add tests on
your own initiative. Existing tests must still pass after any change.

## Three kinds

Unit and integration tests are Go, separated by build tags so each suite runs
on its own. Feature and end-to-end tests use Playwright, in a top-level
`tests/` directory with `tests/feature/` and `tests/e2e/` subdirectories as
needed.

| Kind | Tool and location | Run with | Covers |
| --- | --- | --- | --- |
| Unit | Go, no build tag, next to the code | `go test ./...` | Table-driven tests (`tests := []struct{…}` with `t.Run`), dependencies mocked with moq: service logic, validation, error mapping. |
| Integration | Go, `//go:build integration`, next to the code | `go test -tags integration ./...` | A specific piece of happy-path functionality against real dependencies such as Postgres, for example a repository against a real database. |
| Feature / end-to-end | Playwright, `tests/feature/` or `tests/e2e/` | `npx playwright test` | Black box: calls the API endpoints over HTTP against the running app and its real dependencies, asserting only on responses. Happy paths. |

## Mocks

- Generate mocks with moq from the interface, through a `//go:generate moq …`
  directive, and commit them. Don't hand-write mocks.
- Mock the boundary interfaces (repositories, providers), not the code under
  test.

## Real dependencies

Integration and end-to-end tests get their dependencies from the project's
local stack (such as Docker Compose) through environment variables. Each test
isolates its own data so the suite can run repeatedly.
