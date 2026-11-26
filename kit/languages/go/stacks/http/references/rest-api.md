# Go HTTP: REST APIs

For Go HTTP APIs, adding to the `standards` skill's
[rest-api.md](../../standards/references/rest-api.md). Where this project's own
code or instructions differ, follow the project.

- Each endpoint returns a translated `XResource`, built by
  `NewXResource(model)`, rendered with `github.com/go-chi/render`.
- Requests decode into `XRequest` types with `validate` tags, and reject
  unknown fields.
- Every list endpoint is paginated through `github.com/exalynt/turn`. See
  [pagination.md](pagination.md).
