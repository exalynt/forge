---
name: qa
description: QA a feature end to end against its business rules and stability level, report the issues found, clarify ambiguous rules and assumptions, then confirm which issues to fix before fixing them.
argument-hint: <feature> [focus]
disable-model-invocation: true
---

# QA a Feature

Arguments: `$ARGUMENTS`. Expected format: `<feature> [focus]`, for example
`/qa conversations` or `/qa conversations reconnecting by email`. If no feature
is given, ask.

QA checks that the feature does what it's supposed to, through the paths
people actually take. It judges behavior. Read code to understand or explain
behavior, not to review it. For code review, use `/code-review`.

There are two phases. Steps 1–5 find and report issues and change nothing in
the project. Steps 6–7 fix only what the user confirms.

## 1. Understand the feature

- Find its stability level and declaration as the `stability` skill describes.
  Read `purpose`, `learning`, and `known_gaps` if the declaration has them.
- Find where it lives: pages, routes, endpoints, services, tables, jobs,
  external integrations, and the tests that cover it.
- Find what it's supposed to do. Sources, most authoritative first: the user
  in this conversation, the feature declaration and project docs, specs or
  issues the project links to, existing tests, then the code. Code shows what
  the feature does, not what it should do. Don't treat current behavior as the
  rule just because it's implemented.
- Work out what end to end means here: the primary workflow from a user's
  first action to the result they want, across every layer it crosses (UI,
  API, storage, background jobs, external services).

With a `focus`, go deep on that part but still run the primary workflow once.

For a feature spanning several layers or repositories, research them in
parallel subagents (see the subagents rule), one per layer or source, each
reporting what it found and where.

## 2. Write down the business rules

List the rules the feature should follow, each with its source:

```
R1. <rule> (stated | assumed; source: <user | declaration | doc | test | code path>)
```

Cover who can do what, valid and invalid inputs, states and transitions,
limits, calculations, what the user sees and when, and what happens on
failure. A rule is **stated** when a source says so, and **assumed** when you
inferred it from the code or common sense.

Then list the ambiguities: where sources disagree, where they're silent on a
case that matters, or where the code does something a reasonable person could
read either way. Phrase each as a question with the options you see, which one
the code does now, and which one you'd assume.

## 3. Clarify before testing

Show the rules, then ask about the ambiguities whose answer decides whether
something is a bug. Batch them into one round (AskUserQuestion, up to four per
call), with your recommended option first. Don't ask what you can find out
from the code, the docs, or by trying it.

For ambiguities that don't block testing, or that the user would rather leave
for now, state your assumption, carry on, and list them in the report. Rules
the user answers become **confirmed**.

## 4. Plan and run the checks

Plan checks against the rules, scaled to the stability level (next section).
Typically:

- **Primary workflow**: the happy path end to end, with realistic data.
- **Rules**: each rule exercised at least once, including its boundaries
  (zero, one, the maximum, just over, empty, missing, duplicate).
- **Errors and recovery**: invalid input, failing dependencies, double
  submits, going back, reloading mid-flow.
- **States**: first use, empty, many items, long text, concurrent edits where
  they can happen.
- **Access**: each role, signed out, and another tenant's data.
- **Presentation**, for UI: light and dark, phone width, keyboard, and
  whatever the project's rules ask for.

Run them against the real thing as far as the project allows: the test
suites, lint, type checks, and build; the UI through whichever browser tools
are available; the API; the database. Follow the project's instructions on
what may be run. If it says not to start a dev server, don't, and report what
that left unverified.

- Don't change project files in this phase. Throwaway scripts, test data, and
  scratch tests go in the scratchpad.
- Use local or test environments and test data. Never touch production or
  real customer data, and never send real emails, messages, or payments. Ask
  before anything destructive or outward-facing in a shared environment.
- Reproduce each issue before reporting it, and narrow it to the fewest steps
  that show it.
- When you find an issue, keep going. One bug doesn't mean the rest works.
- Run checks that share nothing in parallel subagents, such as tracing each
  rule through the code or test suites with their own data. Checks that share
  a browser session, a server, a database, or test data run one at a time in
  the main conversation.

## Scale to the stability level

Judge against the expectations in the `stability` skill's reference for the
feature's level:

- **Prototype**: does it answer its learning question? Rough edges outside
  that question aren't issues.
- **Alpha**: the primary workflow must work for real. Edge-case polish can
  wait.
- **Beta** and **GA**: the full set of checks, including breaking changes.

Something only a later level expects isn't an issue. Neither is anything in
`known_gaps`. List it under "Not issues at <level>" only when someone could
reasonably expect it. Tenant isolation and leaked secrets are checked at every
level, and a problem there is always Critical.

## 5. Report

```
Feature: <name>
Stability: <level> (source: <declaration | user | none declared, so GA>)
Scope: <what was QA'd, environment, commit>

Business rules
R1. <rule> (stated | assumed | confirmed; source)

Issues
1. [<severity>] <title>
   Rule: R<n>
   Expected: ...
   Actual: ...
   Steps: 1. ... 2. ...
   Evidence: <output, screenshot path, file:line of the cause if known>
   Suggested fix: <one line, sized to the level>

Open questions
Q1. <ambiguity>. Options: A (current behavior), B. Assumed: A.

Not issues at <level>
- <item>: the level it belongs to

Verified
- <check>: passed

Not verified
- <check>: why (couldn't run, no access, out of scope)
```

Severities:

- **Critical**: security or privacy, data loss or corruption, or the primary
  workflow is blocked.
- **Major**: a rule is broken or a common path fails, with no reasonable
  workaround.
- **Minor**: wrong, but there's a workaround or it's on an uncommon path.
- **Polish**: presentation, wording, consistency.

Order issues by severity. Omit empty sections. Don't pad the report with
things that work or with style preferences presented as issues.

## 6. Confirm what to fix

Ask which issues to fix (AskUserQuestion, multiSelect; group them by severity
if there are many), recommending Critical and Major. Ask any open questions
whose answer changes an issue. An issue that hinges on an unanswered question
can't be fixed yet.

If the user picks nothing, stop.

## 7. Fix and re-verify

- Fix only the confirmed issues, sized to the stability level. Follow the
  `stability` skill on breaking changes.
- If a fix needs a decision that's the user's to make, ask.
- Fixes may run in parallel subagents only when they touch disjoint files.
  Re-verify in the main conversation once they're all in.
- Where the level expects tests, add one that pins the rule the issue broke.
- Re-run each fixed issue's steps and the checks around it to confirm the fix
  and catch regressions. Run the project's checks.
- Anything new found while fixing goes in the report. Don't fix it without
  asking.
- Offer to record confirmed rules where the project keeps them: the feature
  declaration, its docs, or a test. Don't invent a new place without asking.
- Don't commit unless asked.

Report each confirmed issue as fixed (and how the fix was verified) or not
fixed (and why), then the checks you ran with their results, and any new
issues found.
