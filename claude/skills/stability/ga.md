# GA

**Stable.** Generally available and stable. It keeps working as it does today.
Breaking changes are rare, announced ahead of time, and come with time to move
over.

## What people can expect

- **Changes**: rare. When one is needed, we aim to announce it well ahead, with
  a deprecation window.
- **Bugs**: rare, and fixed first.
- **Data**: kept, and recoverable after a loss.
- **People using it**: depend on it, including for business-critical work.
- **Engineers building on it**: build on it. Its interfaces are stable, and
  changes come with notice and time to move.

## Purpose

People are allowed to depend on this feature. GA is a commitment to
reliability, compatibility, maintainability, and long-term operation.

> We are intentionally allowing people to depend on this and will engineer and
> operate it accordingly.

## Expected

Everything in [checklist.md](checklist.md). In short, on top of Beta's: shown
to meet its success in real use; critical workflows checked end to end, a
regression check for every fixed bug, and no untracked flaky or skipped
checks; concurrency, partial failures, and safe repeats handled; data changes
tried at production size, and recovery tried; measured performance targets,
headroom at peak, and a known running cost; a security review where it handles
credentials, payments, personal data, or permissions; least-privilege access
and an audit trail; data and logs kept only as long as needed; legal duties
met; dependencies kept current; workflow success measured and traces of
critical workflows; targets tracked and alerts that fire, are quiet, and say
what to do; rollback tried and rebuildable from scratch; GA bugs fixed before
new work, a known support response, an owner for the long term, and incident
reviews; supported platforms listed and the accessibility standard met; complete user, reference, and operations
documentation; changeable without the author; and public interfaces documented
as stable.

Here: the testing rule's GA expectations, and the engineering docs the
`engineering-docs` skill's level table names for GA, complete enough for an
operator who didn't build it.

## Changing it

Breaking changes are rare and planned. Never make one in place. If a request
seems to need one, clarify the direction with the user first, and propose a new
version:

1. **The change ships as a new version** alongside the current one: a new API
   version (such as `/v2`), a new major version of a library, or a new workflow
   next to the old one behind a flag or setting.
2. **The new version has its own level**, usually Alpha or Beta, with its own
   capability doc, while the current version stays GA and fixes to it are GA
   work.
3. **The current version keeps working** through a deprecation window agreed
   with the client, with notice and a migration path, before it's retired.

Changes that don't break anything, such as new features, new optional fields,
and bug fixes, ship without a new version.

## GA keeps improving

GA doesn't mean finished:

- **Small changes target GA directly.** A new option on a stable feature, or a
  fix, doesn't start over at Prototype.
- **New capabilities start at their own level.** A new capability in a GA
  feature can start as a Prototype or Alpha, opt-in, and move up on its own
  schedule, kept from putting the rest of the feature at risk.
- **Known gaps get closed** in later iterations. An accepted tradeoff can be
  revisited when needs change.

GA rarely moves down. A change that would break it becomes a new version
instead.

## What GA is not

Stability isn't perfection. GA software still has bugs, and it can still have
known gaps, either to close later or accepted as tradeoffs, as long as they're
written down where people can see them. It does not mean every conceivable
edge case is handled, and it does not license speculative extensibility or
scale beyond what real dependence requires. The ceiling still applies: invest
in what the commitment demands, not in everything imaginable.
