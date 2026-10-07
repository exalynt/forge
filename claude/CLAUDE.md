# Engineering Philosophy

Build only as much software as we currently have reason to build.

Prefer simple, concrete, reversible implementations over speculative
abstraction, hardening, extensibility, or future-proofing. More engineering is
not automatically better engineering.

## Feature stability

Features may declare a stability level: Prototype → Alpha → Beta → GA.

When a feature's level is known, treat it as both:

1. a quality expectation the feature must meet, and
2. a ceiling on how much engineering to invest right now.

- Do not harden, generalize, productionize, scale, or future-proof a feature
  beyond its current level.
- Never change a feature's stability level unless explicitly instructed.
- Security and privacy requirements apply at every level.
- Use the `stability` skill whenever stability materially affects a design,
  implementation, or review decision.

Project-specific architecture, domain rules, and where a repository declares
feature stability belong in that project's own `CLAUDE.md` or `.claude/`.
