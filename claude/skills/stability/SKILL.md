---
name: stability
description: Feature stability model (Prototype → Alpha → Beta → GA), with the stability checklist of what each level includes, for deciding how much engineering a feature warrants. Use when designing, implementing, modifying, or reviewing a feature; when adding to, changing, or removing from an existing feature, since its level decides whether the change must be non-breaking; when deciding whether work such as tests, hardening, abstraction, observability, scaling, or documentation is appropriate now; or when discussing whether something is Prototype, Alpha, Beta, or GA work.
---

# Feature Stability

A stability level is an engineering investment ceiling as well as a quality
expectation. Every stability decision asks two questions with equal weight:

1. Does the work meet what the current level expects? (Under-engineering.)
2. Does the work stay within what the current level justifies? (Over-engineering.)

## Levels

| Level     | Stability  | What it is                                                                     | Reference                    |
| --------- | ---------- | ------------------------------------------------------------------------------ | ---------------------------- |
| Prototype | Exploring  | Explores an idea to choose a direction. Data may be fake; may be thrown away.  | [prototype.md](prototype.md) |
| Alpha     | Developing | Works, but still taking shape. Frequent breaking changes and bugs.             | [alpha.md](alpha.md)         |
| Beta      | Refining   | Settled enough for real use. Occasional breaking changes, announced when we can. | [beta.md](beta.md)         |
| GA        | Stable     | Keeps working as it does today. Rare breaking changes, with a deprecation window. | [ga.md](ga.md)            |

