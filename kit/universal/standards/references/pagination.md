# Standards: pagination

For every HTTP API, in any language. Language and stack skills add to this,
with the library that implements it. Where this project's own code or
instructions differ, follow the project.

## Every list endpoint is paginated

- Every endpoint that returns a collection is paginated from its first
  version, at every stability level. Never "return everything for now".
- Use the pagination library the language or stack skill names. Don't
  hand-roll limits, offsets, or cursors.
- Nothing matching is `200` with an empty array, `[]`. Never `null`, and never
  `404`.

## Pagination is never in the body

- The body is a bare JSON array of resources. No envelope, and no `data`,
  `items`, `next`, `cursor`, `has_more`, `page`, or `total` fields.
- Pagination information goes in standard headers:
  - `Link` ([RFC 8288](https://www.rfc-editor.org/rfc/rfc8288)), with
    `first`, `prev`, `next`, and `last` links, as far as each exists. Each
    link keeps the request's other query parameters, such as filters and
    `sort`. With no other pages, there is no `Link` header.
  - `X-Total-Count`, how many items match across every page, only on a
    listing whose clients need a total, such as a table that shows "42
    matching". It costs a `COUNT` query, so don't add it to every listing.
    Like the pages, the count can drift as items change between requests.
- Add both headers to CORS's exposed headers
  (`Access-Control-Expose-Headers`), or browser clients can't read them.

## Cursors by default

- Use cursor pagination, with the `after`, `before`, and `size` query
  parameters. Use numbered pages (`page`, `size`) only for a small or bounded
  listing where a UI needs to jump to page _N_.
- 25 per page by default, 100 at most. A size out of range, a malformed
  cursor, or `after` and `before` together is a `400`, never clamped or
  silently restarted from the first page. Its message says what was wrong,
  without echoing the cursor back.
- Cursors are opaque. Clients take them from the `Link` header and never build
  them.
- A `sort` query parameter names a field, with a leading `-` for descending.
  An unknown field is a `400`. Every order ends with a unique key, such as the
  ID, so the order is total and no item is skipped or repeated between pages.
- Keyset columns are non-null.
- Bind each cursor to its listing and its filters, so a cursor replayed
  against another listing or other filters is rejected, not a wrong page. The
  caller's authorization bounds aren't part of that. They aren't a filter the
  client chose.

## Through the layers

- **Entry points** parse the page selection, call the service, set the
  headers, and render the page's items as resources.
- **Services** take the selection alongside the listing's filter, pass it to
  the repository, and return the page, plus the total when the listing counts.
  An invalid selection becomes the service's own invalid-page error, a `400`.
- **Repositories** own the query and its pagination. The base query, filters,
  and tenant scoping stay the repository's own SQL. When the listing counts,
  the count takes the same query, ignoring the page and sort.

## Documenting

- In the OpenAPI spec, list endpoints use shared `after`, `before`, and `size`
  parameters and a shared `Link` response header, plus `X-Total-Count` where
  the listing has one.
- The endpoint's description says to follow the `Link` header's `next` link
  for the rest.

## Changing a listing

Changing pagination follows [rest-api.md](rest-api.md)'s "Changing
endpoints". At Beta and GA, adding pagination to a listing that returned
everything is breaking, since existing clients would only see the first page.
