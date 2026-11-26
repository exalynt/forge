---
name: implement
description: Implement a feature, improvement, or bug fix (features and improvements at a target stability level) in the current checkout on the current branch, leaving the changes uncommitted. Only when explicitly asked, work on a new branch in its own git worktree instead, so several sessions can work in one repository on the same machine without stepping on each other, then commit, push, and open a pull request against the default branch labeled with that level.
disable-model-invocation: true
---

# Implement

The user's request is the input. It names the task, and may name an effort
(`low`, `medium`, or `high`), a target level (`prototype`, `alpha`, `beta`, or
`ga`), and ask for a worktree, for
example "implement alpha: add reconnecting by email to conversations" or
"implement beta in a worktree: add CSV export to invoices". A bug fix takes no
level, as in "implement: fix the conversation popout overflowing on phones".
If no task is given, ask. If the task isn't a bug fix and the level is
missing, the target is the capability's current level, found in step 1. Tell
the user which level that is. If the level is invalid, ask. If the session
spans several repositories and the task doesn't say which one it's in, ask.

The level is the one the change is aiming for. It sets how much the work does
and what review holds it to: enough to meet that level, and nothing that
belongs to a later one.

The effort sets how much the review in step 5 covers, as the
`stability-review` skill describes. Without one, use `medium`. It doesn't
change how much the work does, which is the level's job, and a bug fix has no
review for it to set. If it could be read as part of the task, as in
"implement high contrast mode", ask.

## Where the work goes

By default the task happens where the session already is: the current
repository or worktree, on the current branch, even if that's the default
branch. Invoking this skill isn't a request for anything more, so don't create
a branch or worktree, commit, push, or open a pull request.

Those happen only when the user explicitly asks:

- **A worktree** ("in a new worktree", "in a worktree"): the task gets its own
  branch, its own worktree, and a pull request, so other sessions working in
  the same repository keep their own working tree, index, and branch. What a
  worktree doesn't separate is covered in step 4.
- **Only a new branch or only a pull request**: do just that part, in the
  current checkout. Create a new branch with `git switch -c <branch>`, named as
  step 3 describes. The `pr` skill refuses the default branch, so if a pull
  request is asked for while on it, ask whether to create a branch first.

If "worktree" could be read as part of the task, ask.

## 1. Settle the target level

A bug fix restores how something should already work, so it has no target
level: skip this step. If a level was given for one, say bug fixes don't take
one and leave it off. The capability's level still decides, as always, whether
the fix may break things.

This only reads, so do it before creating anything.

Find the capabilities the task touches and each one's current level, as
forge's `AGENTS.md` "Where intention lives" says:

- **With Docqet**: run `docqet touches` on the files and folders you expect to
  change. It lists the capabilities involved, each one's level, its rules,
  gaps, tradeoffs, and open decisions. When the task is described in product
  terms, find its capabilities with `docqet list capabilities` first. Read
  each one with `docqet show`. Follow the `docqet` skill's "Before a change".
- **Without Docqet**: the level is the one the user gave. If they gave none
  and the answer would change the work, ask. Otherwise go on without one.

Tell the user which level you found for each capability and where it came
from. Then compare the target against each one:

- **Same level**: go on. More Alpha work on an Alpha capability is the usual
  case.
