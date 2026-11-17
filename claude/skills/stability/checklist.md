# Stability checklist

What each level generally includes, area by area. It mirrors the Exalynt
readme's [Stability checklist](https://readme.exalynt.com/how-it-works/stability-checklist),
which is generated from `src/data/stability-checklist.ts` in the readme
repository. Change an item there first, then here. The **Here:** notes are how
projects using these rules meet an item, and aren't in the readme.

## Reading it

- **Each item applies from its level up.** "Prototype–Alpha" or "Alpha only"
  means a higher level replaces it with its own item.
- **Guidelines, not a gate.** Work fits a level when it generally meets that
  level's items and the ones below, not a perfect score, even at GA. Each item
  is met, not applicable (with a few words on why), or a known gap: something
  people would notice is missing, either to close later or accepted as a
  tradeoff, written down where they can see it. The security and privacy
  baseline is never a known gap.
- **A ceiling as well as a floor.** An item that first appears at a later level
  is UNNECESSARY NOW.
- **Aim for the intent, in the form that fits.** Each area's why, in italics, is
  the intent. Examples in parentheses are only examples. "Errors logged with
  context" means server logs for an API, and a verbose option for a
  command-line tool.
- **Not applicable when the intent doesn't apply.** Backups don't apply to work
  that keeps no data. An example not matching isn't a reason on its own.
- **Feature flags are a recommendation, not a requirement.** A flag, or another
  way to turn work off without a code change, is worth it when new work lands
  in software people already rely on and a problem would hurt, or when
  releasing to some users first. In a new product nobody relies on yet,
  rolling back is usually enough. Once a flag has done its job, removing it
  keeps the code simple, but that's housekeeping, not part of any level. A
  missing flag, or one not yet removed, isn't a known gap, so don't raise or
  ask about it as one.
- **Team practices fit the team.** An owner, a place to report bugs, and
  review by another engineer are recommended, in whatever form the team and
  client already work, such as a client who reports problems by message, or an
  engineer working alone. They usually live outside the code, so they're taken
  as met unless something says otherwise, and one that doesn't fit the team
  isn't a known gap. Don't raise or ask about them as one.
- **A capability can rely on its project.** Items about the whole system, such
  as testing on every change, recovering data, monitoring, and the release
  process, can be met by what the project already has, as long as it covers the
  new work.
- **Software others run**, such as a command-line tool or a library: logs mean
  diagnostics users can turn on and share in a bug report, monitoring and
  alerting usually belong to whoever runs it, and keeping data safe means an
  upgrade never loses or corrupts it.

Before applying it, settle what the items mean for this work:

1. What's getting the level? Usually a capability, sometimes a whole feature or
   project.
2. What do people and other software touch? Screens, endpoints, commands,
   files, and messages are all its interface.
3. What's its main workflow, end to end?
4. What data does it keep, where, and in what format?
5. What does it depend on, and how stable is each?
6. Who runs it, and how does it reach users?

## Building the right thing

### Direction and feedback

_We're building the right thing, and we know it._

- **Prototype**: What it's for is written down: the question it answers, or the
  problem it solves. Here: the capability doc's opening paragraph, and its
  Learning.
- **Prototype**: The client can reach it, try it, and react to it.
- **Prototype**: The client's feedback is recorded, and shapes what comes next.
- **Alpha**: Hooked up end to end. Sample data is fine if the system itself
  stores and serves it.
- **Alpha only**: People who'd want to try it can find out it exists and how to
  opt in.
- **Alpha**: Feedback comes from real use, not just demos.
- **Beta**: The client agrees its shape has settled.
- **Beta**: What success looks like is agreed with the client. Here: the
  capability doc's Success.
- **GA**: Shown to meet that success in real use.

## Works correctly

### Tests

_What's been checked, and what keeps it from breaking again._

- **Prototype**: The main path works when someone tries it by hand. Here: no
  automated tests.
- **Alpha**: The main path is checked automatically (such as unit, integration,
  or end-to-end tests). Here: the testing rule's Alpha suites, an end-to-end
  test of each endpoint's happy path and acceptance tests for the core business
  rules.
- **Alpha**: Checks run on every change, and a failure stops it from shipping
  (such as CI that blocks merging). Here: the stack's CI rule, Alpha jobs.
- **Beta**: Edge cases and error paths are checked. Here: acceptance tests for
  the edge-case business rules, unit and integration tests where they're
  useful, and their CI jobs.
- **GA**: Critical workflows checked end to end.
- **GA**: Every fixed bug has a check that would catch it coming back.
- **GA**: Checks can be trusted: none flaky or skipped without being tracked.

Work whose results vary, such as AI features or ranking, is still checked:
automatically checking it means judging quality against a set of examples
(evals), not exact matches.

### Edge cases and errors

