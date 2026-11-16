---
name: stability
description: Feature stability model (Prototype → Alpha → Beta → GA), with the stability checklist of what each level includes, for deciding how much engineering a feature warrants. Use when designing, implementing, modifying, or reviewing a feature; when adding to, changing, or removing from an existing feature, since its level decides whether the change must be non-breaking; when deciding whether work such as tests, hardening, abstraction, observability, scaling, or documentation is appropriate now; or when discussing whether something is Prototype, Alpha, Beta, or GA work.
---

# Feature Stability

> The goal is NOT to make every feature as robust as possible.

> A stability level is an engineering investment ceiling as well as a quality expectation.

Every stability decision asks two questions with equal weight:

1. Does the work meet what the current level expects? (Under-engineering.)
2. Does the work stay within what the current level justifies? (Over-engineering.)

## Levels

| Level     | Stability  | What it is                                                                           | Reference                    |
| --------- | ---------- | ------------------------------------------------------------------------------------ | ---------------------------- |
| Prototype | Exploring  | Explores an idea to choose a direction. Data may be fake; may be thrown away.        | [prototype.md](prototype.md) |
| Alpha     | Developing | Works, but still taking shape. Frequent breaking changes and bugs.                   | [alpha.md](alpha.md)         |
| Beta      | Refining   | Settled enough for real use. Occasional breaking changes, announced when we can.     | [beta.md](beta.md)           |
| GA        | Stable     | Keeps working as it does today. Rare breaking changes, with a deprecation window.   | [ga.md](ga.md)               |

A level describes how much people can depend on the work, not how good it is.
A prototype that answers its question quickly is excellent work.

These definitions follow the Exalynt readme's
[Stability levels](https://readme.exalynt.com/how-it-works/stability-levels)
and the level wording in `@exalynt/design`. When they change there, change
them here.

Read the reference for the feature's current level before deciding. When the
question is whether some work belongs now or later, also read the next level's
reference. Anything that first appears at a later level is UNNECESSARY NOW by
default.

## The checklist

[checklist.md](checklist.md) lists what each level generally includes, area by
area, from the readme's Stability checklist. Each level includes the items of
the levels below it. Use it to judge whether work meets its level and what
moving up would add.

It's a set of guidelines, not a gate. Aim for each item's intent in the form
that fits what's being built, and skip what doesn't apply: a feature that keeps
no data needs no backups, and a command-line tool's logs are a verbose option.
"Reading it" at the top of the checklist says how.

## Known gaps

When an item applies and isn't met, it's either fixed or written down as a
known gap, of one of two kinds. Either can exist at any level, GA included:

- **To close later**: written down now, and fixed in a later iteration.
- **Accepted tradeoff**: kept on purpose, because closing it isn't worth what
  it would cost, such as no offline mode.

Gaps people would notice go in the capability doc. Technical ones go in
`docs/engineering/`, where a gap to close later is a Known gap with an
**Until**, and an accepted tradeoff is a Tradeoff (see the `engineering-docs`
skill). The security and privacy baseline is never a known gap.

## Determining the level

1. **The user's statement** in the current conversation. The user's word
   wins. If it differs from the project's declaration, say so, and offer to
   update the declaration.
2. **The project's declaration.** The `**Stability:**` line in the feature's
   capability doc, `docs/product/capabilities/<slug>.md` (see the
   `docs` skill), or wherever else the project's `CLAUDE.md` says it's
   recorded.
3. **Otherwise, infer it from the repository.** Look at what its code and
   docs say about the feature:
   - **Where a level is stated**: a stability mark or level in the UI code,
     the project's docs, README, or `CLAUDE.md`, code comments, and the
     `stability:` labels on pull requests that changed its files (find them
     with `git log` on those files, then `gh pr view <number> --json labels`).
   - **What the code shows about its use**: mock or fake data, a feature flag
     or preview gate, versioned public APIs, and whether real users or
     clients depend on it.

   A stated level outweighs what the code suggests. How rough or polished the
   code looks is the weakest signal, and never decides the level alone.

   Tell the user the level you inferred and what it rests on, including any
   evidence that pointed elsewhere, then go on with it. The user can correct
   it, and their word then wins. A subagent puts this in its report. Offer to
   record the level in its capability doc.

The capability doc also says what the feature is for right now. Its opening
paragraph is the purpose, and Learning is the question a Prototype is
answering, which is essential for prototypes. From Beta, Success is what the
client agreed success looks like. Known gaps lists what is intentionally
incomplete that people would notice, and which kind each is. Technical gaps and tradeoffs
are in `docs/engineering/` (see the `engineering-docs` skill). Don't "fix" either
kind of gap opportunistically. A technical gap whose **Until** names this
capability at a level it has reached is due, which makes it REQUIRED NOW.

## Changing existing features

Adding to, adjusting, or removing from an existing feature is governed by that
feature's level. Find the level before changing anything. When the change
touches something several features share, such as a table or an endpoint, the
strictest of their levels applies.

| Level             | Changes                                                                      |
| ----------------- | ---------------------------------------------------------------------------- |
| Prototype, Alpha  | No compatibility guarantee. Take the quickest path, breaking or not.         |
| Beta              | Prefer non-breaking. Make a breaking change only when the user directs it, and aim to announce it ahead of time (a goal, not a guarantee). |
| GA                | Non-breaking only. A breaking change ships as a new version with its own level, while the current one keeps working through a deprecation window ([ga.md](ga.md)). Clarify the direction with the user before making one. |

A breaking change is anything that makes existing clients, data, or
deployments fail or behave differently unless they change something on their
side. Examples: removing or renaming an endpoint, field, or column; changing a
type or meaning; tightening validation; changing what a report contains;
reworking a workflow people have learned; editing a database migration that has
already run instead of adding a new one.

Work on an existing feature is held to its level: small changes to a GA feature
are GA work. A new capability in a GA feature starts at its own level, opt-in,
with its own capability doc.

At Prototype and Alpha, the quickest path is often the breaking one: edit the
existing migration and reset the database, or change the endpoint in place.
The security and privacy baseline still applies.

The stack's rules say how to make specific kinds of changes non-breaking.

## Security and privacy

The security and privacy baseline applies at every level and never relaxes.
It's the Prototype items of the checklist's "Security, privacy, and
compliance":

- secrets kept out of code, version control, sample data, and demos
- secrets and personal data kept out of logs, error messages, and output,
  client bundles and third-party services included
- real client data comes only through production systems, never copied into
  code, sample data, or fixtures
- access control enforced on anything touching client systems or data
- people can only see and change data they're allowed to, and one client or
  customer can never see or change another's
- personal data collected only where needed
- untrusted input can't run as code or commands (such as SQL injection or
  cross-site scripting)

