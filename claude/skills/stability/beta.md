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
- **Engineers building on it**: build on it, and watch for changes. We aim to
  announce them ahead of time.

## Purpose

The client agrees the feature's shape has settled, and people use it for real
work. Now invest in hardening the edges and operating it responsibly: keep
their data, handle the edge cases, see what it's doing in production, and be
able to roll it back.

> The feature's shape is proven enough that hardening it is now worth the
> investment.

## Expected

The items marked Prototype, Alpha, and Beta in [checklist.md](checklist.md).
In short, on top of Alpha's: success agreed with the client; edge cases and
error paths checked and handled; data kept across changes, through migrations
that apply the same everywhere and can be undone; tried at realistic volumes
with bounded work; personal data exportable and deletable; protected against
misuse; legal duties known; reproducible builds with vulnerability scanning;
timeouts on external calls; correlated logs, an error tracker, and metrics on
key operations; automated health checks and a dashboard; released through the
normal pipeline and can be rolled back; bugs triaged and shared; waiting,
empty, and error states handled; usable by everyone it's meant for; how it's
put together documented; release notes; earlier shortcuts removed or tracked;
and a clear public interface.

Here: the testing rule's Beta suites (acceptance tests for the edge-case
business rules, unit and integration tests where useful) and the stack's CI
rule's Beta jobs; an HTTP API's list endpoints paginated, as the stack's
pagination rule describes; and the engineering docs the `engineering-docs`
skill's level table names for Beta.

## Ready for GA when

It has held up in real use, and we're highly confident it solves the right
problem in the right way. The client has tried it along the way, so it doesn't
go past Beta on a guess.

## Changing it

Prefer non-breaking changes. Make a breaking change only when the user directs
it, keep it to a minimum, and aim to announce it to the people using the
feature ahead of time, with how to move to the new behavior. That's the goal,
not a promise: Beta tells people breaking changes are announced ahead of time
when we can, so don't claim more than that in release notes or copy.

## Still avoid

GA's work can wait until the feature is headed there: alerts, latency and
availability targets, traces of critical workflows, tested restores and
rollbacks, runbooks, a security review, least-privilege access and audit
trails, concurrency and partial-failure handling, versioning and deprecation
windows, and support outside working hours.

Beta is not an excuse to implement hypothetical enterprise requirements. Avoid
complexity not tied to real usage: scale beyond realistic projections,
resilience for failure modes that can't plausibly occur, and configurability
nobody needs.
