# Go HTTP: mapping errors in handlers

For Go HTTP APIs, adding to the `go` skill's
[errors.md](../../go/references/errors.md). Where this project's own code or
instructions differ, follow the project.

- Each resource has a `writeXError(w, r, msg, err)` that uses an `errors.Is`
  switch to map each sentinel to a status and a client-safe message.
- The default case is a 500 with a generic message, and the real cause is
  logged.
