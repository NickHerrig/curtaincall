package main

import (
    "database/sql"
    "flag"
    "log"
    "net/http"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main() {

    addr := flag.String("addr", ":8080", "HTTP network address")
    dsn := flag.String("dsn", "./test.db", "Database data source name")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    database, err := sql.Open("sqlite3", *dsn)
    if err != nil {
        errorLog.Fatal(err) 
    }
    defer database.Close()

    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }

    srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  app.routes(),
    }

    infoLog.Printf("Starting Curtain Call on port %s", *addr)
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}
