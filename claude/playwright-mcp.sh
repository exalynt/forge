#!/usr/bin/env bash
# Start Playwright MCP with its browser window out of the way: in the
# bottom-right corner of a secondary monitor, or of the only one, sized to
# about a quarter of it. setup-playwright.sh registers this as the server.
#
# Usage: playwright-mcp.sh <node> <cli.js> [server flags...]
#
# The monitors are read each time the server starts, so it follows a laptop
# being docked or undocked. Where they can't be read (no display, or neither
# xrandr nor macOS), or the browser is headless or not Chrome, the server
# starts with the browser's default window.
set -uo pipefail

node=$1 cli=$2
shift 2

# One line per monitor: "<primary 0|1> <width> <height> <left> <top>".
monitors() {
  if [[ $(uname) == Darwin ]]; then
    # visibleFrame leaves out the menu bar and Dock. macOS measures up from
    # the primary screen's bottom; Chrome measures down from its top.
    osascript -l JavaScript -e '
      ObjC.import("AppKit");
      const s = $.NSScreen.screens, h = s.objectAtIndex(0).frame.size.height, out = [];
      for (let i = 0; i < s.count; i++) {
        const f = s.objectAtIndex(i).visibleFrame;
        out.push([i === 0 ? 1 : 0, f.size.width, f.size.height, f.origin.x,
          h - f.origin.y - f.size.height].map(Math.round).join(" "));
      }
      out.join("\n")'
  elif [[ -n ${DISPLAY-} ]] && command -v xrandr >/dev/null; then
    # " 0: +*eDP-1 1920/344x1080/194+0+0  eDP-1"; "*" marks the primary.
    xrandr --listmonitors | awk 'NR > 1 {
      split($3, g, /[\/x+]/)
      print ($2 ~ /^\+?\*/ ? 1 : 0), g[1], g[3], g[5], g[6]
    }'
  fi
}

placement() {
  # With no primary marked, as under WSLg, the first listed is the primary.
  monitors 2>/dev/null | awk '
    { p[NR] = $1; line[NR] = $2 " " $3 " " $4 " " $5; if ($1) primary = NR }
    END {
      if (!NR) exit 1
      if (!primary) { primary = 1; p[1] = 1 }
      pick = primary
      for (i = 1; i <= NR; i++) if (!p[i]) { pick = i; break }
      split(line[pick], m, " ")
      w = int(m[1] / 2); if (w > 1280) w = 1280
      h = int(m[2] * 2 / 3); if (h > 900) h = 900
      # Clear of the right edge and of a taskbar or dock along the bottom.
      print w "," h " " (m[3] + m[1] - w - 16) "," (m[4] + m[2] - h - 80)
    }'
}

headless=0 browser=chrome prev=
for arg in "$@"; do
  [[ $arg == --headless ]] && headless=1
  [[ $prev == --browser ]] && browser=$arg
  [[ $arg == --browser=* ]] && browser=${arg#--browser=}
  prev=$arg
done

if ((!headless)) && [[ $browser =~ ^(chrome|msedge|chromium) ]] &&
  read -r size position < <(placement); then
  args="\"--window-size=$size\", \"--window-position=$position\""
  # Under Wayland, only X11 lets a window place itself.
  [[ $(uname) != Darwin ]] && args+=', "--ozone-platform=x11"'
  dir=${XDG_CACHE_HOME:-$HOME/.cache}/forge
  mkdir -p "$dir"
  tmp=$(mktemp "$dir/playwright-mcp.XXXXXX")
  printf '{ "browser": { "launchOptions": { "args": [%s] } } }\n' "$args" >"$tmp"
  mv "$tmp" "$dir/playwright-mcp.json"
  exec "$node" "$cli" --config "$dir/playwright-mcp.json" "$@"
fi
exec "$node" "$cli" "$@"
