# Go HTTP: preferred packages

For Go HTTP APIs, adding to the `go` skill's
[packages.md](../../go/references/packages.md). Where this project's own code
or instructions differ, follow the project.

- `github.com/go-chi/chi/v5`: routing and middleware
- `github.com/go-chi/render`: JSON responses and status codes
- `github.com/go-playground/validator/v10`: request validation through struct
  tags, with one process-wide validator
- `github.com/exalynt/turn`: pagination (see [pagination.md](pagination.md))
- `github.com/alexliesenfeld/health`: health check endpoints (see
  [health.md](health.md))
