---
name: ux
description: Try a feature's main flow, through its UI or API, as a specific type of user, report what it's like for them to use (where they hesitate, get stuck, or do more work than they should), then confirm which findings to act on before changing anything.
argument-hint: "[low|medium|high] <feature> [as <type of user>]"
disable-model-invocation: true
---

# UX Walkthrough

Arguments: `$ARGUMENTS`. Expected format:
`[effort] <feature> [as <type of user>]`, for example
`/ux conversations as a support agent on a phone` or
`/ux high billing-api as a partner developer integrating for the first time`.
The type of user can name a persona. The effort is `low`, `medium`, or `high`,
and sets how much time and cost the walkthrough spends (see "Effort"). Without
one, use `medium`. If no feature is given, ask. Anything after the user type
narrows the walkthrough to part of the flow.

A walkthrough asks how it is for this user to get the main flow done. It
judges experience, not correctness. A flow can be bug-free and still slow,
confusing, or easy to get wrong. Bugs found along the way get reported, but
looking for them is `/qa`'s job. Reviewing code is `/code-review`'s. Whether
people can get a flow done at all, especially several people working together,
is `/use`'s.

There are two phases. Steps 1–5 try the feature and report, and change
nothing in the project. Steps 6–7 change only what the user confirms.

## 1. Understand the feature

- Find its stability level and capability doc as the `stability` skill
  describes. Read its purpose, Learning, and Known gaps.
- Find the main flow: what the feature is for, and the path from a user's
  first action to the result they came for. Use the user in this
  conversation, the capability doc, any doc in `docs/product/flows` that
  crosses it, and the project's other docs first. Read code only
  to find where the feature lives and how to reach it, not the steps to take.
  Step 3 works out those steps the way the user would.
- Find what the user would actually have: the screens and navigation for a
  UI; for an API, the public docs, specs, examples, and SDKs an integrator
  would be given.

## 2. Define the user and their goal

Find the persona for this user with the `personas` skill, before trying
anything. When there isn't one, or no user type was given, it asks whether to
create one or use an existing one. Then write down who you're walking through
as:

```
User: <type of user>
Goal: <what they came to get done, in their words>
Knows: <domain knowledge, familiarity with this app, technical skill>
Context: <device, how often they do this, time pressure, interruptions>
Expects: <what they think the first step is, the words they'd look for,
          what "done" looks like to them>
Cares about: <what makes them trust it, and what would make them give up>
Success: <the observable end state that means they got it done>
```

Knows, Context, Expects, and Cares about start from the persona. Fill in the
rest for this walkthrough from the project and common sense, and mark what you
assumed. Ask (AskUserQuestion, one round) only about assumptions that would
change the walkthrough, such as whether they're new or experienced this time,
with your recommended option first.

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
    Thinking: <what's going through this user's head>
