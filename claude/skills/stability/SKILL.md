---
name: stability
description: Feature stability model (Prototype → Alpha → Beta → GA) for deciding how much engineering a feature warrants. Use when designing, implementing, modifying, or reviewing a feature; when adding to, changing, or removing from an existing feature, since its level decides whether the change must be non-breaking; when deciding whether work such as tests, hardening, abstraction, observability, scaling, or documentation is appropriate now; or when discussing whether something is Prototype, Alpha, Beta, or GA work.
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
3. **Otherwise, the feature is GA.** No declared level means GA. Don't infer a
   lower level from how rough the code looks, or a higher one from how polished
   it looks.

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

## Changing existing features

Adding to, adjusting, or removing from an existing feature is governed by that
feature's level. Find the level before changing anything. When the change
touches something several features share, such as a table or an endpoint, the
strictest of their levels applies.

| Level             | Changes                                                                      |
| ----------------- | ---------------------------------------------------------------------------- |
| Prototype, Alpha  | No compatibility guarantee. Take the quickest path, breaking or not.         |
| Beta              | Prefer non-breaking. Make a breaking change only when the user directs it.   |
| GA                | Non-breaking only. If the request seems to need a breaking change, clarify the direction with the user before making it. |

A breaking change is anything that makes existing clients, data, or
deployments fail or behave differently. Examples: removing or renaming an
endpoint, field, or column; changing a type or meaning; tightening validation;
editing a database migration that has already run instead of adding a new one.

At Prototype and Alpha, the quickest path is often the breaking one: edit the
existing migration and reset the database, or change the endpoint in place.
The security and privacy baseline still applies.

The stack's rules say how to make specific kinds of changes non-breaking.

## Security and privacy

The security and privacy baseline applies at every level and never relaxes. It
is only this:

- one tenant or client cannot see or interact with another tenant's or
  client's real data
- real secrets are never leaked anywhere: not in code, version control, logs,
  error messages, client bundles, or third-party services

At Prototype and Alpha, isolation can be as simple as separate deployments or
having only one tenant's real data present. Sample, mock, or fake data and
throwaway placeholder credentials are not covered by the baseline.

Authentication and authorization (permissions) are required only at Beta and
above. They are not part of the Prototype or Alpha baseline, so don't flag
their absence there as REQUIRED NOW or SECURITY/PRIVACY.

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
