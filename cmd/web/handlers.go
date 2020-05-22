package main

import (
	"errors"
	"net/http"
	"strconv"

	"curtaincall.tech/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", &templateData{})
}

func (app *application) retriveAllTheaters(w http.ResponseWriter, r *http.Request) {

	t, err := app.theaters.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Theaters: t}
	app.render(w, r, "all.theaters.page.tmpl", data)
}

func (app *application) retriveTheater(w http.ResponseWriter, r *http.Request) {

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

func (app *application) createTheater(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CREATE THEATER"))
}

func (app *application) deleteTheater(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE THEATER"))
}

func (app *application) patchTheater(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PATCH THEATER"))
}

func (app *application) updateTheater(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UPDATE THEATER"))
}

func (app *application) retriveShow(w http.ResponseWriter, r *http.Request) {

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
