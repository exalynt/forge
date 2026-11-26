# Go: tests

For Go projects, adding to the `testing` skill, which says when to write
tests and what each kind covers. Where this project's own code or instructions
differ, follow the project.

## Levels

Unit and integration tests are Go, separated by build tags so each suite runs
on its own. A stack skill says how end-to-end tests run, such as `go-http`'s;
a library's are below.

| Level | Tool and location | Run with | Covers |
| --- | --- | --- | --- |
| Unit | Go, no build tag, next to the code | `go test ./...` | Table-driven tests (`tests := []struct{…}` with `t.Run`), dependencies mocked with moq: service logic, validation, error mapping. |
| Integration | Go, `//go:build integration`, next to the code | `go test -tags integration ./...` | A specific piece of happy-path functionality against real dependencies such as Postgres, for example a repository against a real database. |

## Acceptance tests

Acceptance tests are Gherkin run with
[godog](https://github.com/cucumber/godog), Cucumber's Go implementation, at
every level (see the `testing` skill):

- Every acceptance test lives in `tests/acceptance/`, one Go package,
  whatever level proves it. The `.feature` files go in
  `tests/acceptance/features/`, one per feature or area of the domain. With
  Docqet, each `Rule:` carries its `docqet:verifies` annotation, as the
  `testing` skill describes.
- One `TestAcceptance` function runs them all, in a test file with the
  `acceptance` build tag, through `godog.TestSuite` with `TestingT: t`. The
  tag keeps them out of the unit and integration suites, and
  `go test -tags acceptance -run '^TestAcceptance$' ./...` runs them without
  the unit tests.
- Step definitions live in the same package, also tagged `acceptance`. They
  call the code at the rule's level: a service, with moq mocks for its
  boundaries or real dependencies where the rule needs them. A stack skill
  says how they reach the running app, where it has one.

## Libraries

A Go module other projects import has no acceptance tests, and its
end-to-end tests are Go (see the `testing` skill's "Libraries"):

- They live in `tests/e2e/`, tagged `//go:build e2e`, in a package that
  imports the library by its module path and uses only what it exports.
- Each test walks one common consumer flow's happy path through the public
  API, against the real dependencies it supports.
- `go test -tags e2e ./tests/e2e/...` runs them. In a repository of several
  modules, `tests/e2e/` may be its own module so it can import the nested
  ones.

## Mocks

- Generate mocks with moq from the interface, through a `//go:generate moq …`
  directive, and commit them. Don't hand-write mocks.
- Mock the boundary interfaces (repositories, providers), not the code under
  test.

## Real dependencies

Integration, end-to-end, and acceptance tests get their dependencies from the
project's local stack (such as Docker Compose) through environment variables.
Each test isolates its own data so the suite can run repeatedly.
