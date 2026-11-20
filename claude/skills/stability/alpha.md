# Alpha

**Developing.** Works, but is still taking shape. Expect frequent breaking
changes and bugs.

## What people can expect

- **Changes**: frequent breaking changes, with no promise of notice.
- **Bugs**: expected, even in the main workflow.
- **Data**: may be reset. Don't keep the only copy of anything there.
- **People using it**: try it with real work, keeping their current way of
  working as a fallback.
- **Engineers building on it**: only to experiment. Its interfaces will change.

## Purpose

Make the happy path real without pretending the shape is settled. The main
workflow works end to end on the real system and is used for real, opt-in, so
feedback comes from real use. Alpha validates the real workflow; it doesn't
harden every edge.

## Expected

The Prototype and Alpha outcomes in [checklist.md](checklist.md).

Here, the main path's automated checks are the testing rule's Alpha suites, an
end-to-end test of each endpoint's happy path and acceptance tests for the core
business rules, run by the stack's CI rule's Alpha jobs. That's all the testing
Alpha needs. A library's end-to-end tests walk common consumer flows through
its public API (see the testing rule's "Libraries").

## Ready for Beta when

Its shape has settled, so big changes are unlikely, and real work has gone
through the main workflows end to end.

## Changing it

No compatibility guarantee, and no promise of notice. Take the quickest path.

## Often premature

Don't add these automatically; a client's requirement, a law, a real incident,
or a security need can call for one now: unit and integration tests, and
acceptance tests for edge-case rules the work isn't touching; exhaustive
edge-case handling; down migrations, restore drills, or rollback plans (a
simple forward migration is fine when people still act on the records);
pagination; error trackers, metrics, dashboards, tracing, or alerts beyond what
the host gives for free; load tests, caches, or scaling; complete docs or
accessibility audits; per-caller rate limits; deprecation windows or versioned
copies; security reviews or certifications nobody requires; and interfaces,
plugin points, settings, or layers with a single use. `checklist/<quality>.md`
has more.

The security and privacy baseline isn't on this list. If it touches real client
data, access control is required now.
