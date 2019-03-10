FROM postgres:alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB neoway-db

ADD init.sql /docker-entrypoint-initdb.d/01-init.sql