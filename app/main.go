package main

import (
    "database/sql"
    "fmt"
    "os"
    // "strings"
    _ "github.com/lib/pq"
    "gopkg.in/Nhanderu/brdoc.v1"
)

type Connect struct {
    HOST, PORT, USER, PASSWORD, DBNAME, TABLENAME  string    
}

func main() {
    c := Connect {
        os.Getenv("HOST"),
        os.Getenv("PORT"),
        os.Getenv("USER"),
        os.Getenv("PASSWORD"),
        os.Getenv("DBNAME"),
        os.Getenv("TABLENAME"),
    }

	// p := &c

    c.dbConnect();

    // // Connect to DB
    // dbConnect()
    
    // Verify and copy values to table
    copyToDB()
    
    // Validate CPF & CNPJ
    cpfIsValid()
    cnpjIsValid()

    // for i:=0; i<10 ; i++{
    //     insertData("placeholder", "placeholder", "placeholder", "placeholder", "placeholder", "placeholder", i, i+10)
    // }
}


func insertData(cpf, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra string, private, incompleto int){
    

    // Get env variables
    host     := os.Getenv("HOST")
    port     := os.Getenv("PORT")
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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


func copyToDB(){
    // Get env variables
    host     := os.Getenv("HOST")
    port     := os.Getenv("PORT") //convert port to int
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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

    sqlStatement := `SELECT cpf FROM banco_tutorial;`
    var returnedCpf string

    row := db.QueryRow(sqlStatement)

    switch err := row.Scan(&returnedCpf); err {
        case sql.ErrNoRows:
            fmt.Println("Table empty")
            fmt.Println()
            sqlStatement := `
                COPY banco_tutorial (
                    cpf, 
                    private, 
                    incompleto, 
                    data_da_ultima_compra, 
                    ticket_medio, 
                    ticket_da_ultima_compra, 
                    loja_mais_frequente, 
                    loja_da_ultima_compra
                ) FROM '/home/base_teste_sem_header.csv' DELIMITERS ',' CSV;
            `
            _,err = db.Exec(sqlStatement)

            if err != nil {
                panic(err)
            }
            fmt.Println("Database successful copy")
        case nil:
            fmt.Println("Database already copy")
        default:
            panic(err)
    }
}

func importCSV(){
    // Get env variables
    // host     := os.Getenv("HOST")
    // port,_   := strconv.Atoi(os.Getenv("PORT")) //convert port to int
    // user     := os.Getenv("USER")
    // password := os.Getenv("PASSWORD")
    // dbname   := os.Getenv("DBNAME")

    // psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    //   "password=%s dbname=%s sslmode=disable",
    //   host, port, user, password, dbname)

    // db, err := sql.Open("postgres", psqlInfo)

    // if err != nil {
    //   panic(err)
    // }
    // defer db.Close()
    
    // err = db.Ping()
    // if err != nil {
    //     panic(err)
    // }
    
    // fmt.Println("Successfully connected!")

    // // Copy the CSV file into the /tmp directory, so the server has access to it
    // tempfile, err := CopyFile("base_teste_sem_header.csv", "./assets")
    // if err != nil  {
    //     fmt.Errorf("cannot make temporary copy of CSV file: %v", err)
    // }

    // // Load the CSV file directly into PostgreSQL, into factual.places
    // _, err = db.Exec("copy banco_tutorial2 " +
    //                     "( cpf, private, incompleto, data_da_ultima_compra, ticket_medio,"+ 
    //                     "ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra) " +
    //                     "from " + tempfile + " with csv;")
    // if err != nil {
    //     fmt.Errorf("cannot copy CSV file into database: %v")
    // }

    // fmt.Println("Copia feita com sucesso!")
}

func (c *Connect) dbConnect() {

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
      "password=%s dbname=%s sslmode=disable",
      c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
      panic(err)
    }
    defer db.Close()
    
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected inside method!")
}

func cpfIsValid(){
    // Get env variables
    host     := os.Getenv("HOST")
    port     := os.Getenv("PORT")
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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

    sqlStatement := `SELECT cpf, id FROM banco_tutorial;`

    rows,err := db.Query(sqlStatement)

    if err != nil {
        // handle this error better than this
        panic(err)
    }

    defer rows.Close()

    i:=0
    for rows.Next() {
        var returnedCpf string
        var returnedID int

        err = rows.Scan(&returnedCpf, &returnedID)
        if err != nil {
            // handle this error
            panic(err)
        }
        // fmt.Printf("CPF: %v - ID: %v - É valido? %v\n", returnedCpf, returnedID, brdoc.IsCPF(returnedCpf))

        if !(brdoc.IsCPF(returnedCpf)){
            sqlStatement := `
                DELETE FROM banco_tutorial
                WHERE id=$1
            `
            _,err = db.Exec(sqlStatement, returnedID)

            if err != nil {
                panic(err)
            }

            fmt.Printf("Successfully delete rows %v with unvalid CPF values\n", returnedID )
            i++
        }
    }
    if i>0{
        fmt.Printf("Deleted %v rows with unvalid CPF values\n", i )
    }


    // get any error encountered during iteration
    err = rows.Err()
    if err != nil {
        panic(err)
    }
}

func cnpjIsValid(){
    // Get env variables
    host     := os.Getenv("HOST")
    port     := os.Getenv("PORT")
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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

    sqlStatement := `SELECT loja_mais_frequente, loja_da_ultima_compra, id FROM banco_tutorial;`

    rows,err := db.Query(sqlStatement)

    if err != nil {
        // handle this error better than this
        panic(err)
    }

    defer rows.Close()

    i:=0
    for rows.Next() {
        var moreFrequentCNPJ string
        var lastBuyCNPJ string
        var returnedID int

        err = rows.Scan(&moreFrequentCNPJ, &lastBuyCNPJ, &returnedID)
        if err != nil {
            // handle this error
            panic(err)
        }
        // fmt.Printf("CNPJ mais frequente: %v - É valido? %v / CNPJ ultima compra: %v - É valido? %v- ID: %v\n", moreFrequentCNPJ, brdoc.IsCNPJ(moreFrequentCNPJ), lastBuyCNPJ, brdoc.IsCNPJ(lastBuyCNPJ), returnedID)
        
        if !(brdoc.IsCNPJ(moreFrequentCNPJ) || brdoc.IsCNPJ(lastBuyCNPJ)){
            sqlStatement := `
                DELETE FROM banco_tutorial
                WHERE id=$1
            `
            _,err = db.Exec(sqlStatement, returnedID)

            if err != nil {
                panic(err)
            }

            fmt.Printf("Successfully delete rows %v with unvalid CNPJ value\n", returnedID )
            i++
        }
    }

    if i>0{
        fmt.Printf("Deleted %v rows with unvalid CNPJ values\n", i )
    }
    // get any error encountered during iteration
    err = rows.Err()
    if err != nil {
        panic(err)
    }
}