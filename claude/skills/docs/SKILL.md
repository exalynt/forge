---
name: docs
description: Standard location and layout for a project's product docs in docs/product (personas, capabilities, flows, domain rules, domain concepts). Use when looking for what a product does, who uses it, how its work flows, its business rules, or what a domain term means; when finding a capability's stability level or known gaps; and whenever creating or updating any of these docs.
---

# Product Docs

Product docs describe what a product is and does, in the domain's words, for
both people and agents. They live in `docs/product/` at the root of the
repository. When a product spans several repositories, one of them holds the
docs, and the others' `CLAUDE.md` says which. How the software is built goes in
`docs/engineering/` instead (see the `engineering-docs` skill), and the two never
repeat each other.

| What you're documenting                                 | Where                                    | Format                         |
| ------------------------------------------------------- | ---------------------------------------- | ------------------------------ |
| Who uses the product                                    | `docs/product/personas/<slug>.md`        | [persona.md](persona.md)       |
| What people can do, its stability level, and known gaps | `docs/product/capabilities/<slug>.md`    | [capability.md](capability.md) |
| How work gets done from start to finish                 | `docs/product/flows/<slug>.md`           | [flow.md](flow.md)             |
| What must or can't happen                               | `docs/product/domain/rules/<slug>.md`    | [rules.md](rules.md)           |
| What something in the domain is                         | `docs/product/domain/concepts/<slug>.md` | [concept.md](concept.md)       |

Read the format for a type before writing one. A folder exists once it has a
file in it, and a project only needs the docs it has a use for.

`docs/product/README.md` tells people what these docs are for and what goes
where. Create it with the first doc in `docs/product/`, as
[readme.md](readme.md) describes.

## Every doc

- **Name**: a kebab-case slug for what it's about, such as
  `capacity-purchasing.md`.
- **Layout**: the title, then the format's sections, with those exact names and
  in that order. Leave out only the sections the format says can be left out.
- **Words**: the domain's language. Docs name no pages, buttons, endpoints,
  tables, or functions, so a redesign doesn't make them wrong.
- **One place for each fact**: a rule is written in a rule doc, a definition in
  a concept doc, and everything else links to it.
- **Nothing private**: docs are checked in, so they never hold a real person's
  details, an account, a password, or any other secret.

## Links

Links are relative and point one way, so there's one place to update:

```text
flows → capabilities, personas
capabilities → rules, concepts
rules, concepts → concepts
```

To go the other way, such as finding which capabilities a rule applies to,
search `docs/product` for the file's name.

## Keeping them current

Update product docs in the same change as what they describe:

- **What people can do changed**: update the capability doc. A new capability
  gets a new doc.
- **A gap opened or closed**: update the capability's Known gaps if people using
  the product would notice it. One only someone working on the code would see
  goes in the engineering docs instead, and a gap is never in both.
- **A rule was decided or changed**: update its rule doc, along with its
  acceptance tests, which need approval as the testing rule describes.
- **A decision was left open**: add it to the capability's Open decisions. It
  isn't a rule yet.
- **The stability level**: changes only by `/promote` or when the user says so.

Create a doc when there's something to record, not to fill in the structure.
Don't invent product decisions to complete one: ask, or leave it out and say
so. Don't commit docs unless asked.

If a project keeps any of these somewhere else, don't add to the old location.
Offer to move them here, in the format for their type.
