# System

`docs/engineering/architecture/systems/<slug>.md`, one per mechanism or system.

A system doc explains how something overarching works, when several features or
repositories rely on it and you need the whole picture to change it safely. It's
the map for reading the code, not a tour of it. Under about 150 lines.

## Layout

```markdown
# <Name>

<What this is responsible for and where it sits, in a sentence or two.>

**Code:** <where to start reading: packages, files, or symbols, per repository
when it spans several>

## How it works

<The pieces, how they fit, and what flows through them.>

## Tradeoffs

## Gotchas

## Known gaps

## Changing it

**<Adding a thing>:** <the steps, and what fails if one is missed.>
```

## Sections

- **Code**: required.
- **How it works**: required. Only what no single file shows, such as the flow
  across packages, or what must stay true between them.
- **Tradeoffs**: only if there are some.
- **Gotchas**: only if there are some.
- **Known gaps**: required. Write "None." when there are none.
- **Changing it**: only for changes that are easy to get partly wrong, such as
  adding a permission or a table. Leave it out otherwise.

The format of Tradeoffs, Gotchas, and Known gaps is in
[SKILL.md](SKILL.md#tradeoffs-gotchas-and-known-gaps).
