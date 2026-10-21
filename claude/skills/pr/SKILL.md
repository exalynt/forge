---
name: pr
description: Commit the current branch's changes, push the branch, and open a pull request against the default branch, labeled and described with the stability level it should be reviewed under (bug fixes have none), or update the open one with new commits.
argument-hint: "[prototype|alpha|beta|ga] [draft]"
disable-model-invocation: true
---

# Open a Pull Request

Arguments: `$ARGUMENTS`. Expected format: `<level> [draft]`, for example
`/pr alpha` or `/pr beta draft`. The level is the stability level the change
aims for, and so the one review holds it to. A bug fix takes no level, as in
`/pr` or `/pr draft`. Pass `draft` to open the pull request as a draft.
Otherwise it opens ready for review.

Without a level, and unless the branch is a bug fix, use the one `/implement`
settled earlier in this conversation, or the open pull request's `stability:`
label. Failing those, use the feature's current level, found as the
`stability` skill describes, and tell the user which level that is.

Invoking this skill is the request to push the current branch and open, or
update, its pull request. Do this in the main conversation, without subagents.

## 1. Check where you are

- Run `git branch --show-current`, `git status`, and
  `git symbolic-ref --short refs/remotes/origin/HEAD` for the default branch.
- If the current branch is the default branch, or HEAD is detached, stop.
  Changes for a pull request belong on their own branch, for example from
  `/implement`.
- Run `gh auth status`. If it isn't authenticated, stop and tell the user
  (`gh auth login`, or replace an invalid `GITHUB_TOKEN` in the environment).
  Don't try to log in yourself.
- See whether the branch already has a pull request:
  `gh pr view --json url,state,isDraft,labels`. If it's merged or closed,
  stop and ask. New work needs a new branch.
- If `/implement` didn't already settle the level, find the current level of
  the features the changes touch and compare the level against it, as
  `/implement` step 1 (`../implement/SKILL.md`) describes. If it's above their
  current level without recording the move, or below it, ask before going on.

## 2. Commit

Group and commit the uncommitted changes as `../commit/SKILL.md` (from this
skill's directory) describes in its steps 1–6:

- In a worktree created for this task, every uncommitted change in it is in
  scope.
- Anywhere else, only the changes this session made are in scope.

If there's nothing to commit and nothing unpushed, say so and stop.

## 3. Check

Run the project's checks (lint, format, build, tests, as its `CLAUDE.md`
lists them) on what's committed. If they fail, fix the cause and commit the
fix. If you can't, stop and report. Don't push failing work unless the user
says to, and then say so in the pull request.

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

Every pull request gets exactly one label, `stability: <level>`, except a
bug fix. A fix restores how something should already work, so it has no
stability level and gets no label.

Check the repository has the labels with `gh label list --search "stability:"`.
If they're missing, open the pull request without the label, and tell the user.
[Marking stability levels](https://readme.exalynt.com/engineers/marking-stability)
has the commands that set them up. Offer to run them.

### Description

Write the body to a temporary file, in your scratchpad directory if you have
one. Follow `.github/pull_request_template.md` if the repository has one, and
add the stability section to it. Otherwise use the template below. For a bug
fix, replace the stability section with a first line saying it's a bug fix
and what it restores.

```markdown
## Stability

Review as **<Level>**: <what the level expects, in a sentence>. See the
[stability checklist](https://readme.exalynt.com/how-it-works/stability-checklist)
for what each level includes.

- **Feature**: <name>, currently <level>, per `<declaration path>` or where
  else the level came from. <"Moves <feature> to <level>." for a move.>
- **Not expected at this level**: <work a reviewer might ask for that belongs
  to a later level, and which level>.

## Summary

What changes and why, in a few sentences or bullets.

## Verification

Each check run and its result, the stability review's outcome if there was
one, and anything not verified.

## Known gaps

What was deliberately left out, if anything.
```

Mention the feature's `known_gaps` the change leaves in place under "Not
expected at this level". Leave out other sections with nothing to say. Keep any pull request attribution your configuration asks for.

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

Leave out `--label` for a bug fix. The title says what the pull request does,
in the style of the repository's commit subjects. For a single commit, use its
subject.

Don't add reviewers, other labels, or assignees, and don't merge, unless
asked.

## 6. Report

Give the pull request URL, its stability label (or why it has none), the
commits pushed (`git log --oneline origin/<default>..HEAD`), the checks with
their results, and whether the default branch has moved on. Note anything left
uncommitted and why.
