# Compatibility

**Does it fit with what's around it?** Compatibility is about the systems and people depending on it, and not breaking them as it changes.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Kept apart, so it can't affect anything people rely on._

### compatibility/prototype/opt-in

It reaches only those who chose to try it, and from Alpha, the people the client has agreed it can reach. _(Prototype–Alpha)_

How it might look:

- Behind a feature flag that's off by default.
- Opt-in for invited people or systems.
- Turned on per tenant, only for the tenants who agreed to try it. _(Web app, HTTP API)_
- A separate environment, or a preview deploy. _(Web app, HTTP API, Background job)_
- A preview endpoint, or an experimental command or flag. _(HTTP API, CLI)_
- An experimental module. _(Library)_
- A pre-release version on the registry it's meant for, such as under an npm `next` tag on the client's private registry. _(Library, CLI)_
- A test build, through TestFlight or a Google Play testing track, not the public store listing. _(Mobile app)_
- A separate build or channel, not the public download or its auto-update. _(Desktop app)_
- Emails, messages, and other effects that reach people sent only to an allowlist or a sandbox, or logged without personal data instead of sent.
- At Alpha, real messages sent only to people the client has agreed it can reach, such as a pilot client's customers.
- A setting that stops sending to real people, for one tenant or everyone, without a release. _(Web app, HTTP API, Background job)_
- Where managed devices block unsigned apps, shown on the builder's machine or a VM the client approves. _(Desktop app)_
- Writes to a system others use go to its test environment or a test account.
- Output written to its own tables or schema, and messages read from its own queue or subscription, never one something else relies on. _(Background job)_

Overengineering:

- Supporting several platforms.
- A feature flag system for a new product nobody relies on yet.
- A public store listing, screenshots, or preparing for public review while it reaches only testers. _(Mobile app)_
- Auto-update and signed installers for every platform while builds can be handed straight to the few people trying it. _(Desktop app)_
- Signing builds only to get past managed devices, while it can be shown another way. _(Desktop app)_
- Packaging for every install channel, when one is enough to try it. _(CLI)_

### compatibility/prototype/marked

It's marked with its stability level wherever it's used. _(Prototype–Beta)_

How it might look:

- The stability mark on what it adds. _(Web app, Mobile app, Desktop app)_
- The mark shown to those trying it, never in messages sent to their customers.
- In a public store, the level shown inside the app rather than in its name, since stores may reject a listing marked beta. _(Mobile app)_
- Its level in the window title, About screen, or test build's notes, when the whole app is at that level. _(Mobile app, Desktop app)_
- Its level in the API reference and the OpenAPI description, or a preview path. _(HTTP API)_
- Its level in `--help` and the README, and a notice on stderr when an experimental command runs, so it doesn't break output scripts read. _(CLI)_
- A pre-release version until Beta, then a regular `0.x` release until GA, with `@experimental` or `@beta` on exports newer than the rest. _(Library)_
- Its level in the docs or a comment on each table or file it writes. _(Background job)_

### compatibility/prototype/removable

It can be removed without affecting anything people rely on. _(Prototype–Alpha)_

How it might look:

- Kept to its own code, data, and settings, so taking it out touches nothing else.
- What it created in other systems, such as deployed services, records, or files in people's repos, listed or marked so it can be cleaned up.
- Its own tables or columns, not changes to ones others read. _(Web app, HTTP API, Background job)_
- Reads from systems people rely on with a read-only account, in small batches or off-hours, and writes nothing back. _(Background job)_
- In an app people already have installed, behind a remote flag, with its data kept apart from what the rest of the app keeps. _(Mobile app, Desktop app)_
- Those trying it know where it keeps its data, and how to delete it. _(Mobile app, Desktop app, CLI)_
- New config keys optional, or in a section older versions ignore. _(CLI, Library)_
- No new runtime dependencies, and no code that runs on import, in a package others already install. _(Library)_

Overengineering:

- Compatibility promises, such as stable versions, or announcing changes.
- Versioned URLs, paging, filtering, and a complete error format before those calling it have reacted to its shape. _(HTTP API)_
- A remote flag in a new app nobody relies on yet, when a test build already keeps it apart. _(Mobile app, Desktop app)_
- A version field or migrations for its config file. _(CLI, Library)_

