# Go: CI

For Go projects, adding to the `standards` skill's
[ci.md](../../standards/references/ci.md). Where this project's own code or
instructions differ, follow the project.

## Setup

Set up Go with `actions/setup-go` and `go-version-file: go.mod`.

## Jobs

| Job | Added when | Steps |
| --- | --- | --- |
| `check` | CI is added | `go vet ./...`; golangci-lint through `golangci/golangci-lint-action`, using the project's `.golangci.yml` if there is one; `go fmt ./...` and `go fix ./...`, then `git diff --exit-code` so the job fails if either changed anything; `go build ./...`. |
| `docker-build` | CI is added, only if the project has a Dockerfile | `docker build .`, without pushing. |
| `unit-test` | Unit tests are expected | `go test ./...` |
| `integration-test` | Integration tests are expected | Start the dependencies the tests need with `docker compose up -d --wait <services>`, then `go test -tags integration ./...`. |
| `acceptance-test` | Acceptance tests are expected | Start the dependencies with `docker compose up -d --wait`, then `go test -tags acceptance -run '^TestAcceptance$' ./...`. A stack skill adds what else it needs, such as the running app. |
| `e2e-test` | End-to-end tests are expected | For a library: start the dependencies its tests need with `docker compose up -d --wait <services>`, then `go test -tags e2e ./tests/e2e/...`. A stack skill says how its own end-to-end tests run. |

A library has no `acceptance-test` job.
