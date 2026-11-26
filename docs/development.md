# Developing forge

## Layout

```text
forge/
├── kit/                      # what forge installs, each skill a folder
│   ├── AGENTS.md             # the block written into a repository's instructions
│   ├── workflows/            # in every install
│   │   ├── implement/        # build a feature, improvement, or fix; a worktree and PR when asked
│   │   ├── commit/           # small, related commits
│   │   ├── pr/               # commit, push, and open or update a PR labeled with its level
│   │   ├── qa/               # QA a capability end to end against its rules
│   │   ├── ux/               # walk a main flow as a type of user
│   │   ├── use/              # usability test, one player per person (role.md is their brief)
│   │   ├── stability-review/ # review work against a level; needs Docqet
│   │   └── promote/          # move a capability up a level; needs Docqet
│   ├── universal/            # in every install: conventions for any language
│   │   ├── standards/        # layers, errors, config, APIs, pagination, health, CI, UI
│   │   ├── testing/          # how we write tests
│   │   └── observability/    # how we add traces, logs, and metrics
│   └── languages/            # chosen per repository
│       ├── go/               # skill go; stacks/http/ is skill go-http
│       └── typescript/       # skill typescript; stacks/react/ is skill typescript-react
├── assets/logo/              # the forge logo and how to use it; not part of the binary
├── docs/                     # user docs the README links to
├── forge.go                  # embeds kit/ in the binary
├── cmd/forge/                # runs the forge command
├── internal/commands/        # each command in <name>.go, its subcommands in <name>/
├── internal/install/         # where each agent's files go, and how they're written
├── internal/playwright/      # Playwright MCP for qa, ux, and use, for forge browser
└── scripts/release/          # builds the release archives
```

Forge's block, `kit/AGENTS.md`, under "Where intention lives", defines how
skills behave with and without Docqet, and every skill follows it.

## CI

CI checks every push and pull request: formatting, vet, the scripts, and an
install into a new repository that's checked, updated and removed.

## Releasing

Every commit on `main` that passes CI is built and published as the `latest`
pre-release by the **Latest** workflow, so there's nothing to do to get a
build of `main`.

A versioned release is a version tag, such as `v0.2.0`, and the builds for
it. In the repository's Actions tab, run the **Release** workflow with the
version: it tags `main`'s latest commit, builds forge for each platform, and
publishes them as a GitHub release with notes from the commits since the last
one. Pushing a tag yourself does the same:

```bash
git tag v0.2.0 && git push origin v0.2.0
```

The tag is the version `forge version` prints, what `go install …@latest`
installs, and what `forge version --check` compares with. Releases follow
semantic versioning: before 1.0, a breaking change to how forge installs, such
as a lock it can't read, bumps the minor version.
