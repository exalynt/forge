# Feature Stability

Features may declare a stability level, Prototype → Alpha → Beta → GA:

- **Prototype** (exploring): explores an idea so we can choose a direction.
  Data may be fake; it may change completely or be thrown away.
- **Alpha** (developing): works, but still taking shape. Frequent breaking
  changes and bugs, and data may be reset.
- **Beta** (refining): settled enough for real use. Occasional breaking changes
  and bugs, data kept, and breaking changes announced when we can.
- **GA** (stable): keeps working as it does today. Breaking changes are rare,
  announced, and come with time to move over.

A level describes how much people can depend on the work, not how good it is.
GA can still have bugs and known gaps.

## Rules

- A level is a quality expectation and an engineering investment ceiling.
  Over-engineering is a stability problem, just like under-engineering. Flag
  both.
- Work beyond the current level needs a concrete reason, such as a request, a
  client's requirement, a law, a real incident, or a security need. "We'll need
  it at GA" isn't one.
- The `stability` skill's checklist is a guideline, not a gate, applied with
  engineering judgment and the current project as context. An outcome that
  applies and isn't met is fixed or written down as a known gap.
- The security and privacy baseline holds at every level and is never a known
  gap: secrets stay out of code, version control, sample data, logs, and
  output; real client data is never copied into code or fixtures; anything
  touching real client systems or data has access control, and one client
  never sees another's; personal data is collected only where needed; and
  untrusted input can't run as code.
- A prototype answers a question. Put realism into the part being learned
  about, and keep everything else simple.
- Before changing a feature, find its level (the `**Stability:**` line in
  `docs/product/capabilities/<slug>.md`, unless the user says otherwise).
  Prototype and Alpha: take the quickest path. Beta: non-breaking unless the
  user directs otherwise. GA: non-breaking only; a breaking change is a new
  version, so clarify with the user first. A shared part takes the strictest
  level among the features using it.
- Never promote or demote a feature unless explicitly instructed. Promotion is
  done with `/promote`.
