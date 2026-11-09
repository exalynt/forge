---
name: use
description: Usability test a product flow the way real people would, with one agent per person, each in their own role and browser, working together through the product. Report whether each person could do what they needed and whether the whole flow got done, along with the friction they ran into.
argument-hint: "[low|medium|high] <flow> [as <role>, <role>...]"
disable-model-invocation: true
---

# Usability Test

Arguments: `$ARGUMENTS`. Expected format:
`[effort] <flow> [as <role>, <role>...]`, for example
`/use client onboarding`,
`/use high capacity as a client and an engineer`, or
`/use low invoices as a finance admin`. A role can name a persona. The effort
is `low`, `medium`, or `high`, and sets how much time and cost the test spends
(see "Effort"). Without one, use `medium`. If no flow is given, ask.

A usability test asks: can each person actually do what they need to do, and
does the whole flow get done? It's run like a session with real people. You're
the facilitator. Each person is a `use-role` agent that knows only what that
person would know, uses the product through their own browser, and works with
the others as real colleagues would.

How easy it was is `/ux`'s question, and whether the rules hold is `/qa`'s.
Friction and bugs get reported here when people run into them, but this test
doesn't look for them. One person is a valid test: skip what's about handoffs
and messages, and the rest is the same.

Steps 1–7 test and report, and change nothing in the project. Step 8 only
decides what happens next. This skill doesn't change code.

## 1. Understand the flow

- Find the flow from the user in this conversation, its doc in
  `docs/product/flows` if it has one, and the project's other docs: what it's
  for, who's involved, and what's true when it's finished. A flow doc's steps
  are for you, and never reach the people.
- Find the stability level of each feature it crosses, as the `stability`
  skill describes. Read Known gaps and Learning in their capability docs.
- Find what the test needs: each app's URL, test accounts and how they're
  seeded, and test modes for payments and other external services. Code and
  seeders are fine for this. You're the facilitator, and they never reach the
  people.

## 2. Define the people and the scenario

Take the roles from the arguments, the flow doc, the capability docs, or the
project's other docs. There are four browsers, so four people at most. Find each
person's persona with the `personas` skill. When there isn't one, it asks whether
to create one or use an existing one, with all the missing roles in one round.

Then write the scenario in the people's terms, with goals and no steps:

```
Scenario: <name>
Done when: <the observable end state, across everyone>

<Name>, <role> (<persona file, if saved>)
  Goal: <what they need to get done, in their words>
  Knows: <beyond the persona: the other people in the scenario, this situation>
  Starts when: <what prompts them to begin>
  Has: <the app's URL, an account or how they'd get one>

Preconditions: <what must already be true, and what you'll hand to people
                because it can't happen for real, such as an invite link>
```

Mark what you assumed, then show it to the user and confirm it (AskUserQuestion,
one round, your recommended option first). Don't run a scenario the user hasn't
seen.

## 3. Set up

- The user starts the stack. Check that each app's URL responds and that it's
  a local or test environment. If not, ask the user to start it and wait. Don't
  start servers yourself unless they ask.
- Set up the preconditions: seeded accounts, test mode, and any data the
  scenario starts with. Use a new organization or new accounts for each run
  where the flow creates them, so runs don't collide.
- Run `browser_close` on `playwright-1` through `playwright-4`, so no sign-in
  or page carries over from an earlier run. If those servers aren't there, say
  to run `claude/scripts/playwright/setup.sh` and start a new session, then stop.

### Rules for the whole test

- Don't change project files, except to git-ignore the QA folder and to save a
  persona the user confirmed.
- The QA folder is `.playwright-mcp/` at the project root. Everything created
  goes there, and the file names include the folder, such as
  `.playwright-mcp/use-<person>-<step>-1280.png`. A bare name lands in the
  project root. People's screenshots go there too, through their browser tools.
