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

func (app *application) createTheaterForm(w http.ResponseWriter, r *http.Request) {
    app.render(w, r, "create.theater.page.tmpl", &templateData{
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

    app.session.Put(r, "flash", "Theater successfully created!")

    http.Redirect(w, r, fmt.Sprintf("/theater/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
    app.render(w, r, "signup.page.tmpl", &templateData{
        Form: forms.New(nil),
        })
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.PostForm)
    form.Required("name", "email", "password")
    form.MaxLength("name", 255)
    form.MaxLength("email", 255)
    form.MatchesPattern("email", forms.EmailRX)
    form.MinLength("password", 10)

    if !form.Valid() {
        app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
        return
    }

    err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
    if err != nil {
        if errors.Is(err, models.ErrDuplicateEmail) {
            form.Errors.Add("email", "Address is already in use")
            app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
        } else {
            app.serverError(w, err)
        }
        return
    }
    app.session.Put(r, "flash", "Your signup was successful. Please log in.")
    http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
    app.render(w, r, "login.page.tmpl", &templateData{
        Form: forms.New(nil),
        })
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.PostForm)
    id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
    if err != nil {
        if errors.Is(err, models.ErrInvalidCredentials) {
            form.Errors.Add("generic", "Email or Password is incorrect")
            app.render(w, r, "login.page.tmpl", &templateData{Form: form})
        } else {
            app.serverError(w, err)
        }
        return
    }

    app.session.Put(r, "authenticatedUserID", id)
    http.Redirect(w, r, "/theater/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
    app.session.Remove(r, "authenticatedUserID")
    app.session.Put(r, "flash", "You've been logged out successfully.")
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}
