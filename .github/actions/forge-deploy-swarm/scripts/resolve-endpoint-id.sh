#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_NAME="${1:?endpoint_name is required}"
API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

ENDPOINT_ID="$({
  curl --fail-with-body -sS "${API}/endpoints" -H "${AUTH_HEADER}" \
  | jq -r --arg name "${ENDPOINT_NAME}" '.[] | select(.Name==$name) | .Id' \
  | head -n1
} || true)"

if [ -z "${ENDPOINT_ID}" ]; then
  echo "Endpoint '${ENDPOINT_NAME}' not found in Portainer" >&2
  exit 1
fi

echo "endpoint_id=${ENDPOINT_ID}" >> "$GITHUB_OUTPUT"
