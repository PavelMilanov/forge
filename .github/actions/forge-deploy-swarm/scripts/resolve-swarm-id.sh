#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_ID="${1:?endpoint_id is required}"
API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

SWARM_ID="$({
  curl --fail-with-body -sS "${API}/endpoints/${ENDPOINT_ID}/docker/info" -H "${AUTH_HEADER}" \
  | jq -r '.Swarm.Cluster.ID'
} || true)"

if [ -z "${SWARM_ID}" ] || [ "${SWARM_ID}" = "null" ]; then
  echo "Swarm ID not found for endpoint_id=${ENDPOINT_ID}" >&2
  exit 1
fi

echo "swarm_id=${SWARM_ID}" >> "$GITHUB_OUTPUT"
