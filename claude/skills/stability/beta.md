# Beta

**Refining.** Settled enough for real use. Breaking changes and bugs are
occasional. We aim to announce breaking changes ahead of time, but can't
always.

## What people can expect

- **Changes**: occasional breaking changes, announced ahead of time when we can.
- **Bugs**: occasional, mostly at the edges, and fixed as they're found.
- **Data**: kept across changes.
- **People using it**: use it for real work, with a fallback for anything
  critical.
- **Engineers building on it**: build on it, and watch for changes.

## Purpose

The client agrees the shape has settled, and people use it for real work. Now
hardening is worth the investment: keep their data, handle the edges, see what
it's doing in production, and be able to undo a bad release.

## Expected

The Prototype, Alpha, and Beta outcomes in [checklist.md](checklist.md).

Here: the testing rule's Beta suites (acceptance tests for the edge-case
business rules, unit and integration tests where useful) and the stack's CI
rule's Beta jobs; an HTTP API's list endpoints paginated, as the stack's
pagination rule describes; and the engineering docs the `engineering-docs`
skill's level table names for Beta.

## Ready for GA when

It has held up in real use, and we're highly confident it solves the right
problem in the right way.

## Changing it

Prefer non-breaking changes. Make a breaking change only when the user directs
it, keep it to a minimum, and aim to announce it ahead of time, with how to
move. That's a goal, not a promise, so don't claim more in release notes or
copy.

## Still avoid

GA's outcomes can usually wait until it's headed there, such as alerts against
targets, runbooks, tested restores, load tests at peak, a security review, an
audit log, and deprecation windows. A client's requirement, a law, or a real
incident can bring one forward. Beta isn't an excuse for hypothetical
enterprise requirements: scale beyond realistic projections, resilience for
failures that can't plausibly happen, or configurability nobody needs.
`checklist/<quality>.md` has more.
