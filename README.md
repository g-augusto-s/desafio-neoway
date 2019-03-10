docker-compose up

docker ps 

docker exec -it [container-id] bash

psql -U postgres

\c neoway-db

CREATE TABLE banco_tutorial (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(255),
    private INT,  
    incompleto INT,
    data_da_ultima_compra VARCHAR(255),
    ticket_medio VARCHAR(255),
    ticket_da_ultima_compra VARCHAR(255),
    loja_mais_frequente VARCHAR(255),
    loja_da_ultima_compra VARCHAR(255)
);

##Outside docker container
$ docker cp app/assets/base_teste_sem_header.csv 5cadd68cd755:home/
$ docker cp app/assets/base_teste_min_sem_header.csv 5cadd68cd755:home/


##Inside docker again
COPY base_teste FROM '/home/base_teste.csv' DELIMITERS ',' CSV;

COPY banco_tutorial (
    cpf, 
    private, 
    incompleto, 
    data_da_ultima_compra, 
    ticket_medio, 
    ticket_da_ultima_compra, 
    loja_mais_frequente, 
    loja_da_ultima_compra
) FROM '/home/base_teste.csv' DELIMITERS ',' CSV;


<!-- ALTER TABLE base_teste ADD COLUMN id SERIAL PRIMARY KEY; -->

<!-- DROP NULL ROWS -->
<!-- UPDATE base_teste SET data_da_ultima_compra = NULL where data_da_ultima_compra = 'NULL'; -->

<!-- elect cast(i as text),cast(t as int)from test; -->