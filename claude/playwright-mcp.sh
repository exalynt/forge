#!/usr/bin/env bash
# Start Playwright MCP with its browser windows arranged across the user's
# monitors, through playwright-chrome.sh, which picks each window's spot when it
# opens. setup-playwright.sh registers this as the server.
#
# Usage: playwright-mcp.sh <node> <cli.js> [server flags...]
#
# If the browser is headless or not Chrome, the server starts with the browser's
# default window.
set -uo pipefail

node=$1 cli=$2
shift 2

headless=0 browser=chrome prev=
for arg in "$@"; do
  [[ $arg == --headless ]] && headless=1
  [[ $prev == --browser ]] && browser=$arg
  [[ $arg == --browser=* ]] && browser=${arg#--browser=}
  prev=$arg
done

if ((!headless)) && [[ $browser == chrome ]]; then
  exec "$node" "$cli" --executable-path "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/playwright-chrome.sh" "$@"
fi
exec "$node" "$cli" "$@"
