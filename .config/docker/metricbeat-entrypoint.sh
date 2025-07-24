#!/bin/sh

# Wait for Elasticsearch to be available
until curl -s http://elasticsearch:9200 >/dev/null; do
  echo "Waiting for Elasticsearch..."
  sleep 2
done

# Get cluster UUID
CLUSTER_ID=$(curl -s http://elasticsearch:9200 | grep cluster_uuid | awk -F'"' '{print $4}')

if [ -z "$CLUSTER_ID" ]; then
  echo "Failed to fetch cluster UUID. Exiting."
  exit 1
fi

echo "Cluster UUID: $CLUSTER_ID"

# Replace the placeholder in the template with actual cluster UUID
sed "s/__CLUSTER_ID__/${CLUSTER_ID}/" /usr/share/metricbeat/metricbeat.template.yml > /usr/share/metricbeat/metricbeat.yml

# Start Metricbeat with the generated config
exec metricbeat -e -c /usr/share/metricbeat/metricbeat.yml
