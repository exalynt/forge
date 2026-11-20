# Reliability

**Does it keep working when something goes wrong?** Reliability is about failure, such as a bad release, a slow dependency, or lost data, and how it recovers.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Made-up, hard-coded, or in-memory data, local files of made-up data, or a throwaway database.
- Those trying it with real work keep their usual way of recording it, so nothing real exists only in the prototype.
- A local database deleted and recreated when its schema changes, with those trying it told their data won't survive a new build. _(Mobile app, Desktop app)_

Overengineering:

- Migrations, backups, retries, failover.
- Resuming or undoing an operation that fails partway.
- Offline sync with conflict resolution, background sync, and retry queues, unless syncing is the question. _(Mobile app, Desktop app)_
- A delivery queue, dead-letter handling, or delivery guarantees, when sending once is enough to see it work. _(HTTP API, Background job)_

## Alpha

_Nothing expected._

How work on it might look:

- Real storage, with data that may be reset, and those using it told before a reset when it holds their real work.
- A managed database's default backups left on when it holds real work, since they cost nothing to keep.
- A simple forward migration, rather than a reset, when people still act on the records.
- A script to reset or seed its data.
- Before a build changes how data is stored on the device, work not yet sent is sent or carried over, and data it can't read is set aside, not deleted. _(Mobile app, Desktop app)_
- Output it can rebuild by running again from the source, so a reset loses nothing, and a time limit on each run so one that hangs shows as failed. _(Background job)_
- Files it writes into people's repos never overwrite their edits without asking. _(CLI)_

Overengineering:

- Down migrations, restore drills, rollback plans.
- Failing over to a second provider.
- Dead-letter queues and retry policies, when a failed message can be seen in the logs and resent by hand. _(HTTP API, Background job)_
- Background sync, resumable uploads, and retry schedules, where syncing when it opens works for those trying it. _(Mobile app, Desktop app)_
- Repair tools or automatic backups for data kept on the device. _(Mobile app, Desktop app)_

## Beta

_Enough that people's data is safe and a bad release can be undone._

### reliability/beta/data-kept

Data that people or systems put in is kept across changes.

How it might look:

- Versioned migrations run by the release, never by hand.
- A down migration, or a backup taken before a risky change.
- Automated backups, such as a managed database's defaults, even before a restore has been tested.
- Queued and scheduled work, such as tomorrow's reminders, kept across deploys and restarts.
- Output that can be rebuilt from its source counts as kept while the source still has it. History only it holds is backed up.
- Local data migrated on upgrade, from a copy taken first and kept until the new version has opened it. _(Mobile app, Desktop app)_
- Work captured on the device kept where the OS won't clear it until the server confirms it has it, and safe to send again. _(Mobile app, Desktop app)_
- Where its local data is kept documented, so people's own backups include it. _(Desktop app, CLI)_
- Config files people wrote for it, and files it wrote, still work after an upgrade. _(CLI, Library)_

Overengineering:

- Restore drills.
- Snapshots of every run's output kept just in case. _(Background job)_
- A cloud backup or sync service for data on devices, when people's own backups cover where it's kept. _(Mobile app, Desktop app)_
- Versioned config schemas with automatic migration, when older config simply keeps loading. _(CLI, Library)_

### reliability/beta/undo-release

A bad release can be undone.

How it might look:

- A rollback to the previous release. _(Web app, HTTP API, Background job)_
- Migrations that add before they remove, so the previous release still runs after a rollback. _(Web app, HTTP API, Background job)_
- It can be switched off quickly without a release, such as with a kill switch. _(Web app, HTTP API, Background job)_
- A quick fixed release. _(Mobile app, Desktop app, CLI, Library)_
- A fixed release can open data the bad release already wrote or migrated. _(Mobile app, Desktop app, CLI)_
- The previous version made the default install again, such as by moving npm's latest tag back or reverting the Homebrew formula, and the bad one deprecated. _(CLI, Library)_
- A remote switch and a minimum-version check in the first build people rely on, since builds already out can't gain them later. _(Mobile app)_
- A phased or staged rollout that can be paused if crash reports climb. _(Mobile app, Desktop app)_
- Processing paused without a release, with queued work held until it resumes, and output a bad release wrote rebuilt by rerunning those runs. _(Background job)_

Overengineering:

- A remote kill switch, and a service to host it, for installed software with no backend. _(Desktop app, CLI, Library)_
- Automatic rollback to the previous version on people's devices. _(Desktop app)_
- Over-the-air code updates to get around store review. _(Mobile app)_

### reliability/beta/fails-clearly

When something it depends on is slow or down, it fails clearly or degrades gracefully instead of hanging.

How it might look:

