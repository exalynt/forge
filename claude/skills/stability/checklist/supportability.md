# Supportability

**Who's responsible, and where do people turn?** Supportability is about ownership, and what happens when something breaks.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Whoever is exploring the idea looks after it.
- Feedback through the test track's own tools, or a message to whoever built it. _(Mobile app, Desktop app)_

Overengineering:

- Bug tracking, support.
- An in-app feedback or bug-report form. _(Web app, Mobile app, Desktop app)_

## Alpha

_Someone owns it, and problems get written down._

### supportability/alpha/owner

It has an owner who decides its direction.

How it might look:

- A named person or team.

Overengineering:

- Support commitments, on-call.

### supportability/alpha/reporting

Problems can be reported, and aren't lost.

How it might look:

- Bugs recorded in the team's tracker.
- Those trying it know where to report a problem.
- Where to report a problem shown in `--help`, the README, the package's metadata, and with unexpected errors. _(CLI, Library)_
- A Help menu item, or the test track's feedback tools, with the version filled in. _(Mobile app, Desktop app)_

Overengineering:

- A help desk or ticketing system for those trying it.

## Beta

_Known problems get fixed, and people can see them._

### supportability/beta/fixed

Known problems are fixed as they're found.

How it might look:

- Bugs triaged regularly.

Overengineering:

- Support response times, on-call.

### supportability/beta/visible

Known problems are visible to those depending on it.

How it might look:

- Known issues listed in the docs, the release notes, or its known gaps.

## GA

_Long-term ownership, with clear ways to get help._

### supportability/ga/long-term

The owner is responsible for it long term.

How it might look:

- On-call, or someone to reach outside working hours, if it's critical enough.
- Where failures outside working hours can wait, a check at the start of each day and a way to catch up, instead of on-call.
- Reviews of serious incidents, so they're less likely to happen again.
- Dependencies, runtime, and database kept on supported versions, and upgraded before their end of life.
- Time planned for keeping up with each year's OS release and the stores' requirements, and for renewing signing certificates. _(Mobile app, Desktop app)_
- A support policy saying which versions get fixes, and for how long. _(Mobile app, Desktop app, CLI, Library, HTTP API)_

Overengineering:

- Round-the-clock on-call for work nobody needs outside working hours.
- Out-of-hours on-call for a nightly job whose output isn't read until morning, when a morning check and a rerun are enough. _(Background job)_
- On-call for a library, when its problems reach the team through its consumers during working hours. _(Library)_

### supportability/ga/help

Those depending on it know where to get help, and what response to expect.

How it might look:

- A support channel, with expected response times.
- When it serves a client's own customers, which questions go to the client and which reach the team is clear.
- A status page or announcement during an outage. _(Web app, Mobile app, HTTP API)_
- A notice to those who read its output when data is late or incomplete, and again when it's fixed. _(Background job)_
- A way to reach support from inside it, with the version and the diagnostics people choose to share. _(Mobile app, Desktop app)_
- A named contact at each team or system that integrates with it, so notices reach a person. _(HTTP API, Library)_

Overengineering:

- A status page for software that keeps working when its services are down, or whose outages whoever runs those services announces. _(Desktop app, CLI, Library)_
- A self-serve developer portal with sign-up and ticketing, for a handful of callers onboarded by hand. _(HTTP API)_

### supportability/ga/not-only-author

Someone other than the author can change and operate it.

How it might look:

- More than one engineer able to work on it.
- Where one engineer works alone, a handover note and shared access, so someone else could take over.
- Building, signing, and publishing done by CI, with certificates, store accounts, and publish rights held by the client or team, not one person. _(Mobile app, Desktop app, CLI, Library)_
