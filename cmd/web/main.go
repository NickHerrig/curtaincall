package main

import (
    "crypto/tls"
    "database/sql"
    "flag"
    "html/template"
    "log"
    "net/http"
    "os"
    "time"

    "curtaincall.tech/pkg/models/sqlite"

    _ "github.com/mattn/go-sqlite3"
    "github.com/golangcollege/sessions"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
    errorLog      *log.Logger
    infoLog       *log.Logger
    session       *sessions.Session
    theaters      *sqlite.TheaterModel
    templateCache map[string]*template.Template
    users         *sqlite.UserModel
}

func main() {

    addr := flag.String("addr", ":8080", "HTTP network address")
    dsn := flag.String("dsn", "./test.db", "Database data source name")
    dr := flag.String("dr" ,"sqlite3", "The Database driver to use")
    secret := flag.String("secret" ,"akdjiekdjfldjfhuejdkiofadsalfjckj", "Secret")
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

    session := sessions.New([]byte(*secret))
    session.Lifetime = 12 * time.Hour
    session.Secure = true
    session.SameSite = http.SameSiteStrictMode

    app := &application{
        errorLog:      errorLog,
        infoLog:       infoLog,
        session:       session,
        theaters:      &sqlite.TheaterModel{DB: db},
        templateCache: templateCache,
        users:         &sqlite.UserModel{DB: db},
    }

    tlsConfig := &tls.Config{
        PreferServerCipherSuites: true,
        CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
    }

    srv := &http.Server{
        Addr:         *addr,
        ErrorLog:     errorLog,
        Handler:      app.routes(),
        TLSConfig:    tlsConfig,
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    infoLog.Printf("Starting Curtain Call on port %s", *addr)
    err = srv.ListenAndServeTLS("./tls/cert.pem","./tls/key.pem")
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
