# Agents and what forge writes

Each agent reads its own place, so forge installs to **targets**, one per
place, and `--agent` names the agents you use. A first install with no
`--agent` installs the `agents` target. On an existing install, `install
--agent` adds targets. `forge agents` lists every name.

| Target   | `--agent` names                                              | Skills                   | Block                                                                      |
| -------- | ------------------------------------------------------------ | ------------------------ | -------------------------------------------------------------------------- |
| `agents` | `codex`, `cursor`, `gemini`, `windsurf`, `copilot`, `agents` | `.agents/skills/<name>/` | `AGENTS.md`                                                                |
| `claude` | `claude`                                                     | `.claude/skills/<name>/` | `CLAUDE.md`, or `AGENTS.md` through an `@AGENTS.md` line with `agents` too |

- With the `agents` target, forge's block goes in `AGENTS.md`. With `claude`
  as well, `CLAUDE.md` gets an `@AGENTS.md` line, and is created if it's
  missing. Claude Code reads `AGENTS.md` itself only in a repository with no
  `CLAUDE.md`, so a repository's own `CLAUDE.md` gets the line with `agents`
  alone too.
- With `claude` alone, the block goes in `CLAUDE.md`, and `AGENTS.md` isn't
  touched.
- The block says where the skills are: `.agents/skills` when the `agents`
  target is installed, otherwise `.claude/skills`.

## What forge writes

All of it is committed with the repository:

| Written                                                           | What for                                                                            |
| ----------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `AGENTS.md` or `CLAUDE.md`                                        | Forge's block, between `<!-- forge:start -->` and `<!-- forge:end -->` markers      |
| `.agents/skills/<name>/`                                          | A copy of each skill, for the `agents` target                                       |
| `.claude/skills/<name>/`                                          | A copy of each skill, for the `claude` target                                       |
| `CLAUDE.md`                                                       | An `@AGENTS.md` line, with both targets, or with `agents` when it exists            |
| `.agents/forge.lock`, or `.claude/forge.lock` with `claude` alone | Forge's version, the targets, the languages and stacks, and each skill it installed |

An existing `AGENTS.md` or `CLAUDE.md` keeps everything it had: forge's block
is added at the end, or replaced where it already is. If the markers are
unbalanced, forge stops and changes nothing.

Forge changes only what's its own: its block, the `@AGENTS.md` line when it
added it, and the skill folders it installed, which it makes match the binary
exactly. It never touches another skill, such as `docqet`, or another block,
such as Docqet's. If a skill folder with one of forge's names is already in a
target and forge didn't install it, forge stops rather than overwrite it.
Copies an earlier forge left in `.claude/skills/`, from before the `claude`
target, are removed, while a project's own skill of the same name stays.

## Supported agents

Last checked against each agent's docs on 2026-11-25. These change often.
`forge agents` prints the same list from the code.

| Agent                                      | `--agent`    | Reads                                                     | Target   | What you do                                                                          | Can't do yet                                                                                         |
| ------------------------------------------ | ------------ | --------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------- |
| Claude Code                                | `claude`     | `CLAUDE.md`, `.claude/skills/`                            | `claude` | Nothing                                                                              | Reading `.agents/skills/`, so without the `claude` target it reads skills only through the block    |
| Codex                                      | `codex`      | `AGENTS.md`, `.agents/skills/`                            | `agents` | Nothing                                                                              |                                                                                                      |
| Cursor                                     | `cursor`     | `AGENTS.md`, `.agents/skills/`                            | `agents` | Nothing                                                                              |                                                                                                      |
| Gemini CLI                                 | `gemini`     | `GEMINI.md` by default, `.agents/skills/`                 | `agents` | Add `"context": {"fileName": ["AGENTS.md", "GEMINI.md"]}` to `.gemini/settings.json` | Reading `AGENTS.md` without that setting                                                             |
| GitHub Copilot (CLI, cloud agent, VS Code) | `copilot`    | `AGENTS.md`; `.agents/skills/` in the CLI and cloud agent | `agents` | Nothing                                                                              | In VS Code, reading skills anywhere but `.github/skills/`, so there it reads them through the block |
| Windsurf                                   | `windsurf`   | `AGENTS.md`, `.agents/skills/`                            | `agents` | Nothing                                                                              |                                                                                                      |
| Any agent that reads `AGENTS.md`           | `agents`     | `AGENTS.md`, and `.agents/skills/` if it loads skills     | `agents` | Nothing                                                                              | Loading skills, for one that doesn't: it reads them through the block                               |
| Aider                                      | use `agents` | `AGENTS.md` only when it's passed in                      | `agents` | Run it with `--read AGENTS.md`, or add that to `.aider.conf.yml`                     | Loading skills: it reads them through the block                                                      |
| Zed                                        | use `agents` | The first rules file it finds, `AGENTS.md` among them     | `agents` | Make sure no other rules file, such as a stray `.rules`, shadows `AGENTS.md`         | Loading skills: it reads them through the block                                                      |

## Known gaps

- **Agents without their own target read skills through the block.** Claude
  Code doesn't read `.agents/skills/`, so where the `claude` target isn't
  installed, it reads a skill only because forge's block says where it is, and
  `/implement`, `/qa` and the rest aren't in its slash menu: ask for one by
  name, such as "use implement". Install `--agent claude` to give it its own
  copies. The same holds for Aider, Zed, and others that don't load skills.
- **Gemini CLI reads `AGENTS.md` only with its `context.fileName` setting.**
  Forge doesn't write that setting: `forge status` and `forge agents` say what
  to add.
- **Copilot in VS Code reads skills only from `.github/skills/`.** There's no
  target for it yet, so in VS Code it reads skills through the block.
