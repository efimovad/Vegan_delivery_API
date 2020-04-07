FROM postgres:11.7

COPY internal/database/sql/init_schema.sql /docker-entrypoint-initdb.d/
