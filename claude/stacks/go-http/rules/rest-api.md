---
paths:
  - "**/handlers/**/*.go"
  - "**/http/**/*.go"
  - "**/api/**/*.go"
  - "**/server/**/*.go"
  - "**/middleware/**/*.go"
  - "**/*handler*.go"
  - "**/routes*.go"
---

# Go HTTP projects: REST APIs

General conventions for Go HTTP API projects. Where this project's own code or
CLAUDE.md differs, follow the project.

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
- Never serialize raw models. Each endpoint returns a translated `XResource`,
  built by `NewXResource(model)`.
- Requests decode into `XRequest` types with `validate` tags and reject unknown
  fields.

## Pagination

Prefer cursor pagination (`after`, `before`, `size`) through `exalynt/turn`,
with a `Link` header.

## Status codes and errors

- `200` for reads and updates, `201` for creates, `204` for deletes.
- `400` for invalid input, `401` unauthenticated, `403` forbidden, `404` not
  found (including out of scope), `409` conflict, and `503` when a dependency
  is unconfigured or unreachable.
- Every error has the same body: `{"error": "<human-readable message>"}`.
- 5xx responses never expose internals.
