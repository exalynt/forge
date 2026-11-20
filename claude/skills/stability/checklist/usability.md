# Usability

**Can its users learn and use it on their own?** Usability is about the people, or programs, using it: whether they can get their work done without help.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Nothing expected._

How work on it might look:

- Whoever built it walks those trying it through it.
- Shown on the device those judging it will really use, such as a front-desk tablet or a phone. _(Web app, Mobile app, Desktop app)_
- A request collection or a few curl commands for the main calls. _(HTTP API)_

Overengineering:

- User docs, or polishing flows beyond what trying the idea needs.
- Layouts for every screen size, when those judging it use one device. _(Web app, Mobile app, Desktop app)_
- Shell completions, colored output, spinners, and interactive prompts. _(CLI)_
- An SDK, a developer portal, generated reference docs, or self-serve sign-up. _(HTTP API, Library)_

## Alpha

_The main path makes sense without the builder explaining it._

### usability/alpha/main-path

Those trying it can use its main path without help from whoever built it.

How it might look:

- A short note for those trying it, covering what it does, how to reach it, and what's missing.
- Clear labels and an obvious flow through the main screens. _(Web app, Mobile app, Desktop app)_
- Example requests for the main endpoints, and a sample of each event it sends. _(HTTP API)_
- `--help` text, and a quick start from installing it to the first real run, including how to set credentials. _(CLI)_
- A README with a quick start. _(Library)_
- A note for those reading its output: what each table or file holds, how often it's refreshed, and what's missing. _(Background job)_

Overengineering:

- Complete docs, accessibility audits.
- Translations, theming, or layouts for devices those trying it don't use. _(Web app, Mobile app, Desktop app)_
- A developer portal, SDKs, or generated client libraries. _(HTTP API)_
- A data catalog, or a full dictionary of every column. _(Background job)_

## Beta

_People can use all of it on their own, mistakes included._

### usability/beta/self-serve

Its users can learn and use it on their own.

How it might look:

- In-app help, or a guide to the main tasks. _(Web app, Mobile app, Desktop app)_
- A reference for every endpoint, such as from an OpenAPI spec, and a sandbox or test keys to try every call without real effects. _(HTTP API)_
- `--help` for every command, and a README. _(CLI)_
- A reference for its config file: every key, its type, its default, and an example. _(CLI, Library)_
- Reference docs for its public API, with types and doc comments editors can show. _(Library)_
- Its inputs, outputs, schedule, and what to do when it fails, with each output table and column described where its readers look. _(Background job)_
- A few real users watched trying the main tasks, and what tripped them up fixed. _(Web app, Mobile app, Desktop app, CLI)_

Overengineering:

- A hosted docs site, interactive tutorial, or setup wizard, when `--help`, a README, and generated reference cover it. _(CLI, Library)_
- A self-serve portal for callers to manage keys or replay webhooks. _(HTTP API)_

### usability/beta/clear-results

Its users can tell what happened after everything they do, and what to do next.

How it might look:

- Loading, empty, and error states handled on purpose. _(Web app, Mobile app, Desktop app)_
- Feedback for every action, showing that it's in progress, done, or failed. _(Web app, Mobile app, Desktop app)_
- Forms that validate as you go, and say what to fix. _(Web app, Mobile app, Desktop app)_
- Each item shows whether it's only on the device, waiting to sync, synced, or failed, and why. _(Mobile app, Desktop app)_
- After an import, a summary of what was read, what was skipped, and why.
- Consistent status codes, and error bodies that say what went wrong without exposing internals. _(HTTP API)_
- Clear error messages, distinct exit codes, output that scripts can parse, and what a long step produced, such as a URL, printed at the end. _(CLI)_
- Errors with a stable code and a message naming the input that was wrong. _(Library)_
- Those reading its output can tell when it last ran, and whether that run was complete. _(Background job)_

### usability/beta/recover

Its users can recover from mistakes.

How it might look:

- Confirmation, undo, or a dry run before anything destructive.
- A confirmation before anything destructive, skipped with `--yes` in scripts. _(CLI)_
- A way to cancel or correct what a call created. _(HTTP API)_
- A run can be repeated for a chosen day or batch, and replaces what that run wrote instead of adding to it. _(Background job)_

### usability/beta/everyone

It works for everyone it's meant for.

How it might look:

- Keyboard access, labels, and contrast. _(Web app, Mobile app, Desktop app)_
- Layouts that work at the screen sizes it's used on. _(Web app, Mobile app)_
- Usable where it's really used, such as in bright light, with gloves, or with one hand. _(Mobile app)_
- Runs without a terminal: no prompts, spinners, or color when output isn't one, and every prompt answerable with a flag or environment variable. _(CLI)_
- Dates, times, numbers, and languages shown the way its users expect.
- Units, currencies, and time zones explicit, such as a currency code with each amount.

Overengineering:

- A formal accessibility audit.
- Translations into languages its users haven't asked for.

## GA

_Polished and consistent everywhere, for everyone it's meant for._

### usability/ga/all-learnable

Everything it does can be learned without asking anyone.

How it might look:

- Complete docs, updated with each release.
- How to install, upgrade, and troubleshoot it. _(Mobile app, Desktop app, CLI, Library)_
- A guide to running it in CI. _(CLI)_
- Docs for every webhook: its events, payloads, how to check its signature, and how retries work. _(HTTP API)_
- A description of every table and column it writes, with units, time zone, and how fresh the data is. _(Background job)_
- How well people complete the main tasks is measured, such as success rates or time taken. _(Web app)_

Overengineering:

- Docs for internals its users never touch.
- In-app analytics to measure how people complete tasks, where watching a few people or the records they submit would do. _(Mobile app, Desktop app)_
- SDKs in several languages, when callers can work from the reference with any HTTP client. _(HTTP API)_

### usability/ga/consistent

It behaves consistently across its whole interface.

How it might look:

- A shared design system, so screens look and behave the same way. _(Web app, Mobile app, Desktop app)_
- Each platform's own conventions, such as back navigation, menus, shortcuts, and file dialogs. _(Mobile app, Desktop app)_
- Naming, paging, and errors consistent across every endpoint. _(HTTP API)_
- The same flag names and output styles across every command, such as `--json` and `--yes` meaning the same everywhere. _(CLI)_
- Names, argument order, units, and error types consistent across every export. _(Library)_
- Table and column names, units, time zones, and what null means kept the same across every table it writes. _(Background job)_

Overengineering:

- A design system of its own, when an established component library already gives it consistency. _(Web app, Mobile app, Desktop app)_
- Screens made identical on iOS and Android against each platform's conventions. _(Mobile app)_

### usability/ga/all-accessible

All of it works for everyone it's meant for.

How it might look:

- Meets WCAG 2.2 AA, or the standard the project has adopted. _(Web app, Mobile app, Desktop app)_
- Works with VoiceOver and TalkBack, and at the largest text sizes the platform offers. _(Mobile app)_
- Plain output when it isn't writing to a terminal or `NO_COLOR` is set, with nothing conveyed by color alone. _(CLI)_
- For software without screens, its docs meet the accessibility standard the project uses. _(HTTP API, Library, Background job)_
- Emails it sends have a plain-text version, and read well with screen readers.

Overengineering:

- WCAG AAA throughout, or a formal third-party audit, unless the client or a law requires it. _(Web app, Mobile app, Desktop app)_
