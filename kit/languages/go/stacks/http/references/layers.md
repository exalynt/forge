# Go HTTP: handlers

For Go HTTP APIs, adding to the `go` skill's
[layers.md](../../go/references/layers.md). Where this project's own code or
instructions differ, follow the project.

Handlers are the entry points. Each decodes the request into an `XRequest`,
validates it with the process-wide validator, calls a service, maps its error
(see [errors.md](errors.md)), and renders an `XResource` with
`github.com/go-chi/render` (see [rest-api.md](rest-api.md)). No business
rules.
