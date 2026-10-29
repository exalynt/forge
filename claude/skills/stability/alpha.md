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
- each endpoint's happy path has a black-box end-to-end test, and the core
  business rules, the primary workflow included, have acceptance tests (in
  Gherkin, run with Cucumber); that is all the testing Alpha needs. A library
  has no acceptance tests, and its end-to-end tests walk common consumer
  flows through its public API (see the testing rule's "Libraries")
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
  - an e2e-test job that starts the needed dependencies and runs the
    black-box end-to-end tests
  - an acceptance-test job that starts what the acceptance tests need and
    runs them; a library has none
- changes receive normal code review
- known limitations are documented

## Changing it

No compatibility guarantee. Take the quickest path, even if it breaks existing
behavior, data, or clients.

## Avoid (premature Beta/GA work)

Don't automatically introduce:

- authentication and authorization (permissions); they are first required at
  Beta
- unit and integration tests, or CI jobs for them, and acceptance tests for
  edge-case business rules the work isn't implementing, changing, or
  clarifying; deeper testing starts at Beta
- pagination of list endpoints; it is first required at Beta
- exhaustive error-path tests
- comprehensive dashboards, sophisticated alerting, or elaborate SLOs
- large-scale load testing or production-scale capacity planning
- exhaustive accessibility review
- compatibility guarantees or long-term migration strategies
- extensive architecture documentation
- generalized plugin or extension frameworks
- complex rollback systems
- infrastructure needed only for hypothetical future scale