- Before creating anything, check that the folder is ignored with
  `git check-ignore -q .playwright-mcp/`. If it isn't, add `/.playwright-mcp/`
  to the root `.gitignore`, and say so in the report. When you're done,
  `git status` should show nothing new besides that line and any saved persona.
- Never touch production or real customer data, and never send real emails,
  messages, or payments.
- Everyone shares the running app on purpose, because the test is about them
  affecting each other. That's why this runs in parallel subagents despite the
  subagents rule.
- You assign each person a numbered browser, and no two people share one.
  Subagents share this session's browsers, so two on one server would land on
  the same page and sign-in. Another Claude session has its own browsers, so a
  test or `/qa` running there can't touch yours through its browser tools, but
  it does use the same app and database. That's another reason each run uses
  its own organization or accounts.
- Close a browser with `browser_close`. Never kill Chrome or Playwright
  processes by name, such as with `pkill chrome`, even if a browser seems stuck.
  That takes down other sessions' browsers too. Report a browser you can't
  recover, and let the user decide.
- Stop and tell the user if anyone sees another tenant's real data or a real
  secret.

## 4. Run the scenario

Start each person when their part of the story begins, as a background
`use-role` agent, on the model "Effort" gives. At high, set no model, so they
use the session's. Their prompt has their persona's full text, since they can't
read files, their section from step 2, their browser (`playwright-<n>`), how
often to think aloud and look at the page (see "Effort"; at high, at every
step), and the agent ids of the people they'd already know how to reach, if any. An agent can't see its own id, so when someone new becomes
reachable, send it to the people who'd know them.

Then run the session:

- **Give them the world, not the answer.** Wake someone with what happens in
  their life: "It's the next morning, and you open the app." Never say what
  someone else did, what to do next, or where to go.
- **Let them work together.** They use the product wherever it lets them. They
  can also message each other directly, as they would by email or chat, and a
  message wakes the person it's sent to. Wake two people at once when the flow
  has them working at the same time.
- **Listen to them think.** Each person thinks aloud to you as they go, a line
  for every step at high effort, and those messages are already shown to the
  user. Don't reply, repeat them, or pass them to anyone else in the test, and
  don't let them change what you tell people. The exception is anything "Rules
  for the whole test" says to stop for.
- **Hand over what can't happen for real.** An invite link, an emailed code,
  anything an external service would deliver in production: tell that person
  it arrived, as it would reach them, and record it under "Not tested".
- **Check in as life would.** When someone's waiting, wake them later to look
  in the way they normally would. Finding out from the product that it's their
  turn is what's being tested.
- **When the flow stalls**, because everyone's waiting or someone's stuck,
  record it as a finding, then do the least that gets it moving, in this order:
  1. Give the person a nudge a real person might get, such as a colleague
     asking how it's going. Record that you intervened.
  2. If the product can't get them past this step, set up the resulting state
     yourself, through the API or database in the test environment, tell them
     what they'd now see, and mark the step as bypassed.

  Stop when the scenario's "Done when" is true, when nothing else can be
  tried, or when a stall has no way through.

Read each person's report as it comes back and keep one timeline: who did
what, what they were thinking when it mattered, every message sent outside the
product and why, every wait, every intervention, and every bypass.

## 5. Debrief

Ask each person, still in character:

- What do you think happened? Is it finished?
- What are you waiting on, and from whom?
- How sure are you that it's right? What would you check?
- Where did you need someone else, or wish you could've asked someone?

Compare their answers with each other and with the real state. Where they
disagree, or someone's sure of something false, that's a finding even when
every screen worked. Then run `browser_close` on each browser.

## Scale to the stability level

Each step is judged by the level of the feature it's in:

- **Prototype**: only what its learning question is about. Rough edges
  elsewhere aren't findings.
- **Alpha** and up: each person must be able to do their part without help and
  without getting it wrong.

A step the flow needs that's in a capability's Known gaps, or not built yet at
that level, isn't a finding. Bypass it and list it under "Later". Everything else about
getting the flow done counts at every level. Tenant isolation and leaked
secrets are always a Blocker.

