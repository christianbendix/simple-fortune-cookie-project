#!/usr/bin/env bash
set -euo pipefail

URL="${1:-http://localhost:8080}"
RETRIES="${2:-30}"

echo "Smoke testing $URL"

for i in $(seq 1 "$RETRIES"); do
  if curl -fsS --max-time 5 "$URL" > /dev/null; then
    echo "OK after ${i}s"
    exit 0
  fi
  sleep 1
done

echo "FAILED: $URL never became reachable"
exit 1
