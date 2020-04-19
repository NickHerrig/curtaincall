package main

import (
    "database/sql"
    "flag"
    "html/template"
    "log"
    "net/http"
    "os"

    "curtaincall.tech/pkg/models/sqlite"

    _ "github.com/mattn/go-sqlite3"
)

type application struct {
    errorLog      *log.Logger
    infoLog       *log.Logger
    theaters      *sqlite.TheaterModel
    templateCache map[string]*template.Template
}

func main() {

    addr := flag.String("addr", ":8080", "HTTP network address")
    dsn := flag.String("dsn", "./test.db", "Database data source name")
    dr := flag.String("dr" ,"sqlite3", "The Database driver to use")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(*dr, *dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer db.Close()

    templateCache, err := newTemplateCache("./ui/html/")
    if err != nil {
        errorLog.Fatal(err)
    }

    app := &application{
        errorLog:      errorLog,
        infoLog:       infoLog,
        theaters:      &sqlite.TheaterModel{DB: db},
        templateCache: templateCache,
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

func openDB(dr, dsn string) (*sql.DB, error) {
    db, err := sql.Open(dr, dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
