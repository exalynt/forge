# README

`docs/engineering/README.md`, one per repository with engineering docs. Create it
with the first doc in `docs/engineering/`.

The README is for people opening these docs without an agent. It says what they
are for and what goes where, in plain words, and never mentions agents, skills,
or commands.

## Layout

Use this text, filling in the placeholders. Keep every row of the table, even for
a folder that doesn't exist yet, since it's how someone knows where a new doc goes.

```markdown
# Engineering docs

How <project> is built, for the people working on it. These docs give you a map
for reading the code and warn you about what it can't show. They don't replace
it: if you could learn something in a few minutes from the code, spec, migrations,
or tests, it isn't here. What the product does is in the
[product docs](<relative path to docs/product/README.md>).

## What goes where

| Folder                   | What it holds                                                  | Add a doc when                                                     |
| ------------------------ | -------------------------------------------------------------- | ------------------------------------------------------------------ |
| `architecture/systems/`  | How something overarching works, such as authentication        | Several features or repositories rely on it                        |
| `architecture/features/` | How one feature works, and what to know before changing it     | The feature needs a map, or has tradeoffs, gotchas, or gaps to record |
| `providers/`             | A third-party service: how it's set up and how it behaves      | It needs setup outside this repository, or was checked by hand     |
| `runbooks/`              | What to do when something has to be fixed by hand              | A recovery can only be done by hand                                |

A folder appears once it has a doc in it. Most features have no doc, and that's
expected. A feature gets one only when there's something the code can't say.

## What you'll find in a doc

| Section    | What it tells you                                               |
| ---------- | --------------------------------------------------------------- |
| Tradeoffs  | What we chose, what it costs, and why that's worth it           |
| Gotchas    | What's correct but will surprise whoever changes it             |
| Known gaps | What isn't done yet, and what closes it                         |

A gap we intend to close is a known gap, and says what it's waiting on. Something
we've decided to live with for good is a tradeoff.

## Deciding where something goes

- **Several features would trip over it**: a system doc.
- **Only someone working on one feature would**: that feature's doc. If a second
  feature comes to need it, move it to a system doc and link to it.
- **Something missing that someone using the product would notice**: the product
  docs, under its capability's Known gaps. If only someone working on the code
  would notice, it goes here. It's never in both.
- **Something true of one function**: a code comment.
- **A business rule**: the product docs.
- **An environment variable**: the repository's README.
- **Something true only while a change rolls out**, such as deploy order: its pull
  request.

These docs never hold real secrets, accounts, or anyone's personal details.

## Keeping them current

A doc changes along with whatever it describes. When a gap is closed, delete it:
the history is in git.
```

Add these sections at the end only when they apply:

```markdown
## Across repositories

<Which repository holds the whole picture for what spans several, and which other
repositories' docs to read.>

## The system at a glance

<One screen on what the system is made of and how a request or job moves through
it, with a link to each doc in `architecture/systems/`.>
```

## Filling it in

- **`<project>`**: the name of what this repository builds.
- **The link to the product docs**: `../product/README.md` when they're in the same
  repository. Otherwise link to the repository that holds them, such as
  `../../../api/docs/product/README.md`.
- **Across repositories**: only when this repository's code is part of something
  spanning several.
- **The system at a glance**: once there are a few system docs to find. It's the
  only part of the README that lists docs, so a system doc added or removed updates
  it.
