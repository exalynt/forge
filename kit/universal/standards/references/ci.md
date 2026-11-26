# Standards: CI

For every project, in any language. Language and stack skills add to this,
with the setup and jobs each needs. Where this project's own code or
instructions differ, follow the project.

## When

CI lives in `.github/workflows/ci.yml`. Add it once a capability's level
expects CI (Docqet's rubric, when the repo has it), and add each job once a
capability's level expects what it runs. Don't add jobs for tests no
capability's level expects yet.

## Workflow

- Run on pushes to `main` and on pull requests.
- Set `permissions: contents: read`.
- CI only checks. It never commits, pushes, opens pull requests with fixes,
  publishes images, or deploys. Deployment, if any, is its own workflow.
- Separate jobs, so each failure shows on its own.

## Dependencies

- Jobs that need real dependencies start them with the project's local stack,
  such as `docker compose up -d --wait <services>`.
- Pass dependency connection details and the app's URL to tests through the
  same environment variables used locally, so CI and local runs behave the
  same way.
