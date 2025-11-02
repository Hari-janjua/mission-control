#!/bin/bash
set -e

echo "Starting Mission Control environment..."
docker-compose up -d --build
sleep 15

echo "Environment up. Issuing a mission..."
MISSION_ID=$(curl -s -X POST http://localhost:8080/missions \
  -H "Content-Type: application/json" \
  -d '{"command": "Secure area Alpha"}' | jq -r .mission_id)

echo "Mission ID: $MISSION_ID"

for i in {1..10}; do
  STATUS=$(curl -s http://localhost:8080/missions/$MISSION_ID | jq -r .status)
  echo "Status after $i sec: $STATUS"
  if [[ "$STATUS" == "COMPLETED" || "$STATUS" == "FAILED" ]]; then
    echo "Mission completed with status: $STATUS"
    break
  fi
  sleep 3
done;

docker-compose down
