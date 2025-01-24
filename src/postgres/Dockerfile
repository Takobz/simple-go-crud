FROM postgres:latest

# Postgres container will run any scripts in docker-entrypoint-initdb.d on startup.
COPY ./scripts/create-initial-tables-if-not-exists.sql /docker-entrypoint-initdb.d/

# Add config if you want this will be used in a compose file.