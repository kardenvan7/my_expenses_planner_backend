FROM mysql:8.0.29

COPY ./app/database/*.sql /docker-entrypoint-initdb.d/