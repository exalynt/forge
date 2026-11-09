# Capability

`docs/product/capabilities/<slug>.md`, one per capability.

A capability is something people can accomplish with the product, not a page, a
button, or an endpoint. It's the thing that has a stability level, so it's the
size of a feature, such as "Capacity purchasing" or "Fulfillment review". Where
other skills say "feature", this doc is that feature's declaration. If parts of
one are really at different levels, split it.

## Layout

```markdown
# <Name>

**Stability:** <Prototype | Alpha | Beta | GA>

<What people can accomplish with this and why it exists, in a sentence or two.>

## What people can do

### <Role>

- <Something they can accomplish, starting with a verb, in their words.>

## Learning

- <A question this prototype exists to answer.>

## Known gaps

- <What's intentionally incomplete>: <how it's handled for now.>

## Open decisions

- **<Topic>:** <the question, and the options if there are any yet.>

## Rules

- [<Area>](../domain/rules/<slug>.md)

## Concepts

- [<Concept>](../domain/concepts/<slug>.md)
```

## Sections

- **Stability**: required, exactly one of the four levels. This line is the
  capability's declared level. `grep '^\*\*Stability' docs/product/capabilities/*`
  lists them all.
- **What people can do**: required, with one heading per role.
- **Learning**: only at Prototype. Leave it out otherwise.
- **Known gaps**: required. Write "None." when there are none. Only what someone
  using the product would notice, in their words, so never paging, load testing,
  test coverage, or how a deploy behaves. A gap only someone working on the code
  would see goes in the engineering docs, and a gap is never in both (see the
  `engineering-docs` skill). What's listed here is left alone by other work and
  isn't reported as a problem by `/qa`, `/ux`, or `/use`. A mock or stub that's
  allowed and that people can see, such as placeholder data, is named here.
- **Open decisions**: questions nobody has settled. Nothing here is a rule.
  Once one is decided, it moves to a rule doc. Leave the section out if there
  are none.
- **Rules** and **Concepts**: links only, never the text of the rule or
  definition. Link a rule's own heading when only part of a rule doc applies.
  Leave them out if there are none.