## Alpha

_Breaking changes are fine, as long as those trying it hear about them._

### compatibility/alpha/standards

Its interfaces follow the common standards for what it does, so others can work with it without surprises.

How it might look:

- Standard formats, such as UTF-8 text and RFC 3339 timestamps with a time zone.
- Standard codes, such as ISO 4217 for currencies, ISO 3166 for countries, and E.164 for phone numbers.
- Methods and status codes as HTTP defines them (RFC 9110), and errors as problem details (RFC 9457). _(HTTP API)_
- Output on stdout, errors on stderr, meaningful exit codes, and the platform's usual flag style, with `--help` and `--version`. _(CLI)_
- Semantic HTML, with native controls before custom ones. _(Web app)_
- The platform's conventions for where data is kept, navigation, permissions, and sharing. _(Mobile app, Desktop app)_
- The language's packaging conventions, and semantic versioning. _(Library)_
- Standard schedule syntax, such as cron expressions. _(Background job)_
- Email sent from a domain set up with SPF, DKIM, and DMARC. _(Web app, HTTP API, Background job)_

Overengineering:

- Testing on every platform.
- Builds for operating systems, architectures, or module formats nobody trying it uses. _(Mobile app, Desktop app, CLI, Library)_

### compatibility/alpha/told-of-changes

Those trying it are told about breaking changes, which can still be made freely. _(Alpha only)_

How it might look:

- A message to those trying it, or a note in its changelog.
- A message to those whose reports or queries read its tables, before a column is renamed or removed. _(Background job)_
- Each caller's technical contact known, so breaking changes reach whoever maintains the integration. _(HTTP API)_
- Those trying it in CI pin an exact version, so a breaking release reaches them only when they update. _(CLI, Library)_

Overengineering:

- Deprecation windows.
- Versioned paths, tables, or exports kept alongside the old ones before anyone depends on them staying the same.
- Reading old config formats, or a command that migrates them. _(CLI, Library)_
- A forced-update check, where expiring old test builds is enough. _(Mobile app)_

## Beta

_Breaking changes are rare, and announced._

### compatibility/beta/available

It's available to everyone it's meant for, without opting in, on the platforms and environments it supports.

How it might look:

- The project's normal release pipeline, with no opting in.
- Released to some of those it's meant for first, where the client wants that, then to everyone, with no one needing to opt in.
- Settings that differ between environments kept out of code.
- Supported browsers listed. _(Web app)_
- Supported operating systems, runtime versions, and architectures listed. _(Mobile app, Desktop app, CLI, Library)_
- Installs and upgrades cleanly on each supported platform. _(Mobile app, Desktop app, CLI, Library)_
- The same version on every channel it's offered through, such as npm and Homebrew, released together. _(CLI)_
- Released through the channel its users get apps from, such as the public store, a business store, or the client's device management. _(Mobile app)_
- An installer the client's IT can deploy, such as an MSI or a signed package. _(Desktop app)_
- Writes to the tables its readers use, and sends to every real recipient. _(Background job)_

Overengineering:

- Supporting browsers, OS versions, devices, or runtimes nobody using it has.
- Packages for every package manager before anyone asks for one. _(CLI)_

### compatibility/beta/changes-heard

Those depending on it hear about changes.

How it might look:

- Release notes or a changelog.
- Those depending on it known, such as the teams or services that use it, so changes reach them directly.
- What's new shown in the app after it updates. _(Mobile app, Desktop app)_
- A notice when a newer version is out, checked at most once a day and turned off in CI. _(CLI)_

### compatibility/beta/breaking-rare

Breaking changes are rare, and announced ahead of time where possible. _(Beta only)_

How it might look:

