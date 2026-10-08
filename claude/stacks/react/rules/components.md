---
paths:
  - "src/**/*.tsx"
---

# React projects: reuse components before creating them

General conventions for React projects that use MUI. Where this project's
CLAUDE.md differs, follow it.

## Before creating a component

Look for one that already does the job, in this order:

1. The project's own components: the `components/` directories near the code
   you're working on, then `src/components/`.
2. MUI: `@mui/material`, and any other `@mui/*` package the project already
   depends on.
3. A composition of the above, written inline in the page.

Create a new component only when none of these fits. If an existing project
component almost fits, extend it with a prop rather than writing a
near-duplicate, as long as the change stays small and fits what the component
is for.

## Using MUI

- Style through the theme and the `sx` prop, using theme values for color,
  spacing, and typography rather than hard-coded ones.
- Set app-wide defaults for an MUI component in the theme. Don't wrap an MUI
  component just to rename it or fix its props.
- Ask before adding another UI or component library.
