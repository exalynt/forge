# Standards: mobile responsive

For every web UI, in any language or framework. Stack skills add to this,
with how their framework does it. Where this project's own code or
instructions differ, follow the project.

Every page and component works on a phone, from 360px wide up through desktop.
This applies at every stability level.

## How

- Start from the small screen and add for larger ones, using the theme's
  breakpoints.
- No fixed widths or heights that can overflow a phone screen, and the page
  never scrolls sideways.
- Change the structure for a screen size only when it must change, not just
  its styling.
- Tables and data grids: on small screens, show rows as a list of cards, or
  hide secondary columns and let the table scroll sideways inside its own
  container.
- Dialogs go full screen on small screens. Persistent side navigation becomes
  a temporary drawer.
- Don't shrink touch targets below the framework's defaults.

## Before calling UI work done

Check the changed screens at phone width (375px) and at desktop width.
