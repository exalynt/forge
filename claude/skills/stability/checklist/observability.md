# Observability

**When something goes wrong, does someone find out and see why?** Observability is about seeing what it's doing while it runs.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Console output is enough.
- Outbound calls and messages printed to the console without credentials or personal data, so those watching can see they happened.
- Counts of what each run read, skipped, and wrote, printed at the end. _(Background job)_

Overengineering:

- Logging setup, error tracking, metrics.
- Product analytics, when watching people try it answers the question. _(Web app, Mobile app, Desktop app)_
- Crash reporting or analytics SDKs, which bring consent and privacy-label duties with them. _(Mobile app, Desktop app)_
- A logging dependency or debug hooks in the package. _(Library)_

## Alpha

_Someone who looks can see what went wrong._

### observability/alpha/see-errors

When an error happens, someone looking can see what happened.

How it might look:

- Error logs that include the operation, the IDs of what it was working on, and a request or run ID.
- Logs that name the file, row, and step, not the values in it.
- The tenant in every error log. _(Web app, HTTP API, Background job)_
- Each run logs when it started and finished, how many records it read, wrote, and skipped, and why it failed. _(Background job)_
- A request ID returned with every response. _(HTTP API)_
- The request ID from a failed call to a service shown in the error, so the service's owners can find it in their logs. _(Mobile app, Desktop app, CLI)_
- A `--debug` flag whose output, with secrets redacted, can be shared in a bug report. _(CLI)_
- Crash logs a developer can get at, and a way for someone trying it to send its recent log, with tokens and personal data left out. _(Mobile app, Desktop app)_
- Errors thrown name what failed and the input or config field that caused it, so the logs of the software using it show why. _(Library)_

Overengineering:

- Dashboards, tracing, and alerts beyond what its host or scheduler sends for free.
- Logging inside a library, when the errors it throws carry what's needed. _(Library)_
- Sending logs from people's devices to a central service, when they can send a log with a bug report. _(Mobile app, Desktop app, CLI)_
- Data-quality frameworks, lineage tools, or freshness monitors. _(Background job)_
- Delivery dashboards or per-message tracking for outgoing notifications or webhooks. _(Web app, HTTP API, Background job)_

## Beta

_Errors reach whoever runs it, without looking._

### observability/beta/healthy

Whoever runs it can tell it's healthy.

How it might look:

- A health endpoint, with an uptime check. _(HTTP API, Web app)_
- How often key operations run and fail, and how long they take, is recorded. _(Web app, HTTP API, Background job)_
- A dashboard of its key metrics, such as traffic, errors, latency, and queue depth. _(Web app, HTTP API, Background job)_
- Row counts or totals checked against the source after each run. _(Background job)_
- A `doctor` or diagnostics command that checks its setup and connections. _(CLI, Desktop app)_
- A diagnostics screen showing what's waiting to sync and the last error. _(Mobile app)_
- Error rates for the calls it makes, from the logs of the service it calls. _(Mobile app, Desktop app, CLI)_
- Which versions are in use, and how many update downloads fail, from the update server's logs. _(Desktop app)_
- Usually not applicable: the software using it watches its own health. _(Library)_

Overengineering:

- Service level objectives.
- Metrics for every endpoint and caller, instead of the few that show whether it's healthy.
- Usage telemetry from people's devices. _(Mobile app, Desktop app, CLI, Library)_
- Tracing, or a metrics stack of its own, when the host's metrics and an error tracker show whether it's healthy. _(Web app, HTTP API, Background job)_
- Data lineage or data-quality tooling with checks on every column. _(Background job)_

### observability/beta/hears-of-errors

Whoever runs it hears about errors, and failed or missed runs, without searching for them.

How it might look:

- An error tracker that notifies people.
- Failures that providers report back, such as undelivered texts or bounced emails, recorded and counted.
- A notification when a run fails or doesn't happen. _(Background job)_
- Crash reporting, with consent. _(Mobile app, Desktop app)_
- Where crash reports can't be sent, a Send diagnostics action that saves redacted logs to attach to a bug report. _(Mobile app, Desktop app)_
- An unexpected error prints a short message and how to report it, with the version and redacted `--debug` output, instead of a bare stack trace. _(CLI)_
- Problems reach the issue tracker, with version details, and where to report one is stated in the package itself. _(CLI, Library)_

Overengineering:

- Paging, distributed tracing.
- A notification for every rejected record, instead of a count per run. _(Background job)_

## GA

_Problems are caught before users report them, with a plan for each._

### observability/ga/finds-out-first

Whoever runs it finds out about failures that matter to its users before they're reported.

How it might look:

- Automated alerts on error rate, latency, and availability against agreed targets. _(Web app, HTTP API, Background job)_
- Where the client runs the service, failed syncs sent as reports without the work's content, and failure rates by version shared by the client. _(Mobile app, Desktop app)_
- Where nothing may leave people's devices, each release goes to a few of the client's users first, with the update server's logs watched. _(Mobile app, Desktop app)_
- Whether its main workflows succeed is measured, not just whether it's up.
- A scheduled check that runs a critical workflow in production, such as signing in to a test account. _(Web app, HTTP API)_
- Traces that follow a critical workflow across services, where it spans more than one. _(Web app, HTTP API, Background job)_
- Failures a provider reports back later, such as undelivered messages or bounces, watched and alerted on.
- Alerts on lateness and backlog, not just failure, and checks on what each run wrote, such as row counts and totals against the source. _(Background job)_
- Crash and update-failure rates for each release, from reports people agreed to send, watched after every release. _(Mobile app, Desktop app)_
- Its version sent with every call, and failure rates by version from the service's logs, watched after each release. _(CLI)_
- Release candidates tried in the consuming services' tests before they're the default install. _(Library)_
- A pre-release channel, so problems show up first with those who chose it. _(Mobile app, Desktop app, CLI, Library)_
- A reminder well before signing certificates, developer accounts, and domains expire. _(Mobile app, Desktop app, CLI, Library)_
- An agreed way to hear from whoever runs a service it depends on when that service has problems.

Overengineering:

- Distributed tracing for a single service.
- Alerts on every error or metric, so the ones that matter get lost.
- Error budgets and burn-rate alerts, where alerts on the agreed targets catch the same failures. _(Web app, HTTP API, Background job)_
- Alerts on client errors, such as 4xx responses, that come from callers' own mistakes. _(HTTP API)_
- A data observability platform, when a few checks on counts, totals, and freshness catch what matters. _(Background job)_
- Telemetry from people's devices beyond the crash and failure reports they agreed to, or built into a library. _(Mobile app, Desktop app, CLI, Library)_

### observability/ga/knows-what-to-do

Whoever runs it knows what to do about each failure.

How it might look:

- Each alert links to a runbook.
- Runbooks for deploying, rolling back, and diagnosing it. _(Web app, HTTP API, Background job)_
- A runbook for when something it depends on is down, covering what users see and how to catch up afterwards.
- A plan for a suspected data leak: who decides, who's told, and by when. _(Web app, HTTP API, Background job)_
- A runbook for rerunning a failed run, backfilling a range, and replaying or discarding messages that were set aside. _(Background job)_
- A way to replay failed deliveries once the receiver is fixed. _(HTTP API, Background job)_
- A runbook for a bad release: stopping it spreading, publishing a fix through each channel, and telling those affected. _(Mobile app, Desktop app, CLI, Library)_
- Runbooks for pausing a rollout, getting a fix through store review quickly, and recovering work from a device that won't sync. _(Mobile app)_

Overengineering:

- Runbooks for failures that haven't happened and aren't likely.
