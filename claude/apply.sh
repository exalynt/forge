#!/usr/bin/env bash
# Apply a stack's rules to a project: copies stacks/<stack>/rules/*.md into
# <project>/.claude/rules/<stack>/, to be committed with the project.
#
# Usage: apply.sh <stack> [project-dir]
#
# Safe to re-run. Only writes inside .claude/rules/<stack>/, and never deletes:
# files there that the stack doesn't have are reported and left in place.
set -euo pipefail

src="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/stacks"
stack=${1:-}
project=${2:-.}

if [[ -z $stack || ! -d $src/$stack/rules ]]; then
  echo "usage: apply.sh <stack> [project-dir]" >&2
  echo "stacks: $(cd "$src" && echo */ | tr -d /)" >&2
  exit 1
fi

dest="$project/.claude/rules/$stack"
mkdir -p "$dest"

for file in "$src/$stack/rules"/*.md; do
  name=$(basename "$file")
  if [[ ! -e $dest/$name ]]; then
    cp "$file" "$dest/$name" && echo "added     $dest/$name"
  elif cmp -s "$file" "$dest/$name"; then
    echo "ok        $dest/$name"
  else
    cp "$file" "$dest/$name" && echo "updated   $dest/$name"
  fi
done

for file in "$dest"/*; do
  [[ -e $src/$stack/rules/$(basename "$file") ]] || echo "not in forge, left alone: $file"
done
