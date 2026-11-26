# Go HTTP: tests

For Go HTTP APIs, adding to the `go` skill's
[testing.md](../../go/references/testing.md). Where this project's own code or
instructions differ, follow the project.

## End-to-end tests

| Level | Tool and location | Run with | Covers |
| --- | --- | --- | --- |
| End-to-end | Playwright, `tests/e2e/` | `npx playwright test` | Black box: one endpoint's happy path per test, called over HTTP against the running app and its real dependencies, asserting only on its response. Setup calls are fine, but a series of calls that verifies a flow is an acceptance test in `tests/acceptance/`. |

## Acceptance tests

A step definition for a rule that only shows across a flow of calls calls the
running app's API over HTTP.
