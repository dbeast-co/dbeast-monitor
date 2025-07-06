#!/bin/sh

# Start Grafana in the background
/run.sh "$@" &
GRAFANA_PID=$!

# Wait for Grafana to be ready
echo "Waiting for Grafana to start..."
sleep 10

# Enable the plugin via API
echo "Enabling plugin..."
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"enabled": true, "pinned": true}' \
  -u admin:admin \
  http://localhost:3000/api/plugins/dbeast-dbeastmonitor-app/settings

# Bring Grafana to the foreground
wait $GRAFANA_PID
