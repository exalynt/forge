# Rules

`docs/product/domain/rules/<slug>.md`, one per area of the domain. Where the
area has a `.feature` file, the rule doc has the same slug.

A rule says what must or can't happen, such as "A block is fulfilled only once
it's approved". Capabilities link to the rules that apply to them, and one rule
can apply to several capabilities.

## Layout

```markdown
# <Area> rules

<The part of the domain these rules govern, in a sentence.>

## <Short name of the rule>

<The rule, in a sentence or two.>

- **Kind:** <Core | Edge case>
- **Why:** <the reason behind it>
- **Acceptance test:** [<file> › <Rule>](<relative path>) | None yet
```

## Sections

One `##` per rule. Capabilities link to these headings, so keep the name stable.

- **Kind**: required. A core rule is one the primary workflow depends on, and an
  edge-case rule covers a boundary, an unusual state, or a failure around one
  (see the testing rule).
- **Why**: when the reason isn't obvious. Leave it out otherwise.
- **Acceptance test**: required. "None yet" says the rule is decided but not
  tested.

## What belongs here

Only decided rules: ones the user stated or confirmed, or that an approved
acceptance test states. A rule inferred from the code or from common sense stays
out until the user confirms it. Something still undecided goes in the
capability's Open decisions.

Acceptance tests remain the specification. Changing a rule that has one means
changing the test, with approval, as the testing rule describes. If a rule doc
and its test disagree, the test holds until the user says otherwise. Report the
disagreement.