Sample, mock, or fake data and throwaway placeholder credentials are not
covered by the baseline. So a Prototype or Alpha that holds only sample data
needs no authentication or authorization, and don't flag their absence there.
Once it touches real client systems or data, at any level, access control is
required. Keeping clients apart can be as simple as separate deployments, or
having only one client's real data present.

Beta and GA add more on top, such as exporting and deleting personal data at
Beta, and a security review at GA.

## Classifying work

When reviewing or planning, classify each item:

- **REQUIRED NOW**: necessary to appropriately meet the current level.
- **UNNECESSARY NOW**: reasonable engineering, but not justified at the current
  level. Name the level where it would become appropriate. If it is already
  built and carries real cost, prefer removing or simplifying it.
- **ACCEPTABLE**: appropriate for the current level, including known gaps the
  level allows.
- **SECURITY/PRIVACY**: required regardless of level.

A REQUIRED NOW item is resolved by meeting it or, with the user's agreement,
by recording it as a known gap. A SECURITY/PRIVACY item can only be met.

When reviewing several features, run one `stability-reviewer` per feature in
parallel (see the subagents rule).

## Moving between levels

- **Moving up is a judgment call**, not a checklist pass: the direction has
  firmed up and the work fits the next level well enough, with its gaps named.
  It can take a day or months. Each level's reference says when it's ready for
  the next.
- **Moves that change what the client relies on**, such as putting it in front
  of real users at Beta or calling it GA, need the client's agreement first.
  Others, the client is just told.
- **Moving down is fine before GA**, when the direction changes a lot or the
  work no longer fits its level. GA rarely moves down: a breaking change
  becomes a new version instead.
- **Dropped work is cleaned up**, flags and data included.

## Rules

- Never change a feature's level as a side effect of other work. Promotion is
  done deliberately with `/promote`.
- Work above the current level needs explicit, concrete justification.
- Don't suggest promoting a feature unless asked.
