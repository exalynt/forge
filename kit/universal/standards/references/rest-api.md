# Standards: REST APIs

For every HTTP API, in any language. Language and stack skills add to this.
Where this project's own code or instructions differ, follow the project.

## Paths

- Resource-based, plural, kebab-case: `/appointments`, `/payment-methods/{id}`.
- An action that isn't natural CRUD on a resource gets an action sub-path:
  `POST /appointments/{id}/reschedule`.
- Version the API under `/v1`.

## Methods

- `GET` reads, and `POST` on a collection creates.
- `PUT /things/{id}` **creates or replaces** the whole resource at that ID.
- `PATCH` **partially updates**: only the fields sent change.
- `DELETE` removes.

Use the method whose semantics match. Never use `PUT` for a partial update.

## IDs

Clients may supply a resource's primary ID (a UUID). When it's omitted, the
server generates one (UUIDv7).

## Bodies

- Responses contain only the resources read, created, or modified. No
  envelopes, no metadata.
- A list is a bare JSON array. Pagination and other metadata go in headers
  (`Link`).
- Never serialize raw models. Each endpoint returns a resource translated from
  the domain model.
- Requests decode into request types that are validated, and reject unknown
  fields.

## Pagination

Every list endpoint is paginated, with pagination information in headers,
never the body. See [pagination.md](pagination.md).

## Status codes and errors

- `200` for reads and updates, `201` for creates, `204` for deletes.
- `400` for invalid input, `401` unauthenticated, `403` forbidden, `404` not
  found (including out of scope), `409` conflict, and `503` when a dependency
  is unconfigured or unreachable.
- Every error has the same body: `{"error": "<human-readable message>"}`.
- 5xx responses never expose internals.

## Changing endpoints

How to change an endpoint depends on the stability level of the capabilities
it serves: from Docqet when the repo has it, otherwise the level the user
gave. The strictest of them applies.

- **Prototype and Alpha**: change, rename, or remove endpoints and fields in
  place.
- **Beta and GA**: keep changes non-breaking. Adding endpoints, optional request
  fields, and response fields is fine. Removing or renaming endpoints or
  fields, changing a field's type or meaning, making an optional field
  required, tightening validation, or changing status codes is breaking.
- A breaking change at Beta needs the user's direction. At GA, clarify the
  direction with the user first, for example a new endpoint alongside the old
  one or a new version under `/v2`.
