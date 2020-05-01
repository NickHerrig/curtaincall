package main

import (
    "bytes"
    "net/http"
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

            if code != tt.wantCode{
                t.Errorf("want %d; got %d", tt.wantCode, code)
            }
            if !bytes.Contains(body, tt.wantBody) {
                t.Errorf("want %d; got %d", tt.wantBody, body)
            }
        })
    }
}
