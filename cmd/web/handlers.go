package main

import (
    "errors"
    "fmt"
    "html/template"
    "net/http"
    "strconv"

    "curtaincall.tech/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) showTheater(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
    s, err := app.theaters.Get(id) 
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
    fmt.Fprintf(w, "%v", s)
}

func (app *application) createTheater(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Alow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    name := "The Fabulous Fox Theater"

    id, err := app.theaters.Insert(name)
    if err != nil {
        app.serverError(w, err)
    }

    http.Redirect(w, r, fmt.Sprintf("/theater?id=%d", id), http.StatusSeeOther)

}
