# Prototype

**Exploring.** Explores an idea so we can choose a direction. Its data may be
fake, and it may change completely or be thrown away.

## What people can expect

- **Changes**: anything can change, or it may be thrown away.
- **Bugs**: only what's being shown is expected to work, and its data may be
  fake.
- **Data**: may be sample data, and may be wiped.
- **People using it**: react to it and help choose the direction. Don't use it
  for real work, or trust the data it shows.
- **Engineers building on it**: don't build on it. Expect it to be rewritten.

## Purpose

A prototype exists to learn, and may be rewritten or deleted entirely. Success
means learning enough to decide what to do next. A prototype doesn't have to
move on: many are done once they've answered their question, and "don't build
it" is a successful outcome. Work that's dropped is cleaned up, flags and data
included, so it doesn't get in the way of what comes next.

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
- getting the client's reaction, early and often
- fast iteration
- low cost of change and of deletion
- proving or disproving assumptions

## Expected

The items marked Prototype in [checklist.md](checklist.md), the security and
privacy baseline above all. In short: what it's for is written down, the client
can try it and their feedback is recorded, the main path works by hand, sample
data is clearly marked, its level is clear wherever people see it, it isn't
available to users by default, and it can be removed cleanly without affecting
anything people rely on.

## Ready for Alpha when

The client likes the direction, and it's built to actually work, not just to
show the idea.

## Changing it

No compatibility guarantee, and no promise of notice. Take the quickest path,
even if it breaks existing behavior, data, or clients.

## Avoid unless it is what we're learning about

Don't add these merely because they'd be useful later:

- generalized abstractions, speculative interfaces, plugin or extension mechanisms
- speculative configuration
- exhaustive edge-case handling
- automated tests and CI workflows (they start at Alpha)
- elaborate retries or sophisticated failure recovery
- logging beyond what the experiment needs, dashboards, alerts, or deep
  observability infrastructure
- production-scale architecture or architecture built around hypothetical scale
- performance optimization
- compatibility mechanisms or public API stability guarantees
- documentation beyond what it does, what it doesn't handle yet, and how to run
  it, or engineering docs beyond recording a stub `/qa` should allow
- elaborate deployment infrastructure
- abstraction to support future implementations
