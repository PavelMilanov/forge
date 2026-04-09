#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_ID="${1:?endpoint_id is required}"
STACK_ID="${2:?stack_id is required}"

API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

RESOURCE_CONTROL_ID="$({
  curl --fail-with-body -sS "${API}/stacks/${STACK_ID}?endpointId=${ENDPOINT_ID}" -H "${AUTH_HEADER}" \
  | jq -r '.ResourceControl.Id // .resourceControl.Id // empty'
} || true)"

if [ -z "${RESOURCE_CONTROL_ID}" ]; then
  RESOURCE_CONTROL_ID="$({
    curl --fail-with-body -sS "${API}/stacks" -H "${AUTH_HEADER}" \
    | jq -r --arg sid "${STACK_ID}" '.[] | select((.Id|tostring)==$sid) | (.ResourceControl.Id // .resourceControl.Id // empty)' \
    | head -n1
  } || true)"
fi

if [ -z "${RESOURCE_CONTROL_ID}" ]; then
  echo "ResourceControl ID not found for stack_id=${STACK_ID}" >&2
  exit 1
fi

echo "resource_control_id=${RESOURCE_CONTROL_ID}" >> "$GITHUB_OUTPUT"
