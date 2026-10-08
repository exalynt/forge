---
paths:
  - "src/**/*.tsx"
---

# React projects: mobile responsive

General conventions for React projects that use MUI. Where this project's
CLAUDE.md differs, follow it.

Every page and component works on a phone, from 360px wide up through desktop.
This applies at every stability level.

## How

- Lay out with MUI's `Grid`, `Stack`, and `Container`, and use responsive
  `sx` values with the theme's breakpoints (`{ xs: …, md: … }`). Start from
  the small screen and add for larger ones.
- No fixed widths or heights that can overflow a phone screen, and the page
  never scrolls sideways.
- Reach for `useMediaQuery(theme.breakpoints.down(…))` only when the
  structure must change, not just its styling.
- Tables and data grids: on small screens, show rows as a list of cards, or
  hide secondary columns and let the table scroll sideways inside its own
  container.
- Dialogs go `fullScreen` on small screens. Persistent side navigation
  becomes a temporary `Drawer`.
- Don't shrink touch targets below MUI's defaults.

## Before calling UI work done

Check the changed screens at phone width (375px) and at desktop width.
