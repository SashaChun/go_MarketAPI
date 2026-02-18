#!/bin/bash

set -e

echo "🚀 Starting deployment..."

cd ~/untitled9

echo "📥 Pulling latest changes..."
git pull origin main

echo "🛑 Stopping existing containers..."
docker-compose -f docker-compose.prod.yml down

echo "🏗️  Building and starting containers..."
docker-compose -f docker-compose.prod.yml up -d --build

echo "🧹 Cleaning up old images..."
docker system prune -af

echo "✅ Deployment complete!"
echo "📊 Container status:"
docker-compose -f docker-compose.prod.yml ps
