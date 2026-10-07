---
name: promote
description: Deliberately promote a feature to a higher stability level (Prototype → Alpha → Beta → GA), implementing only the delta the target level requires.
argument-hint: <feature> <target-level>
disable-model-invocation: true
---

# Promote a Feature

Arguments: `$ARGUMENTS`. Expected format: `<feature> <target-level>`, for
example `/promote conversations alpha`.

Promotion increases engineering investment only enough to satisfy the next
deliberate commitment. Use the `stability` skill for the level definitions.

## 1. Establish current and target levels

- Find the feature's stability declaration as described in the `stability`
  skill. If there isn't one, ask for the current level and where to record it.
- The target must be one of `prototype`, `alpha`, `beta`, `ga`. If it is
  missing or invalid, ask.
- If the target is not above the current level, stop. This isn't a promotion.
- If the target skips a level (for example prototype → beta), confirm that's
  intended. If it is, the delta covers each intermediate level up to the
  target, and never beyond it.

## 2. Review against the target

Invoke the `stability-reviewer` subagent. Give it the feature, where its code
lives if known, the current level, and tell it to review against the
**target** level. Its REQUIRED NOW findings are what the target requires. Its
UNNECESSARY NOW findings exceed even the target.

## 3. Present the delta

Before changing anything, present:

```
Current stability: <level>
Target stability: <level>

What changes are required:
- <change>: why, meaning which target-level expectation it satisfies

What we are intentionally NOT doing yet:
- <tempting work that belongs to a later level, or is speculative>

Known gaps that remain acceptable at <target>:
- <gap>
```

"What we are intentionally NOT doing yet" is mandatory. Name the specific
hardening, abstraction, or operational work someone might expect, and say
which later level (if any) it belongs to.

If the reviewer flagged existing over-engineering, list it separately. Removing
it is not part of the promotion unless the user agrees.

If the delta is large or contains decisions that are the user's to make, such
as choosing a monitoring stack or a public contract, ask before implementing.
Otherwise, proceed.

## 4. Implement only the delta

- Make only the changes listed under "What changes are required."
- Don't redesign unrelated parts of the feature or fix unrelated issues.
- Don't implement requirements from levels beyond the target.
- If you discover another required change, add it to the delta and say so.

## 5. Verify

- Run the relevant tests, linters, and type checks. Fix failures your changes caused.
- Invoke `stability-reviewer` again against the target. Resolve remaining
  REQUIRED NOW and SECURITY/PRIVACY findings, or explain why you can't.

## 6. Record the new level

Only when the second review passes for the target: update the feature's
declared stability level, and update `known_gaps` (or the project's
equivalent) to reflect the gaps that remain. If it doesn't pass, leave the
declared level unchanged.

## 7. Report

Finish with the same format as step 3, reflecting what was actually done. Add
the checks you ran with their results, and, if the promotion is incomplete,
what is still outstanding.
