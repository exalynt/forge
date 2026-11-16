# Feature Stability

Features may declare a stability level:

Prototype → Alpha → Beta → GA

- **Prototype** (exploring): explores an idea so we can choose a direction. Its
  data may be fake, and it may change completely or be thrown away.
- **Alpha** (developing): works, but is still taking shape. Expect frequent
  breaking changes and bugs, and data that may be reset.
- **Beta** (refining): settled enough for real use. Breaking changes and bugs
  are occasional, and data is kept. We aim to announce breaking changes ahead
  of time, but can't always.
- **GA** (stable): generally available and stable. It keeps working as it does
  today. Breaking changes are rare, announced ahead of time, and come with time
  to move over.

A level describes how much people can depend on the work, not how good it is.
Stability isn't perfection: GA software can still have bugs and known gaps.

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
- What each level includes is a set of guidelines, not a gate: the `stability`
  skill's checklist. An item that applies and isn't met is either fixed or
  written down as a known gap, to close later or accepted as a tradeoff, where
  people can see it. Known gaps can exist at any level, GA included.
- Never promote or demote a feature unless explicitly instructed. Promotion is
  a deliberate decision, made with `/promote`.
- The security and privacy baseline applies at every level, including
  Prototype, and is never a known gap: secrets stay out of code, version
  control, sample data, logs, and output; real client data is never copied
  into code or fixtures; anything touching real client systems or data has
  access control, so people only see and change what they're allowed to and
  one client or customer never sees another's; personal data is collected only
  where needed; and untrusted input can't run as code.

## Prototype the uncertainty

A prototype exists to answer a question. Know what that question is before
deciding what to build. Put realism into the part of the system being learned
about, and keep everything else as simple as reasonably possible. Mocking the
backend is fine when the question is whether a UI workflow makes sense. It
defeats the purpose when the question is whether that backend integration can
do what we need.

## Where levels are declared

A feature's level is the `**Stability:**` line in its capability doc,
`docs/product/capabilities/<slug>.md`, along with the known gaps people would
notice. Technical known gaps are in `docs/engineering/`, each saying what closes
it. A project that
also records levels somewhere else says where in its `CLAUDE.md`. The user's
word on a feature's level wins over the declaration. Without
either, infer the level from the repository's code and docs, as the
`stability` skill describes, and tell the user what you inferred.

## Changing existing features

Before adding to, adjusting, or removing from a feature, find its level. It
decides whether the change must be non-breaking:

- **Prototype and Alpha**: no compatibility guarantee. Take the quickest path.
- **Beta**: prefer non-breaking changes. Break only when the user directs it,
  and aim to announce the break to the people using it ahead of time. That's
  the goal, not a guarantee.
- **GA**: non-breaking only. A change that would break something ships as a
  new version with its own level, while the current one keeps working through
  a deprecation window. Clarify the direction with the user first.

A new capability in a GA feature starts at its own level, opt-in.

When a change touches something several features share, the strictest level
applies.

The detailed expectations for each level are in the `stability` skill.
