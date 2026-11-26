# Go: errors

For Go projects, adding to the `standards` skill's
[errors.md](../../standards/references/errors.md). Where this project's own
code or instructions differ, follow the project.

- Export sentinel errors from the services package, named for the outcome:
  `var ErrAppointmentNotFound = errors.New("services: appointment not found")`.
  Shared ones include `ErrUnauthenticated` and `ErrForbidden`.
- Services translate repository errors with `errors.Is`, for example
  `repositories.ErrNotFound` becomes `ErrAppointmentNotFound`.
- To add detail without breaking `errors.Is`, wrap both:
  `fmt.Errorf("%w: %w", ErrInvalidInput, err)`. Wrap unexpected failures with
  context.
- Each method's doc comment names the errors it returns.
