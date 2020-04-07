#!/bin/bash

MIGRATIONS="/migrations"


# Make sure it will accept connections
until nc ${DB_HOST} ${DB_PORT} < /dev/null
do
    echo "waiting for postgres container..."
    sleep 2
done

# Make sure the DB is ready to accept commands
until echo ${DB_PASSWORD} | psql -p ${DB_PORT} -h ${DB_HOST} -U ${DB_USER} -W -d ${DB_NAME} -c "select 1"
do
    echo "waiting for postgres to accept connections..."
    sleep 2
done

# a wrapper for find that looks for migration files
find_migrations(){
  find -L "$MIGRATIONS" -maxdepth 1 -mindepth 1 -type f -name "*.sql" "$@"
}

migrate(){
  local fname=""
  find_migrations "$@" | sort | while read -r fname
  do
    echo ${DB_PASSWORD} | psql -p ${DB_PORT} -h ${DB_HOST} -d ${DB_NAME} -U ${DB_USER} -W ${DB_PASSWORD} -f "${fname}"
  done
}

migrate "$@"
