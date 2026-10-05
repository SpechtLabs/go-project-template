#!/usr/bin/env bash
#
# Turns a fresh copy of go-project-template into the project NAME:
#
#   ./scripts/init.sh NAME
#
# It replaces the PROJECT_NAME placeholder in every tracked file, renames
# cmd/PROJECT_NAME to cmd/NAME, drops the README's template section, and
# deletes this script. Run it once, in a clean checkout, and review the
# result with `git status` and `git diff` before committing.

set -euo pipefail

name="${1:-}"
placeholder="PROJECT_NAME"

if [[ ! "${name}" =~ ^[a-z][a-z0-9-]*$ ]]; then
  echo "usage: $0 NAME" >&2
  echo "NAME is the repository name, e.g. my-tool: lowercase letters, digits and hyphens, starting with a letter." >&2
  exit 1
fi

cd "$(git rev-parse --show-toplevel)"

if [[ -n "$(git status --porcelain)" ]]; then
  echo "The working tree has uncommitted changes; commit or stash them first, so the result is easy to review." >&2
  exit 1
fi

# Every tracked file that mentions the placeholder, except this script.
git grep -lz -e "${placeholder}" -- ':!scripts/init.sh' |
  xargs -0 perl -pi -e "s/${placeholder}/${name}/g"

git mv "cmd/${placeholder}" "cmd/${name}"

# The README's "Using this template" section is for the template only.
perl -0pi -e 's/<!-- template:begin -->.*?<!-- template:end -->\n*//s' README.md

git rm -q scripts/init.sh

echo "Renamed the project to ${name}. Next:"
echo "  1. Review the changes: git status && git diff"
echo "  2. Replace the placeholder descriptions in README.md, cmd/${name}, internal/cli, .goreleaser.yaml and docs/"
echo "  3. Run the gates: mise trust && mise install && mise run check"
