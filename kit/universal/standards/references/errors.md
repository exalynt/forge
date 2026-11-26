# Standards: errors

For every project, in any language. Language and stack skills add to this.
Where this project's own code or instructions differ, follow the project.

## Services return explicit errors

- Services declare their errors, named for the outcome, such as "appointment
  not found". Shared ones cover cross-cutting outcomes, such as
  unauthenticated and forbidden.
- Services translate repository and vendor errors into their own, for example
  a repository's not-found becomes "appointment not found". Entry points never
  key off repository or vendor errors.
- Something outside the caller's scope returns the same not-found error as
  something missing.
- Adding detail to an error never hides which outcome it is. Unexpected
  failures are wrapped with context, and become 500s.
- Each operation's documentation names the errors it returns.

## Entry points key off them

- Each entry point maps each service error to a status and a client-safe
  message, in one place per resource.
- Anything unmapped is a 500 with a generic message, and the real cause is
  logged.
- Never match on error strings.
