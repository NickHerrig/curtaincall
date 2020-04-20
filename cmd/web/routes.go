package main

import (
    "net/http"

    "github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

    mux := pat.New()

    mux.Get("/", http.HandlerFunc(app.home))

    mux.Get("/theater/create", http.HandlerFunc(app.createTheaterForm))
    mux.Post("/theater/create", http.HandlerFunc(app.createTheater))
    mux.Get("/theater/:id", http.HandlerFunc(app.showTheater))

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static", fileServer))

    return secureHeaders(mux)
}
