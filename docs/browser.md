# Browser tools for qa, ux, and use

`qa` and `ux` drive a capability through its UI with whichever browser tool
the agent has. Without one, they report the UI as not verified. `use` needs a
browser for each person in a test.
[Playwright MCP](https://github.com/microsoft/playwright-mcp) gives them that.

## Setting it up

Once per machine:

```bash
forge browser setup                  # headed Chrome
forge browser setup -- --headless    # where there's no display
```

It installs `@playwright/mcp` globally with npm, installs Google Chrome if it's
missing (the server's default browser; this step may ask for `sudo`), and
records where Node and the package are in `setup.json`, in
`~/.local/share/forge/playwright/` (`$XDG_DATA_HOME/forge/playwright/` where
that's set).

Then it prints the command to register, `forge browser server` by forge's
absolute path, under five names: `playwright`, the one used by default, and
`playwright-1` through `playwright-4` for when a skill calls for more, such as
several sign-ins at once or the people in a `use` test. Each is its own
browser and sign-in.

## Registering the servers

Register all five with your agent, with the same command, once:

- **Claude Code**: setup prints a loop that runs
  `claude mcp add --scope user <name> -- /path/to/forge browser server` for
  each name. Replacing one needs `claude mcp remove --scope user <name>` first.
- **Codex**: a table for each name in `~/.codex/config.toml`:

  ```toml
  [mcp_servers.playwright]
  command = "/path/to/forge"
  args = ["browser", "server"]
  ```

- **Cursor**: an entry for each name under `mcpServers` in `.cursor/mcp.json`
  for one project, or `~/.cursor/mcp.json` for all of them, with the same
  `command` and `args`.

Let the browser tools run without asking each time, but keep asking before
`browser_run_code_unsafe`, which runs arbitrary code on your machine, and
before `browser_file_upload` and `browser_drop`, which can send local files to
a page. Set that in your agent's permissions.

## Changing the setup

`forge browser server` reads `setup.json` and runs the package with Node in
its place. So after switching Node versions, or to change the server's flags,
run `forge browser setup` again: the servers use the new setup the next time
they start, without registering them again. Re-registering is needed only if
forge's binary moves. Re-running setup also updates the package.

`forge browser check` exits 1 when something setup recorded is gone, such as a
Node version that was uninstalled, and says what to run.

Flags after `--` are passed to the server. The browser is headed by default,
which works under WSLg, so add `--headless` where there's no display. Pass
`--browser firefox` to use Firefox instead of Chrome. Playwright's bundled
Chromium isn't one of the server's options.

## How it behaves

- **Sessions are isolated.** The server always gets `--isolated`, so each
  session starts with a fresh, in-memory profile. A sign-in from an earlier
  run doesn't carry into first-use, signed-out, or other-tenant checks. The
  cost is signing in again each session.
- **Each agent session gets its own browsers.** Agents that start MCP servers
  per session, as Claude Code does, give every session its own five browsers,
  so separate sessions never share one. Subagents share their session's, so
  each subagent driving a browser needs a numbered server of its own. Browsers
  don't separate data: every agent and session uses the same app and
  database. Killing Chrome by name closes every session's browser, so the
  skills close only their own.
- **Output stays in the project.** The server writes screenshots, snapshots,
  and console logs to `.playwright-mcp/` in the project, and refuses paths
  outside the project or that folder. `qa`, `ux`, and `use` keep everything
  they create there and add `/.playwright-mcp/` to the project's `.gitignore`
  if it's missing.
- **It doesn't depend on your shell's PATH.** Setup records Node and the
  package by absolute path, so the server starts even when the agent's
  environment lacks nvm's PATH.

### Window placement

Browser windows arrange themselves across your monitors, so you can follow
along without them covering your work. Playwright opens Chrome through
`chrome.sh`, which setup writes beside `setup.json` and each server rewrites as
it starts, so it's always the one forge carries.

It reads the monitors each time a browser opens, with `xrandr` on Linux and
WSLg or AppKit on macOS, and gives the window an eighth of a monitor, or more
where that would be smaller than Chrome's minimum window size. Windows fill
from the top right, right to left and then down, on secondary monitors first
and the primary last. Once every spot is taken, new ones stack on the same
spots in order. A spot frees up when its browser closes. Spots are shared
across sessions, but only count browsers it opened, not your own windows.

Where it can't read the monitors, the browser opens as it otherwise would, and
with `--headless` or another browser it isn't used at all. On Windows,
`chrome.sh` needs bash, so Chrome opens as the server would open it anyway.
The skills set each test width with the browser's resize tool, which doesn't
move or resize the window.

### Focus

Under WSLg, a browser opening doesn't keep the focus. `keep-focus.ps1` runs on
the Windows side for the browser's first 10 seconds, and sends focus straight
back to the window you were in. Clicking into a browser window during those
seconds sends you back too. On macOS and native Linux, a new browser can still
take the focus.
