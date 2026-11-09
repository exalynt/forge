# Concept

`docs/product/domain/concepts/<slug>.md`, one per concept, such as
`capacity-block.md`.

A concept doc defines something in the domain model: what it is and what the
term means. It's where people and agents go to find out what a word means here.

## Layout

```markdown
# <Concept>

<What it is, in a sentence or two.>

## Details

<What it's made of, what it belongs to, and how it relates to other
[concepts](<slug>.md).>

## Lifecycle

<State> → <State> → <State>

<What moves it from one state to the next.>

## Not to be confused with

- [<Other concept>](<slug>.md): <the difference.>
```

## Sections

The definition under the title is required. Leave out any section with nothing
to say.

A concept doc says what something is. What must or can't happen to it goes in a
rule doc, and what people can do with it goes in a capability doc.
