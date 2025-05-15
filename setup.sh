#!/bin/bash

set -e

APP_NAME="Replicant"

echo "🚫 Stopping and cleaning old containers..."
docker-compose down --volumes --remove-orphans

echo "🛠️ Building $APP_NAME..."
docker-compose build

echo "🚀 Starting $APP_NAME containers..."
docker-compose up -d

echo "📦 Ensuring model 'mistral' is downloaded..."
docker exec ollama ollama pull mistral || true

echo "✅ $APP_NAME is now running."
echo "To chat with your replicant, attach to the container:"
echo "  docker attach replicant"