## Effort

Effort controls how much of the flow is covered, and so the time and token
cost. Each person is a long-running agent reading full pages, so this adds up
quickly. The steps above describe **high**, and the table says what **medium**
and **low** do less of:

| Part           | Low                                       | Medium                             |
| -------------- | ----------------------------------------- | ---------------------------------- |
| The people (2) | Only those the main flow can't do without | Everyone involved in the main flow |
| Check-ins      | One for each wait                         | As needed                          |
| Debrief        | Only "Is it finished?" and "How sure?"    | All the questions                  |
| Thoughts (4)   | Only at a surprise, wrong turn, or doubt  | At each new screen and each choice |
| Model (4)      | Sonnet                                    | Sonnet                             |
| Looking (4)    | Snapshot new pages; screenshot if stuck   | Also screenshot each new page      |
| Stalls (4)     | One nudge, then stop and report           | As described in step 4             |

Low and medium run the main flow once. **High** also runs the variations a
person would hit: the mistake each person is most likely to make, someone
coming back after an interruption, and the second time through. List anything
skipped under "Not tested".

## 6. Check what you can't judge alone

Some findings depend on things only the user knows, such as how a handoff
really happens today or whether something outside the product is intended.
Phrase each as a question with the options you see, carry on under the most
likely one, and ask with the report.

## 7. Report

```
Flow: <name>
Effort: <low | medium | high>
Scope: <people, environment, commit, widths>

Verdict: <done without help | done only with help | not done>
<Two or three sentences: could everyone do what they needed, where it broke
down if it did, and what it took.>

People
<Name> (<role>): <could do their part? yes | only with help | no>
  <steps, wrong turns, stalls; a line on how it went for them>

Timeline
1. <Name>: <what they did or tried> (bypassed | intervened, if so)
   Thinking: <their words, where it shaped what they did or showed a doubt>
2. <Name> → <Name>, outside the product: <message, and why they needed it>

Handoffs
- <from> → <to>: <how the next person found out, what they had to act on, and
  whether the first person could tell where it stood>

Findings
1. [<severity>] <title>
   Who: <person>. Step: <timeline number>
   What happened: <in their terms>
   Why it matters: <the effect on getting the flow done>
   Evidence: <screenshot path and width, or a quote of the text>
   Suggestion: <one line, sized to the feature's level>

Understanding
- <where people's beliefs differed from each other or from the real state>

Friction along the way
- <person, step>: <what slowed them down but didn't stop them>

Bugs along the way
- <behavior that's wrong, not just hard>

Open questions
Q1. <question>. Options: A, B. Assumed: A.

Later
- <step>: <in Known gaps or not built yet, at this level>

Not tested
- <what was handed over or bypassed, or skipped at this effort>
```

Severities:

- **Blocker**: someone can't do their part, thinks it's done when it isn't,
  gets the wrong result, or the flow stalls with no way for anyone to know
  it's their turn.
- **Major**: it only gets done with help, a guess that could easily go wrong,
  or a workaround outside the product where the product should have done it.

Friction has no severity here and doesn't compete with findings. If there's a
lot of it around a step, suggest `/ux` for that step, and suggest `/qa` for the
bugs.

Agents aren't people. They're patient, they read every label, and they don't
get bored or distracted, so this says more about whether something can be done
than about how it feels. Report steps, wrong turns, and stalls, not timings.
Say so in the report if a finding depends on how a person would really behave.

## 8. Decide what's next

Ask which findings to act on (AskUserQuestion, multiSelect), recommending
Blockers and Majors, along with the open questions whose answer changes a
finding. This skill stops there. Suggest `/implement` for the chosen findings,
and `/use` again afterwards, with the same scenario, to confirm the flow got
through.

If a flow now works, you can also propose it as an acceptance test, as the
testing rule describes, which needs the user's explicit approval.
