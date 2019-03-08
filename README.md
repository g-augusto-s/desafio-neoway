docker-compose up

docker ps 

docker exec -it [container-id] bash

psql -U postgres

\c neoway-db

CREATE TABLE base_teste2 (
    cpf VARCHAR(255) UNIQUE NOT NULL,
    private INT,  
    incompleto INT,
    data_da_ultima_compra VARCHAR(255),
    ticket_medio INT,
    ticket_da_ultima_compra INT,
    loja_mais_frequente VARCHAR(255),
    loja_da_ultima_compra VARCHAR(255)
);

docker cp base_teste.csv d07352f24cb2:/home

COPY base_teste FROM '/home/base_teste.csv' DELIMITERS ',' CSV;