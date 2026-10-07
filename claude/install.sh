#!/usr/bin/env bash
# Expose this directory's Claude Code config through ~/.claude using symlinks.
#
# Safe to re-run. Never overwrites or deletes anything: conflicts and stale
# links are reported and left for you to review. Edits to linked files take
# effect immediately. Re-run after adding a new skill, rule, or agent if that
# directory had to be linked entry by entry.
set -euo pipefail

src="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
dest="${CLAUDE_CONFIG_DIR:-$HOME/.claude}"
added=0 ok=0 stale=0 conflicts=0

link() {
  local from=$1 to=$2
  if [[ -L $to ]]; then
    if [[ $(readlink -f "$to") == "$(readlink -f "$from")" ]]; then
      echo "ok        $to"
      ok=$((ok + 1))
    else
      echo "CONFLICT  $to -> $(readlink "$to") (expected $from)"
      conflicts=$((conflicts + 1))
    fi
  elif [[ -e $to ]]; then
    echo "CONFLICT  $to exists and is not a symlink; review and merge it manually"
    conflicts=$((conflicts + 1))
  else
    ln -s "$from" "$to"
    echo "added     $to -> $from"
    added=$((added + 1))
  fi
}

mkdir -p "$dest"
link "$src/CLAUDE.md" "$dest/CLAUDE.md"

# Link each directory whole when ~/.claude doesn't have its own copy. Otherwise
# link its entries one by one so existing content (e.g. other skills) is kept.
for dir in rules skills agents; do
  if [[ -d $dest/$dir && ! -L $dest/$dir ]]; then
    for entry in "$src/$dir"/*; do
      link "$entry" "$dest/$dir/$(basename "$entry")"
    done
    # Links to entries since removed or renamed here.
    for to in "$dest/$dir"/*; do
      if [[ -L $to && ! -e $to && $(readlink "$to") == "$src/$dir/"* ]]; then
        echo "STALE     $to -> $(readlink "$to") (no longer in forge; remove it)"
        stale=$((stale + 1))
      fi
    done
  else
    link "$src/$dir" "$dest/$dir"
  fi
done

echo
echo "$added added, $ok ok, $stale stale, $conflicts conflicts"
[[ $stale -eq 0 && $conflicts -eq 0 ]]
