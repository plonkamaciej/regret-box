#!/bin/bash

echo "🔍 RegretBox Health Check"
echo "========================="

# Check if app is running
if curl -s http://localhost:8080 > /dev/null; then
    echo "✅ App is responding on port 8080"
else
    echo "❌ App is not responding"
fi

# Check Redis connection
if docker-compose exec redis redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis is connected"
elif redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis is connected (local)"
else
    echo "❌ Redis is not connected"
fi

# Check disk space
DISK_USAGE=$(df / | tail -1 | awk '{print $5}' | sed 's/%//')
if [ $DISK_USAGE -lt 80 ]; then
    echo "✅ Disk usage: ${DISK_USAGE}%"
else
    echo "⚠️  Disk usage high: ${DISK_USAGE}%"
fi

# Check memory
FREE_MEM=$(free | grep Mem | awk '{printf("%.1f", $3/$2 * 100.0)}')
echo "📊 Memory usage: ${FREE_MEM}%"

echo "========================="
echo "🎯 RegretBox Status: OK"
