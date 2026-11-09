---
name: engineering-docs
description: Standard location and layout for a repository's engineering docs in docs/engineering (system and feature architecture docs, provider docs, runbooks). Use before changing code that may have documented tradeoffs, gotchas, or technical known gaps; when looking for how something is built or set up; when a change alters how something is built, opens or closes a technical gap, or adds a recovery done by hand; and whenever creating or updating any of these docs.
---

# Engineering Docs

Engineering docs say how software is built, in technical words, for the
engineer or agent about to change it. They give enough context to read the code
with a map and avoid the traps, and they leave the rest to the code. What the
product does is in the product docs (see the `docs` skill), and the two never
repeat each other.

| What you're documenting                                | Where                                                 | Format                     |
| ------------------------------------------------------ | ----------------------------------------------------- | -------------------------- |
| How something overarching works                        | `docs/engineering/architecture/systems/<slug>.md`     | [system.md](system.md)     |
| How one feature works, and what to know before changing it | `docs/engineering/architecture/features/<slug>.md` | [feature.md](feature.md)   |
| A third-party service: setup and how it behaves        | `docs/engineering/providers/<type>/<vendor>.md`       | [provider.md](provider.md) |
| What to do when something goes wrong                   | `docs/engineering/runbooks/<situation>.md`            | [runbook.md](runbook.md)   |

Read the format for a type before writing one. A folder exists once it has a
file in it.

`docs/engineering/README.md` is the map, once there are a few docs to find: one
screen on what the system is made of, how a request or job moves through it,
and a link to each system doc.

## When to write one

Create a doc, or add to one, only when at least one of these is true:

- Understanding something takes several files, and no single file shows it.
- A reasonable engineer would undo a decision without knowing why it was made.
- Something surprising will bite whoever changes it next.
- A technical gap needs recording.
- A provider needs setup outside the repository, or its behavior had to be
  checked by hand.
- A recovery is done by hand.

Don't write one because a feature was built or promoted, or to fill in the
structure. Most plain create-read-update-delete code has nothing to call out, so
most features have no doc.

## What to leave out

The test: could a competent reader learn this in a few minutes from the code,
spec, migrations, or tests? Then leave it out. That covers lists of endpoints,
columns, permissions, statuses, and metrics, and environment variables, which
go wherever the stack's config rule puts them.

- A fact about one function is a code comment, not a doc.
- A business rule is a rule doc. Link it rather than restating it.
- Name the file or symbol to start reading at, without describing what it does
  line by line. Don't use line numbers, which go stale.
- Something true only while a change rolls out goes in its pull request, such as
  deploy order or a one-time command. What rolling back a migration destroys goes
  in a comment in that migration. A constraint that lasts is a gotcha.
- Behavior checked outside the code, such as how a vendor's page acts, says when
  it was checked.
- Don't invent a reason, a cost, or a trigger to complete an entry. Ask, or leave
  it out and say so.
- Docs are checked in, so they never hold a real person's details, an account, a
  password, or any other secret.

## System or feature

- **System**: several features or repositories rely on it, and you need the
  whole picture to change it safely. Expect a handful per repository, such as
  `authentication`, `authorization`, `tenant-isolation`, or `telemetry`.
- **Feature**: only someone working on that feature needs it. A feature doc can
  be a high-level map, a few gotchas, or both.

Put a fact where the person who'd trip over it will be looking. If it bites
while working on one feature, it's in that feature's doc. If it can bite from
anywhere, it's in the system doc. When a second feature needs the same fact, move
it to a system doc and link to it.

## Naming

- **Systems**: the mechanism or system, in the code's words.
- **Features**: what the code calls it. Scope follows the code, not the
  capability, so one `bookings.md` can serve several capabilities that share its
  code. A slug matches a capability only when the code calls it that too.
- **Providers**: type and vendor, such as `payments/stripe.md`.
- **Runbooks**: the situation, starting with a verb, such as
  `refund-a-booking.md`.
- All of them are kebab-case.

## Tradeoffs, Gotchas, and Known gaps

System, feature, and provider docs share these sections, with these exact names.

| Section        | The question it answers                                         |
| -------------- | --------------------------------------------------------------- |
| **Tradeoffs**  | What did we choose, what does it cost, and why is that worth it? |
| **Gotchas**    | What's correct but will surprise whoever changes it?            |
| **Known gaps** | What isn't done yet, and what closes it?                        |

The test between a gap and a tradeoff is whether we intend to close it. A gap
that will exist forever is a tradeoff, so it records its cost instead.

