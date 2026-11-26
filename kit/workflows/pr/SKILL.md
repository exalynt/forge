---
name: pr
description: Commit the current branch's changes, push the branch, and open a pull request against the default branch, labeled and described with the stability level it should be reviewed under (bug fixes have none), or update the open one with new commits.
disable-model-invocation: true
---

# Open a Pull Request

The user's request is the input. It may name a level (`prototype`, `alpha`,
`beta`, or `ga`) and ask for a draft, for example "open a PR for alpha" or "a
draft PR for beta". The level is the stability level the change aims for, and
so the one review holds it to. A bug fix takes no level. Asked for a draft, open
the pull request as a draft. Otherwise it opens ready for review.

Without a level, and unless the branch is a bug fix, use the one the
`implement` skill settled earlier in this session, or the open pull request's
`stability:` label. Failing those:

- **With Docqet** (see forge's `AGENTS.md`, "Where intention lives"), run
  `docqet touches --base origin/<default>` and use the strictest level among
  the capabilities the branch touches. Tell the user which level that is and
  which capability it comes from.
- **Without Docqet**, ask the user for the level. If they don't want one, open
  the pull request without a stability label.

Invoking this skill is the request to push the current branch and open, or
update, its pull request. Do this in the main session, without subagents.

## 1. Check where you are

- Run `git branch --show-current`, `git status`, and
  `git symbolic-ref --short refs/remotes/origin/HEAD` for the default branch.
- If the current branch is the default branch, or HEAD is detached, stop.
  Changes for a pull request belong on their own branch, for example from the
  `implement` skill in a worktree.
- Run `gh auth status`. If it isn't authenticated, stop and tell the user
  (`gh auth login`, or replace an invalid `GITHUB_TOKEN` in the environment).
  Don't try to log in yourself.
- See whether the branch already has a pull request:
  `gh pr view --json url,state,isDraft,labels`. If it's merged or closed,
  stop and ask. New work needs a new branch.
- With Docqet, if the `implement` skill didn't already settle the level,
  compare the level against the current level of the capabilities the
  branch touches, as `implement` step 1 (`../implement/SKILL.md`) describes.
  If it's above their current level without recording the move, or below it,
  ask before going on.

## 2. Commit

Group and commit the uncommitted changes as `../commit/SKILL.md` (from this
skill's directory) describes in its steps 1–6:

- In a worktree created for this task, every uncommitted change in it is in
  scope.
- Anywhere else, only the changes this session made are in scope.

If there's nothing to commit and nothing unpushed, say so and stop.

## 3. Check

Run the project's checks (lint, format, build, tests, as its `AGENTS.md` or
other instructions list them) on what's committed. With Docqet, also run
`docqet check --changed`. If they fail, fix the cause and commit the fix. If
you can't, stop and report. Don't push failing work unless the user says to,
and then say so in the pull request.

Run `git fetch origin <default>` and
`git rev-list --count HEAD..origin/<default>`. If the default branch has moved
on, mention it in the report. Don't rebase or merge it in unless asked.

## 4. Push

`git push -u origin <branch>`.

- Never push to the default branch.
- Never force-push unless asked. If the push is rejected because the remote
  branch has commits this one doesn't, stop and ask.
- Never bypass hooks with `--no-verify`.

## 5. Open or update the pull request

### Stability label

A pull request with a level gets exactly one label, `stability: <level>`. A
bug fix restores how something should already work, so it has no stability
level and gets no label, and neither does one the user didn't want labeled.

Check the repository has the labels with `gh label list --search "stability:"`.
If they're missing, open the pull request without the label, and tell the user.
Offer to create them with `gh label create <name> --color <color>
--description "<description>"`, using these, so every repository has the same
labels:

| Name                    | Color    | Description                                                                         |
| ----------------------- | -------- | ----------------------------------------------------------------------------------- |
| `stability: prototype`  | `6d6a7c` | Prototype: exploring an idea; expect it to change completely or be thrown away      |
| `stability: alpha`      | `9c4f6c` | Alpha: works but still taking shape; expect frequent breaking changes and bugs      |
| `stability: beta`       | `3f6f88` | Beta: refining; occasional breaking changes, announced ahead of time when we can    |
| `stability: ga`         | `4c7a58` | GA: generally available and stable; rare breaking changes, announced with a deprecation window |

The label is the level the change aims for, not where the code is today. A
change spanning levels takes the highest it touches. Refactoring, tooling, and
dependency updates take the level of the code they change, and changes to CI,
deployment, or shared tooling people already depend on are GA.

### Description

Write the body to a temporary file, in a temporary directory. Follow
`.github/pull_request_template.md` if the repository has one, and add the
stability section to it. Otherwise use the template below. For a bug fix,
replace the stability section with a first line saying it's a bug fix and
what it restores. Without a level, leave the stability section out.

```markdown
## Stability

Review as **<Level>**: <what the level expects, in a sentence>.

- **Capabilities**: <each one the change touches>, currently <level>, per
  <Docqet | the user>. <"Moves <capability> to <level>." for a move.>
- **Not expected at this level**: <work a reviewer might ask for that belongs
  to a later level, and which level>.

## Summary

What changes and why, in a few sentences or bullets.

## Verification

Each check run and its result, the stability review's outcome if there was
one, and anything not verified.

## Deploying

What's true only while this change rolls out: deploy order, one-time commands,
and what rolling back does.

## Known gaps

What was deliberately left out, if anything.
```

With Docqet, say what the level expects from `docqet rubric --level <level>`,
mention the capability's gaps the change leaves in place under "Not expected
at this level", and add a short Docqet section from `docqet diff`, with any
`Docqet-Became` trailers the `docqet` skill asks for when the work will be
squash-merged. Leave out other sections with nothing to say. Keep any pull
request attribution your configuration asks for.

### Create or update

**If one is already open**, the push updated it. Make sure it has the one
right stability label, swapping a stale one with
`gh pr edit --add-label "stability: <level>" --remove-label "stability: <old>"`.
If the new commits or a new level change what it says, update the description
with `gh pr edit --body-file <file>`.

**Otherwise**, create it:

```
gh pr create --base <default> --head <branch> --title "<title>" --body-file <file> --label "stability: <level>" [--draft]
```

Leave out `--label` for a bug fix or a pull request without a level. The
title says what the pull request does, in the style of the repository's
commit subjects. For a single commit, use its subject.

Don't add reviewers, other labels, or assignees, and don't merge, unless
asked.

## 6. Report

Give the pull request URL, its stability label (or why it has none), the
commits pushed (`git log --oneline origin/<default>..HEAD`), the checks with
their results, and whether the default branch has moved on. Note anything left
uncommitted and why.