```

Fill in Thinking at every step, once you've looked at the screen or response and
before you act. "Effort" says which steps at medium and low. Keep it in the log,
not in chat as you go, and the report carries it.

- One or two sentences, in the first person, as this user would think them: what
  catches their eye, what they make of it, what they're worried about, what
  they're about to do and why. Let what they care about, what they expect, and
  what they're used to shape it.
- Say the doubts and guesses too, and don't tidy them up afterwards.
- Stick to what this user would know. What you've learned from the code or the
  docs stays out of it.
- Keep your own judgment out of it. Lenses and findings go in the report.

**Through the UI**: use whichever browser tools are available, with realistic
data for this user, such as their volume of items, their names, and their
statuses. Run it at their device's width: 360px or 375px for a phone user,
desktop otherwise, at 1280px or wider. Set it with the browser's resize tool:
the window is sized to fit the user's screen, so its own width isn't a test
width. Take screenshots at the moments that matter and note the
width. Save them in the QA folder (below).

**Through an API**: start from the public docs, as the integrator would.
Write the client a developer would write, with curl or a short script in the
QA folder. Work through getting credentials, making the first successful
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
- **Workload**: steps, clicks, typing, fields, waits, calls, and switching
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

- Don't change project files in this phase, except to git-ignore the QA
  folder and to save a persona the user confirmed.
- The QA folder is `.playwright-mcp/` at the project root. Everything you
  create goes there: screenshots, snapshots, console logs, scripts, and test
  data. Playwright MCP writes its own files there, and it refuses paths
  outside the project or that folder, so a scratchpad path fails. Include the
  folder in every file name you pass, such as
  `.playwright-mcp/ux-<step>-375.png`. A bare name lands in the project root.
- Before creating anything, check that the folder is ignored with
  `git check-ignore -q .playwright-mcp/`. If it isn't, add `/.playwright-mcp/`
  to the root `.gitignore`, and say so in the report. When you're done,
  `git status` should show nothing new besides that line and any saved persona.
  Move anything else you created into the folder, or delete it.
- Use local or test environments and test data. Never touch production or
  real customer data, and never send real emails, messages, or payments.
- Follow the project's instructions on what may be run. If the UI or API
  can't be run, say so. A walkthrough read from code is a guess at the
  experience. Label any finding that comes from one as "from code, not tried".
- When you start a dev server, use a port that's neither the framework's
  default (Vite's 5173) nor one the project sets anywhere: in its config, dev
  scripts, `.env` files, or Compose files, including the ports of other apps
  it runs beside. The user's own instances run there. Pick a free port outside
  all of them, pass it explicitly
  (`npm run dev -- --port <port> --strictPort`), point the browser at it, and
  stop the server when you're done.
- Run the flow in the main conversation, on one browser, `playwright`. Another
  Claude session has its own browsers, so it can't touch yours through its
  browser tools, and you can't touch its. When walking as several user types,
  do them one at a time and report each separately.
- Keep browser output small, since a long page's snapshot runs to tens of
  thousands of characters. With Playwright MCP, take a full `browser_snapshot`
  on each new page and reuse its refs until the page changes. A ref that's gone
  returns an error saying to take a new one. To find or check something on a
  page you've already seen, use `browser_find`, or `browser_snapshot` with a
  `target`. At high effort, take a full snapshot whenever it helps. "Effort"
  says when at medium and low. Below high, once you've seen that a page is long,
  such as a list with many rows, don't take another full snapshot of it.
- Use as few model calls as you can, since each one re-reads the whole
  conversation. Fill a whole form with `browser_fill_form`, and send actions
  together when you don't need to see the result of one before the next.
- Close a browser with `browser_close`. Never kill Chrome or Playwright
  processes by name, such as with `pkill chrome`, even if a browser seems stuck.
  That takes down other sessions' browsers too. Report a browser you can't
  recover, and let the user decide.

## Scale to the stability level

Judge against the expectations in the `stability` skill's reference for the
feature's level:

- **Prototype**: only the experience its learning question is about. Rough
  edges elsewhere aren't findings.
- **Alpha**: this user must be able to complete the main flow without help and
  without getting it wrong. Anything that stops them, or leads them to the
  wrong result, is a finding. Workload and polish can wait unless they're
  severe.
- **Beta** and **GA**: the full set of lenses, and the variations, as far as
  the effort covers them.

Something only a later level expects isn't a finding. Neither is anything in
the capability's Known gaps. List it under "Later, at <level>" only when it's a
real problem this user would hit.

## Effort

Effort controls how much of the experience a walkthrough covers, and so its
time and token cost. It's independent of the stability level. The level sets
which friction counts as a finding. Effort decides how much of the flow is
looked at. A low-effort walkthrough of a GA feature judges what it covers by
GA's bar, but it doesn't cover everything GA expects.

The steps in this skill describe **high**. **Medium** and **low** do less:

| Part           | Low                                                                    | Medium                                                                            |
| -------------- | ---------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| The user (2)   | Assume what the persona doesn't cover and mark it.                     | Ask about assumptions that would change the walkthrough.                          |
| Variations     | None. The main flow once.                                              | A first time and the mistake this user is most likely to make.                    |
| Lenses         | Finding the way, Feedback, and Mistakes: what stops this user.         | All of them on the main flow. Finding the way and Mistakes on the variations.     |
| Through an API | The first successful call and the whole outcome.                       | Plus the mistakes a newcomer would make.                                          |
| Screenshots    | Only as evidence for findings.                                         | At the moments that matter.                                                       |
| Thoughts (3)   | Only on steps with friction.                                           | At each new screen and each choice of where to go next.                           |
| Page reads     | Full snapshot per new page; else `browser_find` or a targeted one.     | Also a full snapshot after an action that changes much of the page.               |
| Re-walk (7)    | That the friction is gone, not what's nearby. Checks once, at the end. | As described in step 7.                                                           |

At every effort, find the persona, and walk the main flow end to end as this
user, through the UI where there is one, at their device's width. List the
variations and lenses this effort skipped under "Not tried", as skipped at this
effort.

## 4. Check what you can't judge alone

Some findings depend on things only the user knows: who this user really is,
how they work, or a design decision already made on purpose. Phrase each as a
question with the options you see, and carry on under the most likely one.
Ask them with the report in step 6, not before.

## 5. Report

```
Feature: <name>
Stability: <level> (source: <declaration | user | inferred from <what>>)
Effort: <low | medium | high>
User: <type of user> (<persona file, if saved>). Goal: <goal>
Scope: <the flow walked, the variations, environment, commit; through UI at
        <widths> | API | from code only>

Verdict
<Two or three sentences: could they get it done, how sure were they, and how
much work it took. Then the steps, clicks or calls, and rough time taken.>

Walkthrough
S1. <intent>: <what happened, in a line> (friction: <what>, or none)
    Thinking: <their words from the step log, where it has them>

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
- <variation or check>: why (couldn't run, skipped at <effort> effort)
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
