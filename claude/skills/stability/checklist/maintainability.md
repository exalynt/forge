# Maintainability

**Is it easy to understand and change?** Maintainability is about the next person to change it, and keeping complexity to what's needed.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Whatever answers the question fastest. Copying, hard-coding, and shortcuts are fine.

Overengineering:

- Abstractions, configurability, cleanup, or structuring for reuse.
- A deploy pipeline, containers, or infrastructure as code before anyone has decided to keep it.
- A general rules engine, or a layer that hides which provider it calls, when the few cases being tried can be hard-coded.
- A state management library, generated API client, or routing setup for a few screens of fake data. _(Web app)_
- A schema, validator, and version field for a config format only its authors write so far. _(Library, CLI)_

## Alpha

_It follows the project's conventions._

### maintainability/alpha/project-way

It's built the way the rest of the project is, so anyone who knows the project can follow it.

How it might look:

- The language's linter, formatter, and type checker, run on every change.
- The language's idioms, and the patterns the project already uses for similar work.
- The project's existing architecture, with its layers, boundaries, and places things go, instead of a new one.
- In a new project, the client's conventions or the framework's defaults, rather than a structure of its own.
- Established libraries for solved problems, such as dates, parsing, and retries, instead of hand-rolled ones.
- A README on setting it up and running it locally, and its configuration.
- A README on running it by hand, for one date or one message. _(Background job)_
- Review by another engineer or an agent.

Overengineering:

- Generalizing for cases nobody has asked for.
- Interfaces, plugin points, or settings with a single use.
- New layers, services, or architectural patterns it doesn't need yet.
- A workflow orchestrator, when a scheduler and one script will do. _(Background job)_
- A plugin system for commands, or support for several config formats. _(CLI)_
- A config format that grows into a language, with conditions, expressions, or includes. _(Library, CLI)_
- A general sync or offline layer for every kind of record, when only one or two kinds go offline. _(Mobile app, Desktop app)_
- Settings, theming, or branding for each tenant before a second client needs something different. _(Web app, Mobile app, Desktop app)_

## Beta

_Its structure fits the shape it settled on, and shortcuts are cleaned up._

### maintainability/beta/structure

Its structure matches the shape it settled on, with clear boundaries, so what a change affects is easy to see.

How it might look:

- Boundaries and dependencies that follow its shape, such as domain logic that doesn't depend on its transport or storage.
- Parts that can be tested on their own.
- Business logic and sync in plain modules, tested without the UI framework. _(Mobile app, Desktop app)_
- Commands kept thin over the logic they call. _(CLI)_
- A note on how it's built, and the key decisions behind it.
- Code and tests marked with the rules they implement and verify, so changing a rule shows what it touches.

Overengineering:

- Ports and adapters, repositories, or an interface for every dependency, when each has one implementation and the domain is small.
- Splitting it into several packages before anyone needs only part of it. _(Library)_

### maintainability/beta/shortcuts

Earlier shortcuts are removed or tracked.

How it might look:

- Code that survived from the prototype or alpha refactored or replaced.
- Shortcuts that remain recorded as known gaps.

### maintainability/beta/needed-complexity

Every piece of complexity has a reason it needs now.

How it might look:

- Real duplication combined once there's a second real use.

Overengineering:

- Hardening or scaling for load it doesn't have.
- Splitting it into separate services before one part has needs the others don't. _(Web app, HTTP API, Background job)_
- A workflow orchestrator, or pluggable sources and destinations for systems it doesn't connect to. _(Background job)_
- A plugin system for commands, when only the team adds them. _(CLI)_
- A rule engine or DSL, where plain functions and config cover the rules it has. _(Library)_

## GA

_Anyone on the team can change it confidently._

### maintainability/ga/others-can-change

Another engineer can change it confidently without the author, and tell what a change will affect.

How it might look:

- Names and structure that explain themselves, with comments only where the reason isn't obvious.
- Dependencies between parts explicit, with no hidden coupling through shared state or globals.
- Someone new to it can set it up and ship a small change from its docs alone.

### maintainability/ga/clean

There's no dead code, untracked TODOs, or unexplained complexity.

How it might look:

- Dead code, and flags left over from earlier levels, removed.

Overengineering:

- Rewriting working code to a newer pattern just to make it look cleaner.
