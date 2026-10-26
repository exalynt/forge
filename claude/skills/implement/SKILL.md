---
name: implement
description: Implement a feature, improvement, or bug fix (features and improvements at a target stability level) on a new branch in its own git worktree, so several sessions can work in one repository on the same machine without stepping on each other, then commit, push, and open a pull request against the default branch labeled with that level.
argument-hint: "[prototype|alpha|beta|ga] <task>"
disable-model-invocation: true
---

# Implement

Arguments: `$ARGUMENTS`. Expected format: `<level> <task>`, where the level is
`prototype`, `alpha`, `beta`, or `ga`, for example
`/implement alpha add reconnecting by email to conversations` or
`/implement beta add CSV export to invoices`. A bug fix takes no level, as in
`/implement fix the conversation popout overflowing on phones`. If no task is
given, ask. If the task isn't a bug fix and the level is missing, the target
is the feature's current level, found in step 1. Tell the user which level
that is. If the level is invalid, ask. If the
session spans several repositories and the task doesn't say which one it's
in, ask.

The level is the one the change is aiming for. It sets how much the work does
and what review holds it to: enough to meet that level, and nothing that
belongs to a later one.

Invoking this skill is the request to create a branch. Each task gets its own
branch and its own worktree, so other sessions working in the same repository
keep their own working tree, index, and branch. What a worktree doesn't
separate is covered in step 4.

## 1. Settle the target level

A bug fix restores how something should already work, so it has no target
level: skip this step. If a level was given for one, say bug fixes don't take
one and leave it off. The feature's level still decides, as always, whether
the fix may break things.

This only reads, so do it in the main checkout before creating anything.

Find the features the task touches and each one's current level as the
`stability` skill describes: the user's word, then the declaration, then what
the repository's code and docs show. Tell the user which level you found for
each and where it came from. Then compare the target against each one:

- **Same level**: go on. More Alpha work on an Alpha feature is the usual
  case.
- **Above the current level**: this moves the feature up. Confirm that's
  intended. If it is, do the work in step 5 as `../promote/SKILL.md` (from
  this skill's directory) describes: only the delta, then a second review,
  then recording the new level.
- **Below the current level**: changes to an existing feature have to hold
  up to its level, so work on a GA feature is GA work. The exception is a task
  that creates a new feature, or a new version of a GA one, which gets its own
  level. Ask which this is. For something new, record its declaration where
  the project keeps them, and ask where if the project has no convention.
  Otherwise use the current level.
- **Shared code at a stricter level**, such as a table, endpoint, or
  component other features use: the whole change is held to the highest level
  it touches. Say so, and offer to split the change if that's easy.

Read the target level's reference in the `stability` skill, and the next
level's, before planning the work.

## 2. Read the repository

In the repository the task is for:

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

Uncommitted changes in the main checkout stay where they are. Don't carry them
into the worktree, and don't stash or commit them.

## 3. Create the worktree

- **Branch name**: `<type>/<slug>`. Match the prefixes the repository's
  existing branches use. Without a pattern, use `feat/`, `fix/`, `improve/`,
  `docs/`, or `chore/`. The slug is a few words in kebab-case naming the
  change, such as `fix/popout-phone-overflow`. If the name is taken, pick
  another one, and don't reuse the existing branch.
- **Path**: `<repo>/.claude/worktrees/<type>-<slug>`.
- Check that path is ignored: `git check-ignore -q .claude/worktrees/x`. If it
  isn't, add `.claude/worktrees/` to `.git/info/exclude`. That file is local
  and shared by all of the repository's worktrees, so nothing gets committed.
- Create it from the base chosen in step 2:
  `git worktree add --no-track -b <branch> <path> <base>`.
- Move the session into it with the `EnterWorktree` tool, passing `path`.
  Claude Code then loads the worktree's `CLAUDE.md` and settings, and blocks
  edits and commands in the main checkout. If the tool rejects the path, which
  happens when the repository isn't the session's launch directory or nested
  inside it, carry on without it: use absolute paths under the worktree for
  every edit and `git -C <path>` for every git command, and check each one
  points into the worktree.

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

Do the task in the worktree the way the project's instructions and the global
rules describe, sized to the target level: meet what it expects, and leave out
what first appears at a later level. Don't fix the feature's `known_gaps`
along the way. Ask when a decision is the user's to make. A task that
implements, changes, or clarifies a business rule adds or updates its
acceptance tests, and a bug fix or performance work doesn't (see the testing
rule). Acceptance tests need the user's explicit approval before you add,
change, or remove any: propose the scenarios before writing them.

When the work is done:

- Run the project's checks (lint, format, build, tests, as its `CLAUDE.md`
  lists them) and fix what fails.
- Except for a bug fix, invoke the `stability-reviewer` subagent on the
  branch's changes
  (`git diff $(git merge-base <base> HEAD)` from the worktree, which
  includes uncommitted work)
  against the target level, at the effort the user asked for, or `medium`.
  Resolve its REQUIRED NOW and SECURITY/PRIVACY
  findings. Remove work this change added that it flags UNNECESSARY NOW.
  Findings about code this change didn't touch go in the report, not in the
  change. For a move, `../promote/SKILL.md` already has this review.

## 6. Open the pull request

Follow `../pr/SKILL.md` from this skill's directory, with the target level and
whether this is a move, or with no level for a bug fix: commit, push the branch, and open a pull
request against the default branch. The pull request is where review happens.
To make follow-up changes, edit in the same worktree and follow
`../pr/SKILL.md` again, which pushes them to the same pull request.

Stay in the worktree when you're done. Don't remove it, merge the pull request,
or switch back to the default branch unless asked.

## Cleaning up

Only when the user asks, usually once the pull request is merged or abandoned:

1. Leave the worktree with `ExitWorktree` and `action: "keep"`. It doesn't
   remove worktrees entered with `path`.
2. Check it has nothing unpushed: `git -C <path> status --short` and
   `git -C <path> log origin/<branch>..HEAD`. If anything is there, show it
   and ask before going on.
3. `git worktree remove <path>`. If it was a `claude -p` run, unlock it first
   with `git worktree unlock <path>`.
4. Delete the local branch. If `gh pr view <branch> --json state` says
   `MERGED`, use `git branch -D <branch>`, since a squash merge makes
   `git branch -d` refuse. Otherwise ask before deleting it. Leave the remote
   branch alone unless asked.