```markdown
## Tradeoffs

- **<What we chose>, not <the alternative>:** <why>. **Costs:** <what we accept>.
  **Revisit when:** <a concrete trigger>.

## Gotchas

- <What's true, and what to do about it.>

## Known gaps

- **<What's missing>:** <what happens today, and how it's handled>.
  **Until:** <what closes it>.
```

- **Tradeoffs** are only for real choices with an alternative. A decision nobody
  would undo doesn't need an entry. `Costs` is required, and links to the
  capability's gap when people can see the cost, rather than rewording it.
  `Revisit when` is left off one we expect to keep forever.
- **Until** is required on every gap, and is one of:
  - a capability reaching a level, linking its doc:
    `[Bookings](../../../product/capabilities/bookings.md) reaches Beta`
  - a concrete event: `the second instance is deployed`
  - `not decided`, so nobody invents a trigger
- **These docs never state a stability level of their own.** Levels belong to
  capabilities, and one doc can serve capabilities at different levels, so a
  level always names its capability.
- **A gap is due** once its capability has reached the level in `Until`, or is
  being promoted to it. `/promote` and the `stability-reviewer` agent find these
  by searching `docs/engineering/` for the capability doc's filename. A gap that
  isn't due is left alone by other work and isn't reported as a problem.
- **Neither a gap nor a tradeoff can excuse the security and privacy baseline.**
  Tenant isolation and secrets can't be accepted either way.

## Technical gaps and product gaps

Each gap is in exactly one place, decided by who can see it:

- If someone using the product would notice it, it's in the capability doc's
  Known gaps, in their words.
- If only someone working on the code would, it's here, in technical words. That
  includes paging, load testing, test coverage, caching, how a deploy behaves,
  and stubs people can't see.

When a gap has a visible effect and a technical cause, the effect is the gap and
lives in the capability doc. The engineering side covers only how it works, as a
gotcha or a runbook, without restating it as a gap. For example:

- **Capability, Known gaps:** "A customer can be charged without getting their
  booking until someone fixes it by hand."
- **Feature doc, Gotchas:** "The charge happens before the booking is written, and
  webhooks for a charge with no booking are ignored."
- **Runbook:** `reconcile-a-charge-with-no-booking.md`.

Don't add a section that only points to the capability's gaps. `None.` under
Known gaps here means no technical gaps, and says nothing about the product side.

## Several repositories

Each repository keeps its own `docs/engineering/`, because a fact lives in the
repository whose code would change it, so the doc updates in the same change.

- **A system or feature spanning repositories** has one whole-picture doc, in
  the repository that enforces or defines it. It covers the flow and the
  contract, including what every other repository must do.
- **The other repositories document only what's theirs**, under the same slug,
  opening with a link to the owner's doc. Most won't need one.
- **Each repository's `CLAUDE.md`** says where the owner's docs are, as it does
  for product docs. Those links are relative to a sibling checkout, such as
  `../api/docs/engineering/architecture/systems/authentication.md`.

## Links

Links are relative and point one way, so there's one place to update:

```text
runbooks → features, systems, providers
features → systems, providers
everything here → product docs
product docs → nothing here
another repository's doc on the same thing → the owner's doc
```

To go the other way, search for the file's name.

## How much each stability level needs

The level of the capabilities a doc serves sets what's required. It's never a
reason on its own to create a doc.

| Level     | Required                                                                              |
| --------- | ------------------------------------------------------------------------------------- |
| Prototype | Nothing, unless a stub `/qa` should allow has to be recorded                          |
| Alpha     | Gotchas, technical gaps, and provider setup done outside the repository, where they exist |
| Beta      | Also how shared systems work, meaningful tradeoffs, and a runbook for each recovery done by hand |
| GA        | Also complete enough for an operator who didn't build it                              |

Where a doc serves capabilities at different levels, the strictest applies.

## Keeping them current

Update engineering docs in the same change as what they describe:

- **How something is built changed**: fix what it made wrong. If a doc would now
  be worth having, as above, create it.
- **A technical gap opened**: add it with its `Until`.
- **A gap closed**: delete it. Git has the history.
- **We decided never to close a gap**: move it to Tradeoffs with its cost. Only
  the user decides this.
- **A tradeoff was reversed**: delete it, or replace it with the new one.
- **A new recovery is done by hand**: add a runbook.

If a doc and the code disagree, find out which is wrong before changing either.
Keep system docs under about 150 lines and feature docs under about 75, and split
one that grows past that. Don't commit docs unless asked.

If a project keeps any of this somewhere else, or has docs that fill in the same
sections for every feature, don't extend the old shape. Offer to move it here, in
the format for its type.
