#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_ID="${1:?endpoint_id is required}"
STACK_ID="${2:?stack_id is required}"
STACK_FILE="${3:?stack_file is required}"
PRUNE="${4:?prune is required}"
PULL_IMAGE="${5:?pull_image is required}"

API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

jq -n \
  --argjson prune "${PRUNE}" \
  --argjson pull_image "${PULL_IMAGE}" \
  --rawfile content "${STACK_FILE}" \
  '{StackFileContent:$content,Env:[],Prune:$prune,PullImage:$pull_image}' \
| curl --fail-with-body -sS -X PUT \
    "${API}/stacks/${STACK_ID}?endpointId=${ENDPOINT_ID}" \
    -H "${AUTH_HEADER}" \
    -H "Content-Type: application/json" \
    --data @-

echo "stack_id=${STACK_ID}" >> "$GITHUB_OUTPUT"
echo "action=updated" >> "$GITHUB_OUTPUT"
