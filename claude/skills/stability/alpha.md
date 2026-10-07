# Alpha

## Purpose

Turn a promising idea into a real working feature. The main workflow is real
and usable, though the design may still change substantially. Alpha optimizes
for validating the real workflow, not for hardening every edge.

> Make the happy path real without pretending the shape of the feature is
> completely settled.

## Expected (where applicable)

- the primary workflow works end-to-end, and users can actually complete it
- real persistence and integrations replace mocks where necessary
- common paths have high-level end-to-end or black-box automated tests; that
  is all the testing Alpha needs
- failures are visible, not swallowed
- reasonable input validation
- database and schema changes use migrations
- temporary resources are cleaned up appropriately
- basic application logging
- dependencies are reasonably maintained
- normal linting, type checking, and static analysis pass
- a CI workflow (`.github/workflows/ci.yml`) runs on pushes and pull requests
  with:
  - the language's standard checks: linting, formatting, static analysis, and
    a build, plus a container image build if the project uses one. CI only
    checks; it never commits fixes
  - a feature-test job that starts the needed dependencies and runs the
    black-box feature tests
- changes receive normal code review
- known limitations are documented

## Changing it

No compatibility guarantee. Take the quickest path, even if it breaks existing
behavior, data, or clients.

## Avoid (premature Beta/GA work)

Don't automatically introduce:

- authentication and authorization (permissions); they are first required at
  Beta
- unit and integration tests, or CI jobs for them; deeper testing starts at
  Beta
- exhaustive error-path tests
- comprehensive dashboards, sophisticated alerting, or elaborate SLOs
- large-scale load testing or production-scale capacity planning
- exhaustive accessibility review
- compatibility guarantees or long-term migration strategies
- extensive architecture documentation
- generalized plugin or extension frameworks
- complex rollback systems
- infrastructure needed only for hypothetical future scale
