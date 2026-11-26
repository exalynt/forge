---
name: react
description: Conventions for React and MUI frontends - project structure, reusing MUI and existing components, mobile responsiveness, light and dark mode, services for external APIs and SDKs, and CI. Use when writing, changing, or reviewing React or TypeScript code under src/, the theme, or the CI workflow of a React project.
paths:
  - "src/**/*.ts"
  - "src/**/*.tsx"
  - "src/**/theme/**"
  - "src/services/**"
  - "package.json"
  - ".github/**"
---

# React projects

General conventions for React projects that use MUI. Where this project's own
instructions differ, follow them. Read the reference for what you're about to
change, from this skill's `references/` directory:

| Reference                                   | Read when                                                              |
| ------------------------------------------- | ---------------------------------------------------------------------- |
| [structure.md](references/structure.md)     | Adding pages, screens, or components, or deciding where code lives     |
| [components.md](references/components.md)   | Before creating a component, and when using MUI                        |
| [responsive.md](references/responsive.md)   | Building or changing any page or component                             |
| [color-modes.md](references/color-modes.md) | Styling anything, or setting up the theme                              |
| [services.md](references/services.md)       | Calling an external API or SDK from the app                            |
| [ci.md](references/ci.md)                   | Adding or changing the CI workflow or the npm scripts it runs          |

Responsiveness and light and dark mode apply at every stability level. How
much else a capability needs, such as CI, is what its level expects (Docket's
rubric, when the repo has it; see forge's `AGENTS.md`, "Where intention
lives"). The `testing` and `observability` skills hold the practice that
applies to every stack.
