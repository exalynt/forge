# README

`docs/product/README.md`, one per repository that holds product docs. Create it
with the first doc in `docs/product/`.

The README is for people opening these docs without an agent. It says what they
are for and what goes where, in plain words, and never mentions agents, skills,
or commands. It describes the kinds of doc and doesn't list them, so adding a doc
doesn't change it.

## Layout

Use this text, filling in the placeholders. Keep every row of the table, even for
a folder that doesn't exist yet, since it's how someone knows where a new doc goes.

```markdown
# Product docs

What <product> is and does, in the words of the people who use it. You don't need
to know the code to read these. How the software is built is in the
[engineering docs](<relative path to docs/engineering/README.md>).

## What goes where

| Folder             | What it holds                                                    | The question it answers                              |
| ------------------ | ---------------------------------------------------------------- | ---------------------------------------------------- |
| `personas/`        | The people who use the product                                   | Who is this for?                                     |
| `capabilities/`    | What people can do, each with its stability level and known gaps | What can someone accomplish, and how finished is it? |
| `flows/`           | How work gets done from start to finish, across capabilities     | What happens, in order, to get the job done?         |
| `domain/rules/`    | What must or can't happen                                        | Is this allowed?                                     |
| `domain/concepts/` | What a term in the domain means                                  | What exactly is this thing?                          |

A folder appears once it has a doc in it.

## Where to start

New to the product? Read the personas to see who it's for, then the flows to see
how their work gets done. Go to a capability to see what's possible today and
what's missing, and to the rules for the details.

## Deciding where something goes

- **Something people can accomplish** is a capability. A page, a button, or an
  endpoint isn't one.
- **Something that must or can't happen** goes in a rule doc, once, and
  capabilities link to it.
- **What a term means** goes in a concept doc, once, and everything else links to
  it.
- **A question nobody has settled** goes under Open decisions in its capability.
  It becomes a rule once it's decided.
- **Something missing** goes under Known gaps in its capability if someone using
  the product would notice. If only someone working on the code would, it goes in
  the engineering docs. It's never in both.

These docs name no pages, buttons, endpoints, or tables, so a redesign doesn't
make them wrong.

## Stability levels

Each capability says how far along it is:

| Level     | What it means                                                      |
| --------- | ------------------------------------------------------------------ |
| Prototype | An experiment to answer a question. It may be rewritten or removed |
| Alpha     | The main workflow is real and usable, but its shape may change     |
| Beta      | The shape has settled, and the edges are being hardened            |
| GA        | People can depend on it                                            |

## Keeping them current

A doc changes along with whatever it describes. Where a rule has an acceptance
test, that test is the specification, so changing the rule means changing the
test too.
```

## Filling it in

- **`<product>`**: the product's name.
- **The link to the engineering docs**: `../engineering/README.md` when they're in
  the same repository. If this repository has none, link to the repository that
  does, or remove the sentence.
- Say so under the title if these docs also cover other repositories, naming them.
- Don't add a list of the docs. The folders are the list.
