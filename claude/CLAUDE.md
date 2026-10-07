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
- At every level, one tenant or client cannot see or interact with another
  tenant's or client's real data, and real secrets are never leaked anywhere.
  That is the security and privacy baseline.
- Use the `stability` skill whenever stability materially affects a design,
  implementation, or review decision.

Project-specific architecture, domain rules, and where a repository declares
feature stability belong in that project's own `CLAUDE.md` or `.claude/`.
