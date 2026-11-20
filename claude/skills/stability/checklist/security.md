# Security

**Is it protected, and does it follow the rules?** Security is about keeping systems and data safe, and meeting the legal duties that apply.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_The baseline, which holds at every level: secrets and real data stay safe._

### security/prototype/secrets

Secrets stay out of code, version control, sample data, logs, and output. _(baseline, every level)_

How it might look:

- Secrets in the environment or a secret store where it runs, with separate credentials for each environment and providers' test keys everywhere but production. _(Web app, HTTP API, Background job)_
- Nothing secret built into what's shipped to people's devices or published, build-time settings included. _(Web app, Mobile app, Desktop app, CLI, Library)_
- A key that has to ship in the app, such as a maps key, treated as public and restricted to the app and the APIs it calls. _(Web app, Mobile app, Desktop app)_
- Each person's own token, kept in the OS keychain or a file only they can read, never in committed config, and redacted from debug output and errors. _(Mobile app, Desktop app, CLI)_
- Tokens read from a prompt, a file, or an environment variable, never a command-line flag, so they stay out of shell history and logs. In CI, from the CI's secret store. _(CLI)_
- Authorization headers, API keys, and tokens left out of request logs and error reports, for calls it receives and calls it makes.
- API keys it issues stored hashed, and shown once. _(HTTP API)_
- Signing keys, store credentials, and publish tokens kept in CI's secret store, never in the repository or the project's `.npmrc`. _(Mobile app, Desktop app, CLI, Library)_
- Only the files meant to be published go into the package, such as with a files allowlist, so no `.env` or `.npmrc` ships. _(Library, CLI)_
- At first, the token read only from an environment variable, so the tool stores nothing and needs no sign-in. _(CLI)_
- Generated or scaffolded files refer to secrets by name, never contain their values.

Overengineering:

- Keychain integration on every OS, when a file only the person can read keeps the token safe. _(CLI, Desktop app)_

### security/prototype/real-data

Real data comes only through the real systems, never copied into code or fixtures. _(baseline, every level)_

How it might look:

- Sample data that's made up, not exported from production.
- Made-up names, phone numbers, and email addresses that can't belong to anyone, such as `example.com` addresses.
- Made-up files in the formats it reads, with the structure of a real one and none of its values.
- When a real file or record shows a problem, a made-up one that reproduces it goes into the tests, never the real one.
- Recorded responses from a real service scrubbed of real names, IDs, and tokens before they're committed.
- Public reference data, such as published tax rates or postal codes, used as is. A client's own records aren't.
- Real data kept out of notebooks, spreadsheets, chat, and local copies outside the systems approved for it.

### security/prototype/access

Only those allowed can reach real systems or data, and one tenant or account never sees or changes another's. _(baseline, every level)_

How it might look:

- Every read and write scoped to the tenant or account making it.
- A test, run on every change, that one tenant can't read or change another's records, even by guessing an ID.
- Work that uses real credentials isn't open to anyone who finds it, such as by running locally, on a private preview, or behind a shared password, even when its own data is made up.
- Real data encrypted in transit, and at rest where it's stored, such as with HTTPS and a managed database's default encryption.
- Sign-in on everything that reaches real data, with the tenant taken from the session on the server, never from what the browser sends. _(Web app, HTTP API)_
- Each API key tied to one caller, with the caller taken from the key, never from the request. _(HTTP API)_
- Webhooks and other messages carry only the receiving tenant's data, and go only to addresses that tenant registered. _(HTTP API, Background job)_
- Callbacks from providers, such as delivery receipts or payment events, checked as coming from that provider before they change real data. _(Web app, HTTP API, Background job)_
- Each person or pipeline uses credentials issued to it, never a shared or built-in one, sent only to the hosts they belong to. _(Mobile app, Desktop app, CLI)_
- The host a token is sent to comes from the tool's settings or a flag, never from a config file in a repo someone cloned. _(CLI)_
- With no sign-in of its own, data kept in the OS user's own profile, where only they can read it, on a device with disk encryption. _(Desktop app, CLI)_
- Where it has a sign-in of its own, data on the device belongs to the signed-in person, and work not yet sent is kept for them, not shown to the next person, when they sign out. _(Mobile app, Desktop app)_
- Its own read-only account on each system it only reads, made for the work and easy to revoke, and separate storage and output for each client it works for. _(Background job)_
- No state kept between calls that could carry one caller's data to another. _(Library)_
- Published only by CI or named people, to the registry it's meant for, so it can't go public by mistake. _(Library, CLI)_
- Where a client's system offers only a shared key, a test account until it issues credentials for each person or pipeline, before real data flows.
- Copies it writes where others read them, such as warehouse tables, readable only by the people and roles allowed to see that data at its source. _(Background job)_
- A leaked API key stopped by disabling its record, until callers can rotate their own. _(HTTP API)_
- Established libraries for authentication and cryptography, never hand-rolled ones.

