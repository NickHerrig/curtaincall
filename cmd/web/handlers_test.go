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
		{"Valid ID", "/theaters/1", http.StatusOK, []byte("Des Moines Civic Center")},
		{"Non-existent ID", "/theaters/2", http.StatusNotFound, nil},
		{"Negative ID", "/theaters/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/theaters/1.23", http.StatusNotFound, nil},
		{"String ID", "/theaters/desmoines", http.StatusNotFound, nil},
		{"Empty ID", "/theaters/", http.StatusNotFound, nil},
		{"Trailing slash", "/theaters/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want %s; got %s", string(tt.wantBody), string(body))
			}
		})
	}
}
