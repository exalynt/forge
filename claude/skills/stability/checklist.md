# Stability checklist

The outcomes each level expects, by quality, mirroring the readme's
[Stability checklist](https://readme.exalynt.com/how-it-works/stability-checklist)
(`src/data/stability-checklist.ts` in the readme repository; change it there
first). Apply it as "Using the checklist" in [SKILL.md](SKILL.md) says. Each
level includes the ones below it; "Prototype–Alpha" or "Beta only" means a
later level replaces it. IDs are `<quality>/<level>/<id>`. **Here:** notes are
how projects using these rules meet an outcome. `checklist/<quality>.md` has
each outcome's examples and overengineering.

## Utility: is it worth building, and is this the right way?

- Prototype `judged`: Those judging the idea can try it, or watch it run, enough to tell whether it's worth doing and whether this is the way to do it. Here: the capability doc's Learning.
- Alpha `end-to-end`: It works end to end on the real system.
- Alpha `real-use`: It's used for real.
- Alpha `learning`: What's learned from using it shapes what comes next.
- Beta `confirmed`: The need and the solution are confirmed, so what changes from here is refinement.
- Beta `success-agreed`: What success looks like is agreed. Here: the capability doc's Success.
- GA `delivers`: It's shown to deliver what it's for in real use.

## Usability: can its users learn and use it on their own?

- Alpha `main-path`: Those trying it can use its main path without help from whoever built it. Here: the capability doc and its Known gaps.
- Beta `self-serve`: Its users can learn and use it on their own.
- Beta `clear-results`: Its users can tell what happened after everything they do, and what to do next.
- Beta `recover`: Its users can recover from mistakes.
- Beta `everyone`: It works for everyone it's meant for.
- GA `all-learnable`: Everything it does can be learned without asking anyone.
- GA `consistent`: It behaves consistently across its whole interface.
- GA `all-accessible`: All of it works for everyone it's meant for.

## Functionality: does it do what it should, and keep doing it?

Here: the testing rule's suites for each level, run by the stack's CI rule.

- Alpha `main-path`: The main path is known to work after every change.
- Alpha `errors-visible`: Errors are visible, never silent.
- Beta `unusual-input`: Unusual input and error paths are handled correctly.
- Beta `stays-correct`: How it handles unusual input and errors is known to keep working as it changes.
- GA `critical-workflows`: Every critical workflow is known to work after every change.
- GA `bugs-stay-fixed`: A fixed bug stays fixed.

## Reliability: does it keep working when something goes wrong?

- Beta `data-kept`: Data that people or systems put in is kept across changes. Here: the stack's rule on changing the schema.
- Beta `undo-release`: A bad release can be undone.
- Beta `fails-clearly`: When something it depends on is slow or down, it fails clearly or degrades gracefully instead of hanging.
- GA `holds-up`: It holds up under overlapping runs, partial failures, and retries.
- GA `availability`: It meets its availability target.
- GA `recovery`: Recovering from a bad release or a loss of data is known to work.

## Efficiency: is it fast enough, at a cost that makes sense?

- Beta `realistic-volumes`: It works at realistic volumes.
- Beta `bounded`: The work it does is bounded. Here: HTTP list endpoints paginated, per the stack's pagination rule.
- GA `speed-targets`: It meets agreed speed targets at peak load, with headroom.
- GA `cost`: What it costs to run is known and acceptable.
- GA `limits-known`: Its limits are known.

## Security: is it protected, and does it follow the rules?

- Prototype `secrets` (baseline): Secrets stay out of code, version control, sample data, logs, and output.
- Prototype `real-data` (baseline): Real data comes only through the real systems, never copied into code or fixtures.
- Prototype `access` (baseline): Only those allowed can reach real systems or data, and one tenant or account never sees or changes another's.
- Prototype `personal-data` (baseline): Personal data is collected only where it's needed.
- Prototype `untrusted-input` (baseline): Untrusted input can't run as code or commands.
- Alpha `stable-dependencies`: It depends on nothing less stable than itself.
- Alpha `maintained-dependencies`: Its dependencies are maintained, and licensed for how it's used.
- Beta `misuse`: It's protected against misuse at the level it's exposed.
- Beta `duties`: The legal and regulatory duties that apply to it are known and met.
- GA `sound`: Its handling of credentials, payments, personal data, and permissions is known to be sound.
- GA `least-access`: Each part has only the access it needs.
- GA `traceable`: Sensitive actions can be traced to who did them.

## Compatibility: does it fit with what's around it?

- Prototype–Alpha `opt-in`: It reaches only those who chose to try it, and from Alpha, the people the client has agreed it can reach.
- Prototype–Beta `marked`: It's marked with its stability level wherever it's used.
- Prototype–Alpha `removable`: It can be removed without affecting anything people rely on.
- Alpha `standards`: Its interfaces follow the common standards for what it does, so others can work with it without surprises.
- Alpha only `told-of-changes`: Those trying it are told about breaking changes, which can still be made freely.
- Beta `available`: It's available to everyone it's meant for, without opting in, on the platforms and environments it supports.
- Beta `changes-heard`: Those depending on it hear about changes.
- Beta only `breaking-rare`: Breaking changes are rare, and announced ahead of time where possible.
- GA `stable-interface`: Its interface is stable: a breaking change comes as a new version, and the old one keeps working for a set time.
- GA `every-platform`: It's known to work on every platform it supports.

## Observability: when something goes wrong, does someone find out and see why?

Here: whatever is added follows the observability rule.

- Alpha `see-errors`: When an error happens, someone looking can see what happened.
- Beta `healthy`: Whoever runs it can tell it's healthy.
- Beta `hears-of-errors`: Whoever runs it hears about errors, and failed or missed runs, without searching for them.
- GA `finds-out-first`: Whoever runs it finds out about failures that matter to its users before they're reported.
- GA `knows-what-to-do`: Whoever runs it knows what to do about each failure.

## Maintainability: is it easy to understand and change?

Here: engineering docs follow the `engineering-docs` skill's level table.

- Alpha `project-way`: It's built the way the rest of the project is, so anyone who knows the project can follow it.
- Beta `structure`: Its structure matches the shape it settled on, with clear boundaries, so what a change affects is easy to see.
- Beta `shortcuts`: Earlier shortcuts are removed or tracked.
- Beta `needed-complexity`: Every piece of complexity has a reason it needs now.
- GA `others-can-change`: Another engineer can change it confidently without the author, and tell what a change will affect.
- GA `clean`: There's no dead code, untracked TODOs, or unexplained complexity.

## Supportability: who's responsible, and where do people turn?

- Alpha `owner`: It has an owner who decides its direction.
- Alpha `reporting`: Problems can be reported, and aren't lost.
- Beta `fixed`: Known problems are fixed as they're found.
- Beta `visible`: Known problems are visible to those depending on it. Here: the capability doc's Known gaps.
- GA `long-term`: The owner is responsible for it long term.
- GA `help`: Those depending on it know where to get help, and what response to expect.
- GA `not-only-author`: Someone other than the author can change and operate it.

## Special cases

- **Software people install** (mobile, desktop, CLI, library): logs are
  diagnostics people turn on and share; finding out about problems means
  consented crash reports, the logs of services it calls, and watching each
  release, not device telemetry; "whoever runs it" is the team that releases
  it; undoing a release means a fixed release, a paused rollout, or restoring
  the previous default; an upgrade never corrupts data, and from Beta never
  loses it, and unsynced work on a device is never reset away; its interface
  includes commands, flags, exit codes, output formats, config and saved files,
  and a library's results; availability targets usually don't apply.
- **Self-operated software, or parts someone else runs**: outcomes about their
  part, such as its monitoring, availability, or backups, are theirs. Check how
  yours behaves when theirs fails, note what you rely on them for, and write
  down a known gap when their part doesn't do what an outcome needs. For the
  baseline, no real data goes through until it does.
- **Results that vary**, such as AI features: checked by judging quality
  against a set of examples (evals), not exact matches.
