# Standards: light and dark mode

For every web UI, in any language or framework. Stack skills add to this,
with how their framework does it. Where this project's own code or
instructions differ, follow the project.

Every app supports both light and dark mode, and every page and component
looks right in both. This applies at every stability level.

## Theme

- Define both schemes in one theme.
- Follow the system preference by default, and let the user switch with a
  toggle.
- Third-party UI that takes its own appearance settings, such as Clerk or
  Stripe Elements, gets them from the active scheme's palette, so it matches
  the app in both modes.

## Styling

- Take every color from the theme, by its role, such as primary text, paper
  background, divider, and primary. Never hard-code hex, rgb, `white`, or
  `black` in a component.
- When a color the palette doesn't have is needed, add it to both schemes in
  the theme rather than to the component.
- Images, logos, and illustrations need a variant or treatment that holds up
  on both backgrounds.

## Before calling UI work done

Check the changed screens in both light and dark mode.
