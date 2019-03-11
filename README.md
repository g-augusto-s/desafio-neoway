# Desafio Técnico - Processo Seletivo Neoway - Analista de Sistemas


## Indice
[Introdução](#introdução)

[Requisitos](#requisitos)

[Abordagem](#abordagem)

[Pré requisitos](#pré-requisitos)

[Instalando](#instalando)

[Deploy](#deploy)

[Estrutura relacional](#estrutura-relacional)

[Arquivo docker-compose.yml](#arquivo-docker-compose.yml)

[Feito com](#feito-com)

[Autor](#autor)

[Melhorias](#melhorias)


## Introdução

Esse projeto foi construido para o processo seletivo da empresa Neoway, vaga de Analista de Sistemas, seguindo os requisitos descritos na seção abaixo.


### Requisitos:
- Criar um serviço em GO que receba um arquivo csv/txt de entrada (Arquivo Anexo)
- Este serviço deve persistir no banco de dados relacional (postgres) todos os dados contidos no arquivo
  Obs: O arquivo não possui um separador muito convencional
 
- Deve-se fazer o split dos dados em colunas no banco de dados
 Obs: pode ser feito diretamente no serviço em GO ou em sql
 
- Realizar higienização dos dados após persistência (sem acento, maiusculo, etc)
- Validar os CPFs/CNPJs contidos (validos e não validos numericamente)
- Todo o código deve estar disponível em repositório publico do GIT
 
**Desejável:**
- Utilização das linguagen GOLANG para o desenvolvimento do serviço
- Utilização do DB Postgres
- Docker Conpose , com orientações para executar (arquivo readme) 

**Você será avaliado por:**
- Utilização de melhores práticas de desenvolvimento (nomenclatura, funções, classes, etc);
- Utilização dos recursos mais recentes das linguagens;
- Boa organização lógica e documental (readme, comentários, etc);
- Cobertura de todos os requisitos obrigatórios.

Nota:
Todo a estrutura relacional dev estar documentada (criação das tabelas, etc)
Criação de um arquivo README com as instruções de instalação juntamente com as etapas necessárias para configuração.
Você pode escolher sua abordagem de arquitetura e solução técnica.
Apresentar-nos apenas o link do Github com o projeto.

___


## Abordagem

Esta seção visa explicar qual a abordagem adotada para a resolução do problema apresentado. 

Segundo conversa inicial com a Rafaela Leal do time de *Gente e Gestão*, essa vaga exige, além das habilidades técnicas, autonomia, aprendizado rápido e resolução de problemas com prazos curtos. 

Com base nisso, e nos requisitos expostos na seção acima, minha abordagem inicial para ganhar tempo, foi transformar o arquivo dado *base_teste.txt* em um formato *base_teste.csv* utilizando-se do [Google Sheets](https://www.google.com/sheets/about/), já que essa tinha apenas 3 dias para aprender toda a linguagem e desenvolver a solução de uma forma funcional. 

- Primeiramente apaguei a primeira linha que continham os nomes das colunas
- Importei o arquivo no [Google Sheets](https://www.google.com/sheets/about/)
- Salvei o novo arquivo (sem a linha com os nomes das tabelas) como **base_teste.csv**
- Com o arquivo **base_teste.csv** em mãos, partimos para as seções seguintes

## Pré requisitos

Para utilizar-se do serviço criado é necessário ter instalado os seguintes programas:

* [Docker CE](https://docs.docker.com/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Instalando

Para instalar o serviço, basta clonar o repositorio

``` shell
$ git clone https://github.com/g-augusto-s/desafio-neoway.git
```

## Deploy

Para fazer o deploy da aplicação, basta mudar para o diretorio criado e rodar os comandos **docker-compose build && docker-compose up**:

Mudar para o diretorio criado (**Linux**)

``` shell
$ cd desafio-neoway/
```

Copiar o arquivo **base_teste.csv** para a pasta app/assets/ (**Linux**)

``` bash
$ cp "path_do_arquivo"/base_teste.csv app/assets/
```
**Trocar HOST no arquivo docker-compose.yml para seu ip local** conforme [Arquivo docker-compose.yml](#arquivo-docker-compose.yml)

Executar os comandos Docker Compose:
``` bash
$ docker-compose build
$ docker-compose up
```

## Estrutura relacional
O arquivo **init.sql** cria a tabela necessaria no banco de dados ao rodar o comando:

``` sql
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
```
## Arquivo docker-compose.yml

O arquivo docker-compose.yml cria baixa as imagens e cria os containers para o banco de dados com Postgres e para o serviço em GO.

### Atenção com as variáveis de ambiente HOST e PORT

**LEMBRA DE TROCAR AS VARIAVEIS HOST E PORT** para o IP Local e porta conveniente.

*tentei substituir o HOST por localhost porém o docker não reconhecia

``` yml
version: '3'

services:
  db:
    image: neoway-db
    container_name: neoway-container-db
    restart: 'always'
    ports:
      - '5432:5432'
    environment:
      - "POSTGRES_USER=postgres"
      - "POSTGRES_DB=neoway-db"
      - "POSTGRES_PASSWORD=postgres"
    build: './db'
    volumes:
      - ./db/init.sql:/docker-entrypoint-initdb.d/01-init.sql
      - ./app/assets/:/home/
      - ./db/postgresql-data:/var/lib/postgresql/data

  backend:
    build: ./app/
    image: go-env
    container_name: go-env-container
    volumes:
      - ./app/:/go/src/app/
    command: bash -c "sleep 3; go run main.go"
    ports:
      - "8080:8080"
    environment:
      - "CSV_PATH=/home/base_teste.csv"
      - "HOST=192.168.25.109"
      - "PORT=5432"
      - "USER=postgres"
      - "PASSWORD=postgres"
      - "DBNAME=neoway-db"
      - "TABLENAME=base_teste"
    restart: on-failure
    depends_on:
      - db
    links:
      - db
```

## Feito com

* [lib/pq](https://github.com/lib/pq) - A pure Go postgres driver for Go's database/sql package
* [BR Doc](https://github.com/Nhanderu/brdoc) - CPF, CNPJ and CEP validator for Go!
* [Docker CE](https://docs.docker.com/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)
* [Golang Docker Image](https://hub.docker.com/_/golang)
* [Postgres](https://hub.docker.com/_/postgres)

## Autor 

* **Guilherme Augusto** -[g-augusto-s](https://github.com/g-augusto-s/)

## Melhorias

* Fazer a conversão do arquivo *base_texte.txt* para *base_texte.csv* com o serviço em GO