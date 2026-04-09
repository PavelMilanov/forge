#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_NAME="${1:?endpoint_name is required}"
ENDPOINT_ID="$(
  forge deploy list \
  | awk -v target="${ENDPOINT_NAME}" '
      BEGIN { id=""; name="" }
      /^[[:space:]]*ID:[[:space:]]*/ {
        line=$0
        sub(/^[[:space:]]*ID:[[:space:]]*/, "", line)
        id=line
      }
      /^[[:space:]]*Name:[[:space:]]*/ {
        line=$0
        sub(/^[[:space:]]*Name:[[:space:]]*/, "", line)
        name=line
      }
      /^[[:space:]]*$/ {
        if (name == target && id != "") {
          print id
          exit 0
        }
      }
      END {
        if (name == target && id != "") {
          print id
        }
      }
    ' \
  | head -n1
)"

if [ -z "${ENDPOINT_ID}" ]; then
  echo "Endpoint '${ENDPOINT_NAME}' not found in forge deploy list output" >&2
  exit 1
fi

echo "endpoint_id=${ENDPOINT_ID}" >> "$GITHUB_OUTPUT"
