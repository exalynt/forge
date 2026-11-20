# Prototype

**Exploring.** Explores an idea so we can choose a direction. Its data may be
fake, and it may change completely or be thrown away.

## What people can expect

- **Changes**: anything can change, or it may be thrown away.
- **Bugs**: only what's being shown is expected to work.
- **Data**: may be sample data, and may be wiped.
- **People using it**: react to it and help choose the direction. Don't use it
  for real work, or trust the data it shows.
- **Engineers building on it**: don't. Expect it to be rewritten.

## Purpose

A prototype exists to answer a question. Success means learning enough to
decide what to do next, and "don't build it" is a successful outcome. Many are
done once they've answered it. Dropped work is cleaned up, flags and data
included.

Prototype doesn't mean "mock everything." Put realism into the part under
question and keep the rest as simple as reasonably possible. Mocking the
backend is fine when the question is whether a UI workflow makes sense; if the
question is whether Microsoft 365 can support the scheduling we need, that
integration must be real.

## Expected

The Prototype outcomes in [checklist.md](checklist.md), the security and
privacy baseline above all. Most qualities expect nothing yet.

## Ready for Alpha when

The client likes the direction, and it's built to actually work, not just to
show the idea.

## Changing it

No compatibility guarantee, and no promise of notice. Take the quickest path.

## Often overengineering

Unless it's what we're learning about, or the circumstances call for it:
polish, abstractions, configuration, or edge-case handling the idea doesn't
depend on; automated tests and CI, unless they're the quickest way to try the
idea or the baseline needs one, such as a test that tenants stay apart;
migrations, retries, or recovery; logging, metrics, or alerts beyond console
output; optimizing or caching; a production-like environment or deploy
pipeline when a local run or one preview lets people try it; real emails or
texts when showing the message would answer the question; and docs beyond what
it does, what it doesn't handle yet, and how to run it. `checklist/<quality>.md`
has more.
