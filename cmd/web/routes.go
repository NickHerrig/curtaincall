package main

import (
    "net/http"

    "github.com/bmizerany/pat"
    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    dynamicMiddleware := alice.New(app.session.Enable, noSurf)

    mux := pat.New()
    mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
    mux.Get("/theater/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createTheaterForm))
    mux.Post("/theater/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createTheater))
    mux.Get("/theater/:id", dynamicMiddleware.ThenFunc(app.showTheater))

    mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
    mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
    mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
    mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
    mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static", fileServer))

    return standardMiddleware.Then(mux)
}
