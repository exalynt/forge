#!/usr/bin/env bash
# Install forge into a repository, to be committed with it, so everyone working
# on it gets Exalynt's way of working, whatever agent they use:
#
#   - forge's AGENTS.md, written into <repo>/AGENTS.md between
#     <!-- forge:start --> and <!-- forge:end --> markers
#   - each skill in skills/, and each stack asked for, copied into
#     <repo>/.agents/skills/<name>/
#   - <repo>/.agents/forge.lock, listing what was installed
#   - an @AGENTS.md line in <repo>/CLAUDE.md, if that file exists without one
#
# Usage: install.sh <repo> [--stack <name>]... [--no-stacks] [--check]
#
# Without --stack or --no-stacks, it keeps the stacks the lock lists. With any
# --stack, it installs exactly those, so passing a different set drops the
# rest. --no-stacks removes every stack.
#
# Safe to re-run, and re-running updates the copies. It changes only what's
# forge's: its block in AGENTS.md, and the skill folders it installed, which it
# makes match forge exactly. Skills it installed before and no longer does are
# removed. Any other skill, such as docket, and the rest of AGENTS.md are left
# alone. Copies an earlier forge put in <repo>/.claude/skills/ are removed:
# Claude Code finds the skills through the AGENTS.md block until it reads
# .agents/skills/. With --check it changes nothing, and exits 1 if anything
# would change.
set -euo pipefail

src="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
start_marker='<!-- forge:start -->'
end_marker='<!-- forge:end -->'

usage() {
  echo "usage: install.sh <repo> [--stack <name>]... [--no-stacks] [--check]" >&2
  echo "stacks: $(cd "$src/stacks" && echo */ | tr -d /)" >&2
  exit 1
}

repo='' check=0 no_stacks=0 stacks=()
while (($#)); do
  case $1 in
    --check) check=1 ;;
    --no-stacks) no_stacks=1 ;;
    --stack) (($# > 1)) || usage; stacks+=("$2"); shift ;;
    --stack=*) stacks+=("${1#--stack=}") ;;
    -h | --help) usage ;;
    -*) echo "unknown option: $1" >&2; usage ;;
    *) [[ -z $repo ]] || usage; repo=$1 ;;
  esac
  shift
done
[[ -n $repo ]] || usage
if ((no_stacks)) && ((${#stacks[@]})); then
  echo "--no-stacks and --stack can't be used together" >&2
  usage
fi
if [[ ! -d $repo ]]; then
  echo "not a directory: $repo" >&2
  exit 1
fi
repo=$(cd "$repo" && pwd)
lock=$repo/.agents/forge.lock

is_name() { [[ $1 =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]]; }

# What the last install recorded: its skills, and its stacks.
old_skills=() old_stacks=() legacy_lock=0
if [[ -f $lock ]]; then
  # A lock without "# skills .agents" is from a forge that also copied each
  # skill into .claude/skills/.
  grep -qxF '# skills .agents' "$lock" || legacy_lock=1
  while IFS= read -r line; do
    case $line in
      '# stacks'*) read -ra old_stacks <<<"${line#\# stacks}" ;;
      '#'* | '') ;;
      *) is_name "$line" && old_skills+=("$line") ;;
    esac
  done <"$lock"
