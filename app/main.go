package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "os"
    "time"
    _ "github.com/lib/pq"
    "gopkg.in/Nhanderu/brdoc.v1"
)

type Connect struct {
    HOST, PORT, USER, PASSWORD, DBNAME, TABLENAME, CSV_PATH  string    
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

func main() {
    c := Connect {
        os.Getenv("HOST"),
        os.Getenv("PORT"),
        os.Getenv("USER"),
        os.Getenv("PASSWORD"),
        os.Getenv("DBNAME"),
        os.Getenv("TABLENAME"),
        os.Getenv("CSV_PATH"),
    }

    fmt.Println()
    fmt.Printf("Main function starts at: %v\n",time.Now())
    fmt.Println("<----------------- START ------------------->")
   
    // Verify and copy values to table
    copyToDB(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)
    // importCSV(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)
    
    // Validate CPF & CNPJ
    cpfIsValid(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)
    cnpjIsValid(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)

    fmt.Println("<----------------- ENDS ------------------->")
    fmt.Printf("Main function stops at: %v",time.Now())
    fmt.Println()
    fmt.Println()
    
    // for i:=0; i<10 ; i++{
        //     insertData("placeholder", "placeholder", "placeholder", "placeholder", "placeholder", "placeholder", i, i+10)
        // }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func copyToDB(host, port, user, password, dbname, table_name, csv_path string){

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    check(err)

    defer db.Close()
    
    err = db.Ping()
    check(err)
    
    fmt.Println("Successfully connected for copy or verify DB!")

    sqlStatement := `SELECT cpf FROM ` + table_name + " ;"
    var returnedCpf string

    row := db.QueryRow(sqlStatement)

    switch err := row.Scan(&returnedCpf); err {
        case sql.ErrNoRows:
            fmt.Println("Table empty")
            fmt.Println()
            sqlStatement := fmt.Sprintf(`
                COPY %s (
                    cpf, 
                    private, 
                    incompleto, 
                    data_da_ultima_compra, 
                    ticket_medio, 
                    ticket_da_ultima_compra, 
                    loja_mais_frequente, 
                    loja_da_ultima_compra
                ) FROM '%s' DELIMITERS ',' CSV;`, 
            table_name, csv_path)

            _,err = db.Exec(sqlStatement)
            check(err)

            fmt.Println("Database successful copy")
        case nil:
            fmt.Println("Database already copy")
        default:
            panic(err)
    }
}

func cpfIsValid(host, port, user, password, dbname, table_name, csv_path string){
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    check(err)
    defer db.Close()
    
    err = db.Ping()
    check(err)
    
    fmt.Println()
    fmt.Println("Successfully connected for CPF validation!")

    sqlStatement := fmt.Sprintf(`SELECT cpf, id FROM %s;`, table_name)

    rows,err := db.Query(sqlStatement)
    check(err)

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

        if !(brdoc.IsCPF(returnedCpf)){
            sqlStatement := fmt.Sprintf(`
                DELETE FROM %s
                WHERE id=%d
            `, 
            table_name, returnedID)

            _,err = db.Exec(sqlStatement)

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
    check(err)
}

func cnpjIsValid(host, port, user, password, dbname, table_name, csv_path string){
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    check(err)
    defer db.Close()
    
    err = db.Ping()
    check(err)
    
    fmt.Println()
    fmt.Println("Successfully connected for CNPJ validation!")

    sqlStatement := fmt.Sprintf(`SELECT loja_mais_frequente, loja_da_ultima_compra, id FROM %s;`, table_name)

    rows,err := db.Query(sqlStatement)
    check(err)

    defer rows.Close()

    i:=0
    for rows.Next() {
        var moreFrequentCNPJ string
        var lastBuyCNPJ string
        var returnedID int

        err = rows.Scan(&moreFrequentCNPJ, &lastBuyCNPJ, &returnedID)
        check(err)

        if !(brdoc.IsCNPJ(moreFrequentCNPJ) || brdoc.IsCNPJ(lastBuyCNPJ)){
            sqlStatement := fmt.Sprintf(`
                DELETE FROM %s
                WHERE id=%d
            `, 
            table_name, returnedID)

            _,err = db.Exec(sqlStatement)
            check(err)

            fmt.Printf("Successfully delete rows %v with unvalid CNPJ value\n", returnedID )
            i++
        }
    }

    if i>0{
        fmt.Printf("Deleted %v rows with unvalid CNPJ values\n", i )
    }
    // get any error encountered during iteration
    err = rows.Err()
    check(err)
}







func importCSV(host, port, user, password, dbname, table_name, csv_path string){

    csv_path = "/home/base_teste_min_sem_header.csv"
    table_name = "banco_tutorial_min"

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    check(err)
    defer db.Close()
    
    err = db.Ping()
    check(err)
    
    fmt.Println()
    fmt.Println("Successfully connected for import CSV!")
    fmt.Println()

    dat, err := ioutil.ReadFile("/go/src/app/assets/base_teste_min_sem_header.csv") // CSV_PATH=/go/src/app/assets/base_teste_min_sem_header.csv
    check(err)

    f, err := os.Open("/go/src/app/assets/base_teste_min_sem_header.csv")
    check(err)

    fmt.Printf("File dat value: %v \n", string(dat))
    fmt.Printf("File dat type: %T\n", dat)
    fmt.Printf("File f value: %v \n File f type: %T\n", *f, *f)
    fmt.Printf("File f value: %v \n File f type: %T\n", f, f)

    // Load the CSV file directly into PostgreSQL, into factual.places
    sqlStatement := fmt.Sprintf(`
        COPY %s (
            cpf, 
            private, 
            incompleto, 
            data_da_ultima_compra, 
            ticket_medio, 
            ticket_da_ultima_compra, 
            loja_mais_frequente, 
            loja_da_ultima_compra
        ) FROM '%s' DELIMITERS ',' CSV;`, 
    table_name, csv_path)

    _,err = db.Exec(sqlStatement)

    if err != nil {
        panic(err)
    }
    fmt.Println("Import CSV with success!")
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