- **Above the current level**: this moves the capability up. Confirm that's
  intended. If it is, do the work in step 5 as `../promote/SKILL.md` (from
  this skill's directory) describes: only the delta, then a second review,
  then recording the new level. Promotion needs Docqet. Without it there's no
  recorded level to move: do the work at the level the user gave.
- **Below the current level**: changes to an existing capability have to hold
  up to its level, so work on a GA capability is GA work. The exception is a
  task that creates a new capability, or a new version of a GA one, which gets
  its own level. Ask which this is. With Docqet, a new capability is written
  the `docqet` skill's way, at the level the user gives. Otherwise use the
  current level.
- **Shared code at a stricter level**, such as a table, endpoint, or
  component other capabilities use: the whole change is held to the highest
  level it touches. Say so, and offer to split the change if that's easy.

With Docqet, read what the target level expects before planning the work,
with `docqet rubric --level <level>`, and the next level's, so you can
recognize work that belongs later.

## 2. Read the repository

In the repository the task is for, run `git branch --show-current` and
`git status --short`. Uncommitted changes already there aren't this task's.
Leave them as they are, and don't stash, revert, or commit them.

Without a worktree, tell the user which branch the work goes on and skip to
step 5. If HEAD is detached, ask first.

With a worktree:

- Find the default branch: `git symbolic-ref --short refs/remotes/origin/HEAD`
  (for example `origin/main`). Fetch it with `git fetch origin <default>`.
- Run `git worktree list` and `git branch -a`. Other sessions' worktrees
  appear there. Never switch, remove, prune, or commit in a worktree that
  isn't this task's.
- If this session is already in a worktree for a different task, stop and ask.
  Don't nest one worktree inside another.
- If the local default branch has commits that aren't on `origin/<default>`,
  say so and ask whether to branch from the local branch (those commits then
  become part of the pull request) or from `origin/<default>`.

Don't carry the main checkout's uncommitted changes into the worktree.

## 3. Create the worktree

Only with a worktree, as are step 4 and the cleanup at the end.

- **Branch name**: `<type>/<slug>`. Match the prefixes the repository's
  existing branches use. Without a pattern, use `feat/`, `fix/`, `improve/`,
  `docs/`, or `chore/`. The slug is a few words in kebab-case naming the
  change, such as `fix/popout-phone-overflow`. If the name is taken, pick
  another one, and don't reuse the existing branch.
- **Path**: `<repo>/.worktrees/<type>-<slug>`.
- Check that path is ignored: `git check-ignore -q .worktrees/x`. If it isn't,
  add `/.worktrees/` to the root `.gitignore`, and say so in the report.
- Create it from the base chosen in step 2:
  `git worktree add --no-track -b <branch> <path> <base>`.
- Work in it from here on. If your agent can move its session into the
  worktree, do. Otherwise use absolute paths under the worktree for every edit
  and `git -C <path>` for every git command, and check each one points into
  the worktree, not the main checkout.

## 4. Set it up without touching shared state

A worktree is a fresh checkout: it has the repository's tracked files and
nothing else.

- **Dependencies**: install them inside the worktree from the lock file, with
  `npm ci`, `go mod download`, or whatever the project uses. Don't reuse or
  modify the main checkout's `node_modules` or build output. Change the lock
  file only when the task needs a dependency change.
- **Ignored local config**: if the project's checks need gitignored files
  from the main checkout, such as `.env` or `.npmrc`, copy just those files
  into the worktree. Don't print their contents, and never commit them.
- **Servers, ports, databases**: follow the project's rules. Some forbid
  starting a dev server. Where the task does need one, don't take a port,
  database, or container another session may be using. Use a free port, and a
  database named for the branch, or ask. Stop what you started when you're
  done with it.

## 5. Do the work

With Docqet, before changing code, read the components and mechanisms
`docqet touches` names for it, with `docqet show`, including another
project's when the product spans several. Work with their tradeoffs, and the
code their `docqet:tradeoff` comments mark. If the task means undoing one, say
so. Without Docqet, nothing like this is written down: read the code.

Do the task the way the project's instructions and forge's `AGENTS.md`
describe, sized to the target level: meet what it expects, and leave out what
first appears at a later level. With Docqet, don't close the capability's gaps
along the way unless that's the task, and keep its docs and annotations
current in the same change, as the `docqet` skill says. Ask when a decision is
the user's to make. A task that implements, changes, or clarifies a business
rule adds or updates its acceptance tests, and a bug fix or performance work
doesn't (see the `testing` skill). A library has no acceptance tests unless
the user asks. Acceptance tests, and with Docqet the rule docs they verify,
need the user's explicit approval before you add, change, or remove any:
propose them before writing them.

When the work is done:

- Run the project's checks (lint, format, build, tests, as its `AGENTS.md` or
  other instructions list them) and fix what fails.
- Except for a bug fix, review the change against the target level:
  - **With Docqet**, run the `stability-review` skill on this task's changes
    against the target level, at the effort the user asked for, or `medium`,
    in a subagent if your agent can run one, otherwise yourself. At low, a
    smaller, cheaper model is fine for it, if your agent lets you choose. In a
    worktree, the changes are `git diff $(git merge-base <base> HEAD)`, which
    includes uncommitted work. Otherwise they're the uncommitted changes and
    new files this task made, not what step 2 found already there. Resolve
    its REQUIRED NOW and SECURITY/PRIVACY findings. A REQUIRED NOW finding
    can instead be recorded as a gap the Docqet kit's way, if the user agrees. Remove work this change
    added that it flags UNNECESSARY NOW. Findings about code this change
    didn't touch go in the report, not in the change. For a move,
    `../promote/SKILL.md` already has this review.
  - **Without Docqet**, there's no level to review against. Check the change
    yourself against the user's request, the security and privacy baseline,
    and forge's engineering philosophy, and remove what the request doesn't
    need.

## 6. Finish

With Docqet, finish the way the `docqet` skill's "Before you finish"
describes: `docqet check --changed` clean for what you changed, `docqet diff`,
and a short Docqet section in the report.

**By default**, leave the changes uncommitted on the current branch. Report what
changed, the checks and the stability review's outcome, and that nothing was
committed. The user can follow up with the `commit` skill, or the `pr` skill
from a branch other than the default.

**When a pull request was asked for**, including with a worktree, follow
`../pr/SKILL.md` from this skill's directory, with the target level and
whether this is a move, or with no level for a bug fix: commit, push the
branch, and open a pull request against the default branch. The pull request
is where review happens. To make follow-up changes, edit in the same place and
follow `../pr/SKILL.md` again, which pushes them to the same pull request.

Stay in the worktree when you're done. Don't remove it, merge the pull request,
or switch back to the default branch unless asked.

## Cleaning up

Only for a worktree, and only when the user asks, usually once the pull request
is merged or abandoned:

1. If your session moved into the worktree, move it back to the main checkout
   first.
2. Check it has nothing unpushed: `git -C <path> status --short` and
   `git -C <path> log origin/<branch>..HEAD`. If anything is there, show it
   and ask before going on.
3. `git worktree remove <path>`. If it's locked, unlock it first with
   `git worktree unlock <path>`.
4. Delete the local branch. If `gh pr view <branch> --json state` says
   `MERGED`, use `git branch -D <branch>`, since a squash merge makes
   `git branch -d` refuse. Otherwise ask before deleting it. Leave the remote
   branch alone unless asked.
