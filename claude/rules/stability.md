# Feature Stability

Features may declare a stability level:

Prototype → Alpha → Beta → GA

- **Prototype**: learn and answer a specific unknown. It may be rewritten or deleted.
- **Alpha**: the primary workflow is real and usable, though the shape may still change.
- **Beta**: the shape has settled. Harden the edges and operate it responsibly.
- **GA**: people may depend on it. We commit to reliability, compatibility, and
  long-term operation.

## Rules

- A stability level is both a quality expectation and an engineering investment
  ceiling. Meeting the level is required. Exceeding it is not a virtue.
- Prefer work appropriate to the current level. Do not opportunistically
  implement a later level's requirements "while we're here."
- Work beyond the current level needs an explicit, concrete justification, such
  as a request, a real incident, or a security or privacy need. "We'll need it
  at GA" is not one.
- Over-engineering is a stability problem, just like under-engineering. Flag
  both when designing, modifying, or reviewing a feature.
- Never promote or demote a feature unless explicitly instructed. Promotion is
  a deliberate decision, made with `/promote`.
- The security and privacy baseline applies at every level, including
  Prototype: one tenant or client cannot see or interact with another's real
  data, and real secrets are never leaked anywhere.

## Prototype the uncertainty

A prototype exists to answer a question. Know what that question is before
deciding what to build. Put realism into the part of the system being learned
about, and keep everything else as simple as reasonably possible. Mocking the
backend is fine when the question is whether a UI workflow makes sense. It
defeats the purpose when the question is whether that backend integration can
do what we need.

## Where levels are declared

Each project records stability in its own way, and its `CLAUDE.md` should say
where. If a feature has no declared level and the level materially affects the
work, ask rather than assume one.

The detailed expectations for each level are in the `stability` skill.