Overengineering:

- Sign-in or tenant separation for work that holds only made-up data and reaches no real system.
- Roles and permissions beyond what keeping real data apart needs, unless the client's rules call for them.
- A separate database or deployment for each tenant, when scoping every query keeps them apart and no client requires more.
- A sign-in of its own, or encrypting local data, on a device the OS account and disk encryption already protect, unless the client's rules call for it. _(Mobile app, Desktop app)_
- Certificate pinning, jailbreak or root detection, or app attestation, unless the client's rules call for them. _(Mobile app)_
- Key scopes, scheduled rotation, or self-serve key management before anyone relies on it. _(HTTP API)_

### security/prototype/personal-data

Personal data is collected only where it's needed. _(baseline, every level)_

How it might look:

- Only the fields the work needs, with nothing kept just in case.
- Functions and APIs take only what they need, such as a postal code instead of a whole customer record.
- Personal data kept out of logs, error reports, crash reports, and analytics.
- Error messages and debug output name what's wrong without repeating the personal data passed in.
- Emails, texts, notifications, and webhooks carry no more personal data than they need.
- Third-party services sent only what they need to do their part.
- Copies it makes, such as reporting tables or exports, carry only the fields their readers need. _(Background job)_
- Location and other metadata removed from photos and files, unless the work needs them. _(Web app, Mobile app, Desktop app)_
- Device permissions, such as the camera or location, asked for only when a feature needs them. _(Mobile app)_
- The legal duties that come with it, such as consent, or an agreement with a provider that processes it, in place before the first real personal data arrives.

### security/prototype/untrusted-input

Untrusted input can't run as code or commands. _(baseline, every level)_

How it might look:

- Parameterized queries, escaped output, and no shell commands built from input.
- Data from other systems treated as untrusted, even the client's own, and escaped when put into emails, messages, or pages.
- Addresses callers give it, such as webhook URLs, checked so they can't reach internal systems or cloud metadata. _(Web app, HTTP API, Background job)_
- Sort and filter parameters matched against a list of allowed fields, never pasted into a query. _(HTTP API, Web app)_
- No raw HTML rendered from input. _(Web app)_
- Config and rule files parsed as data, never evaluated as code.
- Config read with a safe parser, with keys like `__proto__` ignored when merging. _(Library, CLI)_
- Paths from config files and templates kept inside the repo, so they can't write anywhere else. _(CLI)_
- Commands a repo's own config names run like that repo's own scripts: shown first, with values passed as arguments, and without tokens they don't need. _(CLI)_
- Webhook addresses set by the team for each invited caller, instead of taken from requests, until callers can set their own. _(HTTP API)_
- XML read with external entities and DTDs turned off.
- Cells that could run as formulas escaped in exported spreadsheets.
- Anything that updates itself installs only updates signed by the publisher. Until then, builds are installed by hand. _(Desktop app, CLI)_
- In an Electron app, windows load only its own content, with context isolation on and Node.js integration off, and links open in the system browser. _(Desktop app)_
- Deep links, scanned codes, and notification payloads treated as data, and web views that don't give pages from the network access to native code. _(Mobile app)_

