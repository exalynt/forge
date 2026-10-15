---
name: ux
description: Try a feature's main flow, through its UI or API, as a specific type of user, report what it's like for them to use (where they hesitate, get stuck, or do more work than they should), then confirm which findings to act on before changing anything.
argument-hint: <feature> [as <type of user>]
disable-model-invocation: true
---

# UX Walkthrough

Arguments: `$ARGUMENTS`. Expected format: `<feature> [as <type of user>]`, for
example `/ux conversations as a support agent on a phone` or
`/ux billing-api as a partner developer integrating for the first time`. If no
feature is given, ask. Anything after the user type narrows the walkthrough to
part of the flow.

A walkthrough asks how it is for this user to get the main flow done. It
judges experience, not correctness. A flow can be bug-free and still slow,
confusing, or easy to get wrong. Bugs found along the way get reported, but
looking for them is `/qa`'s job. Reviewing code is `/code-review`'s.

There are two phases. Steps 1–5 try the feature and report, and change
nothing in the project. Steps 6–7 change only what the user confirms.

## 1. Understand the feature

- Find its stability level and declaration as the `stability` skill describes.
  Read `purpose`, `learning`, and `known_gaps` if the declaration has them.
- Find the main flow: what the feature is for, and the path from a user's
  first action to the result they came for. Use the user in this
  conversation, the declaration, and the project's docs first. Read code only
  to find where the feature lives and how to reach it, not the steps to take.
  Step 3 works out those steps the way the user would.
- Find what the user would actually have: the screens and navigation for a
  UI; for an API, the public docs, specs, examples, and SDKs an integrator
  would be given.

## 2. Define the user and their goal

Write down who you're walking through as, before trying anything:

```
User: <type of user>
Goal: <what they came to get done, in their words>
Knows: <domain knowledge, familiarity with this app, technical skill>
Context: <device, how often they do this, time pressure, interruptions>
Expects: <what they think the first step is, the words they'd look for,
          what "done" looks like to them>
Success: <the observable end state that means they got it done>
```

Take the user type from the arguments. If none is given, ask, and offer the
user types the project's docs or the feature declaration describe. Fill in the
rest from the project and common sense, and mark what you assumed. Ask
(AskUserQuestion, one round) only about assumptions that would change the
walkthrough, such as whether they're new or experienced, or on a phone or a
desktop, with your recommended option first.

For an API, the user is the developer integrating with it. Their goal is the
outcome their own product needs, not "call the endpoint".

## 3. Walk the main flow as that user

Do the task the way this user would, with only what they'd have. No reading
the code, the database, or internal docs to find the next step. When the user
would be stuck, note it as a finding, then use what you know to get past it
and keep going.

Keep a step log as you go:

```
S1. Trying to: <intent>
    Sees: <what's on the screen or in the response>
    Expects: <what they think will happen>
    Does: <the action>
    Result: <what happened>
    Friction: <hesitation, a guess, a wrong turn, a wait, a re-read, or none>
```

**Through the UI**: use whichever browser tools are available, with realistic
data for this user, such as their volume of items, their names, and their
statuses. Run it at their device's width: 360px or 375px for a phone user,
desktop otherwise. Take screenshots at the moments that matter and note the
width.

**Through an API**: start from the public docs, as the integrator would.
Write the client a developer would write, with curl or a short script in the
scratchpad. Work through getting credentials, making the first successful
call, and finishing the whole outcome. Count the calls and the round trips the
outcome takes. Make the mistakes a newcomer would make, such as a missing
field, the wrong type, or an expired token, and read the error messages as
they'd read them.

Then walk the variations this user would hit: their first time and a repeat
time, a mistake they'd likely make and recovering from it, being interrupted
and coming back, and the empty and full states they'd see. Skip the ones that
don't apply.

Look at each step through these lenses:

- **Finding the way**: can they find where to start and what to do next
  without guessing? Does each step lead to the next?
- **Language**: do labels, messages, field and endpoint names use this user's
  words, or the system's internal ones?
- **Effort**: steps, clicks, typing, fields, waits, calls, and switching
  between screens or tools. What could be skipped, defaulted, or remembered?
- **Feedback**: do they always know what's happening, whether it worked, and
  what changed?
