package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
    "path/filepath"
	"time"

	"curtaincall.tech/pkg/models"
	"curtaincall.tech/pkg/models/sqlite"

	_ "github.com/mattn/go-sqlite3"
    "github.com/coreos/go-systemd/activation"
    "golang.org/x/crypto/acme/autocert"
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
    env := requireEnv("CC_ENV")

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

    if env == "PROD" {
	    listeners, err := activation.Listeners()
	    if err != nil {
		  log.Fatal(err)
	    }
	    if len(listeners) != 2 {
		  log.Fatal("Missing systemd socket listeners for ports 80 and 443")
	    }

	    // assumes that listener order mirrors the .socket file
	    httpListener := listeners[0]
	    httpsListener := listeners[1]

	    // redirect http to https
	    go http.Serve(httpListener, http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		  target := "https://" + r.Host + r.URL.Path
	      if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		  }
		  http.Redirect(w, r, target, http.StatusMovedPermanently)
	    }))

	    // setup the Let's Encrypt autocert manager
	    m := &autocert.Manager{
		  Prompt:     autocert.AcceptTOS,
		  HostPolicy: autocert.HostWhitelist("curtaincall.tech"),
	    }

	    // create dir for autocert cache
	    dir := filepath.Join(os.Getenv("HOME"), ".cache", "golang-autocert")
	    if err := os.MkdirAll(dir, 0700); err != nil {
		  log.Printf("warning: autocert Manager not using a cache: %v", err)
	    } else {
		    m.Cache = autocert.DirCache(dir)
	    }

	    s := &http.Server{
		  Handler:      app.routes(),
		  TLSConfig:    m.TLSConfig(),
		  IdleTimeout:  time.Minute,
		  ReadTimeout:  5 * time.Second,
		  WriteTimeout: 10 * time.Second,
	    }

	    log.Println("Listening on curtaincall.tech:443...")
	    log.Fatal(s.ServeTLS(httpsListener, "", ""))

    } else{

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
