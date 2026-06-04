#!/bin/sh
set -e

DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

echo "Running migrations..."
/app/bin/goose -dir /app/migrations postgres "${DSN}" up

if [ "${RUN_SEED}" = "true" ]; then
  echo "Seeding test data..."
  /app/bin/seed
fi

echo "Starting server..."
exec /app/bin/server