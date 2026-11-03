#!/usr/bin/env bash
# Open Chrome in the next free spot on the user's monitors. setup.sh
# has Playwright launch this in place of Chrome, so the spot is chosen when a
# browser really opens, not when one of a session's servers starts.
#
# Each monitor is a grid of eight spots: four across and two down, or two across
# and four down if it's portrait. Spots fill from the top right, right to left
# and then down, on secondary monitors first and the primary last. Once they're
# all taken, new windows stack on them in the same order.
#
# Taken spots are files named by their number, each holding its Chrome's PID,
# which exec keeps. They're shared by every session, and a spot is free again
# once its Chrome is gone. Where the monitors can't be read, Chrome opens as it
# otherwise would.
set -uo pipefail

chrome=$(command -v google-chrome || echo "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")

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
    # " 0: +*eDP-1 1920/344x1080/194+0+0  eDP-1"; "*" marks the primary. The
    # height leaves room for a taskbar along the bottom.
    xrandr --listmonitors | awk 'NR > 1 {
      split($3, g, /[\/x+]/)
      print ($2 ~ /^\+?\*/ ? 1 : 0), g[1], g[3] - 48, g[5], g[6]
    }'
  fi
}

# Prints "<width>,<height> <left>,<top>" for a spot's number. With no primary
# marked, as under WSLg, the first listed is the primary.
placement() {
  awk -v spot="$1" '
    { w[NR] = $2; h[NR] = $3; x[NR] = $4; y[NR] = $5; if ($1) primary = NR }
    END {
      if (!NR) exit 1
      if (!primary) primary = 1
      for (i = 1; i <= NR; i++) if (i != primary) order[++n] = i
      order[++n] = primary
      for (j = 1; j <= n; j++) {
        i = order[j]
        cols[i] = h[i] > w[i] ? 2 : 4
        rows[i] = 8 / cols[i]
        # Chrome has a minimum window size, so a small monitor gets fewer,
        # larger spots.
        while (cols[i] > 1 && w[i] / cols[i] < 500) cols[i]--
        while (rows[i] > 1 && h[i] / rows[i] < 400) rows[i]--
        total += cols[i] * rows[i]
      }
      left = spot % total
      for (j = 1; j <= n; j++) {
        i = order[j]
        if (left < cols[i] * rows[i]) break
        left -= cols[i] * rows[i]
      }
      cw = int(w[i] / cols[i]); ch = int(h[i] / rows[i])
      print cw "," ch, x[i] + w[i] - (left % cols[i] + 1) * cw "," y[i] + int(left / cols[i]) * ch
    }'
}

# The lowest free spot, claimed for this process. Creating the file fails if
# it exists, so two browsers opening at once can't take the same spot.
claim_spot() {
  local dir=${XDG_CACHE_HOME:-$HOME/.cache}/forge/playwright-windows spot=0
  mkdir -p "$dir"
  while :; do
    if [[ -e $dir/$spot ]] && kill -0 "$(cat "$dir/$spot")" 2>/dev/null; then
      spot=$((spot + 1))
      continue
    fi
    rm -f "$dir/$spot"
    if (set -C && echo $$ >"$dir/$spot") 2>/dev/null; then
      echo "$spot"
      return
    fi
    spot=$((spot + 1))
  done
}

layout=$(monitors 2>/dev/null)
if [[ -n $layout ]] && read -r size position < <(placement "$(claim_spot)" <<<"$layout"); then
  args=("--window-size=$size" "--window-position=$position")
  # Under Wayland, only X11 lets a window place itself.
  [[ $(uname) != Darwin ]] && args+=(--ozone-platform=x11)
  exec "$chrome" "$@" "${args[@]}"
fi
exec "$chrome" "$@"
