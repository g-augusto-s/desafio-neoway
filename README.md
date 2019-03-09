docker-compose up

docker ps 

docker exec -it [container-id] bash

psql -U postgres

\c neoway-db

CREATE TABLE base_teste (
    cpf VARCHAR(255) UNIQUE NOT NULL,
    private INT,  
    incompleto INT,
    data_da_ultima_compra VARCHAR(255),
    ticket_medio VARCHAR(255),
    ticket_da_ultima_compra VARCHAR(255),
    loja_mais_frequente VARCHAR(255),
    loja_da_ultima_compra VARCHAR(255)
);

##Outside docker container
docker cp base_teste.csv d07352f24cb2:/home


##Inside docker again
COPY base_teste FROM '/home/base_teste.csv' DELIMITERS ',' CSV;

ALTER TABLE base_teste ADD COLUMN id SERIAL PRIMARY KEY;

<!-- DROP NULL ROWS -->
<!-- UPDATE base_teste SET data_da_ultima_compra = NULL where data_da_ultima_compra = 'NULL'; -->

<!-- elect cast(i as text),cast(t as int)from test; -->