#!/usr/bin/env bash
set -euo pipefail

RESOURCE_CONTROL_ID="${1:?resource_control_id is required}"
TEAM_IDS_RAW="${2:-}"

if [ -z "${TEAM_IDS_RAW}" ]; then
  echo "team_ids is empty, skip"
  exit 0
fi

API="${PORTAINER_URL%/}/api"
AUTH_HEADER="X-API-Key: ${PORTAINER_TOKEN}"

TEAMS_JSON="$(jq -cn --arg ids "${TEAM_IDS_RAW}" '$ids | split(",") | map(gsub("^\\s+|\\s+$";"")) | map(select(length>0)) | map(tonumber)')"

if [ "${TEAMS_JSON}" = "[]" ]; then
  echo "team_ids produced an empty team list" >&2
  exit 1
fi

jq -cn --argjson teams "${TEAMS_JSON}" '{AdministratorsOnly:false,teams:$teams}' \
| curl --fail-with-body -sS -X PUT \
    "${API}/resource_controls/${RESOURCE_CONTROL_ID}" \
    -H "${AUTH_HEADER}" \
    -H "Content-Type: application/json" \
    --data @-
