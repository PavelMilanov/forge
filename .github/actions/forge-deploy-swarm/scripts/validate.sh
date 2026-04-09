#!/usr/bin/env bash
set -euo pipefail

ENDPOINT_NAME="${1:?endpoint_name is required}"
STACK_NAME="${2:?stack_name is required}"
STACK_FILE="${3:?stack_file is required}"

command -v curl >/dev/null 2>&1
command -v jq >/dev/null 2>&1

test -n "${PORTAINER_URL:-}"
test -n "${PORTAINER_TOKEN:-}"
test -n "${ENDPOINT_NAME}"
test -n "${STACK_NAME}"
test -f "${STACK_FILE}"
