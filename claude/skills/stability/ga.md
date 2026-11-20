# GA

**Stable.** Generally available and stable. It keeps working as it does today.
Breaking changes are rare, announced ahead of time, and come with time to move
over.

## What people can expect

- **Changes**: rare, announced well ahead, with a deprecation window.
- **Bugs**: rare, and fixed first.
- **Data**: kept, and recoverable after a loss.
- **People using it**: depend on it, including for business-critical work.
- **Engineers building on it**: build on it. Its interfaces are stable.

## Purpose

We're intentionally letting people depend on this, and will engineer and
operate it accordingly: reliability, compatibility, maintainability, and
long-term operation.

## Expected

Every outcome in [checklist.md](checklist.md) that applies. Some often don't,
such as an availability target for software on people's own devices, or alerts
for a library. Say why.

Here: the testing rule's GA expectations, and the engineering docs the
`engineering-docs` skill's level table names for GA, complete enough for an
operator who didn't build it.

## Changing it

Never make a breaking change in place. If a request seems to need one, clarify
the direction with the user, and propose a new version:

1. **It ships as a new version** alongside the current one: a new API version
   (such as `/v2`), a new major version, or a new workflow next to the old one.
2. **The new version has its own level**, usually Alpha or Beta, with its own
   capability doc. The current one stays GA, and fixes to it are GA work.
3. **The current version keeps working** through a deprecation window agreed
   with the client, with notice and a migration path.

Non-breaking changes, such as new features, optional fields, and bug fixes,
ship without a new version. Small changes target GA directly. A new capability
can start at Prototype or Alpha, opt-in, without putting the rest at risk.

## What GA is not

Not perfection: GA still has bugs and written-down known gaps. It doesn't mean
every edge case is handled, or license speculative extensibility or scale. GA
has overengineering too, such as tighter targets than agreed, round-the-clock
on-call for work nobody needs out of hours, or a penetration test neither the
risk nor the client's duties call for. `checklist/<quality>.md` has more.
