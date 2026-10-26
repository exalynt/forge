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
only for end-to-end tests of each endpoint's happy path and acceptance tests
of core business rules. Unit and integration tests, and acceptance tests of edge-case
rules, start at Beta. Otherwise, don't add tests on your own initiative.
Existing tests must still pass after any change. The CI rule covers which CI
job runs each kind.

## Three levels

Unit and integration tests are Go, separated by build tags so each suite runs
on its own. End-to-end tests use Playwright, in a top-level `tests/e2e/`
directory.

| Level | Tool and location | Run with | Covers |
| --- | --- | --- | --- |
| Unit | Go, no build tag, next to the code | `go test ./...` | Table-driven tests (`tests := []struct{…}` with `t.Run`), dependencies mocked with moq: service logic, validation, error mapping. |
| Integration | Go, `//go:build integration`, next to the code | `go test -tags integration ./...` | A specific piece of happy-path functionality against real dependencies such as Postgres, for example a repository against a real database. |
| End-to-end | Playwright, `tests/e2e/` | `npx playwright test` | Black box: one endpoint's happy path per test, called over HTTP against the running app and its real dependencies, asserting only on its response. Setup calls are fine, but a series of calls that verifies a flow is an acceptance test in `tests/acceptance/`. |

## Acceptance tests

Acceptance tests are Gherkin run with
[godog](https://github.com/cucumber/godog), Cucumber's Go implementation, at
every level (see the testing rule):

- Each package's acceptance tests run from one `TestAcceptance` function, in
  a test file with the `acceptance` build tag, through `godog.TestSuite` with
  `TestingT: t`. The tag keeps them out of the unit and integration suites,
  and `go test -tags acceptance -run '^TestAcceptance$' ./...` runs them all
  without the unit tests.
- Rules proved at the service or repository go in a `features/` directory
  next to that package, usually the service. Step definitions call it, with
  moq mocks for its boundaries, or real dependencies where the rule needs
  them.
- Rules proved only across a flow of calls go in `tests/acceptance/`, a Go
  package with its own `features/` directory, whose step definitions call the
  running app's API over HTTP.

## Mocks

- Generate mocks with moq from the interface, through a `//go:generate moq …`
  directive, and commit them. Don't hand-write mocks.
- Mock the boundary interfaces (repositories, providers), not the code under
  test.

## Real dependencies

Integration, end-to-end, and acceptance tests get their dependencies from the
project's local stack (such as Docker Compose) through environment variables.
Each test isolates its own data so the suite can run repeatedly.
