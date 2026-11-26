# Standards: accessibility basics

For every web UI, in any language or framework. Stack skills add to this,
with how their framework does it. Where this project's own code or
instructions differ, follow the project. These apply at every stability level.

## Elements

- A button for an action and a link for navigation. Never a clickable `div`.
- Headings in order, with one `h1` per page, chosen by structure rather than
  size.
- Tables have header cells. Lists of things are lists.

## Labels

- Every input has a visible label. A placeholder isn't a label.
- An icon-only button has an `aria-label`, and usually a tooltip saying the
  same.
- Images have `alt` text. Decorative ones have `alt=""`.

## Keyboard

- Everything works with the keyboard alone, in a sensible order, with a
  visible focus outline. Don't remove the outline.
- Dialogs and menus keep focus inside while open and return it when closed.

## Changes on screen

- A region that's loading has `aria-busy="true"`.
- Errors are announced to screen readers. Changes the user should hear about
  that aren't errors, such as a new result count, go in an
  `aria-live="polite"` region.
- Never use color alone to carry meaning, such as a status. Pair it with text
  or an icon.
