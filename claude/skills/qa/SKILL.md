---
name: qa
description: QA a feature end to end against its business rules and stability level, verifying primarily through its acceptance tests and exploring by hand what they don't cover, report the issues found, clarify ambiguous rules and assumptions, then confirm which issues to fix and which business rules to record as acceptance tests before doing either.
argument-hint: "[low|medium|high] <feature> [focus]"
disable-model-invocation: true
---

# QA a Feature

Arguments: `$ARGUMENTS`. Expected format: `[effort] <feature> [focus]`, for
example `/qa conversations`, `/qa low conversations`, or
`/qa high conversations reconnecting by email`. The effort is `low`, `medium`,
or `high`, and sets how much time and cost the run spends (see "Effort").
Without one, use `medium`. If no feature is given, ask.

QA checks that the feature does what it's supposed to, through the paths
people actually take. It judges behavior. Read code to understand or explain
behavior, not to review it. For code review, use `/code-review`.

The feature's acceptance tests are its specification (see the testing rule),
and QA's primary form of verification. QA runs them rather than repeating by
hand what they cover, uses their scenarios to know the rules and what to click
through, and explores by hand the happy path and what they leave out. Business
rules QA clarifies or finds uncovered are recorded as acceptance tests, with
the user's approval.

A library has no acceptance tests unless the user asked for them (see the
testing rule's "Libraries"). There, run its end-to-end tests in their place,
take its rules from its docs and tests, and walk its public API by hand the
way a consumer would. Record rules QA settles in the end-to-end tests for a
core flow and in unit or integration tests for an edge case, at the levels
that require them.

There are two phases. Steps 1–5 find and report issues, propose acceptance
tests, and change nothing in the project. Steps 6–7 fix only the issues the
user confirms and write only the acceptance tests the user approves.

## 1. Understand the feature

- Find its stability level and declaration as the `stability` skill describes.
  Read `purpose`, `learning`, and `known_gaps` if the declaration has them,
  and note any mock or stub it explicitly allows.
- Find where it lives: pages, routes, endpoints, services, tables, jobs,
  external integrations, and the tests that cover it.
- Find its acceptance tests: the `.feature` files for it, the level each
  scenario runs at, and the command that runs them, from the stack's testing
  rule or the project's CI. Read the scenarios. They state the rules, and
  their steps show the flows to click through and the data to set up.
- Find what it's supposed to do. Sources, most authoritative first: the user
  in this conversation, the feature's acceptance tests (its approved
  specification), the feature declaration and project docs, specs or issues
  the project links to, other existing tests, then the code. Code shows what
  the feature does, not what it should do. Don't treat current behavior as the
  rule just because it's implemented.
- Work out what end to end means here: the primary workflow from a user's
  first action to the result they want, across every layer it crosses (UI,
  API, storage, background jobs, external services). When the feature has a
  UI, end to end starts there: the user's path is through its screens, not
  its endpoints.

A `focus` narrows where the effort goes, not how much: the focus gets the
effort's full checks, and the rest of the feature gets the primary workflow
once.

For a feature spanning several layers or repositories, research them in
parallel subagents (see the subagents rule), one per layer or source, each
reporting what it found and where.

## 2. Write down the business rules

List the rules the feature should follow, each with its source:

```
R1. <rule> (stated | assumed; source: <user | acceptance test | declaration | doc | test | code path>; acceptance test: <file › Rule | none>; core | edge case)
```

Cover who can do what, valid and invalid inputs, states and transitions,
limits, calculations, what the user sees and when, and what happens on
failure. A rule is **stated** when a source says so, and **assumed** when you
inferred it from the code or common sense. Note whether each rule has an
acceptance test, and whether it's a core rule or an edge case (see the testing
rule), since that decides whether the level requires one.

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
the user answers become **confirmed**. An answer that contradicts an
acceptance test changes the specification, so it needs that test updated
(step 5), and it's not an issue in the code until then.

## 4. Plan and run the checks

Plan checks against the rules, cut to the effort and judged by the stability
level (the sections after this one). Start with the acceptance tests:

- **Run the acceptance suite** for the feature first. A rule whose acceptance
  tests pass is verified. Don't check it again by hand.
- A failing acceptance test is an issue in the code, rated like any other
  broken rule, unless the user has said the specification is what's wrong.
  Don't change the test.

Then check by hand what the suite doesn't. At high effort:

