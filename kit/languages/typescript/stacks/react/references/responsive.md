# React: mobile responsive with MUI

For React projects that use MUI, adding to the `standards` skill's
[responsive.md](../../standards/references/responsive.md). Where this
project's own code or instructions differ, follow the project.

- Lay out with MUI's `Grid`, `Stack`, and `Container`, and use responsive
  `sx` values with the theme's breakpoints (`{ xs: …, md: … }`).
- Reach for `useMediaQuery(theme.breakpoints.down(…))` only when the
  structure must change, not just its styling.
- Dialogs go `fullScreen` on small screens. Persistent side navigation
  becomes a temporary `Drawer`.
- Don't shrink touch targets below MUI's defaults.
