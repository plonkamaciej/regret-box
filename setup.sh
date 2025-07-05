#!/bin/bash

echo "🔧 Setting up RegretBox for development..."

# Check if Redis is installed and running
if ! command -v redis-server &> /dev/null; then
    echo "❌ Redis not found. Installing Redis..."
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        sudo apt update
        sudo apt install -y redis-server
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        if command -v brew &> /dev/null; then
            brew install redis
        else
            echo "Please install Homebrew first: https://brew.sh"
            exit 1
        fi
    fi
fi

# Start Redis if not running
if ! pgrep -x "redis-server" > /dev/null; then
    echo "🚀 Starting Redis server..."
    redis-server --daemonize yes
fi

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Build and run the application
echo "🏗️ Building RegretBox..."
go build -o regretbox .

echo "✅ Setup complete!"
echo "🚀 Starting RegretBox..."
./regretbox
