# React projects: accessibility basics

General conventions for React projects that use MUI. Where this project's
own instructions differ, follow them. These apply at every stability level.

MUI's components are accessible when used for what they're for. Most of this
is not working against them.

## Elements

- A button for an action and a link for navigation. Never a clickable `div`
  or `Box`.
- Headings in order, with one `h1` per page, chosen by structure rather than
  size. Set the look with `variant` and the level with `component`.
- Tables have header cells. Lists of things are lists.

## Labels

- Every input has a visible label, such as `TextField`'s `label`. A
  placeholder isn't a label.
- An icon-only button has an `aria-label`, and usually a `Tooltip` saying the
  same.
- Images have `alt` text. Decorative ones have `alt=""`.

## Keyboard

- Everything works with the keyboard alone, in a sensible order, with a
  visible focus outline. Don't remove the outline.
- Dialogs and menus keep focus inside while open and return it when closed.
  MUI does this, so don't work around it.

## Changes on screen

- A region that's loading has `aria-busy="true"`.
- Errors use `Alert`, which screen readers announce. Changes the user should
  hear about that aren't errors, such as a new result count, go in an
  `aria-live="polite"` region.
- Never use color alone to carry meaning, such as a status. Pair it with text
  or an icon.
