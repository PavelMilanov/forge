#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_ID="${1:?endpoint_id is required}"
STACK_NAME="${2:?stack_name is required}"
SWARM_ID="${3:?swarm_id is required}"
STACK_FILE="${4:?stack_file is required}"

API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"
STACK_FILENAME="$(basename "${STACK_FILE}")"

CREATE_RESPONSE="$(
  curl --fail-with-body -sS -X POST \
  "${API}/stacks/create/swarm/file?endpointId=${ENDPOINT_ID}" \
  -H "${AUTH_HEADER}" \
  -F "Name=${STACK_NAME}" \
  -F "SwarmID=${SWARM_ID}" \
  -F "file=@${STACK_FILE};filename=${STACK_FILENAME}" \
  -F 'Env=[]'
)"

RESOURCE_CONTROL_ID="$(printf '%s' "${CREATE_RESPONSE}" | jq -r '.ResourceControl.Id // .resourceControl.Id // empty')"

STACK_ID="$({
  curl --fail-with-body -sS "${API}/stacks" -H "${AUTH_HEADER}" \
  | jq -r --arg name "${STACK_NAME}" '.[] | select(.Name==$name) | .Id' \
  | head -n1
} || true)"

if [ -z "${STACK_ID}" ]; then
  echo "Stack '${STACK_NAME}' was created but id lookup failed" >&2
  exit 1
fi

echo "stack_id=${STACK_ID}" >> "$GITHUB_OUTPUT"
echo "action=created" >> "$GITHUB_OUTPUT"
echo "resource_control_id=${RESOURCE_CONTROL_ID}" >> "$GITHUB_OUTPUT"
