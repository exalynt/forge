---
name: stability
description: Feature stability model (Prototype → Alpha → Beta → GA) for deciding how much engineering a feature warrants. Use when designing, implementing, modifying, or reviewing a feature whose stability level is declared or matters to the decision; when deciding whether work such as tests, hardening, abstraction, observability, scaling, or documentation is appropriate now; or when discussing whether something is Prototype, Alpha, Beta, or GA work.
---

# Feature Stability

> The goal is NOT to make every feature as robust as possible.

> A stability level is an engineering investment ceiling as well as a quality expectation.

Every stability decision asks two questions with equal weight:

1. Does the work meet what the current level expects? (Under-engineering.)
2. Does the work stay within what the current level justifies? (Over-engineering.)

## Levels

| Level     | Purpose                                         | Reference                      |
| --------- | ----------------------------------------------- | ------------------------------ |
| Prototype | Learn. Answer a specific unknown.               | [prototype.md](prototype.md)   |
| Alpha     | Make the primary workflow real and usable.      | [alpha.md](alpha.md)           |
| Beta      | Shape has settled. Harden edges, operate well.  | [beta.md](beta.md)             |
| GA        | People may depend on it. Commit accordingly.    | [ga.md](ga.md)                 |

Read the reference for the feature's current level before deciding. When the
question is whether some work belongs now or later, also read the next level's
reference. Anything that first appears at a later level is UNNECESSARY NOW by
default.

"Where applicable" in the references means it: skip expectations that don't
fit the feature (a feature with no UI has no accessibility review, a feature
with no schema has no migrations).

## Determining the level

1. **The project's declaration.** The project's `CLAUDE.md` should say where
   feature stability is recorded. Use that.
2. **The user's statement** in the current conversation.
3. **Otherwise**, if the level materially affects the decision, ask. Don't
   infer it from how polished the code looks. Mature-looking code may be an
   over-built prototype.

A simple convention some projects use is one YAML file per feature, such as
`.exalynt/features/conversations.yaml`. Follow whatever the project uses. Don't
impose this one unless the project has adopted it.

```yaml
name: Prospect Conversations
stability: prototype

purpose: >
  Determine whether allowing a prospect to begin a conversation before
  creating a full account lowers the barrier to starting a relationship.

learning:
  - Will prospects use conversations as the initial point of contact?
  - Is email sufficient identity for reconnecting conversations later?

known_gaps:
  - notification delivery may initially be simple
  - abuse prevention is incomplete
  - conversation search has not been designed
```

When a declaration has them, `purpose` and `learning` say what the feature is
for right now, which is essential for prototypes. `known_gaps` lists what is
intentionally incomplete. Don't "fix" those gaps opportunistically.

## Security and privacy

Security and privacy apply at every level. The baseline never relaxes:

- authentication and authorization on anything touching real users or data
- no secrets in code, logs, or client bundles
- input validation at trust boundaries
- personal data collected minimally and kept out of logs and unintended third parties
- prototype or unfinished behavior not unintentionally exposed to real users or data

The depth of security review grows with the level, but the baseline does not.

## Classifying work

When reviewing or planning, classify each item:

- **REQUIRED NOW**: necessary to appropriately meet the current level.
- **UNNECESSARY NOW**: reasonable engineering, but not justified at the current
  level. Name the level where it would become appropriate. If it is already
  built and carries real cost, prefer removing or simplifying it.
- **ACCEPTABLE**: appropriate for the current level, including known gaps the
  level allows.
- **SECURITY/PRIVACY**: required regardless of level.

## Rules

- Never change a feature's level as a side effect of other work. Promotion is
  done deliberately with `/promote`.
- Work above the current level needs explicit, concrete justification.
- Don't suggest promoting a feature unless asked.
