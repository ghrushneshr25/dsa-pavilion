#!/usr/bin/env bash
# Generate MDX, sidebars.js, GitHub route map, and ../dsa/readme.md
# Default sibling layout: ../dsa next to this repo.
set -euo pipefail
cd "$(dirname "$0")"
go -C docgen run . "$@"
