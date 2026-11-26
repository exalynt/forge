# Go HTTP: health checks

For Go HTTP APIs, adding to the `standards` skill's
[health.md](../../standards/references/health.md). Where this project's own
code or instructions differ, follow the project.

## Pinging providers

Each implementation that reaches an external service has a
`Ping(ctx context.Context) error`, and its doc comment says what it calls.
Don't add `Ping` to the boundary interface. Where the app is wired, a small
`pinger` interface picks up the implementations that have one.

## How checks run

Use `github.com/alexliesenfeld/health`, behind one `HealthHandler` in the HTTP
handlers package:

- `AddServer(name, check)`, `AddDependency(name, provider, check)`, and
  `AddProvider(name, provider, check)` add checks while the servers are built.
  `provider` is who provides it (`postgres`, `stripe`).
- `Start()` builds the checkers once every check is added, before any server
  serves.

## Telemetry

Filter probe requests out of the HTTP instrumentation with
`otelhttp.WithFilter`.
