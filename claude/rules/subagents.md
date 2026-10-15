# Subagents

Use subagents to run independent work in parallel when that saves real time or
keeps a large search out of the main conversation. This is standing permission
to spawn them, in every skill and task, under the conditions below. The
condition that matters most: the pieces must be separate enough that the
agents can't step on each other.

## Good fits

- **Research and exploration**: different areas of a codebase, different
  repositories, docs versus code versus tests, separate web questions.
- **Read-only reviews and checks**: one `stability-reviewer` per feature,
  reading each layer of a feature, test suites that share no state.
- **Changes in disjoint files**, when each piece is substantial and none
  depends on another's result.

## Keep in the main conversation

- **Work that shares mutable state**: the same files, the git index (staging,
  commits, branches), the same database or test data, a running server or
  browser session, ports, package installs and lock files, generated code.
- **Sequential work**, where one step's result decides the next.
- **Small work**, such as a few searches or file reads. A subagent starts cold,
  so it costs more than it saves there.
- **Questions for the user.** Subagents can't ask. The main conversation
  gathers what needs a decision and asks.

When unsure whether two pieces would interfere, keep them in the main
conversation.

## How

- Split the work so each agent owns a distinct area. State that boundary in its
  prompt, and say so when it's read-only.
- Give each agent the context it needs, since it starts cold, and the shape of
  the result you want back.
- Launch independent agents in the same message so they run in parallel.
- Prefer the agent built for the job: `Explore` for read-only searches,
  `stability-reviewer` for stability reviews.
- Their results come back to you, not to the user. Check claims that matter
  before acting on them, and merge them into one answer.
