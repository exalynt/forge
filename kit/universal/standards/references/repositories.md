# Standards: repositories and schema changes

For every project, in any language. Language and stack skills add to this.
Where this project's own code or instructions differ, follow the project.

## Repositories are dumb persistence layers

- No business logic, no authorization, no decisions. They store and retrieve
  domain models.
- Use common names: `Get`, `List`, `Create`, `Update`, `Upsert`, `Delete`, with
  qualifiers where needed (`GetByEmail`). Business verbs belong in services.
- `List` takes a query of filters plus a page selection. A filter left at its
  zero value is not applied.
- Work that spans several writes runs inside the service's transaction. The
  repository doesn't decide on transactions.

## Errors

- Repository errors (not found, conflict, …) are declared with the interface
  they belong to.
- Each implementation translates raw driver or vendor errors into them, such
  as no rows, or unique or foreign-key violations, so services can key off
  them.

## SQL and migrations

- Parameterized SQL. Never build a query by concatenating input.
- Schema changes are migrations, versioned with the code.
- Migrations run from an explicit command or deploy step, never on startup.

## Changing the schema

How to change it depends on the stability level of the capabilities that use
the tables involved: from Docqet when the repo has it, otherwise the level the
user gave. The strictest of them applies.

- **Prototype and Alpha**: take the quickest path. Editing an existing
  migration in place and resetting the database is fine.
- **Beta and GA**: never edit a migration that may already have run. Add a new
  one, and keep it non-breaking for the code already deployed:
  - add columns as nullable or with a default
  - don't drop or rename a column or table the deployed code still uses.
    Expand, then contract: add the new shape, move the code over and backfill,
    and remove the old shape in a later migration
  - don't tighten a constraint until existing rows meet it
- A breaking schema change at Beta needs the user's direction. At GA, clarify
  the direction with the user first.
