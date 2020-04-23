package main

import (
    "net/http"

    "github.com/bmizerany/pat"
    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    dynamicMiddleware := alice.New(app.session.Enable)

    mux := pat.New()

    mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

    mux.Get("/theater/create", dynamicMiddleware.ThenFunc(app.createTheaterForm))
    mux.Post("/theater/create", dynamicMiddleware.ThenFunc(app.createTheater))
    mux.Get("/theater/:id", dynamicMiddleware.ThenFunc(app.showTheater))

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static", fileServer))

    return standardMiddleware.Then(mux)
}
