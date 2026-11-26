# Go: repositories and migrations

For Go projects, adding to the `standards` skill's
[repositories.md](../../standards/references/repositories.md). Where this
project's own code or instructions differ, follow the project.

## Repositories

- `List` takes a query struct of filters plus a page selection.
- A transaction is a unit of work the service carries on `ctx`.
- Repository errors (`ErrNotFound`, `ErrConflict`, …) are declared in the
  interface file, alongside the interface they belong to.
- Each implementation translates raw driver errors into them, such as
  `sql.ErrNoRows` and unique or foreign-key violations.

## SQL and migrations

- Hand-written, parameterized SQL. No ORMs.
- Schema changes are goose SQL migrations (`NNNNN_description.sql`), embedded
  in the binary.
