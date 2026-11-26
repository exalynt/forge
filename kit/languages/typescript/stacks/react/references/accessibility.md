# React: accessibility with MUI

For React projects that use MUI, adding to the `standards` skill's
[accessibility.md](../../standards/references/accessibility.md). Where this
project's own code or instructions differ, follow the project. These apply at
every stability level.

MUI's components are accessible when used for what they're for. Most of this
is not working against them.

- Never a clickable `Box`. Use `Button` for an action and `Link` for
  navigation.
- Set a heading's look with `variant` and its level with `component`.
- Label inputs with the component's own label, such as `TextField`'s `label`.
- An icon-only `IconButton` has an `aria-label`, and usually a `Tooltip`
  saying the same.
- MUI's dialogs and menus keep and return focus, so don't work around them.
- Show errors with `Alert`, which screen readers announce.
