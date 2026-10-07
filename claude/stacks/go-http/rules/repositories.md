---
paths:
  - "**/repositories/**"
  - "**/repository/**"
  - "**/store/**"
  - "**/migrations/**"
  - "**/*.sql"
---

# Go HTTP projects: repositories and migrations

General conventions for Go HTTP API projects. Where this project's own code or
CLAUDE.md differs, follow the project.

## Repositories are dumb persistence layers

- No business logic, no authorization, no decisions. They store and retrieve
  domain models.
- Use common names: `Get`, `List`, `Create`, `Update`, `Upsert`, `Delete`, with
  qualifiers where needed (`GetByEmail`). Business verbs belong in services.
- `List` takes a query struct of filters plus a page selection. A filter left
  at its zero value is not applied.
- Work that spans several writes runs inside the service's transaction (a unit
  of work carried on `ctx`). The repository doesn't decide on transactions.

## Errors

- Repository errors (`ErrNotFound`, `ErrConflict`, …) are declared in the
  interface file, alongside the interface they belong to.
- Each implementation translates raw driver or vendor errors into them, such
  as `sql.ErrNoRows` and unique or foreign-key violations, so services can key
  off them.

## SQL and migrations

- Hand-written, parameterized SQL. No ORMs.
- Schema changes are goose SQL migrations (`NNNNN_description.sql`), embedded
  in the binary.
- Never edit a migration that may already have run. Add a new one.
- Migrations run from an explicit command or deploy step, never on startup.