fi
if ((!no_stacks)) && ((${#stacks[@]} == 0)) && ((${#old_stacks[@]})); then
  stacks=("${old_stacks[@]}")
fi

# The skills to install, by name, each with the folder it comes from.
names=() dirs=()
for dir in "$src"/skills/*/; do
  names+=("$(basename "$dir")") dirs+=("${dir%/}")
done
for stack in ${stacks[@]+"${stacks[@]}"}; do
  if ! is_name "$stack" || [[ ! -f $src/stacks/$stack/SKILL.md ]]; then
    echo "unknown stack: $stack" >&2
    usage
  fi
  names+=("$stack") dirs+=("$src/stacks/$stack")
done

contains() {
  local want=$1 item
  shift
  for item; do [[ $item == "$want" ]] && return 0; done
  return 1
}

files_in() { (cd "$1" && find . -type f | sed 's|^\./||' | LC_ALL=C sort); }

# --- Check everything that would stop the install before changing anything.

agents=$repo/AGENTS.md
starts=0 ends=0
if [[ -f $agents ]]; then
  starts=$(grep -cxF "$start_marker" "$agents" || true)
  ends=$(grep -cxF "$end_marker" "$agents" || true)
  if ((starts + ends > 0)); then
    first_start=$(grep -nxF "$start_marker" "$agents" | head -1 | cut -d: -f1)
    first_end=$(grep -nxF "$end_marker" "$agents" | head -1 | cut -d: -f1)
    if ((starts != 1 || ends != 1)) || ((first_start > first_end)); then
      echo "AGENTS.md: forge's markers are unbalanced ($starts start, $ends end); fix them by hand" >&2
      exit 1
    fi
  fi
fi

# A skill folder forge didn't install is someone else's, so forge won't
# overwrite it, unless it's already exactly forge's copy.
for i in "${!names[@]}"; do
  name=${names[$i]}
  contains "$name" ${old_skills[@]+"${old_skills[@]}"} && continue
  dest=$repo/.agents/skills/$name
  [[ -e $dest || -L $dest ]] || continue
  if [[ -L $dest ]] || ! diff -rq "${dirs[$i]}" "$dest" >/dev/null 2>&1; then
    echo ".agents/skills/$name exists and wasn't installed by forge; move or remove it first" >&2
    exit 1
  fi
done

# --- Install, or with --check, only report.

added=0 updated=0 removed=0 ok=0

report() {
  printf '%-9s %s\n' "$1" "$2"
  case $1 in
    added) added=$((added + 1)) ;;
    updated) updated=$((updated + 1)) ;;
    removed) removed=$((removed + 1)) ;;
    ok) ok=$((ok + 1)) ;;
  esac
}

# Writes the content of file $2 to $1 (relative to the repo) if it differs.
put() {
  local rel=$1 new=$2 path=$repo/$1
  if [[ ! -e $path ]]; then
    ((check)) || { mkdir -p "$(dirname "$path")" && cp "$new" "$path"; }
    report added "$rel"
  elif cmp -s "$new" "$path"; then
    report ok "$rel"
  else
    ((check)) || cp "$new" "$path"
    report updated "$rel"
  fi
}

remove() {
  ((check)) || rm -rf "${repo:?}/$1"
  report removed "$1"
}

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

# AGENTS.md: forge's block, replaced in place or appended.
{
  echo "$start_marker"
  cat "$src/AGENTS.md"
  echo "$end_marker"
} >"$tmp/block"
if [[ ! -f $agents ]]; then
  cp "$tmp/block" "$tmp/AGENTS.md"
elif ((starts == 1)); then
  awk -v s="$start_marker" -v e="$end_marker" -v block="$tmp/block" '
    $0 == s { while ((getline line < block) > 0) print line; skip = 1; next }
    $0 == e && skip { skip = 0; next }
    !skip' "$agents" >"$tmp/AGENTS.md"
else
  cp "$agents" "$tmp/AGENTS.md"
  if [[ -s $agents ]]; then
    [[ $(tail -c1 "$agents") == "" ]] || echo >>"$tmp/AGENTS.md"
    echo >>"$tmp/AGENTS.md"
  fi
  cat "$tmp/block" >>"$tmp/AGENTS.md"
fi
put AGENTS.md "$tmp/AGENTS.md"

# Skills: each one's folder made to match forge's.
for i in "${!names[@]}"; do
  name=${names[$i]} dir=${dirs[$i]} root=.agents/skills
  if [[ -L $repo/$root/$name ]]; then
    remove "$root/$name"
  fi
  while IFS= read -r file; do
    put "$root/$name/$file" "$dir/$file"
  done < <(files_in "$dir")
  if [[ -d $repo/$root/$name ]]; then
    while IFS= read -r file; do
      [[ -e $dir/$file ]] || remove "$root/$name/$file"
    done < <(files_in "$repo/$root/$name")
    ((check)) || find "$repo/$root/$name" -mindepth 1 -type d -empty -delete
  fi
done

# Skills installed before and no longer forge's or asked for.
for name in ${old_skills[@]+"${old_skills[@]}"}; do
  [[ $name == docket ]] && continue
  contains "$name" "${names[@]}" && continue
  if [[ -e $repo/.agents/skills/$name || -L $repo/.agents/skills/$name ]]; then
    remove ".agents/skills/$name"
  fi
done

# Copies an earlier forge put in .claude/skills/. A name an earlier forge's
# lock lists is forge's. Otherwise one is removed only when it's a link or an
# exact copy, so a project's own skill of the same name stays.
claude_names=(${old_skills[@]+"${old_skills[@]}"})
for name in "${names[@]}"; do
  contains "$name" ${claude_names[@]+"${claude_names[@]}"} || claude_names+=("$name")
done
for name in "${claude_names[@]}"; do
  [[ $name == docket ]] && continue
  dest=$repo/.claude/skills/$name
  [[ -e $dest || -L $dest ]] || continue
  if ! ((legacy_lock)) || ! contains "$name" ${old_skills[@]+"${old_skills[@]}"}; then
    if [[ ! -L $dest ]]; then
      dir=''
      for i in "${!names[@]}"; do
        [[ ${names[$i]} == "$name" ]] && dir=${dirs[$i]}
      done
      if [[ -z $dir ]] || ! diff -rq "$dir" "$dest" >/dev/null 2>&1; then
        continue
      fi
    fi
  fi
  remove ".claude/skills/$name"
  claude_removed=1
done
# Folders those removals left empty, and nothing else under .claude/.
if ((!check)) && ((${claude_removed:-0})); then
  if rmdir "$repo/.claude/skills" 2>/dev/null; then
    rmdir "$repo/.claude" 2>/dev/null || true
  fi
fi

# CLAUDE.md: Claude Code reads AGENTS.md only through an import when there's a
# CLAUDE.md. Never created here.
if [[ -f $repo/CLAUDE.md ]]; then
  if grep -qxF '@AGENTS.md' "$repo/CLAUDE.md"; then
    report ok CLAUDE.md
  else
    if ((!check)); then
      if [[ -s $repo/CLAUDE.md && $(tail -c1 "$repo/CLAUDE.md") != "" ]]; then
        echo >>"$repo/CLAUDE.md"
      fi
      echo '@AGENTS.md' >>"$repo/CLAUDE.md"
    fi
    report updated CLAUDE.md
  fi
fi

# The lock. Its forge version is the one that last changed what it lists, so a
# rerun that changes nothing leaves it as it is.
{
  for name in "${names[@]}"; do echo "$name"; done
} >"$tmp/lock.body"
{
  echo "# skills .agents"
  echo "# stacks${stacks[*]+ ${stacks[*]}}"
  cat "$tmp/lock.body"
} >"$tmp/lock.compare"
if [[ -f $lock ]] && cmp -s "$tmp/lock.compare" <(grep -v '^# forge' "$lock"); then
  report ok .agents/forge.lock
else
  version=$(git -C "$src" rev-parse --short HEAD 2>/dev/null || echo unknown)
  {
    echo "# forge $version"
    cat "$tmp/lock.compare"
  } >"$tmp/forge.lock"
  put .agents/forge.lock "$tmp/forge.lock"
fi

echo
changes=$((added + updated + removed))
if ((check)); then
  echo "check: $added to add, $updated to update, $removed to remove, $ok ok"
  ((changes == 0)) || { echo "run install.sh to bring this repository up to date" >&2; exit 1; }
else
  echo "$added added, $updated updated, $removed removed, $ok ok"
fi
