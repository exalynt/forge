#!/usr/bin/env bash
# Set up Playwright MCP for Claude Code, so /qa and /ux can drive a browser.
#
# Installs @playwright/mcp globally with npm, installs Google Chrome (the
# browser it launches) if it's missing, and registers the server for every
# project at user scope, started through playwright-mcp.sh so its window opens
# out of the way. Arguments are passed on to the server, for example
# --headless where there's no display.
#
# Usage: setup-playwright.sh [server flags...]
#
# Safe to re-run: it updates the package and re-registers the server with the
# current Node and flags. Re-run it after switching Node versions or moving
# this repository.
set -euo pipefail

# The VS Code extension bundles its own CLI and doesn't put `claude` on PATH.
# shellcheck disable=SC2012 # ls -t picks the newest extension version
claude=${CLAUDE_BIN:-$(command -v claude || ls -t ~/.vscode{,-server}/extensions/anthropic.claude-code-*/resources/native-binary/claude 2>/dev/null | head -1 || true)}
if [[ -z $claude ]]; then
  echo "claude not found: install Claude Code, or set CLAUDE_BIN to its path" >&2
  exit 1
fi
if ! command -v npm >/dev/null; then
  echo "npm not found: install Node.js first" >&2
  exit 1
fi

echo "Installing @playwright/mcp globally"
npm install -g @playwright/mcp@latest

if command -v google-chrome >/dev/null || [[ -d "/Applications/Google Chrome.app" ]]; then
  echo "Google Chrome is already installed"
else
  echo "Installing Google Chrome (may ask for sudo)"
  npx -y playwright install --with-deps chrome
fi

# Absolute paths, so the server starts even when Claude Code's environment
# doesn't have Node on PATH, as with nvm outside an interactive shell.
node=$(command -v node)
cli=$(npm root -g)/@playwright/mcp/cli.js
if [[ ! -f $cli ]]; then
  echo "not found after install: $cli" >&2
  exit 1
fi

"$claude" mcp remove --scope user playwright >/dev/null 2>&1 || true
launcher="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/playwright-mcp.sh"
"$claude" mcp add --scope user playwright -- "$launcher" "$node" "$cli" --isolated ${1+"$@"}

echo
echo "Done. Start a new Claude Code session, then check it with: $claude mcp list"
