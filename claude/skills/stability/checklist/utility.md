# Utility

**Is it worth building, and is this the right way to build it?** Utility is about value: whether the work meets a real need, in a way that fits.

Examples of how each outcome might look, and of work that may be overengineering, by level. Indications of boundaries, not instructions or complete lists. Tags are kinds of software an example may suit. Generated from the readme's `src/data/stability-checklist.ts`; change it there first.

## Prototype

_Just enough for someone to judge the idea._

### utility/prototype/judged

Those judging the idea can try it, or watch it run, enough to tell whether it's worth doing and whether this is the way to do it.

How it might look:

- The question it's meant to answer written down before it's built, and the answer once it's judged.
- A clickable UI with fake or hard-coded data. _(Web app, Mobile app, Desktop app)_
- Two or three versions of the key screen, to choose between. _(Web app, Mobile app, Desktop app)_
- A stub endpoint or command that returns realistic output, or a mock server from a draft API description. _(HTTP API, CLI)_
- One run by hand on a recent day's records, with its output shown next to the source. _(Background job)_
- A draft of the main calls and the config format, tried by one team that would use it, in a branch of their code. _(Library)_
- A script or notebook that runs the core idea on made-up or sample data.
- A spike against the one real integration the idea depends on, when that's the open question.
- A dry run, a sandbox, or a test account when trying it would otherwise change a real system or reach real people.

Overengineering:

- Polishing anything the idea doesn't depend on.
- Building out the parts of the system that aren't the question.
- Packaging an installed app when a script on sample files answers the question. _(Mobile app, Desktop app)_
- A production-like environment, infrastructure as code, or a deploy pipeline, when a local run or one preview lets people try it.
- Sending real emails or texts when the question is the workflow, and showing the message on screen would answer it.
- Putting it on a schedule, when running it by hand answers the question. _(Background job)_
- Building for several platforms, such as both iOS and Android, when one answers the question. _(Mobile app, Desktop app)_

## Alpha

_Real use on the real system, to learn what it should become._

### utility/alpha/end-to-end

It works end to end on the real system.

How it might look:

- Wired to real storage and services, with sample data where real data isn't needed.
- Installed the way its users will get it, from a packaged build or the real registry, not run from source. _(Mobile app, Desktop app, CLI, Library)_

### utility/alpha/real-use

It's used for real.

How it might look:

- A few people, teams, or systems using it for real work, or a job running on a real schedule.
- A pilot with one client or team doing real work, with their usual way of working kept as a fallback.
- Run alongside what it replaces, on real input, with the results compared before anything acts on them, when a wrong result would cost money or reach customers.
- Each consuming service keeps its old calculation and chooses between the two, so it can go back without a release. _(Library)_
- Whether it's being used, and how, is known, such as from its logs or by asking those using it.

Overengineering:

- Usage tracking built into what people install. _(Mobile app, Desktop app, CLI, Library)_

### utility/alpha/learning

What's learned from using it shapes what comes next.

How it might look:

- Notes on what worked and what didn't, feeding the next round of changes.

Overengineering:

- Settling its shape. Big changes are still expected.

## Beta

_Confidence that the need is real and the approach is right. From here, it's refinement._

### utility/beta/confirmed

The need and the solution are confirmed, so what changes from here is refinement.

How it might look:

- Feedback is about details, not direction.
- Those integrating have run their real flows through it without needing a change to its shape. _(HTTP API, Library)_

### utility/beta/success-agreed

What success looks like is agreed.

How it might look:

- Success measures agreed with whoever asked for it, such as adoption, time saved, error rates, or jobs completed.
- Today's numbers measured before it's widely used, such as the current no-show rate, so there's something to compare with.
- Use measured from the services it calls, such as requests by its version, not from people's devices. _(Mobile app, Desktop app, CLI)_

Overengineering:

- Analytics built into what people install, when asking them or counting what it produces would do. _(Mobile app, Desktop app, CLI, Library)_

## GA

_Proof, from real use, that it delivers what it's for._

### utility/ga/delivers

It's shown to deliver what it's for in real use.

How it might look:

- When it's too early to show, such as for a yearly workflow, a known gap with the date it will be shown by.
- Usage or outcome numbers that meet the measures agreed at Beta.
- The systems it was built for rely on it, and the copies of the logic it replaced are retired. _(Library)_

Overengineering:

- An analytics pipeline built only to prove value, when the numbers can come from its own data or the client's reports.