A level describes how much people can depend on the work, not how good it is.
These follow the Exalynt readme's
[Stability levels](https://readme.exalynt.com/how-it-works/stability-levels)
and the level wording in `@exalynt/design`. When they change there, change
them here.

Read the reference for the feature's level. When the question is whether some
work belongs now or later, also read the next level's. Anything that first
appears at a later level is UNNECESSARY NOW by default, unless something
concrete calls for it now: the user's request, a client's requirement, a law, a
real incident, or a security need.

## Using the checklist

[checklist.md](checklist.md) lists the outcomes each level expects, grouped by
ten qualities. Each level includes the ones below it. For any outcome,
`checklist/<quality>.md` has **How it might look** and **Overengineering**,
tagged by kind of software. Read it when judging how to meet an outcome,
whether it applies, or whether work is too much.

- **A guideline, not a gate.** Work can be called a level, GA included, with
  outcomes unmet. An outcome that applies is **met** in whatever form fits,
  **not applicable** with a few words on why, or a **known gap**.
- **Judgment wins over wording.** What an outcome takes, and whether it applies,
  depends on the software, who uses it, and the circumstances. Examples and
  overengineering are indications of boundaries, not instructions or complete
  lists. Don't fault work for meeting an outcome another way, or call work
  overengineered only because it resembles a note: say why, for this work.
  Overengineering isn't a to-do list for the next level; some is never needed.
- **The project is the first source of context.** Before asking anything, look
  at its `CLAUDE.md`, capability and engineering docs, stack rules, what it
  already has (CI, error tracking, flags, migrations, a design system), and how
  its other features, especially at the same level, met the same outcomes.
  Follow an established way rather than inventing a second or asking which to
  use. A capability can meet whole-system outcomes with what the project
  already has. But a more stable feature's pattern isn't required of a less
  stable one, a pattern that goes past what the work needs is pointed out, not
  extended, and one that breaks the security baseline is never followed.
- **Ask only what the project can't answer** and that changes what gets
  built, such as the client's agreement, a legal duty, or a target nobody has
  set. Otherwise go with what the project shows, and say what you assumed.
- **Feature flags and team practices aren't known gaps.** A flag is worth it
  when new work lands in software people rely on, or for releasing to some
  users first; in a new product, rollback is usually enough. An owner, a place
  to report problems, and review fit however the team already works, and are
  taken as met unless something says otherwise. Don't raise them as gaps.

## Known gaps

A known gap is either **to close later** or an **accepted tradeoff**, kept
because closing it isn't worth the cost. Either can exist at any level. Gaps
people would notice go in the capability doc's Known gaps. Technical ones go in
`docs/engineering/`: a Known gap with an **Until**, or a Tradeoff (see the
`engineering-docs` skill). Don't fix either kind opportunistically. A technical
gap whose **Until** names this capability at a level it has reached is due, so
REQUIRED NOW. The security and privacy baseline is never a known gap.

## Determining the level

1. **The user's word** wins. If it differs from the declaration, say so, and
   offer to update it.
2. **The declaration**: the `**Stability:**` line in
   `docs/product/capabilities/<slug>.md`, or wherever the project's `CLAUDE.md`
   says levels are recorded.
3. **Otherwise, infer it.** A stated level outweighs what the code suggests: a
   stability mark in the UI, an OpenAPI `x-stability`, `--help`, a library's
   pre-release version or `@experimental`/`@beta` tags, docs, comments, or
   `stability:` labels on the pull requests that changed its files. Then what
   the code shows: fake data, a flag or preview gate, versioned APIs, real
   users. How polished the code looks is the weakest signal. Tell the user the
   level and what it rests on, then go on; a subagent puts it in its report.
   Offer to record it.

The capability doc's opening paragraph is the purpose, Learning is the question
a Prototype answers, and from Beta, Success is what the client agreed success
looks like.

## Changing existing features

Find the feature's level before changing it. Work on a feature is held to its
level, so small changes to a GA feature are GA work.

- **Prototype, Alpha**: no compatibility guarantee. Take the quickest path,
  such as editing the migration and resetting the data.
- **Beta**: prefer non-breaking. Break only when the user directs it, and aim
  to announce it ahead of time (a goal, not a guarantee).
- **GA**: non-breaking only. A breaking change ships as a new version with its
  own level, while the current one keeps working through a deprecation window
  ([ga.md](ga.md)). Clarify the direction with the user first.

A breaking change makes existing clients, data, or deployments fail or behave
differently unless they change something: removing or renaming an endpoint,
field, or column; changing a type or meaning; tightening validation; reworking
a learned workflow; editing a migration that already ran.

A part several features share, such as a table, endpoint, or config format,
holds to the strictest of their levels. When something more stable starts
relying on early work, the early work moves up or the stable feature writes
down a known gap. A new capability in a GA feature starts at its own level,
opt-in, with its own capability doc. The stack's rules say how to make specific
changes non-breaking.

## Security and privacy

The baseline applies at every level and never relaxes:

- secrets stay out of code, version control, sample data, logs, and output,
  including what's shipped to devices or published
- real data comes only through the real systems, never copied into code,
  sample data, or fixtures
- only those allowed can reach real systems or data, and one tenant or account
  never sees or changes another's
- personal data is collected only where needed, and kept out of logs, errors,
  and services that don't need it
- untrusted input can't run as code or commands

A real system holds real data, is relied on, or costs money to use, including a
client's own; a sandbox or test environment isn't, but a test account in a live
system is. Work that holds only made-up data and reaches no real system needs
no sign-in or tenant separation, so don't flag their absence. A real credential
it uses is still a secret, and it isn't left open to anyone who finds it. Once
it holds real data or can change a real system, the baseline applies in full.
Keeping clients apart can be as simple as scoping every query, or separate
deployments. Duties that come with real personal data, such as consent, start
with the first real data, whatever the level.

Side effects reach people too: emails, messages, payments, and data written
where others read it. Before Alpha, use an allowlist, a sandbox, or a dry run.
From Alpha, real use can reach the real people it's for, once the client has
agreed and it can be turned off.

## Classifying work

- **REQUIRED NOW**: necessary to meet the current level. Resolved by meeting
  it or, with the user's agreement, recording a known gap.
- **UNNECESSARY NOW**: reasonable, but not justified at this level in these
  circumstances. Name the level where it would fit, or say it may never be
  needed. If already built and costly, prefer removing or simplifying it.
- **ACCEPTABLE**: fits the level, including gaps the level allows.
- **SECURITY/PRIVACY**: required regardless of level. Can only be met.

When reviewing several features, run one `stability-reviewer` per feature in
parallel.

## Moving between levels

- Moving up is a judgment call, not a checklist pass. Each level's reference
  says when it's ready for the next.
- Moves that change what the client relies on, such as real users at Beta or
  calling it GA, need the client's agreement. Others, the client is told.
- Moving down is fine before GA. GA rarely moves down; a breaking change
  becomes a new version.
- Dropped work is cleaned up, flags and data included.
- Never change a level as a side effect of other work, and don't suggest
  promoting unless asked. Promotion is done with `/promote`.
