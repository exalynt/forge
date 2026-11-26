# Go HTTP projects: CI

General conventions for Go HTTP API projects. Where this project's own code or
instructions differ, follow the project.

## When

CI lives in `.github/workflows/ci.yml`. Add it once a capability's level
expects CI (Docket's rubric, when the repo has it), and add each job once a
capability's level expects what it runs. Don't add jobs for tests no
capability's level expects yet.

## Workflow

- Run on pushes to `main` and on pull requests.
- Set `permissions: contents: read`.
- Set up Go with `actions/setup-go` and `go-version-file: go.mod`.
- CI only checks. It never commits, pushes, or opens pull requests with fixes,
  and it never publishes images.

## Jobs

| Job | Added when | Steps |
| --- | --- | --- |
| `check` | CI is added | `go vet ./...`; golangci-lint through `golangci/golangci-lint-action`, using the project's `.golangci.yml` if there is one; `go fmt ./...` and `go fix ./...`, then `git diff --exit-code` so the job fails if either changed anything; `go build ./...`. |
| `docker-build` | CI is added, only if the project has a Dockerfile | `docker build .`, without pushing. |
| `e2e-test` | End-to-end tests are expected | Start the dependencies with the project's local stack (`docker compose up -d --wait`), start the app and wait until it responds, then `npm ci` and `npx playwright test tests/e2e`. Install Playwright browsers only if the tests use a browser. |
| `acceptance-test` | Acceptance tests are expected | Start the dependencies with `docker compose up -d --wait`, start the app and wait until it responds if any step definitions call its API, then `go test -tags acceptance -run '^TestAcceptance$' ./...`. |
| `unit-test` | Unit tests are expected | `go test ./...` |
| `integration-test` | Integration tests are expected | Start the dependencies the tests need with `docker compose up -d --wait <services>`, then `go test -tags integration ./...`. |

A library has no `acceptance-test` job, and its `e2e-test` job starts the
dependencies its tests need with `docker compose up -d --wait <services>`,
then runs `go test -tags e2e ./tests/e2e/...`, with no app to start.

Pass dependency connection details and the app's URL to tests through the same
environment variables used locally (see the testing reference), so CI and local
runs behave the same way.
