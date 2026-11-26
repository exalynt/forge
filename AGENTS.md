## Forge

This repository follows Exalynt's way of working, installed by forge: how we
engineer, and skills for the workflows that apply it. Forge says how to do the
work. How much a capability needs, and what it's meant to do, is the
project's intention (see "Where intention lives").

### Engineering philosophy

Build only as much software as we currently have reason to build. The best
implementation is the simplest one that appropriately meets the capability's
needs and stability level today: appropriate complexity, not minimal code.

- A capability's stability level is both a bar the work must meet and a
  ceiling on how much to invest. Meeting it is required. Exceeding it is not a
  virtue, and over-engineering is as much a problem as under-engineering.
- Build against concrete, current requirements. "We'll probably need it
  eventually" isn't a reason to build it now.
- Complexity needs a reason that exists now: a demonstrated failure, a real
  second use case, a measured performance problem, an explicit requirement, a
  security or privacy need, or what the capability's level expects.
- Add an abstraction for a problem that exists now: real duplication, a second
  real implementation, or a boundary that's genuinely hard to reason about
  without it. No interfaces, extension points, plugin mechanisms, or
  configuration because they could be useful someday. A few similar concrete
  blocks often beat one premature abstraction.
- When uncertain, prefer the decision that's cheapest to reverse, and code
  that's easy to delete or replace.
- Replace early work once it has served its purpose. Don't contort new work to
  stay compatible with code that was always meant to be temporary.
- Optimize, and add operational infrastructure (dashboards, alerting,
  elaborate retries, caches, queues, flag systems, scaling machinery), on
  evidence of a need or when the level expects it, not on intuition about
  future scale.
- Decide what not to build, and say so briefly. A plainly stated known gap is
  part of a complete answer. When existing code already exceeds what's needed,
  point that out instead of extending the pattern.
- Simple doesn't mean sloppy. What we build is correct and clear.

### Security and privacy

Never speculative, the same at every level, and never a known gap:

- One tenant or client can't see or interact with another's real data.
- Anything touching real client systems or data has authentication and
  authorization. Only work that holds nothing but sample data can go without.
- Real secrets are never leaked anywhere: not in code, version control, sample
  data, logs, error messages, output, client bundles, or third-party services.
- Real client data is never copied into code or fixtures.
- Personal data is collected only where it's needed.
- Untrusted input can't run as code.

Sample, mock, or fake data and throwaway placeholder credentials aren't
covered.

### Git

Don't create a branch unless asked. Work on the current branch. The
`implement` skill does too, unless asked for a worktree.

### Subagents

If your agent can run subagents, use them for independent work in parallel
when that saves real time or keeps a large search out of the main session:
research across separate areas, read-only reviews, changes in disjoint files.
Keep work that shares state (the same files, the git index, a database or test
data, a running server or browser, ports, installs) and questions for the user
in the main session. Without subagents, do the same work yourself in sequence.

### Where intention lives

A capability is one result of value, and the thing that carries a stability
level: draft, prototype, alpha, beta, or ga. A feature is a group of
capabilities.

**With Docket**, when there's a `docket.yaml` at the working directory or
above it, within the repository: the Docket kit, its block in this file and
its `docket` skill, is the authority on each capability's level, its business
rules, the personas it serves, its gaps, tradeoffs, and open decisions, and
what each level expects (`docket rubric`). Read them with `docket touches`,
`docket show`, and `docket list`. Forge's skills record intention only the
kit's way, and only after a person confirms it.

**Without Docket**: a capability's level is whatever the user says. Ask only
when the answer changes the work. Nothing about intention is recorded: no
levels, rules, personas, or gaps written anywhere. The only record is the code
and tests the user approves. The `stability-review` and `promote` skills need
Docket, and say so.

### Precedence

On intention, Docket's instructions win over forge's. This project's own
instructions win over both, except for Docket's rule that nothing an agent
infers counts as declared until a person confirms it. Project-specific
architecture and domain rules belong in the project's own instructions,
outside this block.

### Skills

Forge's skills are in `{{skills}}/<name>/SKILL.md`. An agent that doesn't
load skills from there reads a skill's `SKILL.md` when its work comes up, or
when someone asks for it by name ("use implement", "/implement"). The ones marked "on request" run only when someone asks.

- `implement` (on request): build a feature, improvement, or fix on the
  current branch, or in its own worktree and pull request when asked.
- `commit` (on request): commit changes as small, related commits.
- `pr` (on request): commit, push, and open or update a pull request labeled
  with its stability level.
- `qa` (on request): QA a capability end to end against its business rules.
- `ux` (on request): walk a main flow as a type of user and report the
  friction.
- `use` (on request): usability test a flow with one player per person.
- `stability-review`: review work against a level, for what's missing and
  what's premature. Needs Docket.
- `promote` (on request): move a capability to a higher level, doing only the
  delta. Needs Docket.
- `testing`: how we write and organize tests, whenever you write or change
  them.
- `observability`: how we add traces, logs, and metrics.
- Stack skills, such as `go-http` or `react`, where installed: conventions for
  that kind of code.
