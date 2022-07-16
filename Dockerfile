FROM postgres:latest

LABEL image="postgres" type="local"

WORKDIR /app

ARG PORT

EXPOSE $PORT

ARG PG_DATA_LOCATION

ENV PGDATA=${PG_DATA_LOCATION}

ARG PG_PASS

ENV POSTGRES_PASSWORD=${PG_PASS}

CMD ["postgres"]