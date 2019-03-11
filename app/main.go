package main

import (
    "database/sql"
    "fmt"
    "os"
    "time"
    _ "github.com/lib/pq"
    "gopkg.in/Nhanderu/brdoc.v1"
)

// Create Connect struct to hold environment variables for DB connection
type Connect struct {
    HOST, PORT, USER, PASSWORD, DBNAME, TABLENAME, CSV_PATH  string    
}

// Main function
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
    
    // Validate CPF & CNPJ
    cpfIsValid(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)
    cnpjIsValid(c.HOST, c.PORT, c.USER, c.PASSWORD, c.DBNAME, c.TABLENAME, c.CSV_PATH)

    fmt.Println("<----------------- ENDS ------------------->")
    fmt.Printf("Main function stops at: %v",time.Now())
    fmt.Println()
    fmt.Println()
}

// Check any erros returned
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Persist data from $CSV_PATH into PostgresDB
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

    // Verify if table is empty
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

// Verify if the CPF data is valid and delete those rows which not
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
    // Verify all CPF in the table
    for rows.Next() {
        var returnedCpf string
        var returnedID int

        err = rows.Scan(&returnedCpf, &returnedID)
        check(err)
        // If the CPF doenst exists, the row is delete
        if !(brdoc.IsCPF(returnedCpf)){
            sqlStatement := fmt.Sprintf(`
                DELETE FROM %s
                WHERE id=%d
            `, 
            table_name, returnedID)

            _,err = db.Exec(sqlStatement)

            check(err)

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

// Verify if the CNPJ data is valid and delete those rows which not
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

        // If the CNPJ doenst exists, the row is delete
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