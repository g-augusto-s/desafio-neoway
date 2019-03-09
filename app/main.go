package main

import (
    "database/sql"
    "fmt"
    "os"
    "strconv"
    // "strings"
    _ "github.com/lib/pq"
)

func main() {

    for i:=0; i<10 ; i++{
        insertData("placeholder", "placeholder", "placeholder", "placeholder", "placeholder", "placeholder", i, i+10)
    }
}


func insertData(cpf, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra string, private, incompleto int){
    

    // Get env variables
    host     := os.Getenv("HOST")
    port,_   := strconv.Atoi(os.Getenv("PORT"))
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
      panic(err)
    }
    defer db.Close()
    
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Successfully connected!")

    sqlStatement := `
        INSERT INTO banco_tutorial (
            cpf, 
            private, 
            incompleto, 
            data_da_ultima_compra, 
            ticket_medio, 
            ticket_da_ultima_compra, 
            loja_mais_frequente, 
            loja_da_ultima_compra
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `
    id := 0
    err = db.QueryRow(sqlStatement, cpf, private, incompleto, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra).Scan(&id)

    if err != nil {
        panic(err)
    }

    fmt.Println("New record ID is:", id)
}