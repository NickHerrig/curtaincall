package main

import (
    "errors"
    "fmt"
    "net/http"
    "strconv"

    "curtaincall.tech/pkg/forms"
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

    data := &templateData{Theater: t}

    app.render(w, r, "show.page.tmpl", data)

}

func (app *application) createTheaterForm(w http.ResponseWriter, r *http.Request) {
    app.render(w, r, "create.page.tmpl", &templateData{
        Form: forms.New(nil),
    })

}

func (app *application) createTheater(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.PostForm)
    form.Required("name")

    if !form.Valid() {
        app.render(w, r, "create.page.tmpl", &templateData{Form: form})
        return
    }

    id, err := app.theaters.Insert(form.Get("name"))
    if err != nil {
        app.serverError(w, err)
    }

    http.Redirect(w, r, fmt.Sprintf("/theater/%d", id), http.StatusSeeOther)
}
