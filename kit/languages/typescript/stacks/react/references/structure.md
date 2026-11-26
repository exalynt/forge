# React: project structure

For React projects, adding to the `standards` skill's
[structure.md](../../standards/references/structure.md). Where this project's
own code or instructions differ, follow the project.

## Layout

```text
src/
├── components/              # used across the whole app
├── services/acme/           # Acme API calls (see the typescript skill's services)
└── pages/
    └── appointments/
        ├── List.tsx
        ├── View.tsx
        ├── Edit.tsx         # only once it's needed
        └── components/      # used only by the appointments pages
```

## Pages

- Each page area gets a directory under `src/pages/`. Its screens are named
  for what they do: `List.tsx`, `View.tsx`.
- Add `Edit.tsx`, `Create.tsx`, or other screens only when the page needs a
  dedicated one.
- When a screen has several components of its own, give it a directory:
  `src/pages/<page>/List/List.tsx`, with its components in
  `src/pages/<page>/List/components/`.
- Each page renders inside an error boundary, so a page that crashes shows an
  error with a way to reload or go back, and the app's navigation keeps
  working. Use the router's error handling where it has some, otherwise one
  boundary component in `src/components/`. This applies at every stability
  level.

## Where a component lives

Put a component in the `components/` directory closest to everything that uses
it:

- Used by one screen: next to that screen, in `components/` beside it.
- Used by several screens of one page: `src/pages/<page>/components/`.
- Used by several pages: the `components/` directory of their nearest shared
  parent.
- Used across the whole app: `src/components/`.

Don't import from another page's `components/` directory.
