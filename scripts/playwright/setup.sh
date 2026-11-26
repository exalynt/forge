#!/usr/bin/env bash
# Set up Playwright MCP so the qa, ux, and use skills can drive a browser, with
# any agent that supports MCP servers.
#
# Installs @playwright/mcp globally with npm and Google Chrome (the browser it
# launches) if it's missing, then prints the command and arguments to register
# five servers with your agent: playwright, the one used by default, and
# playwright-1 to playwright-4 for when a skill needs more browsers. Chrome
# opens through chrome.sh, so its windows arrange themselves across the
# monitors. Arguments are passed on to the servers, for example --headless
# where there's no display. README.md shows how to register them in each agent.
#
# Usage: setup.sh [server flags...]
#
# Safe to re-run: it updates the package and prints the registration again with
# the current Node and flags. Re-run it, and re-register, after switching Node
# versions or moving this repository.
set -euo pipefail

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

# Absolute paths, so the server starts even when the agent's environment
# doesn't have Node on PATH, as with nvm outside an interactive shell.
node=$(command -v node)
cli=$(npm root -g)/@playwright/mcp/cli.js
if [[ ! -f $cli ]]; then
  echo "not found after install: $cli" >&2
  exit 1
fi

# Playwright launches chrome.sh in place of Chrome. If the browser is headless or
# not Chrome, the server opens the browser's default window instead.
headless=0 browser=chrome prev=
for arg in "$@"; do
  [[ $arg == --headless ]] && headless=1
  [[ $prev == --browser ]] && browser=$arg
  [[ $arg == --browser=* ]] && browser=${arg#--browser=}
  prev=$arg
done
window=()
if ((!headless)) && [[ $browser == chrome ]]; then
  window=(--executable-path "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/chrome.sh")
fi
args=("$cli" --isolated ${window[@]+"${window[@]}"} ${1+"$@"})

# playwright is the one browser used by default. The numbered servers are for
# when a skill calls for more: several sign-ins at once, a subagent that drives
# a browser (subagents share a session's servers), or a person in a usability
# test.
echo
echo "Register these five MCP servers with your agent, each with the same"
echo "command and arguments (README.md shows how for each agent):"
echo
echo "  names:     playwright playwright-1 playwright-2 playwright-3 playwright-4"
echo "  command:   $node"
echo "  arguments:"
for arg in "${args[@]}"; do
  printf '    %q\n' "$arg"
done
echo
echo "As one line:"
echo
printf '  %q' "$node" "${args[@]}"
echo
