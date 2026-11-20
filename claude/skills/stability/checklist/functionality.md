# Functionality

**Does it do what it should, and keep doing it?** Functionality is about correctness, and the checks that keep it correct as the work changes.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Run the main path by hand after each change.
- A few sample inputs and their expected results in a test file, when that's quicker than checking by hand.
- A handful of output rows checked against the source by hand. _(Background job)_

Overengineering:

- Automated tests, unless they're the quickest way to try the idea, or the security baseline needs one, such as a test that tenants stay apart.
- Handling edge cases or bad input, unless they're the question.
- Contract tests, or an API description checked on every change, before its shape has settled. _(HTTP API)_

## Alpha

_The main path is checked, so a change can't quietly break it._

### functionality/alpha/main-path

The main path is known to work after every change.

How it might look:

- End-to-end tests of the main path's happy path, through the public interface, with real dependencies.
- A request to each endpoint gets the response it should, and each webhook it sends arrives with the right payload. _(HTTP API)_
- The main flow, driven through the UI. _(Web app, Mobile app, Desktop app)_
- The main flow done offline and then synced, when working offline is part of what it's for. _(Mobile app, Desktop app)_
- Run each command on a fixture, and check its output and exit code. _(CLI)_
- Run it on sample input, and check what it wrote. _(Background job)_
- The common flows a consumer follows, through its public API. _(Library)_
- A test that installs the packaged build or packed package the way people will, and runs its main flow. _(Mobile app, Desktop app, CLI, Library)_
- Expected results worked out by whoever owns the rules, not copied from what the code returns.
- Providers that send messages or charge money run in their test mode, or replaced by a fake, so a test run reaches no one.
- Tests that change a real system run against a sandbox or test account on it, and clean up after.
- Acceptance tests for the core rules it follows, each marked with the rule it verifies.
- With no sandbox or test environment, a fake built from the API's docs, checked on every change, plus one real run against a test record before each release.
- Tests run on every change, and a failure blocks merging.

Overengineering:

- Unit and integration tests for edge cases, coverage targets, load tests.
- A test for every input or combination, instead of the common ones.
- A full copy of a system it reads from to test against, when a few made-up records in its format will do.
- UI tests across a device farm or many OS versions. _(Mobile app)_

### functionality/alpha/errors-visible

Errors are visible, never silent.

How it might look:

- Input validated where it enters.
- Records or files it can't process rejected with the reason, or counted and logged, never dropped quietly or partly imported without saying so.
- A failed call to an outside service, such as a message that didn't send, recorded where someone can see it.
- A run that fails ends as a failure the scheduler can see, never a success with nothing written. _(Background job)_
- Every failure ends with a non-zero exit code, including one reported by a service it called. _(CLI)_
- Failures return an error status, never a 200 with the error inside. _(HTTP API)_
- Bad input or config throws a clear error, never returns NaN, zero, or a default. _(Library)_
- A sync that fails says so on the device, and what failed stays queued instead of looking sent. _(Mobile app, Desktop app)_

## Beta

_Edge cases and errors are handled, and checked too._

### functionality/beta/unusual-input

Unusual input and error paths are handled correctly.

How it might look:

- Data validated before it's written.
- Empty, missing, duplicate, and oversized input handled.
- A repeated request or redelivered message doesn't do its work twice, such as with an idempotency key or a duplicate check, when doing it twice would create a second record or send a second message.
- Two people changing the same thing at once can't both succeed unnoticed, such as with a unique constraint or a version check, where that would double-book or overwrite someone's work.
- Numbers kept exact where exactness matters, such as money in integer minor units or a decimal type.
- A rule for when the same item changes on the device and elsewhere before it syncs, and the person told when their change didn't win. _(Mobile app, Desktop app)_
- Records that fail validation set aside with the reason, and records changed at the source after loading updated, not added again. _(Background job)_
- Events it sends carry a unique ID and a timestamp, so receivers can drop duplicates and put them in order. _(HTTP API, Background job)_
- A malformed or incomplete config reported with the file, line, and key at fault, before anything is changed. _(CLI, Library)_

Overengineering:

- Guaranteed in-order delivery of events, when each carries a timestamp or version receivers can order by. _(HTTP API, Background job)_
- A general sync engine or conflict-free data types, where one simple rule for conflicts is enough. _(Mobile app, Desktop app)_
- Distributed locks against overlapping runs, when the scheduler only ever starts one. _(Background job)_

### functionality/beta/stays-correct

How it handles unusual input and errors is known to keep working as it changes.

How it might look:

- Unit tests for logic, edge cases, and failure paths.
- Integration tests against the real database or systems it uses.
- Acceptance tests for its edge-case rules too.
- Contract tests for interfaces others depend on.
- A fixed set of inputs with their expected results, checked on every change, so any changed result is seen.
- Tests against a fake or recording of each outside API, kept honest by a few runs against its sandbox.
- End-to-end tests of the main flows and their error states. _(Web app, Mobile app, Desktop app)_
- Tests that drop the connection, or close the app, partway through a sync. _(Mobile app, Desktop app)_
- For output that varies from run to run, such as an AI feature's, quality judged against a set of examples.

Overengineering:

- Concurrency and failure-injection suites, beyond a test for each known conflict.
- End-to-end tests for each edge case, where a unit or integration test checks it faster.
- Tests for every combination of flags. _(CLI)_
- Running every consumer's test suite before each release. _(Library)_

## GA

_Everything critical is checked on every change._

### functionality/ga/critical-workflows

Every critical workflow is known to work after every change.

How it might look:

- The critical workflows agreed with the client, usually those whose failure would stop their work, lose data, or cost money.
- End-to-end tests for every critical workflow.
- A table of real-world cases and their expected results, checked by someone who knows the rules, and run on every change.
- Anything that depends on the date or time tested across time zones and daylight-saving changes.
- Contract tests against a sandbox or recording of each service it calls, plus a scheduled check against the real one.
- Critical workflows tested on the packaged build, installed the way people install it. _(Mobile app, Desktop app, CLI, Library)_
- Property-based tests where inputs vary widely.
- No flaky or skipped tests left untracked.

Overengineering:

- A coverage target, or end-to-end tests of every path instead of the critical ones.

### functionality/ga/bugs-stay-fixed

A fixed bug stays fixed.

How it might look:

- A regression test with each bug fix.
- A made-up input that reproduces each bug, never the file or record someone sent.
