# Go HTTP: pagination with turn

For Go HTTP APIs, adding to the `standards` skill's
[pagination.md](../../standards/references/pagination.md). Where this
project's own code or instructions differ, follow the project.

Pagination goes through `github.com/exalynt/turn`, whose defaults are the
standard's page sizes.

**Handlers** parse the selection, call the service, and set the headers:

- A shared `cursorSelector(w, r)` parses the selection with
  `turnhttp.CursorQuery`, writing the `400` itself when it's invalid.
- After the service returns the page, a shared `setCursorLinks(w, r, page)`
  sets `Link` with `turnhttp.FormatLinks(turnhttp.CursorLinks(page,
  cursorQuery.URL(r.URL))...)`, then the handler sets `X-Total-Count` if the
  listing has one, and renders `page.Items` as resources.
- `ErrInvalidPage` maps to `400`.

**Services** take a `cursor.Selector` alongside the listing's filter, pass it
into the repository's query, and return the `cursor.Page[models.X]`, plus the
total when the listing counts. They translate turn's selector errors into
their own: when `paginator.IsSelectorError(err)`, return
`fmt.Errorf("%w: %w", ErrInvalidPage, err)`. Any other error is a `500`.

**Repositories** own the query and its pagination:

- `List(ctx, query XQuery) (cursor.Page[models.X], error)`, where `XQuery`
  holds the filters, `Sort`, and `Page cursor.Selector`. When the listing
  counts, `Count(ctx, query)` takes the same query.
- Build the paginator once per listing and order, as a package-level value,
  with `store.NewCursor` and `.Reverse()` for the descending order. Render
  the keyset `WHERE`, `ORDER BY`, and `LIMIT` with `turnsql.Render`.
- Bind each cursor to its listing by passing the listing's name and the
  query's filters as scope values. Leave the caller's authorization bounds out
  of the scope (`json:"-"`).
