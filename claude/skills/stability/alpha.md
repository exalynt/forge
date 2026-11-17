# Alpha

**Developing.** Works, but is still taking shape. Expect frequent breaking
changes and bugs.

## What people can expect

- **Changes**: frequent breaking changes, with no promise of notice.
- **Bugs**: expected, even in the main workflow.
- **Data**: may be reset, sometimes without warning. Don't keep the only copy of
  anything there.
- **People using it**: try it with real work and share feedback. Keep your
  current way of working as a fallback.
- **Engineers building on it**: build on it only to experiment. Expect its
  interfaces to change.

## Purpose

Turn a promising idea into a real working feature. The main workflow is hooked
up end to end and a real user can complete it, though the design may still
change substantially. It's opt-in, and feedback comes from real use, not just
demos. Alpha optimizes for validating the real workflow, not for hardening
every edge.

> Make the happy path real without pretending the shape of the feature is
> completely settled.

## Expected

The items marked Prototype and Alpha in [checklist.md](checklist.md). In short,
on top of Prototype's: hooked up end to end, opt-in, the main path checked
automatically on every change, errors visible and logged with context, input
validated where it enters, usable with today's real data and use, dependencies
no less stable than itself, settings kept out of code, an owner and somewhere
to report bugs, conventions and automated checks followed, and reviewed by
another engineer where the team has one.

Here, the automated checks are the testing rule's Alpha suites, an end-to-end
test of each endpoint's happy path and acceptance tests for the core business
rules, run by the stack's CI rule's Alpha jobs. That's all the testing Alpha
needs. A library has no acceptance tests, and its end-to-end tests walk common
consumer flows through its public API (see the testing rule's "Libraries").

## Ready for Beta when

Its shape has settled, so big changes are unlikely, and real work has gone
through the main workflows end to end.

## Changing it

No compatibility guarantee, and no promise of notice. Take the quickest path,
even if it breaks existing behavior, data, or clients.

## Avoid (premature Beta and GA work)

Don't automatically introduce:

- unit and integration tests, or CI jobs for them, and acceptance tests for
  edge-case business rules the work isn't implementing, changing, or
  clarifying; deeper testing starts at Beta
- exhaustive edge-case and error-path handling
- data-preserving migrations; data may still be reset (Beta keeps data)
- pagination of list endpoints (Beta)
- error trackers, metrics, dashboards, or health monitoring (Beta); alerts,
  targets, or traces (GA)
- load testing (Beta) or capacity planning for peak load (GA)
- accessibility review (Beta), or meeting an accessibility standard (GA)
- rate limiting and other misuse protection (Beta)
- changelogs or compatibility guarantees (Beta); versioning or deprecation
  windows (GA)
- architecture documentation (Beta)
- rollback mechanisms (Beta), or rebuilding the environment from scratch (GA)
- generalized plugin or extension frameworks
- infrastructure needed only for hypothetical future scale

The security and privacy baseline isn't on this list. If it touches real client
data, access control is required now.
