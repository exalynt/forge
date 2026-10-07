# Beta

## Purpose

The shape of the feature has largely settled, and people may begin using it in
meaningful real-world situations. Now invest in hardening the edges and
operating it responsibly.

> The feature's shape is proven enough that hardening it is now worth the
> investment.

## Expected (where applicable)

Testing and correctness:

- primary and important secondary workflows are tested
- deeper unit and integration tests are added where they're useful, beyond
  Alpha's end-to-end tests
- CI adds a unit-test job and an integration-test job; the integration job
  starts the dependencies it needs
- important failure paths are handled
- meaningful edge cases are handled
- concurrency and race behavior have been considered

Operations:

- observability sufficient to understand production behavior: useful metrics,
  and tracing where appropriate
- health and readiness behavior where appropriate
- operational dashboards where useful
- alerts for meaningful operational conditions
- rollback or recovery procedures where necessary
- dependencies and external integrations have appropriate resilience
- graceful degradation considered where appropriate

Scale and performance:

- realistic data volumes considered
- realistic performance validated

Review and documentation:

- accessibility receives meaningful review
- authentication and authorization (permissions) are in place; they are first
  required at Beta
- security receives deeper review
- operational documentation exists
- architecture and meaningful tradeoffs are documented
- someone other than the original author can understand and support it

## Still avoid

Beta is not an excuse to implement hypothetical enterprise requirements. Avoid
complexity not tied to real usage: scale beyond realistic projections, resilience
for failure modes that can't plausibly occur, configurability nobody needs,
compatibility commitments before anyone depends on the feature (that's GA).
