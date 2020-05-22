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

	mux.Get("/theaters", http.HandlerFunc(app.retriveAllTheaters))
	mux.Post("/theaters", http.HandlerFunc(app.createTheater))
	mux.Get("/theaters/:id", http.HandlerFunc(app.retriveTheater))
	mux.Del("/theaters/:id", http.HandlerFunc(app.deleteTheater))
	mux.Patch("/theaters/:id", http.HandlerFunc(app.patchTheater))
	mux.Put("/theaters/:id", http.HandlerFunc(app.updateTheater))


	mux.Get("/shows/:id", http.HandlerFunc(app.retriveShow))
	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
