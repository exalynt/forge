# Feature

`docs/engineering/architecture/features/<slug>.md`, only for a feature with
something the code can't say.

A feature doc helps someone working on one feature: a high-level map of how it
works, what they would otherwise miss, or both. Most features don't have one, and
there's no template to fill in, so no model walkthrough, permission list, or
operations section. Under about 75 lines.

## Layout

```markdown
# <Name>

<What this part is and anything needed to read the sections below, in one to
three sentences.>

**Code:** <where to start reading: packages, files, or symbols, per repository
when it spans several>

## How it works

<The high-level flow: the pieces involved, the path through them, and where
its state lives.>

## Tradeoffs

## Gotchas

## Known gaps
```

## Sections

- **Code**: required.
- **How it works**: only when the feature needs a map, because it crosses layers
  or repositories, runs asynchronously, or has a flow no single file shows. Keep
  to the flow and the pieces, and link to rule docs rather than restating them.
- **Tradeoffs**, **Gotchas**, and **Known gaps**: leave out any with nothing in
  it.

A feature doc has at least one of those four sections. One with only a few
gotchas is complete.

The format of Tradeoffs, Gotchas, and Known gaps is in
[SKILL.md](SKILL.md#tradeoffs-gotchas-and-known-gaps).
