#!/bin/bash

echo "🚀 Deploying RegretBox..."

# Build and start the containers
docker-compose down
docker-compose build --no-cache
docker-compose up -d

echo "✅ RegretBox is now running!"
echo "🌐 Visit: http://localhost:8080"
echo "📊 Check status: docker-compose ps"
echo "📋 View logs: docker-compose logs -f"
