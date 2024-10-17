#!/bin/sh

set -e 

echo "RUN DB MIGRATION "

source /app/app.env

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo " start the app"

exec "$@"