- **Primary workflow**: the happy path end to end, with realistic data,
  following the core rules' scenarios through the UI where there is one.
  Passing acceptance tests don't replace this walk: it's what shows the
  pieces work together for a user.
- **Rules**: each rule without an acceptance test exercised at least once.
- **Edge cases**: those the acceptance tests don't cover, checked explicitly,
  not just when they come up (below).
- **Real data**: nothing mocked or stubbed without an explicit exception
  (below).
- **Access**: each role, signed out, and another tenant's data.
- **Presentation**, for UI: light and dark, on a phone (below), keyboard, and
  whatever the project's rules ask for.

Run them against the real thing as far as the project allows: the test
suites, lint, type checks, and build; the UI; the API; the database. Follow the
project's instructions on what may be run. If it says not to start a dev
server, don't, and report what that left unverified.

### Edge cases

For each rule and each step of the workflow, list the edge cases that apply
and aren't already covered by an acceptance test, then check every one. Write
the list down before running it. It goes in the report. Work through:

- **Boundaries**: zero, one, the maximum, just under and just over each limit,
  negative numbers, the first and last item, the ends of a date range.
- **Input**: empty, whitespace only, missing, very long, special characters
  and emoji, non-Latin text, leading and trailing spaces, different case, the
  wrong type or format, pasted content, HTML or script in text fields.
- **Duplicates and conflicts**: creating the same thing twice, double clicks
  and double submits, two users or tabs editing the same item, acting on
  something that was just deleted or changed elsewhere.
- **State**: first use with no data, one item, many items (pagination, long
  lists), an item in every status the feature has, and transitions that
  shouldn't be allowed, such as skipping a step or acting after completion.
- **Time**: time zones, daylight saving changes, midnight and month and year
  ends, past and far-future dates, expiry exactly at the limit, a session
  timing out mid-flow.
- **Interruption and failure**: reloading, going back, or leaving and
  returning mid-flow; losing the network; a slow or failing dependency; a
  multi-step action failing partway; retrying after a failure.
- **Numbers and money**, where the feature has them: rounding, precision,
  currency, very large values, division by zero.

Skip categories that don't apply, and say so in one line. Dig deeper around
rules marked assumed and around the open questions, where the rules are
least certain.

### Real data, no mocks

The feature must run on its real data and integrations. Check that nothing in
its runtime path is mocked, stubbed, faked, or hard-coded:

- **In the code**: search the feature's non-test code for mocks, stubs,
  fakes, fixtures, sample or placeholder data; hard-coded users, IDs, or
  responses; functions that return a fixed value or resolve literal data;
  simulated delays; "not implemented" and TODO paths; and flags or environment
  switches that turn mocks on, such as mock service workers or dev-only
  branches. Check what the build actually ships, not just the defaults.
- **At runtime**: confirm the UI's network requests reach the real API.
  Create or change something through the UI and confirm it's stored and
  survives a reload. Change something in storage and confirm the UI shows it.
  Confirm external services are actually called, in their test or sandbox
  mode.

Test doubles inside tests are fine, and so is the test data QA sets up. This
check is about the feature's own code paths.

A mock or stub is allowed only when an exception names it explicitly: the
user says so in this conversation, or the feature declaration lists it, for
example in `known_gaps`. A Prototype's general allowance for fake data isn't
an exception by itself. List each allowed exception under Verified with its
source. Any other mock or stub is an issue: Major, or Critical when it makes
the primary workflow look like it works when it doesn't, such as a save that
stores nothing or a send that sends nothing.

### Through the UI

When the feature has a UI, run the primary workflow and the rules through it,
as a user would, with whichever browser tools are available. Use the API and
database to set up data and to confirm what the UI did, not instead of it. An
API that behaves correctly doesn't mean the screen does. In the UI, check:

- **What it shows**: values, labels, formatting, counts, dates and times, and
  that they match what's stored, including after a reload.
- **How it responds**: loading, empty, and error states; validation messages
  and where they appear; disabled and hidden actions; confirmation and success
  feedback.
- **Where it goes**: navigation, links, back and forward, deep links, and what
  a signed-out user or a user without access sees.
- **How it looks**: the presentation checks above, plus text that overflows,
  truncates, or overlaps.
