FROM alpine:latest

RUN apk update && \
    apk add bash postgresql-client && \
    apk add --update openssl

COPY internal/database/migrate.sh /usr/local/bin/

RUN chmod u+x /usr/local/bin/migrate.sh

COPY ./internal/database/sql/ /migrations/

ENTRYPOINT ["migrate.sh"]