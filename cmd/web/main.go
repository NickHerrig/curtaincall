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

	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	shows    interface {
		Get(int) (*models.Show, error)
		Latest(int) ([]*models.Show, error)
	}
	theaters interface {
		Insert(string) (int, error)
		Get(int) (*models.Theater, error)
		Latest() ([]*models.Theater, error)
	}
	templateCache map[string]*template.Template
	users         interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
	}
}

func requireEnv(name string) string {
    value, ok := os.LookupEnv(name)
    if !ok {
      log.Fatalf("No enviornment variable: %q", name)
    }

    return value
}

func main() {

    secret := requireEnv("CC_SESSION_SECRET")
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

	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		shows:         &sqlite.ShowModel{DB: db},
		theaters:      &sqlite.TheaterModel{DB: db},
		templateCache: templateCache,
		users:         &sqlite.UserModel{DB: db},
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
