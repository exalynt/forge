---
paths:
  - "src/**/*.ts"
  - "src/**/*.tsx"
---

# React projects: project structure

General conventions for React projects. Where this project's CLAUDE.md
differs, follow it. Put new code in this layout. Don't move existing code into
it unless asked.

## Layout

```text
src/
├── components/              # used across the whole app
├── services/exalynt/        # Exalynt API calls (see the services rule)
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
  dedicated one. Don't create empty or placeholder screens.
- When a screen has several components of its own, give it a directory:
  `src/pages/<page>/List/List.tsx`, with its components in
  `src/pages/<page>/List/components/`.

## Where a component lives

Put a component in the `components/` directory closest to everything that uses
it:

- Used by one screen: next to that screen, in `components/` beside it.
- Used by several screens of one page: `src/pages/<page>/components/`.
- Used by several pages: the `components/` directory of their nearest shared
  parent.
- Used across the whole app: `src/components/`.

When a second page starts using a component, move it up to the level that
covers both. Don't import from another page's `components/` directory.
