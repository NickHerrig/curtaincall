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

    s, err := app.theaters.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    for _, theater := range s {
        fmt.Fprintf(w, "%v\n", theater)
    }

//    files := []string{
//        "./ui/html/home.page.tmpl",
//        "./ui/html/base.layout.tmpl",
//        "./ui/html/footer.partial.tmpl",
//    }
//
//    ts, err := template.ParseFiles(files...)
//    if err != nil {
//        app.serverError(w, err)
//        return
//    }
//
//    err = ts.Execute(w, nil)
//    if err != nil {
//        app.serverError(w, err)
//    }
}

func (app *application) showTheater(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

    files := []string{
        "./ui/html/show.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.Execute(w, s)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) createTheaterForm(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create a new snippet....FORMMMMM"))

}

func (app *application) createTheater(w http.ResponseWriter, r *http.Request) {

    name := "The Fabulous Fox Theater"

    id, err := app.theaters.Insert(name)
    if err != nil {
        app.serverError(w, err)
    }

    http.Redirect(w, r, fmt.Sprintf("/theater/%d", id), http.StatusSeeOther)
}
