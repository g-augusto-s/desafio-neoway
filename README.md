# Desafio Técnico - Processo Seletivo Neoway - Analista de Sistemas

###Requisitos:
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

## Introdução

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Abordagem

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

## Instalando

Explain what these tests test and why

```
Give an example
```

## Deploy

Add additional notes about how to deploy this on a live system

```
Give an example
```

## Estrutura relacional

Add additional notes about how to deploy this on a live system

```
Give an example
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

* A
* B
* C