# Desafio Técnico - Processo Seletivo Neoway - Analista de Sistemas


## Indice

* [Introdução](#introdução)
* [Requisitos](#requisitos)
* [Abordagem](#abordagem)
* [Pré requisitos](#pré-requisitos)
* [Instalando](#instalando)
* [Deploy](#deploy)
* [Estrutura relacional](#estrutura-relacional)
* [Docker Compose](#docker-compose)
* [Feito com](#feito-com)
* [Autor](#autor)
* [Melhorias](#melhorias)


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

Com base nisso, e nos requisitos expostos na seção acima, minha abordagem inicial para ganhar tempo, foi transformar o arquivo dado *base_teste.txt* em um formato *base_teste.csv* utilizando-se do [Google Sheets](https://www.google.com/sheets/about/), já que tinha apenas 3 dias para aprender toda a linguagem e desenvolver a solução de uma forma funcional. 

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

Para fazer o deploy da aplicação, basta mudar para o diretorio criado e rodar o comando **docker-compose up**:

Mudar para o diretorio criado (**Linux**)

``` shell
$ cd desafio-neoway/
```

O repositorio já está com o arquivo **base_teste.csv** na pasta "app/assets/". Mas caso precise copiar : (**Linux**)

``` bash
$ cp "caminho_do_arquivo"/base_teste.csv "pasta_local_onde_foi_baixado_o_repositorio"/desafio-neoway/app/assets/
```
Executar o comando Docker Compose:
``` bash
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
## Docker Compose

O arquivo docker-compose.yml baixa as imagens e cria os containers para o banco de dados com Postgres e para o serviço em GO.


TROCAR A VARIAVEl **PORT** se precisar usar outra porta mais conveniente no arquivo "./docker-compose.yml" linha 32.

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
      - "HOST=TROCAR SEU IP LOCAL AQUI"
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

* **Guilherme Augusto** - [gaugustos](https://www.linkedin.com/in/gaugustos/)

## Melhorias

* Fazer a conversão do arquivo *base_texte.txt* para *base_texte.csv* com o próprio serviço em GO
* Organizar melhor a estrutura do código
* Criar arquivo para teste
* A conversão de txt para para csv acabou perdendo o numero 0 antes de 6 CPF's que estavam em uma formatação diferente, fazendo com que não passassem no teste de validação conforme [imagem](https://drive.google.com/open?id=1-F_JC9AzqyEHh_YIqXU0wvMlm-nN2GLg). Entretanto esses CPF's possuiam valores nulos de CNPJ's nas colunas de *LOJA MAIS FREQUÊNTE* e *LOJA DA ÚLTIMA COMPRA* na qual já seriam excluidas as respectivas linhas pela validação de CNPJ, tornando esse problema não muito prejudicial ao resultado final. Porém, ainda alvo de uma melhoria futura.