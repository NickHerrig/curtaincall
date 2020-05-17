package main

import (
	"errors"
	"net/http"
	"strconv"

	"curtaincall.tech/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	t, err := app.theaters.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Theaters: t}
	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showTheater(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	t, err := app.theaters.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	s, err := app.shows.Latest(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.theater.page.tmpl", &templateData{
		Theater: t,
		Shows:   s,
	})

}

func (app *application) showShow(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.shows.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.show.page.tmpl", &templateData{
		Show: s,
	})

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
