package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"curtaincall.tech/pkg/models"
	"curtaincall.tech/pkg/models/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	shows    interface {
		Get(int) (*models.Show, error)
		Latest(int) ([]*models.Show, error)
	}
	theaters interface {
		Get(int) (*models.Theater, error)
		Latest() ([]*models.Theater, error)
	}
	templateCache map[string]*template.Template
}

func requireEnv(name string) string {
    value, ok := os.LookupEnv(name)
    if !ok {
      log.Fatalf("No enviornment variable: %q", name)
    }

    return value
}

func main() {

    dr := requireEnv("CC_DB_DRIVER")
    dsn := requireEnv("CC_DB_DSN")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dr, dsn)
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
		shows:         &sqlite.ShowModel{DB: db},
		theaters:      &sqlite.TheaterModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:         ":8080",
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting Curtain Call on port :8080")
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
