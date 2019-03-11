CREATE TABLE IF NOT EXISTS base_teste (
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