#!/bin/bash

set -e

IMAGE="olawaleasan/muchtodo:latest"

echo "Pulling latest image..."
docker pull $IMAGE

echo "Stopping existing container..."
docker stop app || true

echo "Removing old container..."
docker rm app || true

echo "Starting new container..."
docker run -d \
  -p 80:8080 \
  --name app \
  -e MONGO_URI="$MONGO_URI" \
  -e DB_NAME="starttech" \
  $IMAGE

echo "Deployment complete."
