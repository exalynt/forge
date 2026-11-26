# Go: layers and boundaries

For Go projects, adding to the `standards` skill's
[layers.md](../../standards/references/layers.md). Where this project's own
code or instructions differ, follow the project.

## Services

Use one `XService` per area, built with `NewXService(...)` from interfaces,
with `ctx` as the first parameter of every method.

## Boundary packages

- A boundary is a package named for the concept, defining the interface:
  `repositories`, `email`, `payments`, `storage`.
- Implementations live in subdirectories named for the technology or vendor:
  `repositories/postgres`, `email/smtp`, `email/noop`, `payments/stripe`.
  Only that subdirectory imports the vendor's SDK or driver.
- Each implementation asserts that it satisfies the interface:
  `var _ email.Sender = (*Sender)(nil)`.
