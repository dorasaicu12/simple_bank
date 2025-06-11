#!/bin/sh

set -e

echo "Running DB migration..."
migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting the app..."
exec /app/main