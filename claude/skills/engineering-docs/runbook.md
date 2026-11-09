# Runbook

`docs/engineering/runbooks/<situation>.md`, one per situation, named so someone
in the middle of an incident can find it, such as `refund-a-booking.md`.

A runbook says what to do when something has to be fixed by hand. If the product
can do it, or a retry fixes it, there's no runbook.

## Layout

```markdown
# <Situation, starting with a verb>

<Why this has to be done by hand, in a sentence.>

## When

<How someone would notice: the symptom, alert, log line, or request.>

## Check

<What confirms it's this situation, using read-only steps.>

## Do

1. <Each step, in order, with the access it needs.>

## Verify

<How to tell it worked, and who to tell.>
```

## Sections

All four are required.

- **Do**: commands use placeholders such as `'<booking id>'`, and never real
  data or secrets. Say what can't be undone before the step that does it.
- Link to the feature or system doc that explains why, rather than explaining
  it here.
