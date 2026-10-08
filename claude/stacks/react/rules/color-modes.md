---
paths:
  - "src/**/*.tsx"
  - "src/**/theme/**"
---

# React projects: light and dark mode

General conventions for React projects that use MUI. Where this project's
CLAUDE.md differs, follow it.

Every app supports both light and dark mode, and every page and component
looks right in both. This applies at every stability level.

## Theme setup

- Define both schemes in one MUI theme with `colorSchemes: { light, dark }`
  and `cssVariables` turned on.
- Follow the system preference by default, and let the user switch with a
  toggle built on `useColorScheme()`.
- Third-party UI that takes its own appearance settings, such as Clerk or
  Stripe Elements, gets them from the active scheme's palette, so it matches
  the app in both modes.

## Styling

- Take every color from the theme: palette roles such as `text.primary`,
  `background.paper`, `divider`, and `primary.main`, through `sx` or
  `theme.vars`. Never hard-code hex, rgb, `white`, or `black` in a component.
- When a color the palette doesn't have is needed, add it to both schemes in
  the theme rather than to the component.
- Write mode-specific styles with `theme.applyStyles("dark", { … })`. Don't
  branch on `theme.palette.mode`.
- Images, logos, and illustrations need a variant or treatment that holds up
  on both backgrounds.

## Before calling UI work done

Check the changed screens in both light and dark mode.
