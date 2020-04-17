package main

import (
    "fmt"
    "html/template"
    "net/http"
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