- **On a phone**: run the primary workflow again at phone width (360px and
  375px), not just a glance at each screen. Check the page never scrolls
  sideways, nothing is cut off or overlapping, tables and grids still read,
  dialogs and menus fit, navigation is reachable, touch targets stay usable,
  and fixed headers, footers, and action bars don't cover content. Check
  desktop width too, and a tablet width when the layout changes there. The
  project's responsive rule, such as the React stack's, says what's expected;
  where it applies at every level, so does this check. Set each width with the
  browser's resize tool, desktop at 1280px or wider. The browser window is
  sized to fit the user's screen, so its own width isn't a test width.

A rule the UI enforces must also hold at the API. Hiding a button doesn't
control access, so call the API directly for each access rule. Capture a
screenshot as evidence for each UI issue, and note the width it was taken at.
Save it in the QA folder (below).

If no browser tools are available or the app can't be run, say so. Don't
report the UI as verified because the API passed. List it under "Not
verified".

### Rules for this phase

- Don't change project files in this phase, except to git-ignore the QA
  folder.
- The QA folder is `.playwright-mcp/` at the project root. Everything you
  create goes there: screenshots, snapshots, console logs, throwaway scripts,
  test data, and scratch tests. Playwright MCP writes its own files there, and
  it refuses paths outside the project or that folder, so a scratchpad path
  fails. Include the folder in every file name you pass, such as
  `.playwright-mcp/qa-<what>-375.png`. A bare name lands in the project root.
- Before creating anything, check that the folder is ignored with
  `git check-ignore -q .playwright-mcp/`. If it isn't, add `/.playwright-mcp/`
  to the root `.gitignore`, and say so in the report. When you're done,
  `git status` should show nothing new besides that line. Move anything else
  you created into the folder, or delete it.
- Use local or test environments and test data. Never touch production or
  real customer data, and never send real emails, messages, or payments. Ask
  before anything destructive or outward-facing in a shared environment.