_What happens with unusual input, and when something goes wrong._

- **Alpha**: Errors fail visibly, never silently.
- **Alpha**: Input validated where it enters.
- **Beta**: Empty, missing, duplicate, and large inputs handled.
- **Beta**: Error messages say what went wrong and what to do next, without
  exposing internals.
- **GA**: Concurrent use and partial failures handled (two at once, a crash
  halfway through), and repeating an operation is safe.
- **GA**: Every known edge case handled or documented as out of scope.

### Data and migrations

_Your data stays correct, and can be recovered if something goes wrong._

- **Prototype**: Sample data is clearly marked wherever it appears.
- **Alpha**: Data it creates can be cleaned up. Data may still be reset, so
  editing a migration in place and resetting is fine.
- **Beta**: Data people put in is kept across changes.
- **Beta**: Changes to the shape of stored data apply the same way everywhere,
  without hand edits (such as migration scripts). Here: the stack's rule on
  changing the schema.
- **Beta**: Data validated before it's written.
- **Beta**: A data change that goes wrong can be undone (such as a reversible
  migration, or a backup taken first).
- **GA**: Data changes tried on production-sized data before they ship.
- **GA**: Stored data can be recovered after a loss, and recovery has been tried
  (such as backups with a tested restore).

### Performance and cost

_It's fast enough for the work, today and at peak, at a cost that makes sense._

- **Prototype**: Responsive enough for the people trying it.
- **Alpha**: Usable with today's real amounts of data and use.
- **Beta**: Tried at realistic volumes, with bounded work (queries, memory,
  input size). Here: an HTTP API's list endpoints are paginated, as the stack's
  pagination rule describes.
- **Beta**: Obvious problems fixed (such as a query or call per item, missing
  indexes, loading everything into memory, slow startup).
- **GA**: Meets performance targets agreed with the client, measured (such as
  response time, throughput, or startup time).
- **GA**: Handles peak load, or the largest expected input, with headroom.
- **GA**: What it costs to run is known, and fits expected use.

## Safe

### Security, privacy, and compliance

_Your systems and data are protected the same way at every level, and the work
follows the rules that apply to it._

The Prototype items are the **security and privacy baseline**. They hold at
every level and are never a known gap.

- **Prototype (baseline)**: Secrets kept out of code, version control, sample
  data, and demos (such as in the secret store).
- **Prototype (baseline)**: Secrets and personal data kept out of logs, error
  messages, and output.
- **Prototype (baseline)**: Real client data comes only through production
  systems, never copied into code, sample data, or fixtures.
- **Prototype (baseline)**: Access control enforced on anything touching client
  systems or data.
- **Prototype (baseline)**: People can only see and change data they're allowed
  to. One client or customer can never see or change another's.
- **Prototype (baseline)**: Personal data collected only where needed.
- **Prototype (baseline)**: Untrusted input can't run as code or commands (such
  as SQL injection or cross-site scripting).
- **Beta**: Personal data can be exported and deleted on request.
- **Beta**: Protected against misuse at the level it's exposed to (such as rate
  limits, abuse of sign-ups, or scraping).
- **Beta**: The legal and regulatory duties that apply are known (such as
  privacy law, industry rules, where data may be stored, and contracts).
- **GA**: Reviewed for security wherever it handles credentials, payments,
  personal data, or permissions.
- **GA**: Each part has only the access it needs, and sensitive actions can be
  traced to who did them.
- **GA**: Stored data kept only as long as it's needed.
- **GA**: The legal and regulatory duties that apply are met.

### Dependencies and integrations

_What it relies on, and what happens when those things change or fail._

- **Alpha**: Depends on nothing less stable than itself (for a third party, a
  release at a matching stage). When that's the only option, it's a known gap.
- **Alpha**: Dependencies maintained, with compatible licenses.
- **Alpha**: Settings that differ between environments kept out of code (such
  as service URLs and keys).
- **Beta**: Builds use the same dependency versions every time, and known
  vulnerabilities in them are found.
- **Beta**: When an external service is slow or down, it fails clearly instead
  of hanging (such as timeouts, and retries where safe).
- **GA**: The rate limits and quotas of what it depends on fit expected use.
- **GA**: Dependencies kept current (such as automated or scheduled updates).

## Runs in production

### Logs, traces, and metrics

_When something goes wrong, we can see what happened and why._ Here: whatever
is added follows the observability rule.

- **Alpha**: Errors logged with context.
- **Alpha**: Whoever operates it can read the logs.
- **Beta**: Everything from one request, job, or run can be found together
  (such as structured logs with a shared ID).
- **Beta**: Whoever operates it hears about errors without searching the logs
  (such as an error tracker).
- **Beta**: How often key operations run and fail, and how long they take, is
  measured.
