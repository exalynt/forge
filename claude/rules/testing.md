# Testing

The three levels of tests, in any language or framework: what each covers, how
fine-grained it is, and the stability level that first requires it. Then
acceptance tests, which test the domain's business rules at whichever of those
levels proves them. A stack's testing rule says which tools and directories a
kind of project uses for each.

## When to write tests

Write tests when the user asks, or when the feature's stability level requires
them, and only the kinds and depth asked for. Don't add tests on your own
initiative otherwise. Existing tests must still pass after any change.

## Three levels

| Level       | Granularity                                   | Dependencies                       | Required from |
| ----------- | --------------------------------------------- | ---------------------------------- | ------------- |
| Unit        | Lowest: one function, method, or component    | All mocked or stubbed              | Beta          |
| Integration | Middle: one piece of code and a real system   | Real database or external system   | Beta          |
| End-to-end  | Highest: a whole flow, through the public API | The running app, all of them real  | Alpha         |

### Unit

- Tests one unit in isolation: a function, a method, a component.
- Everything outside the unit is replaced with mocks or stubs: repositories,
  clients, providers, and the clock. No database, network, or filesystem.
  Mock the boundaries, never the code under test.
- Covers logic in detail: business rules (as acceptance tests, below),
  validation, branches, error mapping, edge cases, and failure paths. This is
  the cheapest place to test them, so most edge cases belong here.
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
  end-to-end tests above it.

### End-to-end

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

## Acceptance tests

An acceptance test checks a business rule of the domain: what the domain does,
never how the code does it. "A room can't be double-booked", "a trial ends
after 14 days", "only an organization's owner can remove a member". Together,
a feature's acceptance tests are its specification: the agreed statement of
the software working as it should.

- Not a fourth level. An acceptance test lives at the lowest level that can
  prove its rule. Most test the domain or service layer with its boundaries
  mocked, as a unit test would. A rule that only the real system enforces,
  such as uniqueness through a database constraint, needs real dependencies.
  A rule only visible across a whole flow goes through the public interface.
- Its own suite and CI job. Acceptance tests are written with one tool,
  whatever level each one sits at, so they run together, apart from the
  unit, integration, and end-to-end suites. The suite starts whatever its
  tests need.
- What, not how. A scenario reads in the domain's language and names no
  endpoints, status codes, selectors, tables, functions, or mocks. If the
  implementation changes and the rule doesn't, the scenario doesn't change.
  Only the step definitions behind it do.
- Not everything is an acceptance test. Behavior with no meaning in the
  domain, such as serialization, query correctness, error mapping, HTTP
  wiring, retries, or how a component renders, gets an ordinary unit,
  integration, or end-to-end test.
- Required for core business rules from Alpha, and for edge-case business
  rules from Beta. A core rule is one the primary workflow depends on: break
  it and the feature doesn't do its job. An edge-case rule covers a boundary,
  an unusual state, or a failure around a core rule.

### When they change

Acceptance tests follow the business rules, not the code:

- From Alpha, implementing a business rule, changing one, or clarifying one,
  whether the user answers a question or QA settles an ambiguity, means
  adding or updating its acceptance tests, with approval (below), whether
  it's a core rule or an edge case. Rules settled this way are recorded as
  acceptance tests, not only in docs or comments. The level's requirements
  decide which rules nobody is touching must already be covered.
- Bug fixes, performance work, refactors, and presentation changes don't
  change acceptance tests. They restore or keep behavior the specification
  already describes. A bug in a rule that has no acceptance test is a sign
  the rule is uncovered: propose one if the level requires it, as its own
  change.

### Changing them takes explicit approval

Acceptance tests are the specification, so they change only with the user's
explicit approval, asked for and given for that change:

- Adding, changing, or removing a scenario, example, or `.feature` file.
  Propose the Gherkin, say why, and wait for a yes before writing it. A
  stability level or a task requiring acceptance tests isn't approval of
  particular scenarios.
- Changing a step definition in a way that changes what its scenarios check.
  Step definitions may follow the code through refactors without approval as
  long as every scenario still checks exactly what it says.
- When an acceptance test fails, the code is wrong until the user says the
  specification is. Fix the code, or ask. Never edit, skip, tag out, or
  delete an acceptance test to get it passing.
- Approval covers the change it was given for, not later ones.

### Cucumber and Gherkin

Write acceptance tests in [Gherkin](https://cucumber.io/docs/gherkin/reference)
and run them with Cucumber, unless the project already does it another way. A
stack's testing rule names the Cucumber implementation and where the files go.

- One `.feature` file per feature or area of the domain. `Feature:` names it,
  a `Rule:` states each business rule, and the `Example:`s under a rule
  illustrate it with concrete cases, edge cases included.
- Steps are declarative: `Given` the context, `When` the event, `Then` the
  outcome. Write "Given the room is booked from 9:00 to 10:00", not "Given I
  POST /bookings with ...". One `When` per example.
- Use a `Scenario Outline` with `Examples` when one rule needs the same
  example across several values. Keep a `Background` to a few `Given`s that
  every example in the file shares.
- Step definitions hold the how. They call the code at the test's level: the
  service with its boundaries mocked, the code against real dependencies, or
  the public API.
- Reuse step definitions across a file and across features rather than
  writing near-duplicates.

## Where a case goes

Test each case at the lowest level that can prove it, and only once:

- A business rule, edge cases included → acceptance test, usually at unit
  level.
- Other logic or an edge case → unit.
- Behavior that depends on the real database or system → integration.
- A workflow a user or client depends on working → end-to-end.

## By stability level

- **Prototype**: no automated tests. The primary demo path works when
  exercised by hand.
- **Alpha**: end-to-end tests for the common paths, and acceptance tests for
  the core business rules. That's all Alpha needs.
- **Beta**: adds acceptance tests for the edge-case business rules, and unit
  and integration tests where they're useful. Primary and important secondary
  workflows are tested, and so are important failure paths.
- **GA**: important workflows and business rules are comprehensively tested,
  and important failure modes have tests that pin their behavior.

CI runs each suite in its own job from the stability level that requires it:
end-to-end and acceptance tests from Alpha, unit and integration tests from
Beta.

## Suites run in isolation and concurrently

The unit, integration, end-to-end, and acceptance suites are independent, so
CI can run their jobs at the same time:

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

- Integration, end-to-end, and acceptance tests get their dependencies from
  the project's local stack, such as Docker Compose, through the same
  environment variables as local runs and CI.
- Each test creates and isolates its own data, so the suite can run
  repeatedly and in any order.
