# Testing

The three kinds of tests, in any language or framework: what each covers, how
fine-grained it is, and the stability level that first requires it. A stack's
testing rule says which tools and directories a kind of project uses for each.

## When to write tests

Write tests when the user asks, or when the feature's stability level requires
them, and only the kinds and depth asked for. Don't add tests on your own
initiative otherwise. Existing tests must still pass after any change.

## Three kinds

| Kind        | Granularity                                   | Dependencies                       | Required from |
| ----------- | --------------------------------------------- | ---------------------------------- | ------------- |
| Unit        | Lowest: one function, method, or component    | All mocked or stubbed              | Beta          |
| Integration | Middle: one piece of code and a real system   | Real database or external system   | Beta          |
| Feature     | Highest: a whole flow, through the public API | The running app, all of them real  | Alpha         |

### Unit

- Tests one unit in isolation: a function, a method, a component.
- Everything outside the unit is replaced with mocks or stubs: repositories,
  clients, providers, and the clock. No database, network, or filesystem.
  Mock the boundaries, never the code under test.
- Covers logic in detail: business rules, validation, branches, error mapping,
  edge cases, and failure paths. This is the cheapest place to test them, so
  most edge cases belong here.
- Fast enough to run on every change.

### Integration

- Tests one piece of code against the real system it talks to: a repository
  against a real database, a client against a provider's sandbox or a local
  emulator, a consumer against a real queue.
- Any test that needs a database or another external system is an integration
  test.
- Covers that the code and the system agree: queries, migrations,
  constraints, transactions, serialization, and the failures only the real
  system produces, such as a unique-constraint violation.
- Not business logic through every layer. That's unit tests below it and
  feature tests above it.

### Feature

- Black box. Drives the running app only through its public interface, the
  API over HTTP or the UI in a browser, exactly as a caller would, with its
  real dependencies.
- Exercises a whole flow from start to finish, such as signing up, then
  signing in, then booking.
- Asserts only on what a caller can see: responses, the UI, and the results of
  follow-up calls. Never on database rows or internals.
- Mostly happy paths: the main workflows working end to end. Edge cases belong
  in unit tests.
- Third-party services run in their sandbox or test mode, not mocked.

## Where a case goes

Test each case at the lowest level that can prove it, and only once:

- A rule or an edge case → unit.
- Behavior that depends on the real database or system → integration.
- A workflow a user or client depends on working → feature.

## By stability level

- **Prototype**: no automated tests. The primary demo path works when
  exercised by hand.
- **Alpha**: feature tests for the common paths. That's all Alpha needs.
- **Beta**: adds unit and integration tests where they're useful. Primary and
  important secondary workflows are tested, and so are important failure
  paths.
- **GA**: important workflows are comprehensively tested, and important
  failure modes have tests that pin their behavior.

CI runs each kind in its own job from the level that requires it: feature
tests from Alpha, unit and integration tests from Beta.

## Suites run in isolation and concurrently

The unit, integration, and feature suites are independent, so CI can run their
jobs at the same time:

- Each suite runs with its own command and starts what it needs itself. It
  never relies on another suite having run first, or on anything another
  suite leaves behind.
- Suites never interfere when they run at the same time, even against the same
  database or service. None of them resets, truncates, or reseeds shared
  state another suite might be using. Each works only with its own data, or
  with its own database or schema.
- No CI job waits on another test job, unless one needs an artifact the other
  builds, such as the app's image.

## Test data and dependencies

- Integration and feature tests get their dependencies from the project's
  local stack, such as Docker Compose, through the same environment variables
  as local runs and CI.
- Each test creates and isolates its own data, so the suite can run
  repeatedly and in any order.