- **GA**: Whether the main workflows succeed is measured, not just whether the
  system is up.
- **GA**: A failed critical workflow can be followed from start to finish (such
  as traces linked from logs).
- **GA**: Logs are kept no longer than privacy needs allow.

### Monitoring and alerting

_Whoever operates it finds out something is wrong before your users do._

- **Beta**: Health checked automatically (such as uptime checks for services,
  or success checks for scheduled jobs).
- **Beta**: Key metrics visible at a glance (such as a dashboard).
- **GA**: Targets set and tracked for latency, error rate, and availability.
- **GA**: Whoever operates production is alerted to failures users would
  notice.
- **GA**: Every alert says what to do about it (such as a linked runbook).
- **GA**: Alerts are known to fire, and quiet enough that people act on them.

### Release and rollback

_How it reaches your users, and how quickly it can be undone._

- **Prototype**: Its stability level is clear wherever people see it (a
  stability mark, or no mark once it's GA).
- **Prototype only**: Not available to users by default (such as outside
  production, or behind a flag that's off).
- **Prototype–Alpha**: Can be removed cleanly, and can't affect anything people
  rely on, including production data.
- **Alpha only**: Opt-in only (such as only for invited users, or behind a
  feature flag).
- **Beta**: Released the same way as the rest of the product (such as the
  normal deploy pipeline).
- **Beta**: Can be rolled back to the previous version.
- **GA**: Rollback has been tried.
- **GA**: It can be set up again from scratch without hand steps, such as after
  losing its environment.
- **GA**: For software others install, users can verify a release came from us
  (such as signed builds).

### Ownership and support

_Someone is responsible for it, and you know what happens when it breaks._

- **Alpha**: It has an owner who decides what happens to it.
- **Alpha**: Bugs recorded where the team tracks its work.
- **Alpha**: The client and users know how to report a problem.
- **Beta**: Bugs triaged and fixed as found.
- **Beta**: Known bugs shared with the client.
- **GA**: GA bugs fixed before new work.
- **GA**: The client and users know where to get help, and what response to
  expect.
- **GA**: The owner is responsible for keeping it running long term, including
  who responds outside working hours, if it's critical enough to need that.
- **GA**: During a serious outage, users can find out what's happening (such as
  a status page or an announcement).
- **GA**: Serious incidents are reviewed so they're less likely to happen again.

## Usable and lasting

### User experience

_The people using it can get their work done, without surprises._

- **Prototype**: Its main idea is clear to the people using it.
- **Alpha**: A real user can complete the main workflow.
- **Beta**: Waiting, empty, and error states handled on purpose (such as
  loading, no results, and clear error messages).
- **Beta**: Usable by everyone it's meant for, the way they use it (such as
  keyboard access, labels, and contrast in a UI, or help text in a command-line
  tool).
- **GA**: Supported platforms, languages, and regions listed (such as browsers,
  devices, and time zones), and it works the same on each.
- **GA**: Meets the project's accessibility standard.

### Documentation

_You, your users, and the next engineer can use, run, and change it without
asking the author._ Here: what people would notice goes in the capability doc,
and engineering docs follow the `engineering-docs` skill's level table.

- **Prototype**: What it does, and what it doesn't handle yet, is written down
  where the client can see it. Here: the capability doc and its Known gaps.
- **Prototype**: Developers: how to run it.
- **Alpha only**: Users: how to try it.
- **Alpha**: Developers: its configuration.
- **Beta only**: Users: how to use and configure it.
- **Beta**: Developers: how it's put together, and the key decisions behind it.
- **Beta**: Users told what changed in each release (such as release notes or a
  changelog).
- **GA**: Users: complete, current documentation.
- **GA**: Developers: reference docs for its public interface.
- **GA**: Operations: how to deploy, roll back, and diagnose it, or for
  software others install, how to install, upgrade, and troubleshoot it.

### Code and maintainability

_It can keep changing, by any engineer, without getting harder to change._

- **Alpha**: Follows the project's conventions, and passes its automated checks
  (such as lint and type checks).
- **Alpha**: Reviewed by another engineer, where the team has one.
- **Beta**: Earlier shortcuts removed or tracked.
- **GA**: Another engineer can change it without the author.
- **GA**: No untracked TODOs or dead code.

### Breaking changes

_What you build on it keeps working, or you hear first._

- **Beta**: What's part of its public interface, and what's internal, is clear.
- **Beta**: People building on it have somewhere to hear about changes (such as
  release notes or a changelog).
- **Beta only**: Breaking changes kept to a minimum, and announced ahead of time
  where we can, with how to move to the new behavior.
- **GA**: Public interfaces documented as stable.
- **GA**: Breaking changes announced ahead of time, with a deprecation window
  long enough to move (such as a new version, with the previous one kept
  working for a set time).
