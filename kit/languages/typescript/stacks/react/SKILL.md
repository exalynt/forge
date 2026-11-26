---
name: typescript-react
description: Conventions for React and MUI frontends, building on the typescript and standards skills - project structure and error boundaries, reusing MUI and existing components, mobile responsiveness, light and dark mode, accessibility, calling services, loading states and fetching data, search, filters, sorting, and pagination, forms, and CI. Use when writing, changing, or reviewing React or TypeScript code under src/, the theme, or the CI workflow of a React project.
paths:
  - "src/**/*.ts"
  - "src/**/*.tsx"
  - "src/**/theme/**"
  - "src/services/**"
  - "package.json"
  - ".github/**"
---

# React and MUI frontends

Conventions for React projects that use MUI: a stack layer, building on the
`typescript` skill, which builds on the `standards` skill.

Each topic is one file name in every layer. Read the rows for what you're
about to change, left to right: the general rule first, then what TypeScript
adds, then what this stack adds. Where they differ, the more specific one
wins, and this project's own code and instructions win over all of them.

| Topic         | Read when                                                      | Read, in order                                                                                                                                     |
| ------------- | -------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| structure     | Adding pages, screens, or components, or deciding where code lives | [standards](../standards/references/structure.md), [typescript-react](references/structure.md)                                                 |
| components    | Before creating a component, and when using MUI                | [typescript-react](references/components.md)                                                                                                       |
| responsive    | Building or changing any page or component                     | [standards](../standards/references/responsive.md), [typescript-react](references/responsive.md)                                                   |
| color-modes   | Styling anything, or setting up the theme                      | [standards](../standards/references/color-modes.md), [typescript-react](references/color-modes.md)                                                 |
| accessibility | Building or changing any page or component                     | [standards](../standards/references/accessibility.md), [typescript-react](references/accessibility.md)                                             |
| services      | Calling an external API or SDK from the app                    | [standards](../standards/references/layers.md), [typescript](../typescript/references/services.md), [typescript-react](references/services.md)     |
| data          | Fetching data or showing it, including pages with several parts | [typescript-react](references/data.md)                                                                                                            |
| lists         | Building a list with search, filters, sorting, or pagination   | [typescript-react](references/lists.md)                                                                                                            |
| forms         | Building a form, or anything that saves or deletes             | [typescript-react](references/forms.md)                                                                                                            |
| ci            | Adding or changing the CI workflow or the npm scripts it runs  | [standards](../standards/references/ci.md), [typescript](../typescript/references/ci.md)                                                           |
| testing       | Writing or running tests                                       | [testing](../testing/SKILL.md)                                                                                                                     |

Responsiveness, light and dark mode, accessibility, error boundaries, and how
data loads, lists behave, and forms save apply at every stability level. How
much else a capability needs, such as CI, is what its level expects (Docqet's
rubric, when the repo has it; see forge's block, "Where intention lives").
Topics this stack and TypeScript add nothing to are in the `standards` skill's
index.