Overengineering:

- A self-update mechanism of its own, when a package manager installs and updates it. _(CLI)_

## Alpha

_What it depends on is safe to depend on._

### security/alpha/stable-dependencies

It depends on nothing less stable than itself.

How it might look:

- Nothing it relies on is at an earlier stability level, such as a prototype service, or a third party's preview API.
- A third party or client system with no stated level judged by what it promises, such as a preview label, a deprecation policy, or a support agreement.
- A widely used package that stays at 0.x judged by how it's maintained and how it changes, not by its version number.
- Endpoints or services still being built alongside it counted as early, and written down as a known gap until they settle.

### security/alpha/maintained-dependencies

Its dependencies are maintained, and licensed for how it's used.

How it might look:

- Licenses checked when a dependency is added, against how it's distributed, such as copyleft code in an app-store app.
- Dependency versions pinned, such as with a lock file. _(Web app, HTTP API, Background job, Mobile app, Desktop app)_
- A lock file published with it, such as `npm-shrinkwrap.json`, or a bundled build, so people install the versions that were tested. _(CLI)_
- A lock file for its own tests, with runtime dependencies as ranges so the software using it can resolve them, and as few of them as it can manage. _(Library)_
- A supported Electron release, since older ones stop getting security fixes. _(Desktop app)_

Overengineering:

- Security reviews, compliance audits.
- Certifications, such as SOC 2, that no client or law requires yet.
- Rewriting a well-maintained dependency in-house just to have fewer dependencies.

## Beta

_Protected against misuse, as far as it's exposed._

### security/beta/misuse

It's protected against misuse at the level it's exposed.

How it might look:

- Rate limits and input size limits. _(HTTP API, Web app)_
- Protection against cross-site request forgery. _(Web app)_
- Sign-in attempts limited, session cookies secure and HTTP-only, and sessions that expire. _(Web app)_
- The checks in OWASP ASVS level 1. _(Web app, HTTP API)_
- The OWASP API Security Top 10 checked, such as object-level authorization and mass assignment. _(HTTP API)_
- What it receives is checked for who sent it, such as signed webhooks. _(HTTP API, Background job)_
- Webhooks it sends are signed. _(HTTP API, Background job)_
- API keys that can be rotated and revoked. _(HTTP API)_
- Only its own services can put work on its queue, and a message can trigger only what it's for. _(Background job)_
- Links sent in messages use unguessable tokens that expire, and open only what they're for.
- A lost or stolen device can't be used to reach the data, through the platform's storage protection and sessions that can be revoked. _(Mobile app, Desktop app)_
- Windows reach the system only through a few named calls, such as an Electron preload script, each checking its arguments, under a content security policy. _(Desktop app)_
- Signed releases, so the platform or package manager can verify them. _(Mobile app, Desktop app, CLI, Library)_
- Known vulnerabilities in dependencies found and fixed.

Overengineering:

- A web application firewall, bot detection, or a bug bounty, before its exposure or a real incident calls for one. _(Web app, HTTP API)_
- OAuth, mutual TLS, or request signing for callers when keys over TLS are enough and nobody asks for more. _(HTTP API)_
- Encrypting local data with a key of its own, when disk encryption and the OS account already protect it and no duty asks for more. _(Mobile app, Desktop app)_
- Certificate pinning, jailbreak detection, or code obfuscation without a specific threat that calls for them. _(Mobile app)_
- Sessions so short that people have to sign in again while offline. _(Mobile app)_
- Rate limiting inside the tool, when the service it calls enforces its own limits. _(CLI)_

### security/beta/duties

The legal and regulatory duties that apply to it are known and met.

How it might look:

