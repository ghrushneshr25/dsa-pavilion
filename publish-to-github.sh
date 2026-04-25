#!/usr/bin/env bash
# Create the dsa-pavilion repository on GitHub and push (requires: brew install gh && gh auth login)
set -euo pipefail
cd "$(dirname "$0")"

if ! command -v gh &>/dev/null; then
  echo "Install GitHub CLI: brew install gh" >&2
  exit 1
fi

if ! gh auth status -h github.com &>/dev/null; then
  echo "Log in to GitHub (browser)…"
  gh auth login -h github.com -p https -w
fi

# Optional: first argument = full repo name, e.g. ghrushneshr25/dsa-pavilion
REPO_ARG="${1:-}"
if [ -n "$REPO_ARG" ]; then
  FULL_REPO="$REPO_ARG"
else
  LOGIN=$(gh api user -q .login 2>/dev/null || true)
  if [ -z "${LOGIN}" ]; then
    echo "Could not read GitHub login." >&2
    exit 1
  fi
  FULL_REPO="${LOGIN}/dsa-pavilion"
  echo "Using repository: ${FULL_REPO} (pass 'org/dsa-pavilion' as the first arg to override)"
fi

if git remote get-url origin &>/dev/null; then
  echo "Remote 'origin' already set — pushing only."
  git push -u origin "$(git branch --show-current)"
  exit 0
fi

gh repo create "${FULL_REPO}" \
  --public \
  --source=. \
  --remote=origin \
  --push \
  --description "DSA Pavilion: Docusaurus site for DSA in Go, generated with docgen from the dsa code repo"

echo "Done. Remote: https://github.com/${FULL_REPO}"
