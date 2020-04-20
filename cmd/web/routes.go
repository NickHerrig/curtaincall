package main

import (
    "net/http"

    "github.com/bmizerany/pat"
    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := pat.New()

    mux.Get("/", http.HandlerFunc(app.home))

    mux.Get("/theater/create", http.HandlerFunc(app.createTheaterForm))
    mux.Post("/theater/create", http.HandlerFunc(app.createTheater))
    mux.Get("/theater/:id", http.HandlerFunc(app.showTheater))

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static", fileServer))

    return standardMiddleware.Then(mux)
}
