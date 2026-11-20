# Efficiency

**Is it fast enough, at a cost that makes sense?** Efficiency is about speed, load, and cost, measured against the work it really does.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Fast enough that those trying it can judge the idea.

Overengineering:

- Optimizing, caching, load testing.
- Caching a dependency's responses before its speed, cost, or limits get in the way of trying the idea.
- Incremental loads or change capture, when a full pull of a sample answers the question. _(Background job)_
- Server-side rendering, code splitting, or a CDN for something a few people open. _(Web app)_
- Bundling into a single binary, or trimming startup time. _(CLI)_

## Alpha

_Nothing expected._

How work on it might look:

- Fix what's slow enough to get in the way of real use.
- Reads from systems others rely on kept light, such as off-hours, in pages, or from a replica. _(Background job)_
- Photos and large files resized before they're stored or sent. _(Mobile app)_

Overengineering:

- Load tests, caches, scaling work.
- Incremental loads, parallel workers, or a distributed engine, when one process and a full reload finish in time. _(Background job)_
- A message queue or worker fleet, when a scheduled task handles today's volume. _(Web app, HTTP API, Background job)_
- Rate limits and quotas for each caller while only invited ones use it. _(HTTP API)_
- Memoizing, benchmarking, or bundle-size budgets before anyone has found it slow or large. _(Library)_

## Beta

_Fine at realistic volumes, with no runaway work._

### efficiency/beta/realistic-volumes

It works at realistic volumes.

How it might look:

- Tried with realistic amounts of data and use.
- Finishes well within its schedule at realistic volumes. _(Background job)_
- Starts quickly. _(Mobile app, Desktop app, CLI)_
- Long work, such as importing a large file, runs off the main thread with progress shown. _(Mobile app, Desktop app)_
- Tried on the oldest supported device and a slow connection, with a day's work captured offline and then synced. _(Mobile app)_

Overengineering:

- Performance targets, capacity planning.
- Caches, read replicas, or a search engine before a real query is shown to be slow.
- Incremental loads or change data capture when a full reload finishes well within its window. _(Background job)_
- Rewriting it, or bundling it into a native binary, to shave startup time nobody has noticed. _(CLI)_
- Delta updates, or shrinking the download, before its size is a problem. _(Desktop app)_

### efficiency/beta/bounded

The work it does is bounded.

How it might look:

- No query or call for each item in a list, and indexes where queries need them.
- Paging, input size limits, and streaming instead of loading everything into memory.
- Limits on what it accepts and returns are documented.
- A cap on what it can send or spend, such as texts per hour, so a bug or abuse can't flood people or run up a bill.
- Retries capped and spaced out, with work that keeps failing set aside instead of retried forever. _(Background job)_
- A limit on retries and concurrent deliveries to each receiver, so one slow one can't hold up the rest. _(HTTP API, Background job)_
- Reads from a source in pages or batches, within the rate limits it sets.

Overengineering:

- Autoscaling workers ahead of a measured backlog. _(Web app, HTTP API, Background job)_

## GA

_Targets for speed and cost are set, and met with room to spare._

### efficiency/ga/speed-targets

It meets agreed speed targets at peak load, with headroom.

How it might look:

- Measured targets for response time, throughput, or how long a job takes.
- A load test at expected peak, with room to spare.
- A run on the busiest day's volume, such as a seasonal peak, finishing with time to spare before its data is due. _(Background job)_
- Page load targets measured from real users' browsers, such as Core Web Vitals. _(Web app)_
- Benchmarks of the slow parts, on the devices or data sizes it's used with. _(Mobile app, Desktop app, CLI, Library)_
- A startup-time budget for the common commands, checked in CI. _(CLI)_

Overengineering:

- Tighter targets than those agreed.
- Real-user monitoring for a tool a known few people use on known machines. _(Web app)_
- Load tests of software that runs on one person's device. _(Mobile app, Desktop app, CLI, Library)_
- Load tests of every endpoint, instead of the busiest and the critical ones. _(HTTP API)_
- Parallelizing or scaling out a run that already finishes well before its data is due. _(Background job)_

### efficiency/ga/cost

What it costs to run is known and acceptable.

How it might look:

- What it costs per unit of work is known.
- Fees charged per use by what it depends on, such as per message sent, counted in its cost, with an alert on unusual spend.
- Memory, battery, and data use fit the devices it runs on. _(Mobile app, Desktop app)_
- The load it puts on the services it calls, such as calls per command and how often it polls, is known. _(Mobile app, Desktop app, CLI)_

Overengineering:

- Cutting costs beyond what's agreed to be acceptable.
- Cost tracked per tenant or customer, when nobody bills or decides by it.
- Tracking running costs for code that runs inside other teams' services. _(Library)_

### efficiency/ga/limits-known

Its limits are known.

How it might look:

- The rate limits and quotas of what it depends on fit peak use.
- Alerts before it reaches a known capacity limit.
- The largest tenant's or customer's data and traffic, not just the average, within what's been tested.
- The largest amounts, precision, and input sizes it handles exactly, documented and tested at those limits. _(Library)_
- Retries that back off with jitter, so many devices reconnecting at once don't swamp the service they call. _(Mobile app, Desktop app, CLI)_