- Versioning, so a breaking change is visible. _(HTTP API, Library, CLI)_
- An OpenAPI description, checked on every change. _(HTTP API)_
- A deliberate change in results for the same input, such as a new rounding rule, released as breaking, even when no signature changed. _(Library)_
- Messages and events checked against their schemas.
- Files it reads keep being accepted, and exports others read change only by adding to them.
- Tables others read checked against their schemas on every change, and changed by adding to them. A change in what a column means is breaking too. _(Background job)_
- Tests that pin `--json` output and exit codes. _(CLI)_
- Its public exports and types checked on every change, such as with an API report. _(Library)_
- A warning when a flag, config key, or export that's going away is used, in a release before it's removed. _(CLI, Library)_

Overengineering:

- Deprecation windows, compatibility guarantees.
- Versioned copies of every output table kept side by side. _(Background job)_
- Running old and new versions of an endpoint side by side. _(HTTP API)_
- Reading every config format it has ever had, instead of a command that migrates old ones. _(CLI, Library)_
- Versioning or documenting the local database's schema, when only the app reads it. _(Mobile app, Desktop app)_

## GA

_Stable: a breaking change is a new version, with time to move over._

### compatibility/ga/stable-interface

Its interface is stable: a breaking change comes as a new version, and the old one keeps working for a set time.

How it might look:

- Semantic versioning, with deprecation warnings before removal. _(HTTP API, CLI, Library)_
- A new version alongside the old one. _(HTTP API)_
- What counts as breaking written down for those integrating, such as that new fields and enum values may appear. _(HTTP API)_
- Deprecation and Sunset headers on what's being retired, with its use logged by caller so those still on it can be told. _(HTTP API)_
- Webhook and event payloads versioned and kept stable like the endpoints. _(HTTP API, Background job)_
- Its commands, flags, exit codes, `--json` output, environment variables, and config format all kept stable, with a deprecated flag still working and warning before it goes. _(CLI)_
- A deliberate change in results for the same input, such as a new rounding rule, treated as breaking, while a fix that makes results match the documented rules is a patch. _(Library)_
- The previous major version still gets bug and security fixes until its deprecation window ends. _(Library, CLI)_
- A changed table or file, or a changed meaning such as amounts that now exclude tax, comes as a new one alongside the old. _(Background job)_
- A reworked workflow offered alongside the old one, with exports, links, and integrations kept working, and pages people already have open still working after a release. _(Web app)_
- In software that updates itself, a reworked flow or format offered alongside the old one until the agreed date. _(Mobile app, Desktop app)_
- Exports marked `@experimental` or `@beta` are outside the semver promise, and the README says so. _(Library)_
- Work already saved, such as a finished report or an exported file, opens and reads the same after an update.
- Older installed versions keep working against the services they use. _(Mobile app, Desktop app, CLI)_

Overengineering:

- Keeping old versions working longer than the agreed window.
- An old flow kept alongside a reworked one in an app only the client's staff use, when the client agrees a switch date and trains them. _(Mobile app, Desktop app)_
- A major version for every changed result, including bug fixes and rule changes everyone must adopt. _(Library)_
- Keeping every old build working forever, instead of a minimum supported version raised with notice. _(Mobile app)_
- A new version of a table for a change that only adds to it. _(Background job)_
- Keeping human-readable output stable, when `--json` is offered for scripts. _(CLI)_
- Keeping the internal layout of its local data stable, when only the app reads it. _(Mobile app, Desktop app)_
- Treating a visual change as breaking, when no workflow, link, export, or integration changes. _(Web app)_

### compatibility/ga/every-platform

It's known to work on every platform it supports.

How it might look:

- Upgrading from every supported version is tested. _(Mobile app, Desktop app, CLI, Library)_
- Tests run on each supported platform and runtime version, installed through each channel it ships on. _(Mobile app, Desktop app, CLI, Library)_
- For software it runs itself, tests on the one runtime it's deployed to, with the browsers or clients it supports listed. _(Web app, HTTP API, Background job)_
- The critical workflows run in each supported browser. _(Web app)_
- Installs and updates on the client's managed machines, such as without admin rights or through their proxy. _(Desktop app)_

Overengineering:

- Supporting or testing platforms nobody using it runs.
- A test matrix of runtime versions for software that only ever runs on the one it's deployed to. _(Web app, HTTP API, Background job)_
- Testing on every device model, instead of the supported OS versions and the models its users carry. _(Mobile app)_
