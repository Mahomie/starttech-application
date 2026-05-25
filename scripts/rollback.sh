#!/bin/bash

PREVIOUS_IMAGE="olawaleasan/muchtodo:previous"

echo "Stopping current container..."
docker stop app || true

echo "Removing current container..."
docker rm app || true

echo "Starting previous stable version..."

docker run -d \
  -p 80:8080 \
  --name app \
  -e MONGO_URI="$MONGO_URI" \
  -e DB_NAME="starttech" \
  $PREVIOUS_IMAGE

echo "Rollback complete."