- Timeouts on every call to a dependency, with retries only where they're safe.
- A message or call that can't go through retried later or marked failed where someone will see it, never dropped.
- Messages to others' systems, such as webhooks, retried with backoff when the receiver is down, and given up on after a set time with the failure recorded. _(HTTP API, Background job)_
- A run that can't finish leaves the last complete output in place, or marks its own as incomplete. _(Background job)_
- A clear message and a non-zero exit when a service it calls is down or too slow, so a CI job fails instead of hanging. _(CLI)_
- An unreachable update server or service never blocks the work people can do without it. _(Mobile app, Desktop app)_

Overengineering:

- Chaos testing, or running in several regions.
- A transactional outbox or worker service for webhooks, when sending after the write and logging failures is enough. _(HTTP API, Background job)_
- A second provider to fail over to, before the first one's outages are shown to matter.
- A message broker or separate queue service, when a table in its own database is enough. _(Web app, HTTP API, Background job)_
- Timeouts, retries, or fallbacks inside a library that calls nothing, or whose caller already retries. _(Library)_
- Making screens work offline that aren't used offline. _(Mobile app, Desktop app)_

## GA

_Enough to depend on: it rides out failures, and recovery is proven._

### reliability/ga/holds-up

It holds up under overlapping runs, partial failures, and retries.

How it might look:

- Operations safe to repeat, such as with idempotency keys or upserts.
- Tests for two runs at once, and for retrying after a crash halfway through.
- It fails safe where failing could cause harm, stopping rather than charging twice, sending twice, or acting on bad data.
- A write that takes several steps done in one transaction, so a crash halfway leaves nothing half saved.
- Rerunning any time range rebuilds what it produced, a run that starts while another is going waits or exits, and it stops when the source's data no longer has the shape it expects. _(Background job)_
- Deliveries to endpoints others run retried for a set time, then set aside for replay, so one receiver's outage doesn't hold up the rest. _(HTTP API, Background job)_
- Sync picks up where it left off after the app is killed or the connection drops, with IDs made on the device so a retry isn't new work. _(Mobile app, Desktop app)_
- Cancelling it, or the CI job running it, leaves nothing half-done, or a rerun finishes the job. _(CLI)_

Overengineering:

- Exactly-once delivery, where safe retries and duplicate checks are enough.
- A workflow orchestrator or distributed lock service for one job on one schedule, where the scheduler's no-overlap setting or a database lock is enough. _(Background job)_
- Retries, locking, or idempotency handling in pure functions that do no I/O. _(Library)_
- A general-purpose sync engine or CRDTs, where one person owns each record and a simple rule settles conflicts. _(Mobile app, Desktop app)_

### reliability/ga/availability

It meets its availability target.

How it might look:

- An agreed availability target, such as 99.9%, that's measured. _(Web app, HTTP API)_
- Its data ready by an agreed time, and queued work handled within an agreed time. _(Background job)_
- A target that fits when it's used, such as the hours its users work, and that allows for what it depends on.
- Usually not applicable on people's own devices: what matters is that its main work keeps going when the services it calls are down. _(Mobile app, Desktop app, CLI)_
- Usually not applicable: the software using it holds the target. _(Library)_

Overengineering:

- Running in several regions, or failover, beyond what the agreed target needs.
- A round-the-clock target, or more nines, for work used only in business hours.
- A target higher than what it depends on can deliver.
- An availability target for software that runs on people's own devices. _(Mobile app, Desktop app, CLI, Library)_
- Offline support for parts its users only use when connected. _(Mobile app, Desktop app)_

### reliability/ga/recovery

Recovering from a bad release or a loss of data is known to work.

How it might look:

- A rollback that's been practiced, including upgrading, downgrading, and upgrading again. _(Web app, HTTP API, Background job, Desktop app, CLI)_
- A staged rollout that can be paused, and a minimum supported version raised with notice, after work not yet sent has finished sending. _(Mobile app)_
- A bad version deprecated and the previous one made the default again, tried once before it's needed. _(CLI, Library)_
- Automated backups, with a restore that's been tested. _(Web app, HTTP API, Background job)_
- How much data a loss could cost, and how long recovery may take, agreed and shown in a restore test. _(Web app, HTTP API, Background job)_
- Output it can rebuild restored by rerunning from the source, with how far back the source keeps data known. _(Background job)_
- A copy of the local data taken before each migration, with a tested way to put it back. _(Mobile app, Desktop app)_
- For data kept only on people's devices, recovery through their own backups, shown once by copying the data folder to a fresh machine and opening it. _(Desktop app, CLI)_
- How much unsent work a lost device could cost is stated, and kept small by sending work as soon as there's a connection. _(Mobile app)_
- Migrations tried on data the size of production, made up where real data can't be copied.

Overengineering:

- Tooling to restore a single tenant or record before anyone has needed it, when restoring to a copy and fixing by hand would do.
- Backups and restore drills for output it can rebuild from a source that still has the data. _(Background job)_
- Backups and restore drills for software that keeps no data. _(Library)_
- Downgrade migrations for every release of software that updates itself, where an older version can refuse newer data instead. _(Mobile app, Desktop app)_
