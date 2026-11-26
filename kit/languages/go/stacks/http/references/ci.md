# Go HTTP: CI

For Go HTTP APIs, adding to the `go` skill's
[ci.md](../../go/references/ci.md). Where this project's own code or
instructions differ, follow the project.

| Job | Added when | Steps |
| --- | --- | --- |
| `e2e-test` | End-to-end tests are expected | Start the dependencies with the project's local stack (`docker compose up -d --wait`), start the app and wait until it responds, then `npm ci` and `npx playwright test tests/e2e`. Install Playwright browsers only if the tests use a browser. |
| `acceptance-test` | Acceptance tests are expected | As the `go` skill's, and start the app and wait until it responds if any step definitions call its API. |
