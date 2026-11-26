---
name: promote
description: Deliberately promote a capability to a higher stability level (prototype, alpha, beta, ga) in a repository that has Docqet, implementing only the delta the target level requires and recording the new level last.
disable-model-invocation: true
---

# Promote a Capability

The user's request is the input. It names the capability and the target
level, and may name an effort, for example "promote room-booking to alpha" or
"promote room-booking to beta, high effort". The effort is `low`, `medium`, or
`high`, and sets how much the stability reviews cover, as the
`stability-review` skill describes. Without one, use `medium`.

Promotion increases engineering investment only enough to satisfy the next
deliberate commitment. It happens only when a person asks for it.

## Needs Docqet

A capability's level, and what each level expects, live in Docqet (see
forge's `AGENTS.md`, "Where intention lives"). Look for a `docqet.yaml` at the
working directory or above it, within the repository. If there isn't one,
stop and say promotion needs Docqet: the repository can adopt it with
`docqet init` and `docqet kit`.

The `docqet` skill's stability reference, "Promoting a capability", is the
procedure. Follow it. This skill adds forge's review, the shape of the delta,
and its approvals around those steps, and doesn't restate them.

## 1. Confirm the capability and levels

Confirm the capability, its current level, and the target, as the kit's
first step says. If the target isn't above the current level, stop: this
isn't a promotion. A skip covers each level in between, and never beyond the
target, so confirm it's intended.

If the capability plainly isn't ready for the target, such as a shape that's
still changing, say so before going on. A move to Beta or GA changes what the
client relies on, so the client should have agreed to it. If the user hasn't
said they have, note it in the delta.

## 2. Review against the target

Run the `stability-review` skill, in a subagent if your agent can run one,
otherwise yourself. Give it the capability, the current level, the effort,
and tell it to review against the **target** level, using
`docqet rubric --for capability:<id> --to <target>`. At low effort, a
smaller, cheaper model is fine for it, if your agent lets you choose. Its
REQUIRED NOW findings are what the target requires. Its UNNECESSARY NOW
findings exceed even the target. The gaps on the components and mechanisms the
capability relies on, which the rubric's `--for` output and `docqet touches`
show, are part of the delta when they affect an outcome the target expects.

## 3. Present the delta

Before changing anything, present:

```
Capability: <id>
Current stability: <level>
Target stability: <level>

What changes are required:
- <change>: which outcome of the target it meets

Proposed as gaps instead:
- <outcome left unmet for now>: close later | accepted as a tradeoff, and why;
  with its impact, likelihood, and reach

What we are intentionally NOT doing yet:
- <tempting work that belongs to a later level, or is speculative>

Gaps that remain acceptable at <target>:
- <gap already written in Docqet>
```

The rubric is a guideline, not a gate. Each REQUIRED NOW finding is either in
"What changes are required" or, where closing it now isn't worth it, proposed
as a gap. Proposed gaps need the user's agreement. The security and privacy
baseline is never a gap.

"What we are intentionally NOT doing yet" is mandatory. Name the specific
hardening, abstraction, or operational work someone might expect, and say
which later level (if any) it belongs to.

If the review flagged existing over-engineering, list it separately. Removing
it is not part of the promotion unless the user agrees.

If the delta is large or contains decisions that are the user's to make, such
as choosing a monitoring stack or a public contract, ask before implementing.
If it adds or changes acceptance tests or rule docs, show the Gherkin and the
rule and get explicit approval for them, whatever the delta's size (see the
`testing` skill). Which outcomes become gaps is the user's call too.
Otherwise, proceed.

## 4. Implement only the delta

- Make only the changes listed under "What changes are required."
- Don't redesign unrelated parts of the capability or fix unrelated issues.
- Don't implement what levels beyond the target expect.
- If you discover another required change, add it to the delta and say so.
- Delta items in disjoint files that don't depend on each other may run in
  parallel subagents, if your agent can run them. Anything touching shared
  files, migrations, or installs stays in the main session.

## 5. Verify

- Run the relevant tests, linters, and type checks. Fix failures your changes
  caused.
- Run `stability-review` again to verify the fixes, not to review from
  scratch. Give it the target level and the first review's findings you set
  out to fix. Resolve whatever it reports NOT RESOLVED, or explain why you
  can't.

## 6. Record

Below high effort, the first review didn't check everything the target
expects. Show its NOT REVIEWED list and ask before recording the new level.

Only when the verification passes for the target: record the agreed gaps the
kit's way, then change the capability's `stability` last, and finish as the
`docqet` skill's "Before you finish" describes. If it doesn't pass, leave the
level unchanged.

## 7. Report

Finish with the same format as step 3, reflecting what was actually done. Add
the checks you ran with their results, the Docqet section from `docqet diff`,
and, if the promotion is incomplete, what is still outstanding.
