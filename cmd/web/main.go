package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main() {

    addr := flag.String("addr", ":8080", "HTTP network address")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  mux,
    }

    infoLog.Printf("Starting Curtain Call on port %s", *addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
