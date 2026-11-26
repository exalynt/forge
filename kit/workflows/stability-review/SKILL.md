---
name: stability-review
description: Read-only review of a capability's work against its stability level (prototype, alpha, beta, ga) in a repository that has Docqet, flagging both what the level expects that's missing and premature engineering beyond it. Use to review a capability, a diff, or a promotion against a level, or, after fixes, to verify that an earlier review's findings are resolved. Can run inline or in a subagent.
---

# Stability Review

A stability level is both a quality expectation and an engineering investment
ceiling. Every review asks two questions with equal weight:

1. Does the work meet what the level expects? (Under-engineering.)
2. Does the work stay within what the level justifies? (Over-engineering.)

## Needs Docqet

What a level expects comes from Docqet's rubric, and a capability's level from
its Docqet docs (see forge's `AGENTS.md`, "Where intention lives"). Look for a
`docqet.yaml` at the working directory or above it, within the repository. If
there isn't one, stop and say the review needs Docqet: the repository can
adopt it with `docqet init` and `docqet kit`. If `docqet` isn't installed, ask
the user to install it.

## Running it

It can run inline, or in a subagent when your agent can run one, which keeps a
large review out of the main session. Whoever starts it passes the
capabilities or scope, the level to review against if it's not the current
one, the effort, and for a verification, the findings to check. Reviewing
several capabilities, run one review per capability, in parallel where your
agent can.

Don't modify files. Use only read-only commands: `docqet touches`,
`docqet show`, `docqet list`, `docqet rubric`, `git status`, `git diff`,
`git log`, `git show`, listing files, and existing read-only analysis such as
linters or type checkers. Never run commands that change the working tree,
install packages, or touch external systems.

The usual flow has two runs, with fixes made outside the review between them:

1. **First review**: what does this capability need to meet its level, and
   where are the gaps? Follow "Process" through "Report" below.
2. **Verification**: are the flagged findings resolved? Follow "Verification"
   at the end, not the full review.

## Process

1. **Find what the work touches.** Run `docqet touches` on the changed files,
   or `docqet touches --base <base>` for a branch. It lists the capabilities,
   each one's level, rules, gaps, tradeoffs, and open decisions. Read each
   capability with `docqet show capability:<id>`, and its rules with
   `docqet show`.
2. **Determine the stability level.** If the caller names a level to review
   against (for example a promotion target), use it. Otherwise use the user's
   word, then the capability's `stability` as Docqet reports it, in the
   domain the change is in. A part several capabilities share holds to the
   strictest of their levels. Say in the report which level you reviewed
   against and where it came from, including any evidence that pointed
   elsewhere.
3. **Settle the effort.** Use the one the caller gives, `low`, `medium`, or
   `high`, or `medium` without one (see "Effort").
4. **Read what the level expects.** Run `docqet rubric --level <level>`, or
   for a promotion `docqet rubric --for capability:<id> --to <level>`, which
   lists what the target expects that the current level doesn't, with the
   gaps, tradeoffs, decisions, mechanisms, and components on each outcome. At
   medium and high, also read the next level's rubric, so you can recognize
   work that belongs later, and each outcome's `overengineering` examples,
   with `docqet rubric <outcome ref>` where you need them.
5. **Read the capability's context**: its description, rules, gaps,
   tradeoffs, and open decisions. For a prototype, identify the question it's
   answering from its description and decisions, and expect realism there and
   simplicity elsewhere.
6. **Read the technical docs** for the code you're reviewing, as far as the
   effort reaches: the components and mechanisms `docqet touches` names, and
   `docqet list mechanisms`, each with `docqet show`. Note their gaps and
   tradeoffs. A `docqet:tradeoff` comment marks code that's deliberate.
7. **Look at how the project already meets outcomes**, as far as the effort
   reaches: the mechanisms and components that meet them (the rubric's
   `--for` output lists them), and how other capabilities at the same level
   meet them. That's the default to judge against. Flag departures without a
   reason, and patterns that go past the level or break the baseline.
8. **Review the relevant implementation and current changes** (`git diff`,
   `git diff --staged`, or the range the caller gives you), as far as the
   effort reaches.

"Where applicable" means it: an outcome that doesn't fit the capability,
such as accessibility for an API with no screens, counts as met. Put it under
ACCEPTABLE, with a few words on why.

## Effort

Effort sets how much you review, and so the review's time and cost. It
doesn't lower the bar: whatever you review is judged by the full level.

- **Low**: the diff only, or without one, the capability's primary workflow
  path. Check security and privacy, breaking changes, and the level's
  expectations the diff plainly touches. Flag over-engineering only where the
  diff shows it.
- **Medium**: the diff and the code it touches, such as its callers and the
  layers it crosses. Check every expectation of the level that applies to
  what changed.
- **High**: the whole capability, not just the diff: every layer, its tests,
  docs, and operations. Check every outcome the level's rubric lists, and
  look for premature work anywhere in the capability.

Some outcomes can't be seen in the repository, such as the client's agreement
or feedback, or an alert known to fire. Don't fail them for lack of evidence:
list them under NOT REVIEWED for the caller to confirm.

At every effort, check the security and privacy baseline on everything you
review.

## Look for

- what the level expects that's missing
- security and privacy violations: one tenant or client seeing or touching
  another's real data, or a real secret leaked anywhere
- breaking changes the level doesn't allow: any at GA, and any at Beta the
  user didn't direct
- implementation that exceeds the level, as the rubric's `overengineering`
  examples describe
- speculative abstractions
- premature generalization
- premature operational hardening
- unnecessary extensibility
- unnecessary scalability work
- code motivated primarily by hypothetical future requirements
- acceptance tests added, changed, or removed in the diff, including step
  definitions that change what a scenario checks and `docqet:verifies`
  annotations on their `Rule:`s, and rule docs changed. They need the user's
  explicit approval, so list each one for the caller to confirm

Over-engineering is a failed stability review, just like under-engineering.
Prefer removing unjustified complexity rather than merely documenting it.

An outcome that a gap, tradeoff, or decision already written in Docqet
affects isn't missing work: it's declared. Don't flag it as REQUIRED NOW, and
don't flag a capability's written gaps, or a component's or mechanism's, as
missing. None of them excuses the security and privacy baseline. Don't suggest
advancing to the next stability level unless explicitly asked.

## Classifying

- **REQUIRED NOW**: necessary to appropriately meet the level, and not
  already declared as a gap, tradeoff, or decision.
- **SECURITY/PRIVACY**: required regardless of level.
- **UNNECESSARY NOW**: reasonable engineering, but not justified at the
  level. Name the level where it would become appropriate. If it's already
  built and carries real cost, prefer removing or simplifying it.
- **ACCEPTABLE**: appropriate for the level, including gaps it declares and
  what the level doesn't expect yet.
- **NOT REVIEWED**: what the level expects that this effort didn't check.

## Report

```
Capability: <id> (<domain>)
Stability reviewed against: <level> (source: <Docqet | caller | user>)
Effort: <low | medium | high>
Scope: <files, diff, or commit range reviewed>

REQUIRED NOW
- <finding> (<file:line>): which outcome of this level it fails

SECURITY/PRIVACY
- <finding> (<file:line>): the risk

UNNECESSARY NOW
- <finding> (<file:line>): why it exceeds this level, where it would belong,
  and the recommended action (remove, simplify, or leave with a reason)

ACCEPTABLE
- brief notes on things appropriately scoped, and the gaps, tradeoffs, and
  decisions already written that cover what's unmet

NOT REVIEWED
- <outcome of this level, or part of the capability>: not reviewed at
  <effort> effort

Verdict: PASS | FAIL (<effort> effort)
```

Omit empty sections. The verdict is FAIL if there is any REQUIRED NOW or
SECURITY/PRIVACY finding, or any UNNECESSARY NOW finding you recommend removing
or simplifying. A REQUIRED NOW finding can be resolved by recording it as a
gap the Docqet kit's way, with a person's agreement, instead of meeting it, so
say for each whether that would be reasonable. A SECURITY/PRIVACY finding
can't. Below high, a PASS covers only what was reviewed, and NOT REVIEWED
lists what the level expects that you didn't check. Be specific and
cite files. Don't pad the report.

## Verification

When asked to verify fixes, you're given the level and the findings to check.
Check only those findings, at their cited locations and wherever the fix moved
them. Don't review from scratch, raise new findings, or apply an effort level.
If no findings are given, say so instead of starting a full review.

```
Capability: <id> (<domain>)
Stability reviewed against: <level> (source: first review)
Scope: verification of the first review's findings

RESOLVED
- <finding> (<file:line>): what resolved it

NOT RESOLVED
- <finding> (<file:line>): what is still wrong

Verdict: PASS | FAIL (verification)
```

Omit empty sections. The verdict is FAIL if any finding is NOT RESOLVED.
