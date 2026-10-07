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
- the main path has automated tests
- failures are visible, not swallowed
- reasonable input validation
- database and schema changes use migrations
- temporary resources are cleaned up appropriately
- basic application logging
- dependencies are reasonably maintained
- normal linting, type checking, and static analysis pass
- changes receive normal code review
- known limitations are documented

## Avoid (premature Beta/GA work)

Don't automatically introduce:

- exhaustive error-path tests
- comprehensive dashboards, sophisticated alerting, or elaborate SLOs
- large-scale load testing or production-scale capacity planning
- exhaustive accessibility review
- compatibility guarantees or long-term migration strategies
- extensive architecture documentation
- generalized plugin or extension frameworks
- complex rollback systems
- infrastructure needed only for hypothetical future scale
