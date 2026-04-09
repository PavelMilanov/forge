#!/usr/bin/env bash
set -euo pipefail

STACK_NAME="${1:?stack_name is required}"
API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

STACK_ID="$({
  curl --fail-with-body -sS "${API}/stacks" -H "${AUTH_HEADER}" \
  | jq -r --arg name "${STACK_NAME}" '.[] | select(.Name==$name) | .Id' \
  | head -n1
} || true)"

echo "stack_id=${STACK_ID}" >> "$GITHUB_OUTPUT"
