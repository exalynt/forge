# Flow

`docs/product/flows/<slug>.md`, one per flow.

A flow is how a piece of work gets done from start to finish, usually across
several capabilities or several people, such as "Client onboarding". What one
person can do on their own in one capability doesn't need a flow doc.

## Layout

```markdown
# <Name>

<What this flow gets done, and for whom, in a sentence or two.>

- **People:** [<Persona>](../personas/<slug>.md), <role with no persona>
- **Starts when:** <what prompts it>
- **Done when:** <the observable end state, across everyone>

## Steps

1. **<Role>** <does what, or what happens, in the domain's words.>

## Variations

- **<Situation>:** <how the flow goes differently, and how it ends.>

## Capabilities

- [<Capability>](../capabilities/<slug>.md)
```

## Sections

- **People**, **Starts when**, **Done when**: all required. Link a persona where
  there's one that fits.
- **Steps**: required. Say what happens, never which page, button, or endpoint.
  `/ux` and `/use` read this to know what should happen, and the people they
  play still have to find their own way.
- **Variations**: the ways it commonly goes differently. Leave it out if there
  are none.
- **Capabilities**: required, every capability the flow crosses. Their docs
  carry the stability levels and known gaps that apply to it.
