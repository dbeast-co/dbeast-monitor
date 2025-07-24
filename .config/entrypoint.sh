#!/bin/sh

if [ "${DEV}" = "false" ]; then
    echo "Starting test mode"
    /run.sh &   # Start Grafana in background
    GRAFANA_PID=$!
else
    echo "Starting development mode"

    if grep -i -q alpine /etc/issue; then
        /usr/bin/supervisord -c /etc/supervisord.conf &
        GRAFANA_PID=$!
    elif grep -i -q ubuntu /etc/issue; then
        /usr/bin/supervisord -c /etc/supervisor/supervisord.conf &
        GRAFANA_PID=$!
    else
        echo 'ERROR: Unsupported base image'
        exit 1
    fi
fi

# Wait for Grafana to be ready
echo "Waiting for Grafana to start..."
sleep 20

# Enable the plugin via API
echo "Enabling plugin..."
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"enabled": true, "pinned": true}' \
  -u admin:admin \
  http://localhost:3000/api/plugins/dbeast-dbeastmonitor-app/settings

# Keep container running
wait $GRAFANA_PID
