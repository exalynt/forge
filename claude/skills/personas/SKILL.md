---
name: personas
description: Find, choose, or interactively create the reusable user personas stored in a project's docs/personas directory. Use when a skill such as /ux or /use needs someone to walk through a product as, or when asked to create, change, or list personas.
---

# User Personas

A persona is a reusable description of a type of person who uses the product.
It says who they are, not what they came to do. The goal, the scenario, and the
test account belong to each run, so the same persona works for any flow.

## Where they live

One file per persona, at `docs/personas/<slug>.md` from the root of the project
being worked on, such as `docs/personas/client-owner.md`. If the folder doesn't
exist, the project has no personas yet. Create it with the first one.

## Format

```markdown
# <Name>, <role>

Who: <their job, and why they use this product, in a sentence or two>
Knows: <domain knowledge, familiarity with this product, technical skill>
Context: <device, how often they do this, time pressure, interruptions>
Expects: <the words they use, the tools they already know, how they think this
kind of work goes>
Cares about: <what makes them trust it, and what would make them give up>
Based on: <the user's answers, project docs, or research; say what's assumed>
```

Names are invented. A persona file is checked in with the project, so it never
holds a real person's details, an account, a password, or any other secret.

## Choosing a persona

Do this at every effort, for each type of user the calling skill needs:

1. Read everything in `docs/personas/`.
2. If the user named a persona or a role and exactly one persona matches, use it
   and say which. Don't ask.
3. Otherwise ask (AskUserQuestion, up to four people per call). Say which type of
   user you need, and that there's no persona for them. Offer the closest
   existing personas and "Create one now". If the project has no personas to
   offer, offer "Create one now" and "Describe one just for this run", which
   isn't saved.

Never make up a persona without asking, or stretch one that doesn't fit.

## Creating one

Create it with the user, before the calling skill goes on:

1. Read the project's docs for who uses the product, and draft an answer for each
   field from them.
2. Ask about the fields (AskUserQuestion, at most two rounds), with your drafted
   answer as the recommended first option: Who and Knows, then Context, Expects,
   and Cares about. Don't ask what the docs settle or what wouldn't change how
   this person behaves.
3. Show the finished persona, with an invented name, and ask whether to save it.
4. Write the file. Don't commit it unless asked, and mention it in the report so
   the user can.

## Changing one

Change or delete a persona only when the user asks. If a run suggests one is
wrong, such as someone who wouldn't really know what it says they know, say so
in the report.

## Using one

Its fields fill in the calling skill's description of the user. An agent that
can't read files, such as `use-role`, gets the file's full text in its prompt.
