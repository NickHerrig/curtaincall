package main

import "net/http"

func (app *application) routes() *http.ServeMux {

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    return mux
}