- Privacy duties met, such as exporting or deleting someone's personal data on request, with a deletion at the source reaching the copies it made.
- Which duties apply confirmed with the client, such as whether its data counts as health records.
- Rates and rules carry the dates they take effect, so a change in the law ships before it applies. _(Library)_
- Rules about where data may be stored, and for how long.
- Consent and opt-out for the messages it sends, such as honoring STOP replies, where the law requires them. Transactional messages, such as order confirmations, often don't need an opt-out.
- An audit log where a law requires one, such as for health records.
- Where data stays on people's devices, how long it's kept, and how it's deleted written down, so the client can meet its own duties. _(Mobile app, Desktop app, CLI)_
- The app stores' privacy details match what it collects. _(Mobile app)_
- The source of each rule it implements, such as a law or regulation, recorded next to the rule. _(Library)_

Overengineering:

- A certification, such as SOC 2 or HITRUST, before a client or law requires one.
- Column-level encryption or tokenized personal data that no duty or reader calls for.

## GA

_Its sensitive parts are reviewed, and actions can be traced._

### security/ga/sound

Its handling of credentials, payments, personal data, and permissions is known to be sound.

How it might look:

- A security review of sign-in and permission checks.
- The checks in OWASP ASVS level 2. _(Web app, HTTP API)_
- The checks in OWASP MASVS that fit its risk, such as for storage, network, and authentication. _(Mobile app)_
- A review of how it gets, stores, sends, and refreshes tokens, including what debug output shows. _(Mobile app, Desktop app, CLI)_
- A review of what its windows can ask the rest of the app to do, and of how updates are checked before they install. _(Desktop app)_
- A review of how API keys are issued, rotated, and revoked. _(HTTP API)_
- Data and logs kept no longer than they're needed, backups included.
- Builds that can be traced to their source, with a list of what's in them, such as SLSA provenance and an SBOM, where others install or build on it.
- Keys that sign releases kept in CI's secret store or a hardware module, usable by few people, with a plan for if one leaks. _(Mobile app, Desktop app, CLI, Library)_
- A penetration test, where the risk warrants one.

Overengineering:

- A penetration test or formal certification that neither the risk nor the client's duties call for.
- Encrypting individual fields on top of encryption at rest, unless a duty or a specific risk calls for it.
- Root and jailbreak detection, obfuscation, or anti-tampering beyond what the client's risk calls for. _(Mobile app)_
- Mutual TLS or signed requests from callers, when API keys over TLS meet the risk. _(HTTP API)_

### security/ga/least-access

Each part has only the access it needs.

How it might look:

- Where a service offers no narrower token, a separate revocable token for each pipeline, with the missing scopes written down as a gap in that service.
- Service accounts and tokens limited to what each part needs.
- People's access to production limited to those who operate it, and removed when they leave.
- The app's database user can't change the schema, and migrations run as a separate one. _(Web app, HTTP API, Background job)_
- Read-only access to the systems it reads from, and write access only to its own tables. _(Background job)_
- The narrowest token scope each command needs, and a deploy-only token in CI. _(CLI)_
- A publish token that can publish only this package, kept in CI. _(Library, CLI)_
- Only the device permissions it uses. _(Mobile app, Desktop app)_

### security/ga/traceable

Sensitive actions can be traced to who did them.

How it might look:

- The sensitive actions agreed, such as changing permissions, exporting data, or moving money.
- An audit log of sensitive actions.
- Staff reaching a client's data, such as for support, logged and visible to that client.
- Who created, rotated, or revoked each API key, or changed a webhook address, recorded. _(HTTP API)_
- Manual reruns, backfills, and replayed messages recorded with who started them. _(Background job)_
- Who's acting, such as the person or the CI run, sent with each sensitive call, so the service's audit log records it. _(Mobile app, Desktop app, CLI)_
- Who published each version, and from which commit, recorded. _(Library, CLI)_

Overengineering:

- Logging every action, instead of the sensitive ones.
- A separate tamper-proof audit store, when the client's duties don't call for one.
- An audit log of its own on people's devices or inside a library, when the service it calls or the software using it records who did what. _(Mobile app, Desktop app, CLI, Library)_
