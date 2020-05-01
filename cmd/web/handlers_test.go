package main

import (
    "bytes"
    "net/http"
    "net/url"
    "testing"
)

func TestPing(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    code, _, body := ts.get(t, "/ping")

    if code != http.StatusOK {
        t.Errorf("want %d; got %d", http.StatusOK, code)
    }

    if string(body) != "OK" {
        t.Errorf("want body to equal %q", "OK")
    }
}

func TestShowTheater(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    tests := []struct {
        name     string
        urlPath  string
        wantCode int
        wantBody []byte
    }{
         {"Valid ID", "/theater/1", http.StatusOK, []byte("Des Moines Civic Center")},
         {"Non-existent ID", "/theater/2", http.StatusNotFound, nil},
         {"Negative ID", "/theater/-1", http.StatusNotFound, nil},
         {"Decimal ID", "/theater/1.23", http.StatusNotFound, nil},
         {"String ID", "/theater/desmoines", http.StatusNotFound, nil},
         {"Empty ID", "/theater/", http.StatusNotFound, nil},
         {"Trailing slash", "/theater/1/", http.StatusNotFound, nil},
     }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            code, _, body := ts.get(t, tt.urlPath)

            if code != tt.wantCode {
                t.Errorf("want %d; got %d", tt.wantCode, code)
            }
            if !bytes.Contains(body, tt.wantBody) {
                t.Errorf("want %d; got %d", tt.wantBody, body)
            }
        })
    }
}

func TestSignupUser(t *testing.T) {
    app := newTestApplication(t)
    ts := newTestServer(t, app.routes())
    defer ts.Close()

    _, _, body := ts.get(t, "/user/signup")
    csrfToken := extractCSRFToken(t, body)

    tests := []struct {
        name         string
        userName     string
        userEmail    string
        userPassword string
        csrfToken    string
        wantCode     int
        wantBody     []byte
    }{
         {"Valid submission", "nick", "nick@gmail.com", "validPa$$word", csrfToken, http.StatusSeeOther, nil},
         {"Empty name", "", "nick@gmail.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
         {"Empty email", "nick", "", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
         {"Empty password", "nick", "nick@gmail.com", "", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
         {"Invalid email (incomplete domain)", "nick", "nick@gmail.", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
         {"Invalid email (missing @)", "nick", "nickgmail.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
         {"Invalid email (missing local part)", "nick", "@gmail.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
         {"Short password", "nick", "nick@gmail.com", "Pass", csrfToken, http.StatusOK, []byte("This field is too short")},
         {"Duplicate email", "nick", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Address is already in use")},
         {"Invalid CSRF Token", "", "", "", "wrongtoken", http.StatusBadRequest, nil},
     }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            form := url.Values{}
            form.Add("name", tt.userName)
            form.Add("email", tt.userEmail)
            form.Add("password", tt.userPassword)
            form.Add("csrf_token", tt.csrfToken)

            code, _, body := ts.postForm(t, "/user/signup", form)

            if code != tt.wantCode {
                t.Errorf("want %d; got %d", tt.wantCode, code)
            }
            if !bytes.Contains(body, tt.wantBody) {
                t.Errorf("want %q; got %s", tt.wantBody, body)
            }
        })
    }
}
