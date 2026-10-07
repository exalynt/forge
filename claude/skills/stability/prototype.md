# Prototype

## Purpose

Explore an idea and answer important unknowns. A prototype exists primarily to
learn. It may be rewritten or deleted entirely. Success means learning enough to
decide what to do next. Stopping after answering the current question is
success.

## Prototype the uncertainty

First identify the question the prototype is trying to answer. Prototype does
not mean "mock everything."

- If the question is *does this UI workflow make sense?*, mocking the backend
  may be appropriate.
- If the question is *can Microsoft 365 support the scheduling behavior we
  need?*, the Microsoft 365 integration must be real, because that integration
  is the thing we are learning about.

Put realism into the part of the system under question. Keep unrelated parts
as simple as reasonably possible:

- If an integration is not what we're learning about, mocking it may be preferable.
- If persistence is not what we're learning about, use the simplest reasonable persistence.
- If architecture is not what we're learning about, use the simplest
  architecture that demonstrates the idea.

## Optimize for

- making the idea tangible
- answering the current unknown
- getting user, client, or team reaction
- fast iteration
- low cost of change and of deletion
- proving or disproving assumptions

## Expected (where applicable)

- the primary demo path works when manually exercised
- known gaps are documented
- it is clear what is real and what is sample, mock, or fake
- performance is sufficient for the experiment or demo
- the security and privacy baseline is respected: one tenant or client cannot
  see or interact with another's real data, and real secrets are never leaked
  anywhere (authentication and authorization are not required until Beta)
- prototype behavior is not unintentionally exposed as production behavior
- experimental code is reasonably isolated from stable production behavior

## Changing it

No compatibility guarantee. Take the quickest path, even if it breaks existing
behavior, data, or clients.

## Avoid unless it is what we're learning about

Don't add these merely because they'd be useful later:

- generalized abstractions, speculative interfaces, plugin or extension mechanisms
- speculative configuration
- exhaustive edge-case handling
- exhaustive automated testing
- elaborate retries or sophisticated failure recovery
- dashboards, alerts, or deep observability infrastructure
- production-scale architecture or architecture built around hypothetical scale
- large-scale performance optimization
- long-term compatibility mechanisms or public API stability guarantees
- exhaustive documentation
- elaborate deployment infrastructure
- CI workflows (they start at Alpha)
- abstraction to support future implementations
