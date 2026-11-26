# React: light and dark mode with MUI

For React projects that use MUI, adding to the `standards` skill's
[color-modes.md](../../standards/references/color-modes.md). Where this
project's own code or instructions differ, follow the project.

## Theme setup

- Define both schemes in one MUI theme with `colorSchemes: { light, dark }`
  and `cssVariables` turned on.
- Build the toggle on `useColorScheme()`.

## Styling

- Take colors from palette roles such as `text.primary`, `background.paper`,
  `divider`, and `primary.main`, through `sx` or `theme.vars`.
- Write mode-specific styles with `theme.applyStyles("dark", { … })`. Don't
  branch on `theme.palette.mode`.