- **Mistakes**: what's easy to get wrong? Is it prevented, caught early, and
  explained in a way that says how to fix it? Can they undo it?
- **Confidence**: at the end, are they sure it's done and done right? Would
  they trust it with real work?
- **Consistency**: does it work like the rest of the app, and like the tools
  this user already knows?
- **Fit**: does it suit their context, such as one-handed on a phone, in a
  hurry, done fifty times a day, or done once a year?

### Rules for this phase

- Don't change project files in this phase. Scripts and test data go in the
  scratchpad.
- Use local or test environments and test data. Never touch production or
  real customer data, and never send real emails, messages, or payments.
- Follow the project's instructions on what may be run. If the UI or API
  can't be run, say so. A walkthrough read from code is a guess at the
  experience. Label any finding that comes from one as "from code, not tried".
- The flow shares a browser session or a server, so run it in the main
  conversation. When walking as several user types, do them one at a time and
  report each separately.

## Scale to the stability level

Judge against the expectations in the `stability` skill's reference for the
feature's level:

- **Prototype**: only the experience its learning question is about. Rough
  edges elsewhere aren't findings.
- **Alpha**: this user must be able to complete the main flow without help and
  without getting it wrong. Anything that stops them, or leads them to the
  wrong result, is a finding. Effort and polish can wait unless they're
  severe.
- **Beta** and **GA**: the full set of lenses, and the variations.

Something only a later level expects isn't a finding. Neither is anything in
`known_gaps`. List it under "Later, at <level>" only when it's a real problem
this user would hit.

## 4. Check what you can't judge alone

Some findings depend on things only the user knows: who this user really is,
how they work, or a design decision already made on purpose. Phrase each as a
question with the options you see, and carry on under the most likely one.
Ask them with the report in step 6, not before.

## 5. Report

```
Feature: <name>
Stability: <level> (source: <declaration | user | none declared, so GA>)
User: <type of user>. Goal: <goal>
Scope: <the flow walked, the variations, environment, commit; through UI at
        <widths> | API | from code only>

Verdict
<Two or three sentences: could they get it done, how sure were they, and how
much effort it took. Then the steps, clicks or calls, and rough time taken.>

Walkthrough
S1. <intent>: <what happened, in a line> (friction: <what>, or none)

Findings
1. [<severity>] <title>
   Step: S<n>
   Experience: <what this user ran into, in their terms>
   Why it matters: <the effect on this user's goal, or on trust>
   Evidence: <screenshot path and width, response, or quote of the text>
   Suggestion: <one line, sized to the level>

Works well
- <what helped this user and should be kept>

Bugs found along the way
- <behavior that's wrong, not just hard>: suggest `/qa` for it

Open questions
Q1. <question>. Options: A, B. Assumed: A.

Later, at <level>
- <item>

Not tried
- <variation or check>: why
```

Severities:

- **Blocker**: this user can't complete the flow, or thinks they did when
  they didn't, or completes it with the wrong result.
- **Major**: they get there, but only after a wrong turn, a guess they could
  easily get wrong, real wasted effort, or needing help.
- **Minor**: noticeable friction or confusion they get past on their own.
- **Polish**: wording, consistency, and small refinements.

Order findings by severity. Omit empty sections. Keep "Works well" short, and
only for things worth protecting in a later change.

## 6. Confirm what to change

UX changes are often design decisions, so the user decides. Ask which
findings to act on (AskUserQuestion, multiSelect; group them by severity if
there are many), recommending Blockers and Majors. Ask the open questions
whose answer changes a finding. Where a finding has more than one reasonable
fix, offer the options instead of choosing.

If the user picks nothing, stop.

## 7. Change and walk it again

- Change only what's confirmed, sized to the stability level. Follow the
  `stability` skill on breaking changes and the project's own UI and API
  rules.
- If a change needs a decision that's the user's to make, ask.
- Walk the affected steps again as the same user, through the same path and
  at the same width, to confirm the friction is gone and nothing nearby got
  worse. Run the project's checks.
- Anything new found while changing goes in the report. Don't change it
  without asking.
- Don't commit unless asked.

Report each confirmed finding as addressed (and how the walk again confirmed
it) or not (and why), then the checks you ran with their results, and any new
findings.