- When you start a dev server, use a port that's neither the framework's
  default (Vite's 5173) nor one the project sets anywhere: in its config, dev
  scripts, `.env` files, or Compose files, including the ports of other apps
  it runs beside. The user's own instances run there. Pick a free port outside
  all of them, pass it explicitly
  (`npm run dev -- --port <port> --strictPort`), point the browser at it, and
  stop the server when you're done.
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
- **Alpha**: the primary workflow must work for real. An edge case that breaks
  it, or loses or corrupts data, is an issue. Other edge-case polish can wait.
- **Beta** and **GA**: the full set of checks, including breaking changes.

Edge cases are checked at every level. The level decides which failures count
as issues, not whether to look; the effort decides how many are looked at.
Mocks and stubs need an explicit exception at every level.

Something only a later level expects isn't an issue. Neither is anything in
`known_gaps`. List it under "Not issues at <level>" only when someone could
reasonably expect it. Tenant isolation and leaked secrets are checked at every
level, and a problem there is always Critical.

## Effort

Effort controls how much of the feature a run checks, and so its time and
token cost. It's independent of the stability level. The level sets the bar
each check is judged by. Effort decides how many checks run. A low-effort run
on a GA feature judges what it checks by GA's bar, but it doesn't check
everything GA expects, so it can't say the feature meets GA.

The steps in this skill describe **high**. **Medium** and **low** do less:

| Part                       | Low                                                                                   | Medium                                                                                                                                     |
| -------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Understand (1)             | The declaration and the primary workflow's path through the code.                     | Every layer the feature crosses.                                                                                                           |
| Rules and questions (2, 3) | The rules the primary workflow touches. Ask only what blocks running it.              | All the rules. Ask what decides whether something is a bug.                                                                                |
| Workflows                  | The primary workflow, once.                                                           | The primary workflow and the important secondary ones.                                                                                     |
| Rules exercised by hand    | Those the primary workflow hits that no acceptance test covers.                       | Each rule without an acceptance test, once.                                                                                                |
| Edge cases by hand         | Those the primary workflow runs into. No separate list.                               | Boundaries, input, and state for each rule, where no acceptance test covers them, and any case around an assumed rule or an open question. |
| Acceptance tests proposed  | For rules clarified or changed in this run.                                           | Also for every rule the level requires one for that has none.                                                                              |
| Real data                  | Search the primary workflow's path for mocks, and confirm its save survives a reload. | Search the feature's code, and run the runtime checks on the primary workflow.                                                             |
| Presentation               | Desktop and 375px, light mode, on the primary workflow's screens.                     | The primary workflow at desktop, 360px, and 375px, in light and dark.                                                                      |
| Subagents                  | None.                                                                                 | Only for separate repositories, or test suites that run long.                                                                              |

At every effort:

- Run the feature's acceptance tests.
- Run the primary workflow end to end, through the UI where there is one.
- Check tenant isolation: try the primary workflow's data from another tenant,
  at the API. The security and privacy baseline isn't a matter of effort.
- Run the project's checks: tests, lint, type checks, and build.
- Reproduce each issue, with a screenshot for each UI issue, and report
  anything you run into outside the plan.

Report what the run skipped that the stability level expects under "Not
verified", as skipped at this effort, so the report never reads as more than
it checked.

## 5. Report

```
Feature: <name>
Stability: <level> (source: <declaration | user | inferred from <what>>)
Effort: <low | medium | high>
Scope: <what was QA'd, environment, commit; through UI | API | code only>

Business rules
R1. <rule> (stated | assumed | confirmed; source; acceptance test: <file › Rule | none>)

Acceptance tests
- Ran: <command>: <passed / failed counts>
- Failing: <scenario>: Rule R<n>, listed as issue <n>

Issues
1. [<severity>] <title>
   Rule: R<n>
   Expected: ...
   Actual: ...
   Steps: 1. ... 2. ... (as a user's actions in the UI, where there is one)
   Evidence: <output, screenshot path, file:line of the cause if known>
   Suggested fix: <one line, sized to the level>

Open questions
Q1. <ambiguity>. Options: A (current behavior), B. Assumed: A.

Proposed acceptance tests
A1. <add | update | remove> R<n>, <file>: <why: clarified | changed | uncovered>
    <the Gherkin: the Rule and its Examples, or the diff to an existing one>

Not issues at <level>
- <item>: the level it belongs to

Verified
- <check>: passed
- Edge cases: <each one checked that passed, grouped by rule or step>
- Real data: <what was confirmed real; each allowed mock or stub and its
  exception's source>

Not verified
- <check>: why (couldn't run, no access, out of scope, skipped at <effort>
  effort)
```

Severities:

- **Critical**: security or privacy, data loss or corruption, or the primary
  workflow is blocked.
- **Major**: a rule is broken or a common path fails, with no reasonable
  workaround.
- **Minor**: wrong, but there's a workaround or it's on an uncommon path.
- **Polish**: presentation, wording, consistency.

Propose acceptance tests for rules the user clarified or changed in this run,
and for rules the level requires one for that have none: core rules from
Alpha, edge-case rules from Beta. Prototype has no automated tests, so propose
none there. Bug fixes, performance, and presentation issues don't get
acceptance tests of their own: fixing them restores behavior the
specification already describes. Write the Gherkin as the testing rule
describes, in the domain's language.

Order issues by severity. Omit empty sections. Keep things that work and
style preferences out of Issues. What passed belongs under Verified.

## 6. Confirm what to fix

Ask which issues to fix (AskUserQuestion, multiSelect; group them by severity
if there are many), recommending Critical and Major. Ask any open questions
whose answer changes an issue. An issue that hinges on an unanswered question
can't be fixed yet.

Ask separately which proposed acceptance tests to write, with the Gherkin in
front of the user. That answer is the explicit approval the testing rule
requires, for those tests as shown. One the user wants reworded gets shown
again before it's written.

For a mock or stub, the user can make an explicit exception instead of a fix.
Offer to record it in the feature declaration, such as in `known_gaps`, so the
next QA run knows.

If the user picks no issues and no acceptance tests, stop.

## 7. Fix, record the rules, and re-verify

- Write the approved acceptance tests, as shown, in the place the stack's
  testing rule names. Each should pass once the confirmed fixes are in. One
  that still fails is a new issue: report it, and don't fix it without
  asking.
- Fix only the confirmed issues, sized to the stability level. Follow the
  `stability` skill on breaking changes.
- If a fix needs a decision that's the user's to make, ask.
- Fixes may run in parallel subagents only when they touch disjoint files.
  Re-verify in the main conversation once they're all in.
- Re-run each fixed issue's steps through the same path it was found on, in
  the UI and at the width where it showed, and the checks around it, to
  confirm the fix and catch regressions. Run the acceptance tests and the
  project's checks.
- Anything new found while fixing goes in the report. Don't fix it without
  asking.
- Don't commit unless asked.

Report each confirmed issue as fixed (and how the fix was verified) or not
fixed (and why), each approved acceptance test as written and passing or
failing, then the checks you ran with their results, and any new issues
found.
